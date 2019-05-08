package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SoftTicketPartyProduct TODO: NEEDS COMMENT INFO
type SoftTicketPartyProduct struct {
	DoNotSellMoreThan      int64                         `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore        *Timestamp                    `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter         *Timestamp                    `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	ExhibitionName         string                        `json:"exhibitionName" bson:"exhibitionName"`
	Featured               bool                          `json:"featured" bson:"featured"`
	Active                 bool                          `json:"active" bson:"active"`
	MenuTicketID           *bson.ObjectId                `json:"menuTicketId" bson:"menuTicketId"`
	PartyProductID         bson.ObjectId                 `json:"partyProductId" bson:"partyProductId"`
	Batches                []SoftSlaveTicketPartyProduct `json:"batches" bson:"batches"`
	Price                  Price                         `json:"price" bson:"price"`
	AdditionalInformations AdditionalInformations        `json:"additionalInformations" bson:"additionalInformations"`
	CreditConsumable       *float64                      `json:"creditConsumable" bson:"creditConsumable"`
}

// SoftSlaveTicketPartyProduct TODO: NEEDS COMMENT INFO
type SoftSlaveTicketPartyProduct struct {
	DoNotSellMoreThan int64          `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore   *Timestamp     `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter    *Timestamp     `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	ExhibitionName    string         `json:"exhibitionName" bson:"exhibitionName"`
	Active            bool           `json:"active" bson:"active"`
	MenuTicketID      *bson.ObjectId `json:"menuTicketId" bson:"menuTicketId"`
	PartyProductID    bson.ObjectId  `json:"partyProductId" bson:"partyProductId"`
	Price             Price          `json:"price" bson:"price"`
}

// SoftSlaveTicket FKNJSD  iiii qnj q
func (partyP PartyProduct) SoftSlaveTicket() SoftSlaveTicketPartyProduct {
	status := false
	if partyP.Status == "ACTIVE" {
		status = true
	}
	empty := Timestamp(time.Time{})
	product := SoftSlaveTicketPartyProduct{
		PartyProductID:    partyP.ID,
		DoNotSellMoreThan: partyP.DoNotSellMoreThan,
		DoNotSellBefore:   &empty,
		DoNotSellAfter:    &empty,
		ExhibitionName:    partyP.Name,
		MenuTicketID:      partyP.MenuTicketID,
		Price:             partyP.MoneyAmount,
		Active:            status,
	}
	if partyP.DoNotSellAfter != nil {
		product.DoNotSellAfter = partyP.DoNotSellAfter
	}
	if partyP.DoNotSellBefore != nil {
		product.DoNotSellBefore = partyP.DoNotSellBefore
	}
	return product

}
