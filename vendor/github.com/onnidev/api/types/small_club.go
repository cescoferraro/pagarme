package types

import (
	"gopkg.in/mgo.v2/bson"
)

// SmallClub TODO: NEEDS COMMENT INFO
func (club Club) SmallClub() SmallClub {
	return SmallClub{
		ID:            club.ID,
		Name:          club.Name,
		OperationType: club.OperationType,
		Image:         club.Image,
	}

}

// SmallClub TODO: NEEDS COMMENT INFO
type SmallClub struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name          string        `json:"name" bson:"name"`
	OperationType string        `bson:"operationType" json:"operationType"`
	Image         Image         `bson:"image" json:"image"`
}
