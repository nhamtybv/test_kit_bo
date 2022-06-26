package entity

type Customer struct {
	Command            string        `xml:"command" faker:"-" json:"command"`
	CustomerNumber     string        `xml:"customer_number" faker:"cc_number" json:"customer_number"`
	CustomerCategory   string        `xml:"customer_category" faker:"-" json:"customer_category"`
	CustomerExtType    string        `xml:"customer_ext_type" faker:"-" json:"customer_ext_type"`
	CustomerExtID      string        `xml:"customer_ext_id" faker:"-" json:"customer_ext_id"`
	CustomerRelation   string        `xml:"customer_relation" faker:"-" json:"customer_relation"`
	Resident           string        `xml:"resident" faker:"-" json:"resident"`
	Nationality        string        `xml:"nationality" faker:"-" json:"nationality"`
	CreditRating       string        `xml:"credit_rating" faker:"-" json:"credit_rating"`
	MoneyLaundryRisk   string        `xml:"money_laundry_risk" faker:"-" json:"money_laundry_risk"`
	MoneyLaundryReason string        `xml:"money_laundry_reason" faker:"-" json:"money_laundry_reason"`
	CustomerCount      string        `xml:"customer_count" faker:"-" json:"customer_count"`
	CustomerStatus     string        `xml:"customer_status" faker:"-" json:"customer_status"`
	EmploymentStatus   string        `xml:"employment_status" faker:"-" json:"employment_status"`
	EmploymentPeriod   string        `xml:"employment_period" faker:"-" json:"employment_period"`
	ResidenceType      string        `xml:"residence_type" faker:"-" json:"residence_type"`
	MaritalStatus      string        `xml:"marital_status" faker:"-" json:"marital_status"`
	MaritalStatusDate  string        `xml:"marital_status_date" faker:"-" json:"marital_status_date"`
	IncomeRange        string        `xml:"income_range" faker:"-" json:"income_range"`
	NumberOfChildren   string        `xml:"number_of_children" faker:"-" json:"number_of_children"`
	FlexibleField      FlexibleField `xml:"flexible_field" faker:"-" json:"flexible_field"`
	Note               Note          `xml:"note" faker:"-" json:"note"`
	Contract           Contract      `xml:"contract" faker:"-" json:"contract"`
	Person             Person        `xml:"person" faker:"-" json:"person"`
	Company            Company       `xml:"company" faker:"-" json:"company"`
	Contact            Contact       `xml:"contact" faker:"-" json:"contact"`
	Address            Address       `xml:"address" faker:"-" json:"address"`
	Notification       Notification  `xml:"notification" faker:"-" json:"notification"`
	Error              Error         `xml:"error" faker:"-" json:"error"`
}
