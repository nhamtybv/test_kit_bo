package entity

import "encoding/xml"

type Application struct {
	XMLName           xml.Name `xml:"application" faker:"-" json:"xml_name"`
	Xmlns             string   `xml:"xmlns,attr" faker:"-" json:"xmlns"`
	ApplicationID     string   `xml:"application_id" faker:"-" json:"application_id"`
	ApplicationType   string   `xml:"application_type" faker:"-" json:"application_type"`
	ApplicationFlowID string   `xml:"application_flow_id" faker:"-" json:"application_flow_id"`
	ApplicationStatus string   `xml:"application_status" faker:"-" json:"application_status"`
	OperatorID        string   `xml:"operator_id" faker:"-" json:"operator_id"`
	InstitutionID     string   `xml:"institution_id" faker:"-" json:"institution_id"`
	AgentID           string   `xml:"agent_id" faker:"-" json:"agent_id"`
	AgentNumber       string   `xml:"agent_number" faker:"-" json:"agent_number"`
	CustomerType      string   `xml:"customer_type" faker:"-" json:"customer_type"`
	Customer          Customer `xml:"customer" faker:"-" json:"customer"`
	Error             Error    `xml:"error" faker:"-" json:"error"`
}
