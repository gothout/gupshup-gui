package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundMiddleware intercepta rotas não mapeadas e responde com erro 404 padronizado.
func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() == http.StatusNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"erro":   "Rota não encontrada",
				"status": 404,
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			})
		}
	}
}
