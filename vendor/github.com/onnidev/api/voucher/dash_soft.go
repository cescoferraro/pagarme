package voucher

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// DashSoft is commented
func DashSoft(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "partyId")
	party, err := onni.Party(r.Context(), id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := onni.Club(r.Context(), party.ClubID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchers, err := vouchersCollection.GetByParty(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	list := types.VouchersList(vouchers)
	AmountProducts := float64(list.Accountable().FilterDrinks().Sum())
	AmountTickets := float64(list.Accountable().FilterTickets().Sum())
	AmountProductsLiquid := float64(list.Accountable().LiquidSumDrinks(club.PercentDrink))
	AmountTicketsLiquid := float64(list.Accountable().LiquidSumTickets(club.PercentTicket, true))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, struct {
		AmountTotal          float64                        `json:"amountTotal"`
		AmountTotalLiquid    float64                        `json:"amountTotalLiquid"`
		AmountProducts       float64                        `json:"amountProducts"`
		AmountProductsLiquid float64                        `json:"amountProductsLiquid"`
		AmountTickets        float64                        `json:"amountTickets"`
		AmountTicketsLiquid  float64                        `json:"amountTicketsLiquid"`
		Drinks               []types.DashSoftVoucherSummary `json:"drinks"`
		Tickets              []types.DashSoftVoucherSummary `json:"tickets"`
	}{
		AmountTotal:          AmountProducts + AmountTickets,
		AmountTotalLiquid:    AmountProductsLiquid + AmountTicketsLiquid,
		AmountProducts:       AmountProducts,
		AmountProductsLiquid: AmountProductsLiquid,
		AmountTickets:        AmountTickets,
		AmountTicketsLiquid:  AmountTicketsLiquid,
		Drinks:               list.DrinksSummary(),
		Tickets:              list.TicketsSummary(),
	})
}
