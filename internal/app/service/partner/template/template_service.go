package template

import (
	ent "gupshup-gui/internal/app/model/partner/template"
)

type TemplateService interface {
	GetTemplates(appID string) ([]ent.PartnerTemplate, error)
	GetTemplateByID(appID, templateID string) (*ent.PartnerTemplate, error)
	CreateTemplateText(appID string, template ent.TemplateCreateRequest) (*ent.TemplateCreateRequest, error)
	CreateTemplateImage(appID string, imagePath string, template ent.TemplateCreateRequest) (*ent.TemplateCreateRequest, error)
	UploadImageForTemplate(appID string, filePath string) (imageCode string, err error)
}
