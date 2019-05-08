package userClub

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Oauth2 TODO: NEEDS COMMENT INFO
func Oauth2(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	oauth2LoginRequest := r.Context().
		Value(middlewares.UserClubOauth2MailKey).(types.Oauth2LoginRequest)
	fbValidation, err := onni.FacebookStaffAppValidate(oauth2LoginRequest.AccessToken)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(fbValidation)
	info, err := onni.GetFacebookProfile(oauth2LoginRequest.AccessToken)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	user, err := collection.FacebookExtraUserRegex(info.Email)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

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
