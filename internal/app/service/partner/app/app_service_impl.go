package app

import (
	"encoding/json"
	"fmt"
	ent "gupshup-gui/internal/app/model/partner/app"
	authService "gupshup-gui/internal/app/service/auth"
	config "gupshup-gui/package/configuration/config"
	"net/http"
)

type partnerAppServiceImpl struct {
	auth authService.LoginService
}

// NewPartnerAppService cria uma nova instância do serviço de apps
func NewPartnerAppService(auth authService.LoginService) PartnerAppService {
	return &partnerAppServiceImpl{auth: auth}
}

func (s *partnerAppServiceImpl) GetApps() (*ent.PartnerAppsResponse, error) {
	tokenCache, found := s.auth.GetCachedToken()
	if !found {
		return nil, fmt.Errorf("token não encontrado ou expirado")
	}

	url := config.URLPartner + "partner/account/api/partnerApps"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+tokenCache.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %w", err)
	}
	defer resp.Body.Close()

	var result ent.PartnerAppsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	return &result, nil
}
