package types

import (
	"github.com/onnidev/api/shared"
	"gopkg.in/mgo.v2/bson"
)

// Party type for the above middleware
type Party struct {
	ID                         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate               *Timestamp    `json:"creationDate" bson:"creationDate,omitempty"`
	UpdateDate                 *Timestamp    `json:"updateDate" bson:"updateDate,omitempty"`
	Name                       string        `json:"name" bson:"name"`
	NameSearchable             string        `json:"nameSearchable" bson:"nameSearchable"`
	Description                string        `json:"description" bson:"description"`
	ClubID                     bson.ObjectId `json:"clubId" bson:"clubId,omitempty"`
	MainAttraction             string        `json:"mainAttraction" bson:"mainAttraction"`
	MainAttractionSearchable   string        `json:"mainAttractionSeachable" bson:"mainAttractionSeachable"`
	OtherAttractions           string        `json:"otherAttractions" bson:"otherAttractions"`
	OtherAttractionsSearchable string        `json:"otherAttractionsSearchable" bson:"otherAttractionsSearchable"`
	EndDate                    *Timestamp    `json:"endDate" bson:"endDate,omitempty"`
	StartDate                  *Timestamp    `json:"startDate" bson:"startDate,omitempty"`

	PreSaleStartDate *Timestamp `json:"preSaleStartDate" bson:"preSaleStartDate,omitempty"`
	PreSaleEndDate   *Timestamp `json:"preSaleEndDate" bson:"preSaleEndDate,omitempty"`

	AvarageExpeditureTickets  *float64 `json:"avarageExpeditureTickets,omitempty" bson:"avarageExpeditureTickets,omitempty"`
	AvarageExpeditureProducts *float64 `json:"avarageExpeditureProducts,omitempty" bson:"avarageExpeditureProducts,omitempty"`

	AssumeServiceFee bool     `json:"assumeServiceFee" bson:"assumeServiceFee"`
	Location         Location `json:"location" bson:"location"`
	Address          Address  `json:"address" bson:"address"`
	ChangeAddress    bool     `json:"changeAddress" bson:"changeAddress"`

	Status string `json:"status" bson:"status"`

	ClubMenuTicketID  *bson.ObjectId `json:"clubMenuTicketId" bson:"clubMenuTicketId,omitempty"`
	ClubMenuProductID *bson.ObjectId `json:"clubMenuProductId" bson:"clubMenuProductId,omitempty"`

	MusicStyles     []Style `json:"musicStyles" bson:"musicStyles"`
	BackgroundImage Image   `bson:"backgroundImage" json:"backgroundImage"`
	// Optional are filed that might be inject here by Id like clubId
	Club *Club     `json:"club" bson:"club"`
	Tags *[]string `json:"tags" bson:"tags"`
}

// Distance TODO: NEEDS COMMENT INFO
func (party Party) Distance(long, lat float64) float64 {
	return shared.LatLongDistance(
		party.Location.Coordinates[0],
		party.Location.Coordinates[1],
		lat,
		long)
}

// SmallParty TODO: NEEDS COMMENT INFO
type SmallParty struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name             string        `bson:"name" json:"name"`
	Description      string        `bson:"description" json:"description"`
	Location         SoftLocation  `json:"location" bson:"location"`
	StartDate        *Timestamp    `bson:"startDate" json:"startDate,omitempty"`
	EndDate          *Timestamp    `bson:"endDate" json:"endDate,omitempty"`
	MainAttraction   string        `json:"mainAttraction" bson:"mainAttraction"`
	OtherAttractions string        `json:"otherAttractions" bson:"otherAttractions"`
	Address          Address       `json:"address" bson:"address"`
	BackgroundImage  Image         `bson:"backgroundImage" json:"backgroundImage"`
}

// SmallParty TODO: NEEDS COMMENT INFO
func (party *Party) SmallParty() SmallParty {
	return SmallParty{
		ID:          party.ID,
		Name:        party.Name,
		Description: party.Description,
		Location: SoftLocation{
			Longitude: party.Location.Coordinates[0],
			Latitude:  party.Location.Coordinates[1],
		},
		StartDate:        party.StartDate,
		EndDate:          party.EndDate,
		MainAttraction:   party.MainAttraction,
		OtherAttractions: party.OtherAttractions,
		Address:          party.Address,
		BackgroundImage:  party.BackgroundImage,
	}
}

// FullParty dsfjkhn
type FullParty struct {
	Party
	Club Club `json:"club" bson:"club"`
}

// Parties is a list of cards of a giver user
// swagger:response partiesList
type Parties struct {
	// in: body
	Body []Party `json:"body"`
}

// PartyPathParamID are the data you need to send in order to
// swagger:parameters sendMail generateExcel
type PartyPathParamID struct {
	// in: path
	PartyID string `json:"partyId"`
}

// PartyFilter dkfjgndfkgj
type PartyFilter struct {
	From *Timestamp `json:"from,omitempty" bson:"from,omitempty"`
	Till *Timestamp `json:"till,omitempty" bson:"till,omitempty"`
	Lat  float64    `json:"lat" bson:"lat"`
	Long float64    `json:"long" bson:"long"`
}
