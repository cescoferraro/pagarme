package types

import (
	"github.com/onnidev/api/shared"
	"gopkg.in/mgo.v2/bson"
)

// AppParty type for the above middleware
type AppParty struct {
	ID                         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate               *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate                 *Timestamp    `bson:"updateDate" json:"updateDate,omitempty"`
	Name                       string        `bson:"name" json:"name"`
	NameSearchable             string        `bson:"nameSearchable" json:"nameSearchable"`
	Description                string        `bson:"description" json:"description"`
	ClubID                     bson.ObjectId `json:"clubId" bson:"clubId,omitempty"`
	MainAttraction             string        `json:"mainAttraction" bson:"mainAttraction"`
	MainAttractionSearchable   string        `json:"mainAttractionSeachable" bson:"mainAttractionSeachable"`
	OtherAttractions           string        `json:"otherAttractions" bson:"otherAttractions"`
	OtherAttractionsSearchable string        `json:"otherAttractionsSearchable" bson:"otherAttractionsSearchable"`
	StartDate                  *Timestamp    `bson:"startDate" json:"startDate,omitempty"`
	EndDate                    *Timestamp    `bson:"endDate" json:"endDate,omitempty"`
	AssumeServiceFee           bool          `json:"assumeServiceFee" bson:"assumeServiceFee"`
	Location                   Location      `json:"location" bson:"location"`
	Address                    Address       `json:"address" bson:"address"`
	ChangeAddress              bool          `json:"changeAddress" bson:"changeAddress"`

	Status string `json:"status" bson:"status"`

	BackgroundImage Image `bson:"backgroundImage" json:"backgroundImage"`
	// Optional are filed that might be inject here by Id like clubId
	Club *AppClub `json:"club" bson:"club"`
}

// Distance TODO: NEEDS COMMENT INFO
func (party AppParty) Distance(long, lat float64) float64 {
	return shared.LatLongDistance(
		party.Location.Coordinates[0],
		party.Location.Coordinates[1],
		lat,
		long)
}
