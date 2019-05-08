package party

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Patch TODO: NEEDS COMMENT INFO
func Patch(w http.ResponseWriter, r *http.Request) {
	party, err := onni.Party(r.Context(), chi.URLParam(r, "partyID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if party.ClubMenuTicketID != nil {
		id := *party.ClubMenuTicketID
		log.Println("#### party has clubmenuticket", id.Hex())
	}
	if party.ClubMenuProductID != nil {
		id := *party.ClubMenuProductID
		log.Println("#### party has clubmenuproduct", id.Hex())
	}
	req, ok := r.Context().Value(middlewares.SoftPartyPostRequestKey).(types.SoftPartyPostRequest)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	batches, err := onni.MusicStyles(r.Context(), req.MusicStyles)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	clubMenuTicketID, err := onni.PatchClubMenuTicket(r.Context(), req, party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	clubMenuProductID, err := onni.PatchClubMenuProduct(r.Context(), req, party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	patchedParty, err := onni.PatchParty(r.Context(), party, req, batches, clubMenuProductID, clubMenuTicketID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedParty.ID.Hex())
}
