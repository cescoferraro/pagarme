package userClub

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// LoginLeitor sdkjfn
func LoginLeitor(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	loginRequest := r.Context().Value(middlewares.UserClubLoginReq).(types.LoginRequest)
	user, err := collection.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		shared.MakeONNiError(w, r, 403, err)
		return
	}
	if user.Profile != "ATTENDANT" && user.Profile != "ADMIN" {
		shared.MakeONNiError(w, r, 406, err)
		return
	}
	clubsCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	var clubs []types.Club
	clubs, err = clubsCollection.Mine(user.Clubs)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	response, err := onni.Response(r.Context(), user, clubs)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
