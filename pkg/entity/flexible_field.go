package entity

type FlexibleField struct {
	FlexibleFieldName  string `xml:"flexible_field_name" json:"flexible_field_name"`
	FlexibleFieldValue string `xml:"flexible_field_value" json:"flexible_field_value"`
	Error              Error  `xml:"error" json:"error"`
}
