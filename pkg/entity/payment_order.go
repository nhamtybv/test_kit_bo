package entity

type PaymentOrder struct {
	Command                string           `xml:"command" json:"command"`
	PaymentOrderNumber     string           `xml:"payment_order_number" json:"payment_order_number"`
	PaymentPurposeID       string           `xml:"payment_purpose_id" json:"payment_purpose_id"`
	Label                  string           `xml:"label" json:"label"`
	Description            string           `xml:"description" json:"description"`
	PaymentAmount          string           `xml:"payment_amount" json:"payment_amount"`
	Currency               string           `xml:"currency" json:"currency"`
	PaymentAmountAlgo      string           `xml:"payment_amount_algo" json:"payment_amount_algo"`
	PaymentDate            string           `xml:"payment_date" json:"payment_date"`
	PostingDate            string           `xml:"posting_date" json:"posting_date"`
	IsPaymentOrderTemplate string           `xml:"is_payment_order_template" json:"is_payment_order_template"`
	EventType              string           `xml:"event_type" json:"event_type"`
	TemplateStatus         string           `xml:"template_status" json:"template_status"`
	AttemptCount           string           `xml:"attempt_count" json:"attempt_count"`
	PaymentParameter       PaymentParameter `xml:"payment_parameter" json:"payment_parameter"`
	Error                  Error            `xml:"error" json:"error"`
}

type PaymentParameter struct {
	PaymentParameterName  string `xml:"payment_parameter_name" json:"payment_parameter_name"`
	PaymentParameterValue string `xml:"payment_parameter_value" json:"payment_parameter_value"`
}
