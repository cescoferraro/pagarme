package bans

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// BanEndpoint is commented
func BanEndpoint(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "customerId")
	customer, err := onni.Customer(r.Context(), id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	var log types.Log
	bans, err := onni.BanCustomer(r.Context(), customer, &log, &userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, bans)
}
