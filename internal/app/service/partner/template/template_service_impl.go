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
