package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Card type for the above middleware
// swagger:model
type Card struct {
	// required:true
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate" json:"updateDate,omitempty"`
	Deprecated   *string       `json:"deprecated" bson:"deprecated"`
	// required:true
	CustomerID bson.ObjectId `json:"customerId" bson:"customerId,omitempty"`
	// required:true
	CardToken string `bson:"cardToken" json:"cardToken"`
	// required:true
	Last4 string `bson:"last4" json:"last4"`
	// required:true
	Brand string `bson:"brand" json:"brand"`
	// required:true
	Default bool `bson:"defaultCard" json:"defaultCard"`
	// required:true
	GatewayType string `bson:"gatewayType" json:"gatewayType"`
}

// TranferBody sdkfjn
type TranferBody struct {
	Email string `bson:"email" json:"email"`
}

// CardRequest type for the above middleware
type CardRequest struct {
	// The number should be 16 numbers long
	// required:true
	// min length: 12
	// max length: 12
	Number string `bson:"card_number" json:"card_number"`
	// The holdername of the credit card
	// required:true
	HolderName string `bson:"holdername" json:"holdername"`
	// The number should be 4 numbers long
	// required:true
	ExpirationDate string `bson:"card_expiration_date" json:"card_expiration_date"`
	// The number should be 3 numbers long
	// required:true
	Cvv string `bson:"card_cvv" json:"card_cvv"`
}

// PagarmeCardRequest type for the above middleware
type PagarmeCardRequest struct {
	APIKey         string `bson:"api_key" json:"api_key"`
	Number         string `bson:"card_number" json:"card_number"`
	HolderName     string `bson:"card_holder_name" json:"card_holder_name"`
	ExpirationDate string `bson:"card_expiration_date" json:"card_expiration_date"`
	Cvv            string `bson:"card_cvv" json:"card_cvv"`
}

// PagarmeTransactionMeta sdfjkn
type PagarmeTransactionMeta struct {
	ClubID string `json:"clubId" bson:"clubId"`
}

// Transaction dsflm
type Transaction struct {
	Object string `json:"object"`
	Status string `json:"status"`
}

// InnerHits sdkjfn
type InnerHits struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	ID     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Source Transaction `json:"_source"`
}

// Hits sdkjfn
type Hits struct {
	Total     int         `json:"total"`
	MaxScore  float64     `json:"max_score"`
	InnerHits []InnerHits `json:"hits"`
}

// PagarmeTransaction sdkjgn
type PagarmeTransaction struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     Hits `json:"hits"`
}

// PostCardRequestBody are the data you need to send in order to
// create a Credit Card
// swagger:parameters postCard
type PostCardRequestBody struct {
	// in: body
	Body CardRequest `json:"body"`
}

// CardPatch sdfkjd
type CardPatch struct {
	Default bool `bson:"defaultCard" json:"defaultCard"`
}

// PostCardUpdateBody are the data you need to send in order to
// create a Credit Card
// swagger:parameters patchCard
type PostCardUpdateBody struct {
	// in: body
	Body CardPatch `json:"body"`
}

// ClubPathParamID are the data you need to send in order to
// swagger:parameters  getRecipients
type ClubPathParamID struct {
	// in: path
	ClubID string `json:"clubId"`
}

// PathParamID are the data you need to send in order to
// create a Credit Card
// swagger:parameters getCard deleteCard patchCard
type PathParamID struct {
	// in: path
	ID string `json:"id"`
}

// Cards is a list of cards of a giver user
// swagger:response cardsList
type Cards struct {
	// in: body
	Body []Card `json:"body"`
}

// ErrorResponse is the card added to MongoDB
// swagger:response error
type ErrorResponse struct {
	// in: body
	Body string `json:"body"`
}

// CardResponse is the card added to MongoDB
// swagger:response dbCard
type CardResponse struct {
	// in: body
	Body Card `json:"body"`
}
