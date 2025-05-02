package partner

import "gupshup-gui/internal/app/service/partner/app"

// PartnerService define os sub-serviços disponíveis para o domínio `partner`
type PartnerService interface {
	AppService() app.PartnerAppService
	// Futuro: TemplateService() template.PartnerTemplateService
	// Futuro: ChannelService() channel.PartnerChannelService
}
