package entity

// Account type

type Account struct {
	Command       string        `xml:"command" faker:"-" json:"command"`
	AccountNumber string        `xml:"account_number" faker:"-" json:"account_number"`
	Currency      string        `xml:"currency" faker:"-" json:"currency"`
	AccountType   string        `xml:"account_type" faker:"oneof: 130, 110" json:"account_type"`
	AccountStatus string        `xml:"account_status" faker:"-" json:"account_status"`
	AccountCount  string        `xml:"account_count" faker:"-" json:"account_count"`
	PaymentOrder  PaymentOrder  `xml:"payment_order" faker:"-" json:"payment_order"`
	CstAccAgentID string        `xml:"cst_acc_agent_id" faker:"-" json:"cst_acc_agent_id"`
	AccountObject AccountObject `xml:"account_object" faker:"-" json:"account_object"`
	Error         Error         `xml:"error" faker:"-" json:"error"`
	FlexibleField FlexibleField `xml:"flexible_field" faker:"-" json:"flexible_field"`
	Facilitator   Facilitator   `xml:"facilitator" faker:"-" json:"facilitator"`
}

type AccountObjectProperty struct {
	LinkPropertyType string `xml:"link_property_type" json:"link_property_type"`
	LinkProperty     string `xml:"link_property" json:"link_property"`
	LinkPropertyFlag string `xml:"link_property_flag" json:"link_property_flag"`
}

// AccountObject type
type AccountObject struct {
	RefID                 string                `xml:"ref_id,attr" faker:"-" json:"ref_id"`
	AccountLinkFlag       string                `xml:"account_link_flag" faker:"-" json:"account_link_flag"`
	IsPosDefault          string                `xml:"is_pos_default" faker:"-" json:"is_pos_default"`
	IsAtmDefault          string                `xml:"is_atm_default" faker:"-" json:"is_atm_default"`
	AccountSeqNumber      string                `xml:"account_seq_number" faker:"-" json:"account_seq_number"`
	AccountObjectProperty AccountObjectProperty `xml:"account_object_property" faker:"-" json:"account_object_property"`
	Error                 Error                 `xml:"error" faker:"-" json:"error"`
}

type Facilitator struct {
	AccountNumber string `xml:"account_number" json:"account_number"`
	Description   string `xml:"description" json:"description"`
}
