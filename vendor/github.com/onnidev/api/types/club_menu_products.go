package types

import (
	"gopkg.in/mgo.v2/bson"
)

// ClubMenuProduct TODO: NEEDS COMMENT INFO
type ClubMenuProduct struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ClubID       bson.ObjectId `json:"clubId" bson:"clubId"`
	Name         string        `json:"name" bson:"name"`
	MenuDefault  bool          ` json:"menuDefault" bson:"menuDefault"`
	Status       string        `json:"status" bson:"status"`
	Products     []MenuProduct `json:"products" bson:"products"`
}

// MenuProduct TODO: NEEDS COMMENT INFO
type MenuProduct struct {
	ID                 bson.ObjectId  `json:"menuProductId" bson:"_id,omitempty"`
	Price              Price          `json:"price" bson:"price"`
	ExhibitionName     string         `json:"exhibitionName" bson:"exhibitionName"`
	Category           string         `json:"category" bson:"category"`
	DoNotSellMoreThan  int64          `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellAfter     *Timestamp     `json:"doNotSellAfter" bson:"doNotSellAfter,omitempty"`
	DoNotSellBefore    *Timestamp     `json:"doNotSellBefore" bson:"doNotSellBefore,omitempty"`
	ProductID          *bson.ObjectId `json:"productId" bson:"productId"`
	CreditConsumable   *float64       `json:"creditConsumable" bson:"creditConsumable"`
	Featured           bool           `json:"featured" bson:"featured"`
	Active             bool           `json:"active" bson:"active"`
	GeneralInformation string         `json:"generalInformation" bson:"generalInformation"`
	Image              Image          `json:"image" bson:"image"`
	Combo              *[]Combo       `json:"combo" bson:"combo,omitempty"`
}
