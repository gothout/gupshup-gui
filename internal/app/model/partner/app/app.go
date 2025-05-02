package app

type PartnerAppToken struct {
	AppID string
	Token string
}

type PartnerApp struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Phone          string  `json:"phone,omitempty"` // Alguns apps não têm phone
	CustomerID     string  `json:"customerId"`
	Live           bool    `json:"live"`
	PartnerID      int     `json:"partnerId"`
	CreatedOn      int64   `json:"createdOn"`
	ModifiedOn     int64   `json:"modifiedOn"`
	PartnerCreated bool    `json:"partnerCreated"`
	CxpEnabled     bool    `json:"cxpEnabled"`
	PartnerUsage   bool    `json:"partnerUsage"`
	Stopped        bool    `json:"stopped"`
	Healthy        bool    `json:"healthy"`
	Cap            float64 `json:"cap"`
}

type PartnerAppsResponse struct {
	Status          string       `json:"status"`
	PartnerAppsList []PartnerApp `json:"partnerAppsList"`
}
