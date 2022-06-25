package entity

type CardRequest struct {
	Product  Product `json:"product"`
	AgentId  int     `json:"agent_id"`
	Category string  `json:"category"`
}

type CachedCard struct {
	CardID     string `json:"card_id"`
	CardNumber string `json:"card_number"`
}
