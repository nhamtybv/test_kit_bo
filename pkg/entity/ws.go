package entity

import "encoding/xml"

type SoapRequest struct {
	App              *Application      `json:"app"`
	OperationRequest *OperationRequest `xml:"operationRequest" json:"operation_request"`
	Address          string            `json:"address"`
}

type SoapReponse struct {
	XMLName xml.Name         `xml:"Envelope" json:"xml_name"`
	Soap    string           `xml:"soap,attr" json:"soap"`
	Body    *SoapReponseBody `xml:"Body" json:"body"`
}

type SoapReponseBody struct {
	Application      *Application      `xml:"application" json:"application"`
	OperationRequest *OperationRequest `xml:"operationRequest" json:"operation_request"`
	Fault            *Fault            `xml:"Fault" json:"fault"`
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
