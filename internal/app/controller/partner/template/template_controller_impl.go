package template

import (
	ent "gupshup-gui/internal/app/model/partner/template"
	"gupshup-gui/internal/app/service/partner"
)

type templateControllerImpl struct {
	service partner.PartnerService
}

// NewAppController retorna uma instância do AppController
func NewAppController(service partner.PartnerService) TemplateController {
	return &templateControllerImpl{
		service: service,
	}
}

func (c *templateControllerImpl) GetTemplates(appID string) ([]ent.PartnerTemplate, error) {
	return c.service.TemplateService().GetTemplates(appID)
}

func (c *templateControllerImpl) GetTemplateByID(appID, templateID string) (*ent.PartnerTemplate, error) {
	return c.service.TemplateService().GetTemplateByID(appID, templateID)
}

func (c *templateControllerImpl) CreateTemplateText(appID string, template ent.TemplateCreateRequest) (*ent.TemplateCreateRequest, error) {
	return c.service.TemplateService().CreateTemplateText(appID, template)
}
