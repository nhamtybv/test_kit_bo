package entity

type Product struct {
	ID            int                  `json:"product_id"`
	ProductType   string               `json:"product_type"`
	ContractType  string               `json:"contract_type"`
	ParentID      int                  `json:"parent_id"`
	InstID        int                  `json:"inst_id"`
	Status        string               `json:"status"`
	ProductNumber string               `json:"product_number"`
	SplitHash     int                  `json:"split_hash"`
	ProductName   string               `json:"product_name"`
	Services      []ProductService     `json:"product_services,omitempty"`
	AccountTypes  []ProductAccountType `json:"product_account_types,omitempty"`
	CardsTypes    []ProductCardType    `json:"card_types,omitempty"`
	Children      []Product            `json:"children,omitempty"`
}

type ProductService struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
	EntityType  string `json:"entity_type"`
}

type ProductAccountType struct {
	AccountType string `json:"account_type"`
	Currency    string `json:"currency"`
}

type ProductCardType struct {
	CardTypeID int    `json:"card_type_id"`
	CardType   string `json:"card_type"`
}

type ProductList struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}
