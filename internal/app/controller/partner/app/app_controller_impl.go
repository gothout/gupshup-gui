package app

import (
	ent "gupshup-gui/internal/app/model/partner/app"
	"gupshup-gui/internal/app/service/partner"
)

type appControllerImpl struct {
	service partner.PartnerService
}

// NewAppController retorna uma instância do AppController
func NewAppController(service partner.PartnerService) AppController {
	return &appControllerImpl{
		service: service,
	}
}

func (c *appControllerImpl) GetApps() (*ent.PartnerAppsResponse, error) {
	return c.service.AppService().GetApps()
}

func (c *appControllerImpl) GetAppToken(appID string) (*ent.PartnerAppToken, error) {
	return c.service.AppService().GetAppToken(appID)
}
