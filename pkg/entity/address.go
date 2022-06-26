package entity

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
	House       string      `xml:"house" faker:"boundary_start=11, boundary_end=99" json:"house"`
	Apartment   string      `xml:"apartment" faker:"boundary_start=11, boundary_end=99" json:"apartment"`
	PostalCode  string      `xml:"postal_code" faker:"-" json:"postal_code"`
	PlaceCode   string      `xml:"place_code" faker:"-" json:"place_code"`
	RegionCode  string      `xml:"region_code" faker:"-" json:"region_code"`
	Latitude    string      `xml:"latitude" faker:"-" json:"latitude"`
	Longitude   string      `xml:"longitude" faker:"-" json:"longitude"`
	Error       Error       `xml:"error"`
}
