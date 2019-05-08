package cart

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Cart TODO: NEEDS COMMENT INFO
func Cart(w http.ResponseWriter, r *http.Request) {
	partyID := chi.URLParam(r, "partyID")
	customer, err := onni.CustomerFromToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	log.Println("festa tem this many partyProducfts", len(partyProducts))
	voucherCount, err := onni.HowManyVoucheraCustomerHas(r.Context(), customer, partyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	result, err := onni.GroupedPartyProducts(r.Context(), club, party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	ticketStatus, err := onni.TicketStatus(r.Context(), result["_empty"], partyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	drinkStatus, err := onni.DrinkStatus(r.Context(), result, partyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	finalPromotions, err := onni.GetFinalPromotions(r.Context(), partyProducts, party, club, customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.Cart{
		Promotions:      finalPromotions,
		Party:           party.SmallParty(),
		Club:            club.SmallClub(),
		GroupedProducts: result,
		ShoppingCounter: voucherCount,
		DrinkStatus:     drinkStatus,
		TicketStatus:    ticketStatus,
	})
}
