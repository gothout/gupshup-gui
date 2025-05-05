package template

import (
	"net/http"

	"gupshup-gui/internal/app/controller/partner/template"
	"gupshup-gui/package/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	controller template.TemplateController
}

func NewTemplateHandler(ctrl template.TemplateController) *TemplateHandler {
	return &TemplateHandler{controller: ctrl}
}

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
