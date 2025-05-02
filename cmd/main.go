package main

import (
	"github.com/gin-gonic/gin"

	authHandler "gupshup-gui/internal/app/handler/auth"
	appHandler "gupshup-gui/internal/app/handler/partner/app"
	appTemplateHandler "gupshup-gui/internal/app/handler/partner/template"
	serverMiddleware "gupshup-gui/package/middleware/server"
)

func main() {
	// 🌐 Inicia o servidor HTTP com o framework Gin
	r := gin.Default()

	// 🛣️ Cada handler se encarrega de registrar suas rotas e montar controller + service
	authHandler.RegisterAuthRoutes(r)
	appHandler.RegisterAppRoutes(r)
	appTemplateHandler.RegisterAppRoutes(r)

	// ❌ Middleware para rotas não encontradas
	r.NoRoute(serverMiddleware.NotFoundMiddleware())

	// 🚀 Inicia o servidor na porta 8080
	r.Run(":8080")
}
