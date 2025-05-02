package app

import (
	"net/http"

	controller "gupshup-gui/internal/app/controller/partner/app"
	restError "gupshup-gui/package/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	controller controller.AppController
}

// NewAppHandler cria uma nova instância do handler com controller injetado
func NewAppHandler(ctrl controller.AppController) *AppHandler {
	return &AppHandler{controller: ctrl}
}

// GetAppsHandler chama o controller e responde com sucesso
func (h *AppHandler) GetAppsHandler(c *gin.Context) {
	apps, err := h.controller.GetApps()
	if err != nil {
		restError.NewInternalServerError(err.Error(), []restError.Causes{})
		return
	}

	c.JSON(http.StatusOK, apps)
}
