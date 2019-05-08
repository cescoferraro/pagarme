package types

import (
	"gopkg.in/mgo.v2/bson"
)

// PromotionalCustomer type for the above middleware
type PromotionalCustomer struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	PromotionID  bson.ObjectId `json:"promotionId" bson:"promotionId"`
	CustomerID   bson.ObjectId `json:"customerId" bson:"customerId"`
	CustomerName string        `json:"customerName" bson:"customerName"`
	CustomerMail string        `json:"customerMail" bson:"customerMail"`
	PromoterID   bson.ObjectId `json:"promoterId" bson:"promoterId"`
	PromoterName string        `json:"promoterName" bson:"promoterName"`
}

// PromotionalCustomerPost TODO: NEEDS COMMENT INFO
type PromotionalCustomerPost struct {
	Mails        []string `json:"mails" bson:"mails"`
	PromotionIDS []string `json:"promotionIds" bson:"promotionIds"`
}
