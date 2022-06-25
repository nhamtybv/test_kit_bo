package entity

import "encoding/xml"

type SoapRequest struct {
	App              *Application
	OperationRequest *OperationRequest `xml:"operationRequest"`
	Address          string
}

type SoapReponse struct {
	XMLName xml.Name         `xml:"Envelope"`
	Soap    string           `xml:"soap,attr"`
	Body    *SoapReponseBody `xml:"Body"`
}

type SoapReponseBody struct {
	Application      *Application      `xml:"application"`
	OperationRequest *OperationRequest `xml:"operationRequest"`
	Fault            *Fault            `xml:"Fault"`
}

// FaultDetail type
type FaultDetail struct {
	Details string `xml:"details" json:"details"`
}

// Fault type
type Fault struct {
	Faultcode   string      `xml:"faultcode" json:"faultcode"`
	Faultstring string      `xml:"faultstring" json:"faultstring"`
	Detail      FaultDetail `xml:"detail" json:"detail"`
}
