package types

import (
	"gopkg.in/mgo.v2/bson"
)

// JWTRefresh skjdnkk
type JWTRefresh struct {
	Token string `bson:"token" json:"token"`
}

// PagarMeReAuth skjdnkk
type PagarMeReAuth struct {
	Token  string `bson:"token" json:"token"`
	ClubID string `bson:"clubId" json:"clubId"`
}

// AntiTheftResult TODO: NEEDS COMMENT INFO
type AntiTheftResult struct {
	ID                 bson.ObjectId               `json:"id" bson:"_id,omitempty"`
	PartyID            bson.ObjectId               `json:"partyId" bson:"partyId"`
	ClubID             bson.ObjectId               `json:"clubId" bson:"clubId"`
	CreationDate       *Timestamp                  `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate         *Timestamp                  `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Score              int                         `json:"score" bson:"score"`
	ScoreCard          int                         `json:"scoreCard" bson:"scoreCard"`
	Ban                *Ban                        `json:"ban" bson:"ban,omitempty"`
	Bans               *[]Ban                      `json:"bans" bson:"bans,omitempty"`
	Invoices           *Invoices                   `json:"invoices" bson:"invoices,omitempty"`
	Vouchers           *[]bson.ObjectId            `json:"vouchers" bson:"vouchers,omitempty"`
	ScorePicture       int                         `json:"scorePicture" bson:"scorePicture"`
	PG                 *PagarMeTransactionResponse `json:"pg" bson:"pg"`
	ScoreTickets       int                         `json:"scoreTickets" bson:"scoreTickets"`
	Customer           *Customer                   `json:"customer" bson:"customer,omitempty"`
	Club               *Club                       `json:"club" bson:"club,omitempty"`
	Party              *Party                      `json:"party" bson:"party,omitempty"`
	ScoreDrinks        int                         `json:"scoreDrinks" bson:"scoreDrinks"`
	Cards              int                         `json:"cards" bson:"cards"`
	Total              float64                     `json:"total" bson:"total"`
	CustomerID         bson.ObjectId               `json:"customerId" bson:"customerId"`
	Reviewed           bool                        `json:"reviewed" bson:"reviewed"`
	CustomerMail       string                      `json:"customerMail" bson:"customerMail"`
	CustomerName       string                      `json:"customerName" bson:"customerName"`
	CustomerDate       string                      `json:"customerDate" bson:"customerDate"`
	AccoutableDrinks   float64                     `json:"accountablesDrinks" bson:"accountablesDrinks"`
	AccoutableTickets  float64                     `json:"accountablesTickets" bson:"accountablesTickets"`
	ClubTicketsAverage float64                     `json:"clubTicketAvg" bson:"clubTicketAvg"`
	ClubDrinksAverage  float64                     `json:"clubDrinkAvg" bson:"clubDrinkAvg"`
}

// FullTheft TODO: NEEDS COMMENT INFO
type FullTheft struct {
	AntiTheftResult
	Banned bool `json:"banned" bson:"banned"`
}
