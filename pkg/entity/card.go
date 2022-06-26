package entity

type CardRequest struct {
	Product  Product `json:"product"`
	AgentId  int     `json:"agent_id"`
	Category string  `json:"category"`
}

type CachedCard struct {
	CardID     string `json:"card_id"`
	CardNumber string `json:"card_number"`
}

// Card type
type Card struct {
	Command                string        `xml:"command" faker:"-" json:"command"`
	DeliveryAgentNumber    string        `xml:"delivery_agent_number" json:"delivery_agent_number"`
	CardID                 string        `xml:"card_id" json:"card_id"`
	CardNumber             string        `xml:"card_number" json:"card_number"`
	CardType               string        `xml:"card_type" json:"card_type"`
	CardCount              string        `xml:"card_count" json:"card_count"`
	Category               string        `xml:"category" json:"category"`
	CardBlankType          string        `xml:"card_blank_type" json:"card_blank_type"`
	CardDeliveryChannel    string        `xml:"card_delivery_channel" json:"card_delivery_channel"`
	CardDeliveryStatus     string        `xml:"card_delivery_status" json:"card_delivery_status"`
	StartDate              string        `xml:"start_date" json:"start_date"`
	ExpirationDate         string        `xml:"expiration_date" json:"expiration_date"`
	CardIssDate            string        `xml:"card_iss_date" json:"card_iss_date"`
	Pvv                    string        `xml:"pvv" json:"pvv"`
	PvkIndex               string        `xml:"pvk_index" json:"pvk_index"`
	PinBlockFormat         string        `xml:"pin_block_format" json:"pin_block_format"`
	StartDateRule          string        `xml:"start_date_rule" json:"start_date_rule"`
	ExpirationDateRule     string        `xml:"expiration_date_rule" json:"expiration_date_rule"`
	PinRequest             string        `xml:"pin_request" json:"pin_request"`
	PersoPriority          string        `xml:"perso_priority" json:"perso_priority"`
	EmbossingRequest       string        `xml:"embossing_request" json:"embossing_request"`
	PinMailerRequest       string        `xml:"pin_mailer_request" json:"pin_mailer_request"`
	ReissueCardID          string        `xml:"reissue_card_id" json:"reissue_card_id"`
	ReissueCardNumber      string        `xml:"reissue_card_number" json:"reissue_card_number"`
	ReissueExpirationDate  string        `xml:"reissue_expiration_date" json:"reissue_expiration_date"`
	ReissueReason          string        `xml:"reissue_reason" json:"reissue_reason"`
	ReissueCommand         string        `xml:"reissue_command" json:"reissue_command"`
	CloneOptionalServices  string        `xml:"clone_optional_services" json:"clone_optional_services"`
	SequentialNumber       string        `xml:"sequential_number" json:"sequential_number"`
	CardState              string        `xml:"card_state" json:"card_state"`
	CardStatus             string        `xml:"card_status" json:"card_status"`
	FlexibleField          FlexibleField `xml:"flexible_field" json:"flexible_field"`
	Cardholder             Cardholder    `xml:"cardholder" json:"cardholder"`
	CardInstance           CardInstance  `xml:"card_instance" json:"card_instance"`
	PrecedingCard          PrecedingCard `xml:"preceding_card" json:"preceding_card"`
	EmbossedLineAdditional string        `xml:"embossed_line_additional" json:"embossed_line_additional"`
	SupplementaryInfo1     string        `xml:"supplementary_info_1" json:"supplementary_info_1"`
	Error                  Error         `xml:"error" json:"error"`
}

type GetCardInfoResponse struct {
	BatchId  string `xml:"batchId" json:"batchId"`
	CardInfo Card   `xml:"cardInfo" json:"cardInfo"`
}
type CardInstance struct {
	SequentialNumber    string `xml:"sequential_number" json:"sequential_number"`
	CardState           string `xml:"card_state" json:"card_state"`
	CardStatus          string `xml:"card_status" json:"card_status"`
	CardIssDate         string `xml:"card_iss_date" json:"card_iss_date"`
	CardStartDate       string `xml:"card_start_date" json:"card_start_date"`
	ExpirationDate      string `xml:"expiration_date" json:"expiration_date"`
	CardholderName      string `xml:"cardholder_name" json:"cardholder_name"`
	CompanyEmbossedName string `xml:"company_embossed_name" json:"company_embossed_name"`
}

type PrecedingCard struct {
	SequentialNumber string `xml:"sequential_number" json:"sequential_number"`
	CardID           string `xml:"card_id" json:"card_id"`
	CardNumber       string `xml:"card_number" json:"card_number"`
	ExpirationDate   string `xml:"expiration_date" json:"expiration_date"`
	Error            Error  `xml:"error" json:"error"`
}
