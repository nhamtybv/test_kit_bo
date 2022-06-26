package entity

type Error struct {
	ID           string `xml:"id,attr" json:"id"`
	ErrorCode    string `xml:"error_code" json:"error_code"`
	ErrorDesc    string `xml:"error_desc" json:"error_desc"`
	ErrorElement string `xml:"error_element" json:"error_element"`
	ErrorDetails string `xml:"error_details" json:"error_details"`
}
