package entity

// Contract type
type Contract struct {
	Command        string    `xml:"command" faker:"-" json:"command"`
	ContractType   string    `xml:"contract_type" faker:"-" json:"contract_type"`
	ProductID      string    `xml:"product_id" faker:"-" json:"product_id"`
	ProductNumber  string    `xml:"product_number" faker:"-" json:"product_number"`
	ContractNumber string    `xml:"contract_number" faker:"-" json:"contract_number"`
	StartDate      string    `xml:"start_date" faker:"-" json:"start_date"`
	EndDate        string    `xml:"end_date" faker:"-" json:"end_date"`
	Card           Card      `xml:"card" faker:"-" json:"card"`
	Services       []Service `xml:"service" faker:"-" json:"services"`
	Account        Account   `xml:"account" faker:"-" json:"account"`
	Document       Document  `xml:"document" faker:"-" json:"document"`
	Error          Error     `xml:"error" faker:"-" json:"error"`
}
