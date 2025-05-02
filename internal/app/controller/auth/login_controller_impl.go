package auth

import (
	"fmt"
	model "gupshup-gui/internal/app/model/auth"
	service "gupshup-gui/internal/app/service/auth"
	"gupshup-gui/package/configuration/env"
	"os"
)

type loginControllerImpl struct {
	service service.LoginService
}

func NewLoginController(svc service.LoginService) LoginController {
	return &loginControllerImpl{service: svc}
}

func (c *loginControllerImpl) HandleLogin() {

	env.LoadEnv()
	if os.Getenv("EMAIL") == "" && os.Getenv("SENHA") == "" {
		fmt.Println("O usuario e senha nao foram informados no .env")
		os.Exit(1)
	}
	partner := model.Partner{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("SENHA"),
	}

	token, err := c.service.Authenticate(partner)
	if err != nil {
		fmt.Println("Erro ao autenticar:", err)
		return
	}
	fmt.Println("Token atual:", token)
}

func (c *loginControllerImpl) FetchToken() (*model.TokenCache, bool) {
	return c.service.GetCachedToken()
}
