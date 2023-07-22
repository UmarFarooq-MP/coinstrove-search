package dto

// Exchange hold the information regarding a exchange that has different currencies involved
type Exchange struct {
	ExchangeName string     `json:"exchange_name" bson:"exchange_name,omitempty"`
	Currencies   []Currency `json:"Currencies" bson:"Currencies,omitempty"`
}

// Currency holds information regarding a particular currency
type Currency struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Price string `json:"price" bson:"price,omitempty"`
}

// Message Websocket Response struct
type Message struct {
	Data         Exchange `json:"data" bson:"data,omitempty"`
	ErrorMessage string   `bson:"errorMessage" json:"errorMessage"`
}
