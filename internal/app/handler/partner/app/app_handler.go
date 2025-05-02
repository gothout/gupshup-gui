package app

import (
	"net/http"

	binding "gupshup-gui/internal/app/binding/partner/app"
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
		rest := restError.NewInternalServerError(err.Error(), []restError.Causes{})
		c.JSON(rest.Code, rest)
		return
	}

	c.JSON(http.StatusOK, apps)
}

// GetAppTokenHandler faz o bind do appID via JSON e retorna o token do app
func (h *AppHandler) GetAppTokenHandler(c *gin.Context) {
	var input binding.AppIDInput
	input.AppID = c.Param("app_id")

	if input.AppID == "" {
		rest := restError.NewBadRequestValidationError("App ID não informado", []restError.Causes{
			{Field: "app_id", Message: "O ID do app é obrigatório na URL"},
		})
		c.JSON(rest.Code, rest)
		return
	}

	appToken, err := h.controller.GetAppToken(input.AppID)
	if err != nil {
		rest := restError.NewInternalServerError(err.Error(), []restError.Causes{})
		c.JSON(rest.Code, rest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token_app": appToken})
}
