package model

// Exchange hold the information regarding an exchange that holds different currencies involved
type Exchange struct {
	ExchangeName string     `json:"exchange_name" bson:"exchange_name,omitempty"`
	Currencies   []Currency `json:"Currencies" bson:"currencies,omitempty"`
	Latest       bool       `bson:"latest"`
}

// Currency holds information regarding a particular currency
type Currency struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Price string `json:"price" bson:"price,omitempty"`
}
