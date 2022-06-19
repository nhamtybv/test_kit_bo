package entity

import "encoding/xml"

type Application struct {
	XMLName           xml.Name `xml:"application" faker:"-" json:"xml_name"`
	Xmlns             string   `xml:"xmlns,attr" faker:"-" json:"xmlns"`
	ApplicationType   string   `xml:"application_type" faker:"-" json:"application_type"`
	ApplicationFlowID string   `xml:"application_flow_id" faker:"-" json:"application_flow_id"`
	ApplicationStatus string   `xml:"application_status" faker:"-" json:"application_status"`
	OperatorID        string   `xml:"operator_id" faker:"-" json:"operator_id"`
	InstitutionID     string   `xml:"institution_id" faker:"-" json:"institution_id"`
	AgentID           string   `xml:"agent_id" faker:"-" json:"agent_id"`
	AgentNumber       string   `xml:"agent_number" faker:"-" json:"agent_number"`
	CustomerType      string   `xml:"customer_type" faker:"-" json:"customer_type"`
	Customer          Customer `xml:"customer" faker:"-" json:"customer"`
}

// Customer type
type Customer struct {
	Command           string       `xml:"command" faker:"-" json:"command"`
	CustomerNumber    string       `xml:"customer_number" faker:"cc_number" json:"customer_number"`
	CustomerCategory  string       `xml:"customer_category" faker:"-" json:"customer_category"`
	CustomerRelation  string       `xml:"customer_relation" faker:"-" json:"customer_relation"`
	Resident          string       `xml:"resident" faker:"-" json:"resident"`
	Nationality       string       `xml:"nationality" faker:"-" json:"nationality"`
	EmploymentStatus  string       `xml:"employment_status" faker:"-" json:"employment_status"`
	EmploymentPeriod  string       `xml:"employment_period" faker:"-" json:"employment_period"`
	ResidenceType     string       `xml:"residence_type" faker:"-" json:"residence_type"`
	MaritalStatusDate string       `xml:"marital_status_date" faker:"-" json:"marital_status_date"`
	NumberOfChildren  string       `xml:"number_of_children" faker:"-" json:"number_of_children"`
	Contract          Contract     `xml:"contract" faker:"-" json:"contract"`
	Person            Person       `xml:"person" faker:"-" json:"person"`
	Contact           Contact      `xml:"contact" faker:"-" json:"contact"`
	Address           Address      `xml:"address" faker:"-" json:"address"`
	Notification      Notification `xml:"notification" faker:"-" json:"notification"`
}

// Contract type
type Contract struct {
	Command        string    `xml:"command" faker:"-" json:"command"`
	ContractType   string    `xml:"contract_type" faker:"-" json:"contract_type"`
	ProductID      string    `xml:"product_id" faker:"-" json:"product_id"`
	ContractNumber string    `xml:"contract_number" faker:"-" json:"contract_number"`
	StartDate      string    `xml:"start_date" faker:"-" json:"start_date"`
	Card           Card      `xml:"card" faker:"-" json:"card"`
	Service        []Service `xml:"service" faker:"-" json:"service"`
	Account        Account   `xml:"account" faker:"-" json:"account"`
}

// SecWord type
type SecWord struct {
	SecretQuestion string `xml:"secret_question" faker:"sentence" json:"secret_question"`
	SecretAnswer   string `xml:"secret_answer" faker:"word" json:"secret_answer"`
}

// LanguagePreference type
type LanguagePreference struct {
	Command       string `xml:"command" faker:"-" json:"command"`
	PreferredLang string `xml:"preferred_lang" faker:"-" json:"preferred_lang"`
}

// Cardholder type
type Cardholder struct {
	Command                 string             `xml:"command" faker:"-" json:"command"`
	CardholderNumber        string             `xml:"cardholder_number" faker:"-" json:"cardholder_number"`
	CardholderName          string             `xml:"cardholder_name" faker:"name" json:"cardholder_name"`
	CardholderPhotoFileName string             `xml:"cardholder_photo_file_name" faker:"-" json:"cardholder_photo_file_name"`
	CardholderSignFileName  string             `xml:"cardholder_sign_file_name" faker:"-" json:"cardholder_sign_file_name"`
	SecWord                 SecWord            `xml:"sec_word" faker:"-" json:"sec_word"`
	Notification            Notification       `xml:"notification" faker:"-" json:"notification"`
	LanguagePreference      LanguagePreference `xml:"language_preference" faker:"-" json:"language_preference"`
}

// Card type
type Card struct {
	ID                  string     `xml:"id,attr" faker:"-" json:"id"`
	Command             string     `xml:"command" faker:"-" json:"command"`
	DeliveryAgentNumber string     `xml:"delivery_agent_number" faker:"-" json:"delivery_agent_number"`
	CardID              string     `xml:"card_id" faker:"-" json:"card_id"`
	ExpirationDate      string     `xml:"expiration_date" faker:"-" json:"expiration_date"`
	CardType            string     `xml:"card_type" faker:"-" json:"card_type"`
	Category            string     `xml:"category" faker:"oneof: 800, 400" json:"category"`
	Cardholder          Cardholder `xml:"cardholder" faker:"-" json:"cardholder"`
	CardNumber          string     `xml:"card_number" faker:"-" json:"card_number"`
}

