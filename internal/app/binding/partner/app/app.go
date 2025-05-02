package app

import (
	"gupshup-gui/internal/app/model/partner/app"
)

type AppIDInput struct {
	AppID string `json:"appid" binding:"required"`
}

func (u *AppIDInput) ToAppToken() *app.PartnerAppToken {
	return &app.PartnerAppToken{
		AppID: u.AppID,
	}
}
