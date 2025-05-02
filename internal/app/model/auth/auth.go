package auth

type Partner struct {
	Email    string
	Password string
}

type TokenCache struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
