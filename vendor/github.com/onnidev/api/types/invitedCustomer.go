package types

import "gopkg.in/mgo.v2/bson"

// InvitedCustomer TODO: NEEDS COMMENT INFO
type InvitedCustomer struct {
	ID             bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp     `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp     `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Done           bool           `json:"done" bson:"done"`
	Mail           string         `json:"mail" bson:"mail"`
	FBID           *string        `json:"fbid" bson:"fbid"`
	LinkedCustomer *bson.ObjectId `json:"linkedCustomer" bson:"linkedCustomer"`
	AssignedMail   *string        `json:"assignedMail" bson:"assignedMail"`
	Customer       *Customer      `json:"customer,omitempty" bson:"customer,omitempty"`
}

// InvitedLinkCustomerPost TODO: NEEDS COMMENT INFO
type InvitedLinkCustomerPost struct {
	Mail string `json:"mail" bson:"mail"`
}
