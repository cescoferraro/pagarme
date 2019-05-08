package types

import "gopkg.in/mgo.v2/bson"

// Style is a struct
type Style struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate" json:"updateDate,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Image        Image         `bson:"image" json:"image"`
}

// Location is the shit
type Location struct {
	Type        string     `json:"type" bson:"type"`
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

// SoftLocation is the shit
type SoftLocation struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

// IDName is the shit
type IDName struct {
	ID   *string `json:"id" bson:"id"`
	Name string  `json:"name" bson:"name"`
}

// Address TODO: NEEDS COMMENT INFO
type Address struct {
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Street  string `json:"street" bson:"street"`
	Number  string `json:"number" bson:"number"`
	Unit    string `json:"unit" bson:"unit"`
}

// Image is a struct
type Image struct {
	FileID       bson.ObjectId `json:"fileId" bson:"fileId,omitempty"`
	MimeType     string        `bson:"mimeType" json:"mimeType"`
	CreationDate *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
}

// ClubPatch sdfkjd
type ClubPatch struct {
	Name          string  `bson:"name,omitempty" json:"name,omitempty"`
	Mail          string  `bson:"mail,omitempty" json:"mail,omitempty"`
	Latitude      float64 `bson:"latitude" json:"latitude"`
	Longitude     float64 `bson:"longitude" json:"longitude"`
	PercentDrink  float64 `bson:"percentDrink,omitempty" json:"percentDrink,omitempty"`
	PercentTicket float64 `bson:"percentTicket,omitempty" json:"percentTicket,omitempty"`
	Description   string  `bson:"description,omitempty" json:"description,omitempty"`
	Status        string  `bson:"status,omitempty" json:"status,omitempty"`
	City          string  `bson:"city,omitempty" json:"city,omitempty"`
	State         string  `bson:"state,omitempty" json:"state,omitempty"`
	Country       string  `bson:"country,omitempty" json:"country,omitempty"`
	Street        string  `bson:"street,omitempty" json:"street,omitempty"`
	Number        string  `bson:"number,omitempty" json:"number,omitempty"`
	Unit          string  `bson:"unit,omitempty" json:"unit,omitempty"`
}
