package types

import (
	"gopkg.in/mgo.v2/bson"
)

// AppClub is a mongo document
type AppClub struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp    `bson:"updateDate" json:"updateDate,omitempty"`
	Name           string        `json:"name" bson:"name"`
	Mail           string        `json:"mail" bson:"mail"`
	OperationType  string        `bson:"operationType" json:"operationType"`
	NameSearchable string        `bson:"nameSearchable" json:"nameSearchable"`
	Description    string        `bson:"description" json:"description"`
	Featured       bool          `bson:"featured" json:"featured"`

	PercentDrink   float64 `bson:"percentDrink" json:"percentDrink"`
	PercentTicket  float64 `bson:"percentTicket" json:"percentTicket"`
	PercentPrePaid bool    `bson:"percentPrePaid" json:"percentPrePaid"`
	TronEndPoint   string  `bson:"tronEndPoint" json:"tronEndPoint"`
	TronLicense    string  `bson:"tronLicense" json:"tronLicense"`
	ProductType    string  `bson:"productType" json:"productType"`

	Location           Location      `json:"location" bson:"location"`
	Address            Address       `json:"address" bson:"address"`
	PagarMeRecipientID bson.ObjectId `json:"pagarMeRecipientId" bson:"pagarMeRecipientId,omitempty"`
	Image              Image         `bson:"image" json:"image"`

	BackgroundImage *Image `bson:"backgroundImage" json:"backgroundImage"`
	// MusicStyles     []Style `bson:"musicStyles" json:"musicStyles"`

	Status         string `bson:"status" json:"status"`
	FlatProducts   bool   `bson:"flatProducts" json:"flatProducts"`
	RegisterOrigin string `bson:"registerOrigin" json:"registerOrigin"`
}
