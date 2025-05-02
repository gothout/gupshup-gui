package auth

import (
	controller "gupshup-gui/internal/app/controller/auth"
	service "gupshup-gui/internal/app/service/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	svc := service.NewLoginService()
	ctrl := controller.NewLoginController(svc)
	ctrl.HandleLogin()
	handler := NewTokenHandler(ctrl)

	authGroup := r.Group("/auth")
	authGroup.GET("/token", handler.GetToken)
}
