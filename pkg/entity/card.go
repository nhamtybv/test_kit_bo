package entity

import "encoding/xml"

type CardRequest struct {
	Product  Product `json:"product"`
	AgentId  int     `json:"agent_id"`
	Category string  `json:"category"`
}

type CachedCard struct {
	CardID        int    `json:"card_id"`
	CardNumber    string `json:"card_number"`
	ApplicationId string `json:"application_id"`
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
	Customer               *Customer     `xml:"customer" json:"customer"`
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

type ProcessIBGDataRequest struct {
	XMLName             xml.Name     `xml:"ins:processIBGDataRequest"`
	CardID              string       `xml:"iss:card_id,omitempty" json:"card_id"`
	CardUid             string       `xml:"iss:card_uid,omitempty" json:"card_uid"`
	CardNumber          string       `xml:"iss:card_number,omitempty" json:"card_number"`
	ExpirationDate      string       `xml:"iss:expiration_date,omitempty" json:"expiration_date"`
	CardSequentalNumber string       `xml:"iss:card_sequental_number,omitempty" json:"card_sequental_number"`
	CardInstanceID      string       `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	State               string       `xml:"iss:state,omitempty" json:"state"`
	CardSecurity        CardSecurity `xml:"iss:card_security,omitempty" json:"card_security"`
	IssueDate           string       `xml:"iss:issue_date,omitempty" json:"issue_date"`
}

type CardSecurity struct {
	PVV            string `xml:"iss:PVV,omitempty" json:"pvv,omitempty"`
	PINOffset      string `xml:"iss:PIN_offset,omitempty" json:"pin_offset,omitempty"`
	PINBlock       string `xml:"iss:PIN_block,omitempty" json:"pin_block,omitempty"`
	KeyIndex       string `xml:"iss:key_index,omitempty" json:"key_index,omitempty"`
	PINBlockFormat string `xml:"iss:PIN_block_format,omitempty" json:"pin_block_format,omitempty"`
}

type ProcessIBGDataResponse struct {
	Text string `xml:",chardata"`
}

type GetCardInfoRequest struct {
	XMLName       xml.Name `xml:"ins:getCardInfoRequest"`
	ApplId        string   `xml:"ins:applId" json:"applId"`
	IncludeLimits string   `xml:"ins:includeLimits" json:"includeLimits"`
	Lang          string   `xml:"ins:lang" json:"lang"`
}

type GetCardInfoResponse struct {
	XMLName  xml.Name `xml:"getCardInfoResponse"`
	BatchId  string   `xml:"batchId" json:"batchId"`
	CardInfo CardInfo `xml:"cardInfo" json:"cardInfo"`
}

type CardInfo struct {
	CardNumber          string     `xml:"card_number" json:"card_number"`
	CardMask            string     `xml:"card_mask" json:"card_mask"`
	CardID              string     `xml:"card_id" json:"card_id"`
	CardIssDate         string     `xml:"card_iss_date" json:"card_iss_date"`
	CardStartDate       string     `xml:"card_start_date" json:"card_start_date"`
	ExpirationDate      string     `xml:"expiration_date" json:"expiration_date"`
	SequentialNumber    string     `xml:"sequential_number" json:"sequential_number"`
	CardStatus          string     `xml:"card_status" json:"card_status"`
	CardState           string     `xml:"card_state" json:"card_state"`
	PINUpdateFlag       string     `xml:"pin_update_flag" json:"pin_update_flag"`
	CardTypeId          string     `xml:"card_type_id" json:"card_type_id"`
	AgentNumber         string     `xml:"agent_number" json:"agent_number"`
	AgentName           string     `xml:"agent_name" json:"agent_name"`
	DeliveryAgentNumber string     `xml:"delivery_agent_number" json:"delivery_agent_number"`
	DeliveryAgentName   string     `xml:"delivery_agent_name" json:"delivery_agent_name"`
	ProductNumber       string     `xml:"product_number" json:"product_number"`
	ProductName         string     `xml:"product_name" json:"product_name"`
	AccountInfo         Account    `xml:"account" json:"account"`
	CardholderInfo      Cardholder `xml:"cardholder" json:"cardholder"`
	CustomerInfo        Customer   `xml:"customer" json:"customer"`
}
