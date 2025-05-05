package template

import (
	templateController "gupshup-gui/internal/app/controller/partner/template"
	authService "gupshup-gui/internal/app/service/auth"
	partnerService "gupshup-gui/internal/app/service/partner"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(r *gin.Engine) {
	authSvc := authService.NewLoginService()
	partnerSvc := partnerService.NewPartnerService(authSvc)

	//appCtrl := appController.NewAppController(partnerSvc)
	templateCtrl := templateController.NewAppController(partnerSvc)

	//appHandler := app.NewAppHandler(appCtrl)
	templateHandler := NewTemplateHandler(templateCtrl)

	group := r.Group("/app")
	group.GET("/apps/:app_id/templates", templateHandler.GetTemplatesHandler)
	group.GET("/apps/:app_id/templates/:template_id", templateHandler.GetTemplateByIDHandler)
	group.POST("/apps/:app_id/templates", templateHandler.CreateTemplateTextHandler)
	group.POST("/upload/image/:app_id", templateHandler.UploadImageHandler)
}
