package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	model "gupshup-gui/internal/app/model/auth"
	config "gupshup-gui/package/configuration/config"
	env "gupshup-gui/package/configuration/env"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
)

var gupshupTokenCache = cache.New(1*time.Hour, 10*time.Minute)

type loginServiceImpl struct {
	cache *cache.Cache
}

func NewLoginService() LoginService {
	/*******************************************************/
	// NewLoginService retorna uma instancia do LoginService
	// que ira  utiliza o cache global para armazenar o token
	// de autenticacao da Gupshup.
	/*******************************************************/
	return &loginServiceImpl{cache: gupshupTokenCache}
}

func (s *loginServiceImpl) Authenticate(p model.Partner) (*model.TokenCache, error) {
	// Verifica se já tem token em cache
	if token, found := s.cache.Get("gupshup_token"); found {
		return token.(*model.TokenCache), nil
	}

	// Monta corpo da requisição
	form := url.Values{}
	form.Add("email", p.Email)
	form.Add("password", p.Password)

	req, err := http.NewRequest("POST", config.URLPartner+"partner/account/login", bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Faz a requisição
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Trata erro HTTP
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro no login: %d", res.StatusCode)
	}

	// Lê o JSON de retorno
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	tokenStr, ok := result["token"].(string)
	if !ok {
		return nil, errors.New("token não encontrado na resposta")
	}

	// Cria o modelo
	token := &model.TokenCache{
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	}

	// Salva no cache em memória
	s.cache.Set("gupshup_token", token, 1*time.Hour)
	return token, nil
}

func (s *loginServiceImpl) GetCachedToken() (*model.TokenCache, bool) {
	token, found := s.cache.Get("gupshup_token")
	if !found {
		return nil, false
	}
	return token.(*model.TokenCache), true
}

func (s *loginServiceImpl) ForceLogin() (*model.TokenCache, error) {
	partner := model.Partner{
		Email:    env.GetEmail(),
		Password: env.GetSenha(),
	}
	return s.Authenticate(partner)
}
