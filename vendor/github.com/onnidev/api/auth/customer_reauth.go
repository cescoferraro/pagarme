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

// Auth TODO: NEEDS COMMENT INFO
func Auth(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middlewares.JWTRefreshRequestKey).(types.JWTRefresh)
	var claim types.CustomerClaims
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
	customerRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customer, err := customerRepo.GetByID(claim.CustomerID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	newToken, err := customer.GenerateToken()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.JWTRefresh{Token: newToken})
}
