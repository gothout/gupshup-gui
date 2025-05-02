package partner

import (
	"gupshup-gui/internal/app/service/auth"
	"gupshup-gui/internal/app/service/partner/app"
	"gupshup-gui/internal/app/service/partner/template"
)

type partnerServiceImpl struct {
	appService      app.PartnerAppService
	templateService template.TemplateService
}

// NewPartnerService cria um agregador dos sub-serviços do domínio `partner`
func NewPartnerService(authSvc auth.LoginService) PartnerService {
	return &partnerServiceImpl{
		appService:      app.NewPartnerAppService(authSvc),
		templateService: template.NewTemplateService(authSvc),
	}
}

// AppService retorna o serviço de manipulação de apps da Gupshup
func (s *partnerServiceImpl) AppService() app.PartnerAppService {
	return s.appService
}

// TemplateService retorna o serviço de templates da Gupshup
func (s *partnerServiceImpl) TemplateService() template.TemplateService {
	return s.templateService
}
