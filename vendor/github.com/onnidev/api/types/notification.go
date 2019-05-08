package types

import (
	"encoding/json"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"

	"gopkg.in/mgo.v2/bson"
)

// Notification sdkjfn
type Notification struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`

	CreatedBy     bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedByName string        `json:"createdByName" bson:"createdByName"`

	Title  string `json:"title" bson:"title"`
	Text   string `json:"text" bson:"text"`
	Type   string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`

	PartyID   *bson.ObjectId   `json:"partyId" bson:"partyId"`
	PartyName *string          `json:"partyName" bson:"partyName"`
	ClubID    *bson.ObjectId   `json:"clubId" bson:"clubId"`
	ClubName  *string          `json:"clubName" bson:"clubName"`
	Customers *[]bson.ObjectId `json:"customers,omitempty" bson:"customers,omitempty"`

	// Creation/Update record
	UpdatedBy     *bson.ObjectId `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	UpdatedByName *string        `json:"updatedByName,omitempty" bson:"updatedByName,omitempty"`

	Party *Party `json:"party,omitempty" bson:"party,omitempty" `
	Club  *Club  `json:"club,omitempty" bson:"club,omitempty" `
}

// GetIOSNotification sdfkjn
func (notification Notification) GetIOSNotification(device string) (*apns2.Notification, error) {
	noti := &apns2.Notification{}
	noti.DeviceToken = device
	noti.Topic = "br.com.onni.app"
	byt, err := json.Marshal(notification)
	if err != nil {
		return noti, err
	}
	noti.Payload = payload.NewPayload().
		Alert(notification.Title).Sound("ping.aiff").
		Badge(1).Custom("notification", string(byt))
	return noti, nil
}

// NotificationPatchRequest skdjfn
type NotificationPatchRequest struct {
	Title   string `json:"title" bson:"title"`
	Text    string `json:"text" bson:"text"`
	PartyID string `bson:"partyId,omitempty" json:"partyId,omitempty"`
	ClubID  string `bson:"clubId,omitempty" json:"clubId,omitempty"`
}

// NotificationPostRequest skdjfn
type NotificationPostRequest struct {
	ID      string `json:"id" bson:"id"`
	Title   string `bson:"title" json:"title"`
	Text    string `bson:"text" json:"text"`
	PartyID string `bson:"partyId,omitempty" json:"partyId,omitempty"`
	ClubID  string `bson:"clubId,omitempty" json:"clubId,omitempty"`
}
