package entity

type Company struct {
	ID                        string       `xml:"id,attr" json:"id"`
	Command                   string       `xml:"command" json:"command"`
	EmbossedName              string       `xml:"embossed_name" json:"embossed_name"`
	IncorpForm                string       `xml:"incorp_form" json:"incorp_form"`
	PresenceOnLocation        string       `xml:"presence_on_location" json:"presence_on_location"`
	AuthorizedCapital         string       `xml:"authorized_capital" json:"authorized_capital"`
	AuthorizedCapitalCurrency string       `xml:"authorized_capital_currency" json:"authorized_capital_currency"`
	CompanyName               CompanyName  `xml:"company_name" json:"company_name"`
	IdentityCard              IdentityCard `xml:"identity_card" json:"identity_card"`
	Error                     Error        `xml:"error" json:"error"`
}

type CompanyName struct {
	Language         string `xml:"language,attr" json:"language"`
	ID               string `xml:"id,attr" json:"id"`
	CompanyShortName string `xml:"company_short_name" json:"company_short_name"`
	CompanyFullName  string `xml:"company_full_name" json:"company_full_name"`
}
