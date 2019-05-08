package userClub

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ByClub is the shit
// swagger:route GET /userClub backoffice gerUserClub
//
// Get all the latest parties.
//
// Get all backoffice users
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
func ByClub(w http.ResponseWriter, r *http.Request) {
	partiesCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	id := chi.URLParam(r, "clubId")

	allUser, err := partiesCollection.ListByClub(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
