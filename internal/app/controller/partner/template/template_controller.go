package template

import (
	ent "gupshup-gui/internal/app/model/partner/template"
)

// AppController define as ações disponíveis para manipular os apps da Gupshup
type TemplateController interface {
	GetTemplates(appID string) ([]ent.PartnerTemplate, error)
}
