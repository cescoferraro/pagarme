package auth

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// LeitorReAuth TODO: NEEDS COMMENT INFO
func LeitorReAuth(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middlewares.JWTRefreshRequestKey).(types.JWTRefresh)
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
	clubsCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	clubs, err := clubsCollection.MineClubs(user)
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
