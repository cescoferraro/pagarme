package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Banner skjdnkk
type Banner struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	CreatedByID  bson.ObjectId `json:"createdById" bson:"createdById"`
	UpdatedByID  bson.ObjectId `json:"updatedById,omitempty" bson:"updatedById,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Description  string        `json:"description" bson:"description"`
	Type         string        `json:"type" bson:"type"`
	Status       string        `bson:"status" json:"status"`
	Action       bson.ObjectId `json:"action" bson:"action"`
	BannerImage  Image         `json:"bannerImage" bson:"bannerImage"`
	Party        *Party        `json:"party,omitempty" bson:"party,omitempty" `
	Club         *Club         `json:"club,omitempty" bson:"club,omitempty" `
}

// BannerPatch skjdnkk
type BannerPatch struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Status      string `bson:"status,omitempty" json:"status,omitempty"`
	Action      string `json:"action,omitempty" bson:"action,omitempty"`
	Image       string `json:"image,omitempty" bson:"image,omitempty"`
}

// BannerPostRequest type for the above middleware
type BannerPostRequest struct {
	ID          string `json:"id" bson:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `json:"description" bson:"description"`
	Type        string `bson:"type" json:"type"`
	Action      string `bson:"action" json:"action"`
	Image       string `json:"image" bson:"image"`
}
