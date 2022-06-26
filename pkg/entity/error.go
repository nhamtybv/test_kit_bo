package entity

type Error struct {
	ID           string `xml:"id,attr,omitempty" json:"id,omitempty"`
	ErrorCode    string `xml:"error_code,omitempty" json:"error_code,omitempty"`
	ErrorDesc    string `xml:"error_desc,omitempty" json:"error_desc,omitempty"`
	ErrorElement string `xml:"error_element,omitempty" json:"error_element,omitempty"`
	ErrorDetails string `xml:"error_details,omitempty" json:"error_details,omitempty"`
}
