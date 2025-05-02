package auth

import (
	controller "gupshup-gui/internal/app/controller/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, ctrl controller.LoginController) {
	handler := NewTokenHandler(ctrl)
	authGroup := router.Group("/auth")
	{
		authGroup.GET("/token", handler.GetToken)
	}
}
