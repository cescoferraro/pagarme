package types

import (
	"gopkg.in/mgo.v2/bson"
)

// ONNiLog type for the above middleware
type ONNiLog struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate" json:"updateDate,omitempty"`
}
