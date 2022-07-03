package entity

import "encoding/xml"

type SoapReponse struct {
	XMLName xml.Name         `xml:"Envelope"`
	Soap    string           `xml:"soap,attr"`
	Body    *SoapReponseBody `xml:"Body"`
}

type SoapReponseBody struct {
	Application                   *Application                   `xml:"application"`
	GetCardInfoResponse           *GetCardInfoResponse           `xml:"getCardInfoResponse"`
	ProcessIBGDataResponse        *ProcessIBGDataResponse        `xml:"processIBGDataResponse"`
	GetApplicationDetailsResponse *GetApplicationDetailsResponse `xml:"getApplicationDetailsResponse"`
	OperationResponse             *OperationResponse             `xml:"operationResponse"`
	Fault                         *Fault                         `xml:"Fault,omitempty"`
}

type OperationResponse struct {
	Operation struct {
		OperID string `xml:"oper_id"`
		Status string `xml:"status"`
	} `xml:"operation"`
}

// FaultDetail type
type FaultDetail struct {
	Details string `xml:"details" json:"details"`
}

type FaultSubcode struct {
	Text  string `xml:",chardata"`
	Value struct {
		Text string `xml:",chardata"`
		Ns1  string `xml:"ns1,attr"`
	} `xml:"Value"`
}
type FaultCode struct {
	Value   string       `xml:"Value"`
	Subcode FaultSubcode `xml:"Subcode"`
}

type FaultReason struct {
	Chardata string `xml:",chardata"`
	Text     struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"Text"`
}

// Fault type
type Fault struct {
	Faultcode   string       `xml:"faultcode" json:"faultcode"`
	Faultstring string       `xml:"faultstring" json:"faultstring"`
	Detail      *FaultDetail `xml:"detail" json:"detail"`
	Code        *FaultCode   `xml:"Code"`
	Reason      *FaultReason `xml:"Reason"`
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
