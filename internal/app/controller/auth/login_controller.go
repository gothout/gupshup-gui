package auth

import (
	model "gupshup-gui/internal/app/model/auth"
)

type LoginController interface {
	HandleLogin()                          // Login automático no startup
	FetchToken() (*model.TokenCache, bool) // Usado pelo handler para retornar token atual
}
