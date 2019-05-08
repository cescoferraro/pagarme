package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Recipient is a mongo document
type Recipient struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate  *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate    *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ClubID        bson.ObjectId `json:"clubId" bson:"clubId"`
	RecipientID   string        `json:"recipientId" bson:"recipientId"`
	BankAccountID int           `json:"bankAccountId" bson:"bankAccountId"`
	Type          string        `json:"type" bson:"type"`
	Status        string        `json:"status" bson:"status"`
	BankingInfo   BankingInfo   `json:"bankingInfo" bson:"bankingInfo"`
}

// BankingInfo sdjknf
type BankingInfo struct {
	BankCode        string  `json:"bankCode" bson:"bankCode"`
	BankBranch      string  `json:"bankBranch" bson:"bankBranch"`
	BankBranchVC    *string `json:"bankBranchVC" bson:"bankBranchVC"`
	BankAccount     string  `json:"bankAccount" bson:"bankAccount"`
	BankAccountVC   string  `json:"bankAccountVC" bson:"bankAccountVC"`
	BankAccountName string  `json:"bankAccountName" bson:"bankAccountName"`
	PersonType      string  `json:"personType" bson:"personType"`
	DocumentNumber  string  `json:"documentNumber" bson:"documentNumber"`
}

// AntecipationPostRequest type for the above middleware
type AntecipationPostRequest struct {
	RecipientID     string `json:"recipient_id" bson:"recipient_id"`
	Timeframe       string `json:"timeframe" bson:"timeframe"`
	Build           bool   `json:"build" bson:"build"`
	PaymentDay      int    `json:"payment_date" bson:"payment_date"`
	RequestedAmount int    `json:"requested_amount" bson:"requested_amount"`
}

// RecipientPost sdfkjd
type RecipientPost struct {
	ID              string `json:"id" bson:"id"`
	PersonType      string `json:"personType" bson:"personType"`
	ClubID          string `json:"clubID" bson:"clubID"`
	Status          string `json:"status" bson:"status"`
	BankCode        string `json:"bankCode" bson:"bankCode"`
	BankBranch      string `json:"bankBranch" bson:"bankBranch"`
	BankBranchVC    string `json:"bankBranchVC" bson:"bankBranchVC"`
	BankAccount     string `json:"bankAccount" bson:"bankAccount"`
	BankAccountVC   string `json:"bankAccountVC" bson:"bankAccountVC"`
	BankAccountName string `json:"bankAccountName" bson:"bankAccountName"`
	BankAccountType string `json:"bankAccountType" bson:"bankAccountType"`
	DocumentNumber  string `json:"documentNumber" bson:"documentNumber"`
}

// RecipientTweaksPatch sdfkjd
type RecipientTweaksPatch struct {
	ID               string `json:"id" bson:"id"`
	TransferInterval string `json:"transfer_interval,omitempty" bson:"transfer_interval,omitempty"`
	TransferDay      int    `json:"transfer_day,omitempty" bson:"transfer_day,omitempty"`
	TransferEnabled  bool   `json:"transfer_enabled,omitempty" bson:"transfer_enabled,omitempty"`
}

// RecipientPatch sdfkjd
type RecipientPatch struct {
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	Status          string `json:"status,omitempty" bson:"status,omitempty"`
	BankCode        string `json:"bankCode" bson:"bankCode"`
	BankBranch      string `json:"bankBranch" bson:"bankBranch"`
	BankBranchVC    string `json:"bankBranchVC" bson:"bankBranchVC"`
	BankAccount     string `json:"bankAccount" bson:"bankAccount"`
	BankAccountVC   string `json:"bankAccountVC" bson:"bankAccountVC"`
	BankAccountName string `json:"bankAccountName" bson:"bankAccountName"`
	PersonType      string `json:"personType,omitempty" bson:"personType,omitempty"`
	DocumentNumber  string `json:"documentNumber,omitempty" bson:"documentNumber,omitempty"`
}

// RecipientWithDraw sdfkjd
type RecipientWithDraw struct {
	Amount      int    `json:"amount" bson:"amount"`
	RecipientID string `json:"recipient_id" bson:"recipient_id"`
}
