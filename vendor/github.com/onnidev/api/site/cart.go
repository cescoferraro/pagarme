package site

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

// TicketCart TODO: NEEDS COMMENT INFO
func TicketCart(w http.ResponseWriter, r *http.Request) {
	partyID := chi.URLParam(r, "partyID")
	party, err := onni.Party(r.Context(), partyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	club, err := onni.Club(r.Context(), party.ClubID.Hex())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := onni.GroupedPartyProducts(r.Context(), club, party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result["_empty"])
}

// PromoCart TODO: NEEDS COMMENT INFO
func PromoCart(w http.ResponseWriter, r *http.Request) {
	partyID := chi.URLParam(r, "partyID")
	party, err := onni.Party(r.Context(), partyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	club, err := onni.Club(r.Context(), party.ClubID.Hex())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	partyProducts, err := onni.PartyPartyProducts(r.Context(), partyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer, ok := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	finalPromotions, err := onni.GetFinalPromotions(r.Context(), partyProducts, party, club, customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, finalPromotions)
}
