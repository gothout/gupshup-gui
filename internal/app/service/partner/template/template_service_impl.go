package template

import (
	"encoding/json"
	"fmt"
	"gupshup-gui/internal/app/model/partner/template"
	"gupshup-gui/internal/app/service/auth"
	appService "gupshup-gui/internal/app/service/partner/app"
	"gupshup-gui/package/configuration/config"
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
