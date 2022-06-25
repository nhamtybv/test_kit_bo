package entity

type CardRequest struct {
	Product  Product `json:"product"`
	AgentId  int     `json:"agent_id"`
	Category string  `json:"category"`
}
