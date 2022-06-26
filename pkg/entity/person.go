package entity

// Person type
type Person struct {
	Command       string        `xml:"command" faker:"-" json:"command"`
	PersonTitle   string        `xml:"person_title" faker:"-" json:"person_title"`
	PersonName    PersonName    `xml:"person_name" faker:"-" json:"person_name"`
	Suffix        string        `xml:"suffix" faker:"-" json:"suffix"`
	Birthday      string        `xml:"birthday" faker:"-" json:"birthday"`
	PlaceOfBirth  string        `xml:"place_of_birth" faker:"-" json:"place_of_birth"`
	Gender        string        `xml:"gender" faker:"-" json:"gender"`
	IdentityCard  IdentityCard  `xml:"identity_card" faker:"-" json:"identity_card"`
	FlexibleField FlexibleField `xml:"flexible_field"`
	Error         Error         `xml:"error"`
}

// PersonName type
type PersonName struct {
	Language   string `xml:"language,attr" faker:"-" json:"language"`
	Surname    string `xml:"surname" faker:"last_name" json:"surname"`
	FirstName  string `xml:"first_name" faker:"first_name" json:"first_name"`
	SecondName string `xml:"second_name" faker:"-" json:"second_name"`
}
