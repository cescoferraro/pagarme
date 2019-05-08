package types

import "time"

// TransanctionFounds ssfjkdn
type TransanctionFounds struct {
	Amount int `json:"amount" bson:"amount"`
}

// TransanctionBalance ssfjkdn
type PagarMeTransactionBalance struct {
	Object        string             `json:"object" bson:"object"`
	WaitingFounds TransanctionFounds `json:"waiting_funds" bson:"waiting_funds"`
	Available     TransanctionFounds `json:"available" bson:"available"`
	Transferred   TransanctionFounds `json:"transferred" bson:"transferred"`
}

// FinanceQuery sdjkfn
type FinanceQuery struct {
	From        *Timestamp `json:"from,omitempty" bson:"from,omitempty"`
	Till        *Timestamp `json:"till,omitempty" bson:"till,omitempty"`
	RecipientID string     `json:"recipient_id" bson:"recipient_id"`
}

// Saldo TODO: NEEDS COMMENT INFO
type Saldo struct {
	Balance int `json:"balance" bson:"balance"`
	In      int `json:"in" bson:"in"`
	Out     int `json:"out" bson:"out"`
}

// Balance TODO: NEEDS COMMENT INFO
type Balance struct {
	Amount Saldo `json:"amount" bson:"amount"`
	Fee    Saldo `json:"fee" bson:"fee"`
}

// PagarmePayablesTimeline skdjnf
type PagarmePayablesTimeline struct {
	Date time.Time `json:"date" bson:"date"`
	Real bool      `json:"real" bson:"real"`
	Paid struct {
		CreditCard PagarMEHEY `json:"credit_card" bson:"credit_card"`
		Boleto     PagarMEHEY `json:"boleto" bson:"boleto"`
		DebitCard  PagarMEHEY `json:"debit_card" bson:"debit_card"`
	} `json:"paid" bson:"paid"`
}

// PagarMEHEY ksjdnf
type PagarMEHEY struct {
	Credit     PagarmePayablesTimelinePricing `json:"credit" bson:"credit"`
	Refund     PagarmePayablesTimelinePricing `json:"refund" bson:"refund"`
	ChargeBack PagarmePayablesTimelinePricing `json:"chargeback" bson:"chargeback"`
}

// PagarmePayablesTimelinePricing TODO: NEEDS COMMENT INFO
type PagarmePayablesTimelinePricing struct {
	Amount        int    `json:"amount" bson:"amount"`
	PaymentMethod string `json:"payment_method"`
	Fee           int    `json:"fee" bson:"fee"`
}

// TransactionDaysBalance skdjnf
type TransactionDaysBalance struct {
	Date         time.Time `json:"date" bson:"date"`
	Available    Balance   `json:"available" bson:"available"`
	WaitingFound Balance   `json:"waiting_funds" bson:"waiting_funds"`
}
