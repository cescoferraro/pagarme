package pagarme

// Founds ssfjkdn
type Founds struct {
	Amount int `json:"amount" bson:"amount"`
}

// RecipientBalance ssfjkdn
type RecipientBalance struct {
	Object        string `json:"object" bson:"object"`
	WaitingFounds Founds `json:"waiting_funds" bson:"waiting_funds"`
	Available     Founds `json:"available" bson:"available"`
	Transferred   Founds `json:"transferred" bson:"transferred"`
}
