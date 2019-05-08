package userClub

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Attendant is the shit
// swagger:route POST /user/login backoffice loginResponse
//
// Login a user
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
//       200: loginResponse
func Attendant(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	loginRequest := r.Context().Value(middlewares.UserClubLoginReq).(types.LoginRequest)
	user, err := collection.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	if user.Profile != "ATTENDANT" {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	clubsCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	var clubs []types.Club
	clubs, err = clubsCollection.Mine(user.Clubs)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	// if len(user.Clubs) == 0 {
	// 	render.Status(r, http.StatusExpectationFailed)
	// 	render.JSON(w, r, err.Error())
	// 	return
	// }
	response, err := onni.Response(r.Context(), user, clubs)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
