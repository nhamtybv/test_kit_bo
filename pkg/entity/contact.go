package entity

// Contact type
type Contact struct {
	Command       string      `xml:"command" faker:"-" json:"command"`
	ContactType   string      `xml:"contact_type" faker:"-" json:"contact_type"`
	PreferredLang string      `xml:"preferred_lang" faker:"-" json:"preferred_lang"`
	JobTitle      string      `xml:"job_title" faker:"-" json:"job_title"`
	Person        Person      `xml:"person" faker:"-" json:"person"`
	ContactData   ContactData `xml:"contact_data" faker:"-" json:"contact_data"`
	Error         Error       `xml:"error" faker:"-" json:"error"`
}

// ContactData type
type ContactData struct {
	CommunMethod  string `xml:"commun_method" faker:"-" json:"commun_method"`
	CommunAddress string `xml:"commun_address" faker:"email" json:"commun_address"`
	StartDate     string `xml:"start_date" faker:"-" json:"start_date"`
	EndDate       string `xml:"end_date" faker:"-" json:"end_date"`
}
