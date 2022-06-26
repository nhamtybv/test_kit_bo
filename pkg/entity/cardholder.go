package entity

// Cardholder type

type Cardholder struct {
	Command                 string             `xml:"command" faker:"-" json:"command"`
	CardholderNumber        string             `xml:"cardholder_number" faker:"-" json:"cardholder_number"`
	CardholderName          string             `xml:"cardholder_name" faker:"-" json:"cardholder_name"`
	CardholderPhotoFileName string             `xml:"cardholder_photo_file_name" faker:"-" json:"cardholder_photo_file_name"`
	CardholderSignFileName  string             `xml:"cardholder_sign_file_name" faker:"-" json:"cardholder_sign_file_name"`
	Person                  Person             `xml:"person" faker:"-" json:"person"`
	Contact                 Contact            `xml:"contact" faker:"-" json:"contact"`
	Address                 Address            `xml:"address" faker:"-" json:"address"`
	SecWord                 SecWord            `xml:"sec_word" faker:"-" json:"sec_word"`
	Notification            Notification       `xml:"notification" faker:"-" json:"notification"`
	EmbossedSurname         string             `xml:"embossed_surname" faker:"-" json:"embossed_surname"`
	EmbossedFirstName       string             `xml:"embossed_first_name" faker:"-" json:"embossed_first_name"`
	EmbossedSecondName      string             `xml:"embossed_second_name" faker:"-" json:"embossed_second_name"`
	EmbossedTitle           string             `xml:"embossed_title" faker:"-" json:"embossed_title"`
	CustomerRelation        string             `xml:"customer_relation" faker:"-" json:"customer_relation"`
	Resident                string             `xml:"resident" faker:"-" json:"resident"`
	Nationality             string             `xml:"nationality" faker:"-" json:"nationality"`
	MaritalStatus           string             `xml:"marital_status" faker:"-" json:"marital_status"`
	FlexibleField           FlexibleField      `xml:"flexible_field" faker:"-" json:"flexible_field"`
	LanguagePreference      LanguagePreference `xml:"language_preference" faker:"-" json:"language_preference"`
}

// SecWord type
type SecWord struct {
	SecretQuestion string `xml:"secret_question" faker:"sentence" json:"secret_question"`
	SecretAnswer   string `xml:"secret_answer" faker:"word" json:"secret_answer"`
}

// LanguagePreference type
type LanguagePreference struct {
	Command       string `xml:"command" faker:"-" json:"command"`
	PreferredLang string `xml:"preferred_lang" faker:"-" json:"preferred_lang"`
}
