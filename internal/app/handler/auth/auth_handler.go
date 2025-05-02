package auth

import (
	controller "gupshup-gui/internal/app/controller/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	controller controller.LoginController
}

func NewTokenHandler(ctrl controller.LoginController) *TokenHandler {
	return &TokenHandler{controller: ctrl}
}

func (h *TokenHandler) GetToken(c *gin.Context) {
	tokenCache, found := h.controller.FetchToken()
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token não encontrado ou expirado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenCache.Token,
		"expires_at": tokenCache.ExpiresAt,
	})
}
