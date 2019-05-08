package types

import (
	"gopkg.in/mgo.v2/bson"
)

// ClubMenuTicket TODO: NEEDS COMMENT INFO
type ClubMenuTicket struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ClubID       bson.ObjectId `json:"clubId" bson:"clubId"`
	Name         string        `json:"name" bson:"name"`
	MenuDefault  bool          ` json:"menuDefault" bson:"menuDefault"`
	Status       string        `json:"status" bson:"status"`
	Tickets      []MenuTicket  `json:"tickets" bson:"tickets"`
}

// MenuTicket TODO: NEEDS COMMENT INFO
type MenuTicket struct {
	ID                     bson.ObjectId `json:"menuTicketId" bson:"_id,omitempty"`
	Price                  Price         `json:"price" bson:"price"`
	ExhibitionName         string        `json:"exhibitionName" bson:"exhibitionName"`
	DoNotSellMoreThan      int64         `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore        *Timestamp    `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter         *Timestamp    `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	CreditConsumable       *float64      `json:"creditConsumable" bson:"creditConsumable,omitempty"`
	Active                 bool          `json:"active" bson:"active"`
	GeneralInformation     string        `json:"generalInformation" bson:"generalInformation"`
	FreeInformation        string        `json:"freeInformation" bson:"freeInformation"`
	AnniversaryInformation string        `json:"anniversaryInformation" bson:"anniversaryInformation"`
	Batches                []MenuTicket  `json:"batches" bson:"batches"`
}
