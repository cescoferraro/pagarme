package auth

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// UserClubReAuth TODO: NEEDS COMMENT INFO
func UserClubReAuth(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middlewares.PagarMeReAuthRequestKey).(types.PagarMeReAuth)
	var claim types.MyClaimsType
	jwtToken, err := jwt.ParseWithClaims(
		token.Token,
		&claim,
		shared.JWTAuth.Options.ValidationKeyGetter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if !jwtToken.Valid {
		err := errors.New("token expired")
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
	userClubCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	user, err := userClubCollection.GetByID(claim.ClientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if user.Profile == "ATTENDANT" {
		render.Status(r, http.StatusExpectationFailed)
		err := errors.New("attendente não loga no backoffice")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	clubsCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	clubs, err := clubsCollection.MineClubs(user)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	newtoken, err := user.GenerateToken()
	if err != nil {
		render.Status(r, http.StatusBadGateway)
		render.JSON(w, r, err.Error())
		return
	}
	response := types.UserClubLoginResponse{
		ID:      user.ID,
		Token:   newtoken,
		Name:    user.Name,
		Mail:    user.Mail,
		Profile: user.Profile,
		Clubs:   clubs,
		ClubID:  token.ClubID,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
