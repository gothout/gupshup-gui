package app

import (
	"encoding/json"
	"errors"
	"fmt"
	ent "gupshup-gui/internal/app/model/partner/app"
	authService "gupshup-gui/internal/app/service/auth"
	config "gupshup-gui/package/configuration/config"
	"net/http"
	"sync"
)

// Struct que representa o token do app

// Estrutura do service
type partnerAppServiceImpl struct {
	auth           authService.LoginService
	appTokenCache  map[string]*ent.PartnerAppToken
	appTokenLocker sync.RWMutex
}

func NewPartnerAppService(auth authService.LoginService) PartnerAppService {
	return &partnerAppServiceImpl{
		auth:          auth,
		appTokenCache: make(map[string]*ent.PartnerAppToken),
	}
}

// Busca todos os apps
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

	resp, err := http.DefaultClient.Do(req)
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

// Busca o token do app, usando cache
func (s *partnerAppServiceImpl) GetAppToken(appId string) (*ent.PartnerAppToken, error) {
	s.appTokenLocker.RLock()
	token, found := s.appTokenCache[appId]
	s.appTokenLocker.RUnlock()

	if found {
		return token, nil
	}

	return s.fetchAppToken(appId)
}

// Força atualização do token
func (s *partnerAppServiceImpl) RefreshAppToken(appId string) (*ent.PartnerAppToken, error) {
	return s.fetchAppToken(appId)
}

// Faz a requisição de token do app e salva no cache
func (s *partnerAppServiceImpl) fetchAppToken(appId string) (*ent.PartnerAppToken, error) {
	tokenCache, found := s.auth.GetCachedToken()
	if !found {
		return nil, errors.New("token principal não encontrado ou expirado")
	}

	url := fmt.Sprintf("%spartner/app/%s/token", config.URLPartner, appId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição de token do app: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenCache.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição de token do app: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, errors.New("limite de requisições atingido (429)")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta da API: %d", resp.StatusCode)
	}

	var result struct {
		Token struct {
			Token string `json:"token"`
		} `json:"token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta do token do app: %w", err)
	}

	if result.Token.Token == "" {
		return nil, errors.New("token do app não encontrado na resposta")
	}

	appToken := &ent.PartnerAppToken{AppID: appId, Token: result.Token.Token}
	s.appTokenLocker.Lock()
	s.appTokenCache[appId] = appToken
	s.appTokenLocker.Unlock()

	return appToken, nil
}
