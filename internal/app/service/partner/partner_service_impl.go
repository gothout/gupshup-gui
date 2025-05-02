package partner

import (
	"gupshup-gui/internal/app/service/auth"
	"gupshup-gui/internal/app/service/partner/app"
)

type partnerServiceImpl struct {
	appService app.PartnerAppService
}

// NewPartnerService cria um agregador dos sub-serviços do domínio `partner`
func NewPartnerService(authSvc auth.LoginService) PartnerService {
	return &partnerServiceImpl{
		appService: app.NewPartnerAppService(authSvc),
	}
}

// AppService retorna o serviço de manipulação de apps da Gupshup
func (s *partnerServiceImpl) AppService() app.PartnerAppService {
	return s.appService
}
