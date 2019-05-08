package club

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)

	id := chi.URLParam(r, "clubId")
	club, err := repo.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	clubPatch := r.Context().Value(middlewares.ClubPatchKey).(types.ClubPatch)
	var patchedCustomer types.Customer
	log.Println(clubPatch.Latitude)
	log.Println(clubPatch.Longitude)
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":    &now,
				"name":          orBlank(club.Name, clubPatch.Name),
				"percentDrink":  orBlankFloat64(club.PercentDrink, clubPatch.PercentDrink),
				"percentTicket": orBlankFloat64(club.PercentTicket, clubPatch.PercentTicket),
				"location": types.Location{
					Type:        "Point",
					Coordinates: [2]float64{clubPatch.Latitude, clubPatch.Longitude},
				},
				// Address is the shit
				"address": types.Address{
					City:    orBlank(club.Address.City, clubPatch.City),
					State:   orBlank(club.Address.State, clubPatch.State),
					Country: orBlank(club.Address.Country, clubPatch.Country),
					Street:  orBlank(club.Address.Street, clubPatch.Street),
					Number:  orBlank(club.Address.Number, clubPatch.Number),
					Unit:    orBlank(club.Address.Unit, clubPatch.Unit),
				},
				"mail":        orBlank(club.Mail, clubPatch.Mail),
				"description": orBlank(club.Description, clubPatch.Description),
				"status":      orBlank(club.Status, clubPatch.Status),
			}},
		ReturnNew: true,
	}
	_, err = repo.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &patchedCustomer)
	if err != nil {

		shared.MakeONNiError(w, r, 400, err)

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedCustomer)
}

func orBlankFloat64(og, sent float64) float64 {
	if sent == 0 {
		return og
	}
	return sent
}

func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}

func orBlanktime(og, sent time.Time) time.Time {
	if sent.IsZero() {
		return og
	}
	return sent
}
