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
	Buttons          string          `json:"buttons,omitempty"`
}

// TemplateCreateRequest representa a estrutura para criação de um template de mensagem.
type TemplateCreateRequest struct {
	ElementName  string `json:"elementName" binding:"required,min=3,max=50,elementname"` // Nome do template (único por namespace da WABA)
	Vertical     string `json:"vertical" binding: "required,min=3,max=50,vertical"`      // Nome bonito do template.
	LanguageCode string `json:"languageCode"`                                            // Código de idioma, ex: en_US, pt_BR
	Category     string `json:"category"`                                                // Categoria: AUTHENTICATION, MARKETING ou UTILITY
	TemplateType string `json:"templateType"`                                            // Tipo do template: TEXT, IMAGE, LOCATION, PRODUCT, CATALOG, CAROUSEL, VIDEO, DOCUMENT

	// ⚠️ Templates do tipo Catalog, Carousel e Product estão disponíveis apenas para categorias MARKETING & UTILITY
	// ⚠️ Templates do tipo Catalog, LTO e Carousel não são suportados na API On-Premises

	// Apenas para templates do tipo CAROUSEL
	Cards []Card `json:"cards,omitempty"` // Lista de cartões, necessário apenas se o tipo for CAROUSEL

	// Apenas para templates LTO (Limited Time Offer)
	IsLTO            bool   `json:"isLTO,omitempty"`            // Deve ser true para criar um template LTO
	LimitedOfferText string `json:"limitedOfferText,omitempty"` // Texto de oferta limitada (máx. 16 caracteres)
	HasExpiration    bool   `json:"hasExpiration,omitempty"`    // Se true, requer botão de cópia de código

	// Estrutura do conteúdo
	Header  string `json:"header,omitempty"` // Cabeçalho (máx. 60 caracteres)
	Content string `json:"content"`          // Corpo do template. Máx. 1028 caracteres
	Footer  string `json:"footer,omitempty"` // Rodapé (máx. 60 caracteres)

	// Botões do template
	Buttons []TemplateButton `json:"buttons,omitempty"` // Lista de botões: URL, PHONE_NUMBER, QUICK_REPLY, COPY_CODE, etc.

	// Exemplos
	Example       string `json:"example"`       // Exemplo do conteúdo com preenchimento dos placeholders
	ExampleHeader string `json:"exampleHeader"` // Exemplo do cabeçalho
	EnableSample  bool   `json:"enableSample"`  // Obrigatório para criação de qualquer template

	// Configurações opcionais
	AllowTemplateCategoryChange bool `json:"allowTemplateCategoryChange"`         // Se true, permite que o Meta altere a categoria automaticamente
	AddSecurityRecommendation   bool `json:"addSecurityRecommendation,omitempty"` // Adiciona aviso de segurança (somente AUTHENTICATION)
	CodeExpirationMinutes       int  `json:"codeExpirationMinutes,omitempty"`     // Somente AUTHENTICATION (1-90 min)
}

type Card struct {
	HeaderType   string       `json:"headerType"` // Ex: IMAGE
	MediaURL     string       `json:"mediaUrl,omitempty"`
	MediaID      string       `json:"mediaId,omitempty"`
	ExampleMedia string       `json:"exampleMedia,omitempty"`
	Body         string       `json:"body"`       // Texto principal do card
	SampleText   string       `json:"sampleText"` // Texto exemplo
	Buttons      []CardButton `json:"buttons"`    // Botões do card
}

type CardButton struct {
	Type        string `json:"type"`                  // Ex: URL, QUICK_REPLY
	Text        string `json:"text"`                  // Texto visível no botão
	URL         string `json:"url,omitempty"`         // Necessário para botão do tipo URL
	ButtonValue string `json:"buttonValue,omitempty"` // URL final (se aplicável)
	Suffix      string `json:"suffix,omitempty"`      // Opcional para complementar a URL
}

type TemplateButton struct {
	Type          string   `json:"type"`                     // PHONE_NUMBER, URL, QUICK_REPLY, COPY_CODE, OTP
	Text          string   `json:"text"`                     // Texto do botão
	PhoneNumber   string   `json:"phone_number,omitempty"`   // Para PHONE_NUMBER
	URL           string   `json:"url,omitempty"`            // Para URL
	Example       []string `json:"example,omitempty"`        // Exemplo de URL com variável
	OTPType       string   `json:"otp_type,omitempty"`       // COPY_CODE ou ONE_TAP
	AutofillText  string   `json:"autofill_text,omitempty"`  // Para ONE_TAP
	PackageName   string   `json:"package_name,omitempty"`   // Para ONE_TAP
	SignatureHash string   `json:"signature_hash,omitempty"` // Para ONE_TAP
}

type TemplateResponse struct {
	Status    string            `json:"status"`
	Templates []PartnerTemplate `json:"templates"`
}
