package template

import (
	ent "gupshup-gui/internal/app/model/partner/template"
)

type TemplateService interface {
	GetTemplates(appID string) ([]ent.PartnerTemplate, error)
	GetTemplateByID(appID, templateID string) (*ent.PartnerTemplate, error)
	CreateTemplateText(appID string, template ent.TemplateCreateRequest) (*ent.TemplateCreateRequest, error)
}
