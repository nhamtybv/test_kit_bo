package entity

type PaymentOrder struct {
	Command                string           `xml:"command,omitempty" json:"command"`
	PaymentOrderNumber     string           `xml:"payment_order_number,omitempty" json:"payment_order_number"`
	PaymentPurposeID       string           `xml:"payment_purpose_id,omitempty" json:"payment_purpose_id"`
	Label                  string           `xml:"label,omitempty" json:"label"`
	Description            string           `xml:"description,omitempty" json:"description"`
	PaymentAmount          string           `xml:"payment_amount,omitempty" json:"payment_amount"`
	Currency               string           `xml:"currency,omitempty" json:"currency"`
	PaymentAmountAlgo      string           `xml:"payment_amount_algo,omitempty" json:"payment_amount_algo"`
	PaymentDate            string           `xml:"payment_date,omitempty" json:"payment_date"`
	PostingDate            string           `xml:"posting_date,omitempty" json:"posting_date"`
	IsPaymentOrderTemplate string           `xml:"is_payment_order_template,omitempty" json:"is_payment_order_template"`
	EventType              string           `xml:"event_type,omitempty" json:"event_type"`
	TemplateStatus         string           `xml:"template_status,omitempty" json:"template_status"`
	AttemptCount           string           `xml:"attempt_count,omitempty" json:"attempt_count"`
	PaymentParameter       PaymentParameter `xml:"payment_parameter,omitempty" json:"payment_parameter"`
	Error                  Error            `xml:"error,omitempty" json:"error"`
}

type PaymentParameter struct {
	PaymentParameterName  string `xml:"payment_parameter_name,omitempty" json:"payment_parameter_name"`
	PaymentParameterValue string `xml:"payment_parameter_value,omitempty" json:"payment_parameter_value"`
}
