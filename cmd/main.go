package main

import (
	"github.com/gin-gonic/gin"

	// Importa os pacotes internos do app
	controller "gupshup-gui/internal/app/controller/auth"
	handler "gupshup-gui/internal/app/handler/auth"
	service "gupshup-gui/internal/app/service/auth"
	serverMiddleware "gupshup-gui/package/middleware/server"
)

func main() {
	// 🔧 Inicializa a camada de serviço (login e cache de token)
	svc := service.NewLoginService()

	// 🎯 Inicializa o controller e injeta o serviço (controlador orquestra a lógica)
	ctrl := controller.NewLoginController(svc)

	// 🔐 Realiza o login assim que a aplicação inicia e guarda o token no cache
	ctrl.HandleLogin()

	// 🌐 Inicia o servidor HTTP com o framework Gin
	r := gin.Default()

	// 🛣️ Registra as rotas relacionadas à autenticação no grupo /auth
	// Ex: GET /auth/token → retorna token atual
	handler.RegisterAuthRoutes(r, ctrl)
	// Middleware para 404
	r.NoRoute(serverMiddleware.NotFoundMiddleware())

	// 🚀 Inicia o servidor na porta 8080
	r.Run(":8080")
}
