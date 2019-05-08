package types

import (
	"gopkg.in/mgo.v2/bson"
)

// SoftParty TODO: NEEDS COMMENT INFO
type SoftParty struct {
	ID               string                    `json:"id"`
	StartDate        *Timestamp                `json:"startDate"`
	EndDate          *Timestamp                `json:"endDate"`
	Description      string                    `json:"description"`
	Name             string                    `json:"name"`
	Status           string                    `json:"status"`
	Address          Address                   `json:"address"`
	ClubID           bson.ObjectId             `json:"clubId" bson:"clubId"`
	BackgroundImage  Image                     `json:"backgroundImage"`
	MainAttraction   string                    `json:"mainAttraction"`
	OtherAttractions string                    `json:"otherAttractions"`
	AssumeServiceFee bool                      `json:"assumeServiceFee"`
	MusicStyles      []string                  `json:"musicStyles"`
	Tickets          []SoftTicketPartyProduct  `json:"tickets" bson:"tickets"`
	Products         []SoftProductPartyProduct `json:"products" bson:"products"`

	ClubMenuTicket  *IDName `json:"clubMenuTicket" bson:"clubMenuTicket"`
	ClubMenuProduct *IDName `json:"clubMenuProduct" bson:"clubMenuProduct"`
}

// SoftPartyPostRequest TODO: NEEDS COMMENT INFO
type SoftPartyPostRequest struct {
	Address          Address                       `json:"address"`
	AssumeServiceFee bool                          `json:"assumeServiceFee"`
	ClubID           string                        `json:"clubId" bson:"clubId"`
	ClubMenuProduct  *IDName                       `json:"clubMenuProduct" bson:"clubMenuProduct"`
	ClubMenuTicket   *IDName                       `json:"clubMenuTicket" bson:"clubMenuTicket"`
	Description      string                        `json:"description"`
	EndDate          *Timestamp                    `json:"endDate"`
	MainAttraction   string                        `json:"mainAttraction"`
	MusicStyles      []string                      `json:"musicStyles"`
	Name             string                        `json:"name"`
	OtherAttractions string                        `json:"otherAttractions"`
	PartyID          string                        `json:"partyId"`
	Products         []SoftProductPostPartyProduct `json:"products" bson:"products"`
	StartDate        *Timestamp                    `json:"startDate"`
	Tickets          []SoftTicketPostPartyProduct  `json:"tickets" bson:"tickets"`
}

// SoftProductPostPartyProduct TODO: NEEDS COMMENT INFO
type SoftProductPostPartyProduct struct {
	DoNotSellMoreThan  int64      `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore    *Timestamp `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter     *Timestamp `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	ExhibitionName     string     `json:"exhibitionName" bson:"exhibitionName"`
	Featured           bool       `json:"featured" bson:"featured"`
	Active             bool       `json:"active" bson:"active"`
	Category           string     `json:"category" bson:"category"`
	MenuProductID      string     `json:"menuProductId" bson:"menuProductId"`
	GeneralInformation string     `json:"generalInformation" bson:"generalInformation"`
	PartyProductID     string     `json:"partyProductId" bson:"partyProductId"`
	ProductID          string     `json:"productId" bson:"productId"`
	Image              string     `json:"image" bson:"image"`
	Price              float64    `json:"price" bson:"price"`
	Combo              *[]Combo   `json:"combo" bson:"combo"`
}

// SoftTicketPostPartyProduct TODO: NEEDS COMMENT INFO
type SoftTicketPostPartyProduct struct {
	ID                     string                            `json:"id" bson:"_id"`
	Batches                []SoftSlaveTicketPostPartyProduct `json:"batches" bson:"batches"`
	Active                 bool                              `json:"active" bson:"active"`
	AdditionalInformations AdditionalInformations            `json:"additionalInformations" bson:"additionalInformations"`
	DoNotSellAfter         *Timestamp                        `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	DoNotSellBefore        *Timestamp                        `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellMoreThan      int64                             `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	ExhibitionName         string                            `json:"exhibitionName" bson:"exhibitionName"`
	PartyProductID         string                            `json:"partyProductId,omitempty" bson:"partyProductId,omitempty"`
	MenuTicketID           string                            `json:"menuTicketId" bson:"menuTicketId"`
	Price                  float64                           `json:"price" bson:"price"`
	Featured               bool                              `json:"featured" bson:"featured"`
	// // CreditConsumable       *float64                          `json:"creditConsumable" bson:"creditConsumable"`
}

// SoftSlaveTicketPostPartyProduct TODO: NEEDS COMMENT INFO
type SoftSlaveTicketPostPartyProduct struct {
	DoNotSellMoreThan      int64                  `json:"doNotSellMoreThan" bson:"doNotSellMoreThan"`
	DoNotSellBefore        *Timestamp             `bson:"doNotSellBefore,omitempty" json:"doNotSellBefore,omitempty"`
	DoNotSellAfter         *Timestamp             `bson:"doNotSellAfter,omitempty" json:"doNotSellAfter,omitempty"`
	ExhibitionName         string                 `json:"exhibitionName" bson:"exhibitionName"`
	Active                 bool                   `json:"active" bson:"active"`
	MenuTicketID           string                 `json:"menuTicketId" bson:"menuTicketId"`
	AdditionalInformations AdditionalInformations `json:"additionalInformations" bson:"additionalInformations"`
	PartyProductID         string                 `json:"partyProductId,omitempty" bson:"partyProductId,omitempty"`
	Price                  float64                `json:"price" bson:"price"`
}
