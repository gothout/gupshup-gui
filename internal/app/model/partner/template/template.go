package template

import "encoding/json"

type Meta struct {
	Example  string `json:"example"`
	MediaID  string `json:"mediaId,omitempty"`
	MediaURL string `json:"mediaUrl,omitempty"`
}

type ContainerMeta struct {
	AppID                       string `json:"appId"`
	Data                        string `json:"data"`
	SampleText                  string `json:"sampleText"`
	EnableSample                bool   `json:"enableSample"`
	MediaID                     string `json:"mediaId,omitempty"`
	MediaURL                    string `json:"mediaUrl,omitempty"`
	EditTemplate                bool   `json:"editTemplate"`
	AllowTemplateCategoryChange bool   `json:"allowTemplateCategoryChange"`
	AddSecurityRecommendation   bool   `json:"addSecurityRecommendation"`
	CorrectCategory             string `json:"correctCategory,omitempty"`
}

type PartnerTemplate struct {
	ID               string          `json:"id"`
	AppID            string          `json:"appId"`
	ElementName      string          `json:"elementName"`
	Category         string          `json:"category"`
	TemplateType     string          `json:"templateType"`
	Status           string          `json:"status"`
	LanguageCode     string          `json:"languageCode"`
	Namespace        string          `json:"namespace"`
	ExternalID       string          `json:"externalId"`
	Data             string          `json:"data"`
	Vertical         string          `json:"vertical"`
	ModifiedOn       int64           `json:"modifiedOn"`
	CreatedOn        int64           `json:"createdOn"`
	MetaRaw          json.RawMessage `json:"meta"`
	ContainerMetaRaw json.RawMessage `json:"containerMeta"`
	WabaID           string          `json:"wabaId"`
	LanguagePolicy   string          `json:"languagePolicy"`
	Priority         int             `json:"priority"`
	Stage            string          `json:"stage"`
	Retry            int             `json:"retry"`
	Quality          string          `json:"quality"`
	InternalType     int             `json:"internalType"`
	InternalCategory int             `json:"internalCategory"`
	ButtonSupported  string          `json:"buttonSupported,omitempty"`
}

type TemplateResponse struct {
	Status    string            `json:"status"`
	Templates []PartnerTemplate `json:"templates"`
}
