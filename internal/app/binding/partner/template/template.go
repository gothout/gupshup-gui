package template

import (
	"gupshup-gui/internal/app/model/partner/template"
)

type CreateTemplateInput struct {
	ElementName               string                    `json:"elementName" binding:"required"`
	LanguageCode              string                    `json:"languageCode" binding:"required"`
	Category                  string                    `json:"category" binding:"required"`
	TemplateType              string                    `json:"templateType" binding:"required"`
	Header                    string                    `json:"header,omitempty"`
	Content                   string                    `json:"content" binding:"required"`
	Footer                    string                    `json:"footer,omitempty"`
	Buttons                   []template.TemplateButton `json:"buttons,omitempty"`
	Example                   string                    `json:"example,omitempty"`
	ExampleHeader             string                    `json:"exampleHeader,omitempty"`
	IsLTO                     bool                      `json:"isLTO,omitempty"`
	LimitedOfferText          string                    `json:"limitedOfferText,omitempty"`
	HasExpiration             bool                      `json:"hasExpiration,omitempty"`
	Cards                     []template.Card           `json:"cards,omitempty"`
	CodeExpirationMinutes     int                       `json:"codeExpirationMinutes,omitempty"`
	AddSecurityRecommendation bool                      `json:"addSecurityRecommendation,omitempty"`
}

func (c *CreateTemplateInput) ToTemplateCreateRequest() *template.TemplateCreateRequest {
	return &template.TemplateCreateRequest{
		ElementName:                 c.ElementName,
		LanguageCode:                c.LanguageCode,
		Category:                    c.Category,
		TemplateType:                c.TemplateType,
		Header:                      c.Header,
		Content:                     c.Content,
		Footer:                      c.Footer,
		Buttons:                     c.Buttons,
		Example:                     c.Example,
		ExampleHeader:               c.ExampleHeader,
		IsLTO:                       c.IsLTO,
		LimitedOfferText:            c.LimitedOfferText,
		HasExpiration:               c.HasExpiration,
		Cards:                       c.Cards,
		CodeExpirationMinutes:       c.CodeExpirationMinutes,
		AddSecurityRecommendation:   c.AddSecurityRecommendation,
		EnableSample:                true,
		AllowTemplateCategoryChange: true,
	}
}
