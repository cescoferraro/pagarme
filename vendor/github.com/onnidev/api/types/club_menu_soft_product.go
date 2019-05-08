package types

import (
	"gopkg.in/mgo.v2/bson"
)

// SoftProductPartyProduct TODO: NEEDS COMMENT INFO
type SoftProductPartyProduct struct {
	DoNotSellMoreThan  int64          `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore    *Timestamp     `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter     *Timestamp     `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	ExhibitionName     string         `json:"exhibitionName" bson:"exhibitionName"`
	Featured           bool           `json:"featured" bson:"featured"`
	Active             bool           `json:"active" bson:"active"`
	GeneralInformation string         `json:"generalInformation" bson:"generalInformation"`
	Category           string         `json:"category" bson:"category"`
	MenuProductID      *bson.ObjectId `json:"menuProductId" bson:"menuProductId"`
	PartyProductID     bson.ObjectId  `json:"partyProductId" bson:"partyProductId"`
	ProductID          *bson.ObjectId `json:"productId" bson:"productId"`
	Image              Image          `json:"image" bson:"image"`
	Price              Price          `json:"price" bson:"price"`
	Combo              *[]Combo       `json:"combo" bson:"combo"`
}

// AdditionalInformations TODO: NEEDS COMMENT INFO
type AdditionalInformations struct {
	General     string `json:"general"`
	Free        string `json:"free"`
	Anniversary string `json:"anniversary"`
}

// SoftComboPartyProduct TODO: NEEDS COMMENT INFO
type SoftComboPartyProduct struct {
	PartyProductID string `json:"partyProductId"`
	ProductName    string `json:"productName"`
	ProductType    string `json:"productType"`
}

// SoftCombo TODO: NEEDS COMMENT INFO
func (partyP PartyProduct) SoftCombo() SoftComboPartyProduct {
	return SoftComboPartyProduct{
		PartyProductID: partyP.ID.Hex(),
		ProductName:    partyP.Name,
		ProductType:    partyP.Type,
	}
}
