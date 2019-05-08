package userClub

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ListUsersClub is the shit
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
func ListUsersClub(w http.ResponseWriter, r *http.Request) {
	partiesCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	allUser, err := partiesCollection.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
