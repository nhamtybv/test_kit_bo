package entity

import "encoding/xml"

type SoapReponse struct {
	XMLName xml.Name         `xml:"Envelope"`
	Soap    string           `xml:"soap,attr"`
	Body    *SoapReponseBody `xml:"Body"`
}

type SoapReponseBody struct {
	Application            *Application            `xml:"application"`
	GetCardInfoResponse    *GetCardInfoResponse    `xml:"getCardInfoResponse"`
	ProcessIBGDataResponse *ProcessIBGDataResponse `xml:"processIBGDataResponse"`
	Fault                  *Fault                  `xml:"Fault"`
}

// FaultDetail type
type FaultDetail struct {
	Details string `xml:"details" json:"details"`
}

// Fault type
type Fault struct {
	XMLName     xml.Name    `xml:"Fault"`
	Faultcode   string      `xml:"faultcode" json:"faultcode"`
	Faultstring string      `xml:"faultstring" json:"faultstring"`
	Detail      FaultDetail `xml:"detail" json:"detail"`
}

type SoapHeader struct {
	XMLName xml.Name `xml:"soap:Header"`
}

type SoapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Request interface{}
}

type SoapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Soap    string   `xml:"xmlns:soap,attr,omitempty"`
	Ins     string   `xml:"xmlns:ins,attr,omitempty"`
	Iss     string   `xml:"xmlns:iss,attr,omitempty"`
	Header  SoapHeader
	Body    SoapBody
}
