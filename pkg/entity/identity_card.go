package entity

// IdentityCard identify data
type IdentityCard struct {
	Command       string        `xml:"command" faker:"-" json:"command"`
	IDType        string        `xml:"id_type" faker:"-" json:"id_type"`
	IDSeries      string        `xml:"id_series" faker:"cc_number" json:"id_series"`
	IDNumber      string        `xml:"id_number" faker:"cc_number" json:"id_number"`
	IDIssuer      string        `xml:"id_issuer" faker:"-" json:"id_issuer"`
	IDIssueDate   string        `xml:"id_issue_date" faker:"-" json:"id_issue_date"`
	IDExpireDate  string        `xml:"id_expire_date" faker:"-" json:"id_expire_date"`
	IDDesc        string        `xml:"id_desc" faker:"-" json:"id_desc"`
	FlexibleField FlexibleField `xml:"flexible_field" faker:"-" json:"flexible_field"`
	Error         Error         `xml:"error" faker:"-" json:"error"`
}
