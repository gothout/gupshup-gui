package app

import (
	ent "gupshup-gui/internal/app/model/partner/app"
)

// PartnerAppService define a interface de funcionalidades relacionadas aos apps da Gupshup
type PartnerAppService interface {
	GetApps() (*ent.PartnerAppsResponse, error)
	GetAppToken(appID string) (*ent.PartnerAppToken, error)
	RefreshAppToken(appID string) (*ent.PartnerAppToken, error)
}