// AttributeNum type
type ApplicationAttributeNum struct {
	Value             string `xml:"value,attr" faker:"-" json:"value"`
	AttributeValueNum string `xml:"attribute_value_num" faker:"-" json:"attribute_value_num"`
}

// AttributeChar type
type ApplicationAttributeChar struct {
	Value              string `xml:"value,attr" faker:"-" json:"value"`
	AttributeValueChar string `xml:"attribute_value_char" faker:"-" json:"attribute_value_char"`
}

// AttributeLimit type
type ApplicationAttributeLimit struct {
	Text             string `xml:",chardata" faker:"-" json:"text"`
	Value            string `xml:"value,attr" faker:"-" json:"value"`
	LimitSumValue    string `xml:"limit_sum_value" faker:"-" json:"limit_sum_value"`
	LimitCheckType   string `xml:"limit_check_type" faker:"-" json:"limit_check_type"`
	Currency         string `xml:"currency" faker:"-" json:"currency"`
	CounterAlgorithm string `xml:"counter_algorithm" faker:"-" json:"counter_algorithm"`
}

// ServiceObject type
type ServiceObject struct {
	RefID          string                    `xml:"ref_id,attr" faker:"-" json:"ref_id"`
	StartDate      string                    `xml:"start_date" faker:"-" json:"start_date"`
	AttributeNum   ApplicationAttributeNum   `xml:"attribute_num" faker:"-" json:"attribute_num"`
	AttributeChar  ApplicationAttributeChar  `xml:"attribute_char" faker:"-" json:"attribute_char"`
	AttributeLimit ApplicationAttributeLimit `xml:"attribute_limit" faker:"-" json:"attribute_limit"`
}

// Service type
type Service struct {
	Value         string        `xml:"value,attr" faker:"-" json:"value"`
	ServiceObject ServiceObject `xml:"service_object" faker:"-" json:"service_object"`
}

// AccountObject type
type AccountObject struct {
	RefID           string `xml:"ref_id,attr" faker:"-" json:"ref_id"`
	AccountLinkFlag string `xml:"account_link_flag" faker:"-" json:"account_link_flag"`
}

// Account type
type Account struct {
	ID            string        `xml:"id,attr" faker:"-" json:"id"`
	Command       string        `xml:"command" faker:"-" json:"command"`
	AccountNumber string        `xml:"account_number" faker:"-" json:"account_number"`
	Currency      string        `xml:"currency" faker:"-" json:"currency"`
	AccountType   string        `xml:"account_type" faker:"oneof: 130, 110" json:"account_type"`
	AccountObject AccountObject `xml:"account_object" faker:"-" json:"account_object"`
}

// PersonName type
type PersonName struct {
	Language   string `xml:"language,attr" faker:"-" json:"language"`
	Surname    string `xml:"surname" faker:"last_name" json:"surname"`
	FirstName  string `xml:"first_name" faker:"first_name" json:"first_name"`
	SecondName string `xml:"second_name" faker:"-" json:"second_name"`
}

// IdentityCard type
type IdentityCard struct {
	Command  string `xml:"command" faker:"-" json:"command"`
	IDType   string `xml:"id_type" faker:"-" json:"id_type"`
	IDSeries string `xml:"id_series" faker:"cc_number" json:"id_series"`
	IDNumber string `xml:"id_number" faker:"cc_number" json:"id_number"`
}

// Person type
type Person struct {
	Command      string       `xml:"command" faker:"-" json:"command"`
	PersonName   PersonName   `xml:"person_name" faker:"-" json:"person_name"`
	IdentityCard IdentityCard `xml:"identity_card" faker:"-" json:"identity_card"`
}

// ContactData type
type ContactData struct {
	CommunMethod  string `xml:"commun_method" faker:"-" json:"commun_method"`
	CommunAddress string `xml:"commun_address" faker:"email" json:"commun_address"`
}

// Contact type
type Contact struct {
	Command       string      `xml:"command" faker:"-" json:"command"`
	ContactType   string      `xml:"contact_type" faker:"-" json:"contact_type"`
	PreferredLang string      `xml:"preferred_lang" faker:"-" json:"preferred_lang"`
	ContactData   ContactData `xml:"contact_data" faker:"-" json:"contact_data"`
}

// AddressName type
type AddressName struct {
	Language string `xml:"language,attr" faker:"-" json:"language"`
	Region   string `xml:"region" faker:"word" json:"region"`
	City     string `xml:"city" faker:"name" json:"city"`
	Street   string `xml:"street" faker:"name" json:"street"`
}

// Address type
type Address struct {
	Command     string      `xml:"command" faker:"-" json:"command"`
	AddressType string      `xml:"address_type" faker:"-" json:"address_type"`
	Country     string      `xml:"country" faker:"-" json:"country"`
	AddressName AddressName `xml:"address_name" faker:"-" json:"address_name"`
	House       int8        `xml:"house" faker:"boundary_start=11, boundary_end=99" json:"house"`
	Apartment   string      `xml:"apartment" faker:"boundary_start=11, boundary_end=99" json:"apartment"`
}

// Notification type
type Notification struct {
	Command         string `xml:"command" faker:"-" json:"command"`
	DeliveryChannel string `xml:"delivery_channel" faker:"-" json:"delivery_channel"`
	DeliveryAddress string `xml:"delivery_address" faker:"e_164_phone_number" json:"delivery_address"`
	IsActive        string `xml:"is_active" faker:"-" json:"is_active"`
}
