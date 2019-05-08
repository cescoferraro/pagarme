package onni

import (
	"context"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	geo "github.com/martinlindhe/google-geolocate"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// ClubFromLead TODO: NEEDS COMMENT INFO
func ClubFromLead(ctx context.Context, lead types.ClubLead, pgid bson.ObjectId, img types.Image, bkimg types.Image) (types.Club, error) {
	club := types.Club{}
	horario := types.Timestamp(time.Now())
	client := geo.NewGoogleGeo(viper.GetString("GOOGLEMAPSTOKEN"))
	point, err := client.Geocode(lead.Address())
	if err != nil {
		if err.Error() == "ZERO_RESULTS" {
			point = &geo.Point{
				Lng: -30.0346471,
				Lat: -51.2176584,
			}
		} else {
			return club, err
		}
	}
	zero := float64(0.0)
	club = types.Club{
		ID:               lead.ID,
		CreationDate:     &horario,
		Name:             lead.ClubName,
		Description:      lead.ClubDescription,
		NameSearchable:   strings.ToLower(lead.ClubName),
		BankLegalAddress: lead.BankLegalAddress,
		Mail:             lead.AdminMail,
		OperationType:    "ONNI_APP",
		Featured:         true,
		MusicStyles:      []types.Style{},

		AverageExpendituresProduct: &zero,
		AverageExpendituresTicket:  &zero,

		PercentDrink:  95.0,
		PercentTicket: 90.0,

		ProductionType: lead.ClubType,

		Location: types.Location{
			Type:        "Point",
			Coordinates: [2]float64{point.Lat, point.Lng},
		},
		Address: types.Address{
			City:    lead.City,
			State:   lead.State,
			Country: lead.Country,
			Street:  lead.Street,
			Number:  lead.Number,
			Unit:    lead.Unit,
		},
		PagarMeRecipientID: pgid,
		Image:              img,
		BackgroundImage:    &bkimg,
		Status:             "PENDING",
		FlatProducts:       false,
		RegisterOrigin:     "SITE",
	}
	err = PersistClub(ctx, club)
	if err != nil {
		return club, err
	}
	return club, nil
}
