package app

import (
	ent "gupshup-gui/internal/app/model/partner/app"
)

// AppController define as ações disponíveis para manipular os apps da Gupshup
type AppController interface {
	GetApps() (*ent.PartnerAppsResponse, error)
}
