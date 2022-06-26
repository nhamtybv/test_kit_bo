package entity

type Document struct {
	DocumentType     string `xml:"document_type" json:"document_type"`
	DocumentObject   string `xml:"document_object" json:"document_object"`
	DocumentNumber   string `xml:"document_number" json:"document_number"`
	DocumentDate     string `xml:"document_date" json:"document_date"`
	DocumentName     string `xml:"document_name" json:"document_name"`
	CustomerEds      string `xml:"customer_eds" json:"customer_eds"`
	UserEds          string `xml:"user_eds" json:"user_eds"`
	UserName         string `xml:"user_name" json:"user_name"`
	DocumentContents string `xml:"document_contents" json:"document_contents"`
	Error            Error  `xml:"error" json:"error"`
}
