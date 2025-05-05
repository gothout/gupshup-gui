package template

import (
	"fmt"
	"net/http"

	binding "gupshup-gui/internal/app/binding/partner/template"
	"gupshup-gui/internal/app/controller/partner/template"
	internalutil "gupshup-gui/internal/app/handler/partner/template/internal"
	"gupshup-gui/internal/app/service/server/upload"
	"gupshup-gui/package/configuration/rest_err"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type TemplateHandler struct {
	controller template.TemplateController
}

func NewTemplateHandler(ctrl template.TemplateController) *TemplateHandler {
	return &TemplateHandler{controller: ctrl}
}

const MaxUploadSize = 5 << 20 // 5 MB (5 * 1024 * 1024)

func (h *TemplateHandler) GetTemplatesHandler(c *gin.Context) {
	appID := c.Param("app_id")
	if appID == "" {
		rest := rest_err.NewBadRequestValidationError("App ID não informado", []rest_err.Causes{
			{Field: "app_id", Message: "O ID do app é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	templates, err := h.controller.GetTemplates(appID)
	if err != nil {
		rest := rest_err.NewInternalServerError(err.Error(), []rest_err.Causes{})
		c.JSON(rest.Code, rest)
		return
	}

	c.JSON(http.StatusOK, templates)
}

func (h *TemplateHandler) GetTemplateByIDHandler(c *gin.Context) {
	appID := c.Param("app_id")
	if appID == "" {
		rest := rest_err.NewBadRequestValidationError("App ID não informado", []rest_err.Causes{
			{Field: "app_id", Message: "O ID do app é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	templateID := c.Param("template_id")
	if templateID == "" {
		rest := rest_err.NewBadRequestValidationError("Template ID não informado", []rest_err.Causes{
			{Field: "template_id", Message: "O ID do template é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	template, err := h.controller.GetTemplateByID(appID, templateID)
	if err != nil {
		rest := rest_err.NewInternalServerError(err.Error(), []rest_err.Causes{})
		c.JSON(rest.Code, rest)
		return
	}

	c.JSON(http.StatusOK, template)
}

func (h *TemplateHandler) CreateTemplateTextHandler(c *gin.Context) {
	// 1. Captura o app_id da URL
	appID := c.Param("app_id")
	//fmt.Println("AppID recebido:", appID)

	if appID == "" || appID == ":app_id" {
		rest := rest_err.NewBadRequestValidationError("App ID não informado", []rest_err.Causes{
			{Field: "app_id", Message: "O ID do app é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	// 2. Faz o binding do JSON recebido com validação
	var input binding.CreateTemplateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var causes []rest_err.Causes

		// Verifica se é erro de validação de campo
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				causes = append(causes, rest_err.Causes{
					Field:   fieldErr.Field(),
					Message: fmt.Sprintf("Campo inválido: %s (condição: %s)", fieldErr.Value(), fieldErr.Tag()),
				})
			}
		} else {
			// Erro geral de binding
			causes = append(causes, rest_err.Causes{
				Field:   "body",
				Message: err.Error(),
			})
		}

		rest := rest_err.NewBadRequestValidationError("Erro ao validar os campos da requisição", causes)
		c.JSON(rest.Code, rest)
		return
	}

	// 3. Converte para modelo do domínio
	req := input.ToTemplateCreateRequest()

	// 4. Preenche automaticamente os exemplos se necessário
	if req.Example == "" {
		req.Example = internalutil.PreencherExemploComVariaveis(req.Content)
	}
	if req.ExampleHeader == "" && req.Header != "" {
		req.ExampleHeader = internalutil.PreencherExemploComVariaveis(req.Header)
	}

	// 5. Valida e envia para o serviço apropriado com base no tipo
	switch req.TemplateType {
	case "TEXT":
		createdTemplate, err := h.controller.CreateTemplateText(appID, *req)
		if err != nil {
			rest := rest_err.NewInternalServerError("Erro ao criar template", []rest_err.Causes{
				{Field: "controller", Message: err.Error()},
			})
			c.JSON(rest.Code, rest)
			return
		}
		c.JSON(http.StatusCreated, createdTemplate)

	default:
		rest := rest_err.NewBadRequestValidationError("Tipo de template não suportado", []rest_err.Causes{
			{Field: "templateType", Message: req.TemplateType},
		})
		c.JSON(rest.Code, rest)
	}
}

// Enviar imagem para o repositorio local da aplicação.
func (h *TemplateHandler) UploadImageHandler(c *gin.Context) {
	appID := c.Param("app_id")
	if appID == "" {
		rest := rest_err.NewBadRequestValidationError("App ID não informado", []rest_err.Causes{
			{Field: "app_id", Message: "O ID do app é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	// Limita o tamanho total da requisição
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20) // até 20MB no total
	if err := c.Request.ParseMultipartForm(20 << 20); err != nil {
		rest := rest_err.NewBadRequestValidationError("Erro ao ler arquivos", []rest_err.Causes{
			{Field: "form", Message: err.Error()},
		})
		c.JSON(rest.Code, rest)
		return
	}

	form := c.Request.MultipartForm
	files := form.File["file"]

	if len(files) == 0 {
		rest := rest_err.NewBadRequestValidationError("Nenhum arquivo enviado", []rest_err.Causes{
			{Field: "file", Message: "Envie pelo menos um arquivo"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	var savedPaths []string
	for i, file := range files {
		// Valida o tamanho individual
		if file.Size > MaxUploadSize {
			rest := rest_err.NewBadRequestValidationError("Arquivo excede o tamanho permitido", []rest_err.Causes{
				{Field: fmt.Sprintf("file[%d]", i), Message: "Tamanho máximo: 5MB"},
			})
			c.JSON(rest.Code, rest)
			return
		}

		path, err := upload.SaveUploadedFile(file)
		if err != nil {
			rest := rest_err.NewInternalServerError("Erro ao salvar arquivo", []rest_err.Causes{
				{Field: "upload", Message: err.Error()},
			})
			c.JSON(rest.Code, rest)
			return
		}
		savedPaths = append(savedPaths, path)
	}

	c.JSON(http.StatusOK, gin.H{
		"arquivosSalvos": savedPaths,
	})
}
