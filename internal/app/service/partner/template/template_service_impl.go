package template

import (
	"encoding/json"
	"fmt"
	"gupshup-gui/internal/app/model/partner/template"
	"gupshup-gui/internal/app/service/auth"
	appService "gupshup-gui/internal/app/service/partner/app"
	"gupshup-gui/package/configuration/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type templateServiceImpl struct {
	auth auth.LoginService
}

func NewTemplateService(auth auth.LoginService) TemplateService {
	return &templateServiceImpl{
		auth: auth,
	}
}

func (s *templateServiceImpl) GetTemplates(appID string) ([]template.PartnerTemplate, error) {
	appToken, err := appService.NewPartnerAppService(s.auth).GetAppToken(appID)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%spartner/app/%s/templates", config.URLPartner, appID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+appToken.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %w", err)
	}
	defer resp.Body.Close()

	var response template.TemplateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	return response.Templates, nil

}

func (s *templateServiceImpl) GetTemplateByID(appID, templateID string) (*template.PartnerTemplate, error) {
	// Obtém o token da aplicação (Bearer token para autenticação)
	appToken, err := appService.NewPartnerAppService(s.auth).GetAppToken(appID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter token da aplicação: %w", err)
	}

	// Monta a URL da API para buscar o template
	url := fmt.Sprintf("%swa/app/%s/template/%s", config.URLPartner, appID, templateID)

	// Cria a requisição HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %w", err)
	}

	// Define os headers de autenticação e conteúdo
	req.Header.Set("Authorization", "Bearer "+appToken.Token)
	req.Header.Set("Content-Type", "application/json")

	// Envia a requisição
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta HTTP foi bem-sucedida (200 OK)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("falha ao buscar template. Status: %d. Resposta: %s", resp.StatusCode, string(body))
	}

	// Decodifica o JSON da resposta no struct PartnerTemplate
	var result template.PartnerTemplate
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %w", err)
	}

	return &result, nil
}

// CreateTemplateText cria um template de texto para a aplicação especificada pelo appID.
func (s *templateServiceImpl) CreateTemplateText(appID string, tpl template.TemplateCreateRequest) (*template.TemplateCreateRequest, error) {
	// 1. Obtem token da aplicação
	appToken, err := appService.NewPartnerAppService(s.auth).GetAppToken(appID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter token da aplicação: %w", err)
	}

	// 2. Garante os campos obrigatórios
	tpl.EnableSample = true
	tpl.AllowTemplateCategoryChange = true

	// 3. Monta o corpo como form-urlencoded
	form := url.Values{}
	form.Set("elementName", tpl.ElementName)
	form.Set("languageCode", tpl.LanguageCode)
	form.Set("category", tpl.Category)
	form.Set("templateType", tpl.TemplateType)
	form.Set("vertical", tpl.Vertical)
	form.Set("header", tpl.Header)
	form.Set("content", tpl.Content)
	form.Set("footer", tpl.Footer)
	form.Set("example", tpl.Example)
	form.Set("exampleHeader", tpl.ExampleHeader)
	form.Set("enableSample", "true")
	form.Set("allowTemplateCategoryChange", "true")

	// 👇 Serializa os botões se houverem
	if len(tpl.Buttons) > 0 {
		buttonsBytes, err := json.Marshal(tpl.Buttons)
		if err != nil {
			return nil, fmt.Errorf("erro ao serializar botões: %w", err)
		}
		form.Set("buttons", string(buttonsBytes))
	}

	// 4. Cria requisição
	url := fmt.Sprintf("%spartner/app/%s/templates", config.URLPartner, appID)
	req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+appToken.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 5. Envia
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %w", err)
	}
	defer resp.Body.Close()

	// 6. Valida resposta
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("erro na criação do template: %s", string(body))
	}

	// Ignora o decode da resposta — você confia que deu certo
	io.Copy(io.Discard, resp.Body) // consome o body só pra fechar corretamente

	return &tpl, nil // Retorna o que foi enviado
}
