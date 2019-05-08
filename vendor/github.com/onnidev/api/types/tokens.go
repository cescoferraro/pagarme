package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Token is a mongo document
type Token struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	UserID       bson.ObjectId `json:"userId" bson:"userId,omitempty"`
	Token        string        `bson:"token" json:"token"`
}
