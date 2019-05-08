package types

import (
	"gopkg.in/mgo.v2/bson"
)

// PushRegistry skjdnkk
type PushRegistry struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	CustomerID   bson.ObjectId `json:"customerId,omitempty" bson:"customerId,omitempty"`
	DeviceUUID   string        `bson:"deviceUUID" json:"deviceUUID"`
	DeviceToken  string        `bson:"deviceToken" json:"deviceToken"`
	Platform     string        `bson:"platform" json:"platform"`
}

// PushRegistryPostRequest type for the above middleware
type PushRegistryPostRequest struct {
	DeviceUUID  string `bson:"deviceUUID" json:"deviceUUID"`
	DeviceToken string `bson:"deviceToken" json:"deviceToken"`
	Platform    string `bson:"platform" json:"platform"`
}
