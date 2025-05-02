package app

import (
	controller "gupshup-gui/internal/app/controller/partner/app"
	authService "gupshup-gui/internal/app/service/auth"
	partnerService "gupshup-gui/internal/app/service/partner"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(r *gin.Engine) {
	// Monta o service → controller → handler internamente
	authSvc := authService.NewLoginService()
	partnerSvc := partnerService.NewPartnerService(authSvc)
	appCtrl := controller.NewAppController(partnerSvc)
	handler := NewAppHandler(appCtrl)

	group := r.Group("/partner")
	group.GET("/apps", handler.GetAppsHandler)
}
