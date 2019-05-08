package party

import (
	"net/http"

	"github.com/bradfitz/slice"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ReadClubParties is the shit
// swagger:route GET /parties backoffice getPartiesbyID
//
// Get all the latest parties.
//
// By latest we meand parties with startDate grater than
// 30 days before today.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: https
//
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: partiesList
//       404: error
func ReadClubParties(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clubId")
	fromApp := r.Context().Value(interfaces.IsFromAppKey).(bool)
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if fromApp {
		parties, err := partiesCollection.GetByAppClubID(id, fromApp)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		// red carpet
		if id == "5b31b3ab08d1250001d0d000" {
			// coolture sobe a serrra
			party1, err := partiesCollection.GetByAppID("5af8c25acc922d4af19fb9c5")
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			if party1.Status != "CLOSED" {
				parties = append(parties, party1)
			}
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, SortAppParties(fromApp, parties))
		return
	}
	parties, err := partiesCollection.GetByClubID(id, fromApp)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	var hey types.MyClaimsType
	_, err = jwt.ParseWithClaims(
		r.Header.Get("JWT_TOKEN"),
		&hey,
		shared.JWTAuth.Options.ValidationKeyGetter)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClubCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	clubUser, err := userClubCollection.GetByID(hey.ClientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	result := []types.Party{}
	for _, party := range parties {
		if clubUser.Profile == "PRODUCER" {
			// result = append(result, party)
			if party.Tags != nil {
				if clubUser.Tags != nil {
					for _, tag := range *party.Tags {
						if contains(*clubUser.Tags, tag) {
							result = append(result, party)
							continue
						}
					}
				}
			}
			continue
		}
		result = append(result, party)
		continue
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, SortParties(fromApp, result))

}
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// SortParties TODO: NEEDS COMMENT INFO
func SortAppParties(fromApp bool, parties []types.AppParty) []types.AppParty {
	if fromApp {
		slice.Sort(parties[:], func(i, j int) bool {
			return (parties[i].StartDate.Time().Before(parties[j].StartDate.Time()))
		})
	} else {

		slice.Sort(parties[:], func(i, j int) bool {
			return (parties[j].StartDate.Time().Before(parties[i].StartDate.Time()))
		})
	}
	return parties
}

// SortParties TODO: NEEDS COMMENT INFO
func SortParties(fromApp bool, parties []types.Party) []types.Party {
	if fromApp {
		slice.Sort(parties[:], func(i, j int) bool {
			return (parties[i].StartDate.Time().Before(parties[j].StartDate.Time()))
		})
	} else {

		slice.Sort(parties[:], func(i, j int) bool {
			return (parties[j].StartDate.Time().Before(parties[i].StartDate.Time()))
		})
	}
	return parties
}
