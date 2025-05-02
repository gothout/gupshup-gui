package partner

import (
	"gupshup-gui/internal/app/service/partner/app"
	"gupshup-gui/internal/app/service/partner/template"
)

// PartnerService define os sub-serviços disponíveis para o domínio `partner`
type PartnerService interface {
	AppService() app.PartnerAppService
	TemplateService() template.TemplateService
	// Futuro: TemplateService() template.PartnerTemplateService
	// Futuro: ChannelService() channel.PartnerChannelService
}
