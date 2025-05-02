package auth

import "gupshup-gui/internal/app/model/auth"

type LoginService interface {
	Authenticate(p auth.Partner) (*auth.TokenCache, error)
	GetCachedToken() (*auth.TokenCache, bool)
	ForceLogin() (*auth.TokenCache, error)
}
