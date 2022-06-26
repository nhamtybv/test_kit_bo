package entity

type Note struct {
	NoteType    string      `xml:"note_type" json:"note_type"`
	NoteContent NoteContent `xml:"note_content" json:"note_content"`
	Error       Error       `xml:"error" json:"error"`
}

type NoteContent struct {
	Language   string `xml:"language,attr" json:"language"`
	NoteHeader string `xml:"note_header" json:"note_header"`
	NoteText   string `xml:"note_text" json:"note_text"`
	StartDate  string `xml:"start_date" json:"start_date"`
	EndDate    string `xml:"end_date" json:"end_date"`
	Error      Error  `xml:"error" json:"error"`
}
