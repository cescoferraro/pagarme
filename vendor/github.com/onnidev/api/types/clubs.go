package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Club is a mongo document
type Club struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Name           string        `json:"name" bson:"name"`
	Mail           string        `json:"mail" bson:"mail"`
	OperationType  string        `bson:"operationType" json:"operationType"`
	NameSearchable string        `bson:"nameSearchable" json:"nameSearchable"`
	Description    string        `bson:"description" json:"description"`
	Featured       bool          `bson:"featured" json:"featured"`
	MusicStyles    []Style       `json:"musicStyles" bson:"musicStyles"`

	BankLegalAddress string `json:"bankLegalAddress" bson:"bankLegalAddress"`
	// DocumentNumber             string  `json:"documentNumber" bson:"documentNumber"`
	AverageExpendituresProduct *float64 `json:"averageExpendituresProduct" bson:"averageExpendituresProduct"`
	AverageExpendituresTicket  *float64 `json:"averageExpendituresTicket" bson:"averageExpendituresTicket"`
	PercentDrink               float64  `bson:"percentDrink" json:"percentDrink"`
	PercentTicket              float64  `bson:"percentTicket" json:"percentTicket"`
	PercentPrePaid             *float64 `bson:"percentPrePaid" json:"percentPrePaid"`
	// TronEndPoint   string  `bson:"tronEndPoint" json:"tronEndPoint"`
	// TronLicense    string  `bson:"tronLicense" json:"tronLicense"`
	ProductionType string `bson:"productionType" json:"productionType"`

	Location           Location      `json:"location" bson:"location"`
	Address            Address       `json:"address" bson:"address"`
	PagarMeRecipientID bson.ObjectId `json:"pagarMeRecipientId" bson:"pagarMeRecipientId,omitempty"`
	Image              Image         `bson:"image" json:"image"`

	BackgroundImage *Image `bson:"backgroundImage" json:"backgroundImage"`

	Recipients     *[]Recipient `bson:"recipients" json:"recipients"`
	Tags           *[]string    `bson:"tags" json:"tags"`
	Status         string       `bson:"status" json:"status"`
	FlatProducts   bool         `bson:"flatProducts" json:"flatProducts"`
	RegisterOrigin string       `bson:"registerOrigin" json:"registerOrigin"`
	Liability      *string      `bson:"liability" json:"liability"`
}

// ClubPostRequest is a mongo document
type ClubPostRequest struct {
	Name          string   `json:"name" bson:"name"`
	Mail          string   `json:"mail" bson:"mail"`
	Description   string   `bson:"description" json:"description"`
	PercentDrink  float64  `bson:"percentDrink" json:"percentDrink"`
	PercentTicket float64  `bson:"percentTicket" json:"percentTicket"`
	Latitude      float64  `bson:"latitude" json:"latitude"`
	Longitude     float64  `bson:"longitude" json:"longitude"`
	MusicStyles   []string `json:"musicStyles,omitempty" bson:"musicStyles,omitempty"`

	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Street  string `json:"street" bson:"street"`
	Number  string `json:"number" bson:"number"`
	Unit    string `json:"unit" bson:"unit"`

	Recipient struct {
		BankCode           string `json:"bankCode" bson:"bankCode"`
		BankBranch         string `json:"bankBranch" bson:"bankBranch"`
		BankBranchVC       string `json:"bankBrasanchVC" bson:"bankBranchVC"`
		BankAccount        string `json:"bankAccount" bson:"bankAccount"`
		BankAccountVC      string `json:"bankAccountVC" bson:"bankAccountVC"`
		BankAccountName    string `json:"bankAccountName" bson:"bankAccountName"`
		BankAccountType    string `json:"bankAccountType" bson:"bankAccountType"`
		BankLegalAddress   string `json:"bankLegalAddress" bson:"bankLegalAddress"`
		BankDocumentNumber string `json:"documentNumber" bson:"bankDocumentNumber"`
	} `json:"recipient" bson:"recipient"`

	Featured       bool    `bson:"featured" json:"featured"`
	OperationType  string  `bson:"operationType" json:"operationType"`
	PercentPrePaid float64 `bson:"percentPrePaid" json:"percentPrePaid"`
	ProductionType string  `bson:"productionType" json:"productionType"`
	DocumentNumber string  `json:"documentNumber" bson:"documentNumber"`
	Status         string  `bson:"status" json:"status"`
}
