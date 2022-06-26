package entity

// Notification type
type Notification struct {
	Command         string `xml:"command" faker:"-" json:"command"`
	EventType       string `xml:"event_type" faker:"-" json:"event_type"`
	DeliveryChannel string `xml:"delivery_channel" faker:"-" json:"delivery_channel"`
	DeliveryAddress string `xml:"delivery_address" faker:"e_164_phone_number" json:"delivery_address"`
	IsActive        string `xml:"is_active" faker:"-" json:"is_active"`
	Error           Error  `xml:"error" faker:"-" json:"error"`
}
