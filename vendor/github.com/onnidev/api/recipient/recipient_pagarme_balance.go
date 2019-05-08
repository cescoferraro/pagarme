package recipient

import (
	"net/http"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// BalanceTransactions is commented
func BalanceTransactions(w http.ResponseWriter, r *http.Request) {
	query, err := onni.GetFinanceQueryFromBody(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	balance, err := onni.GetPagarmeRecipientBalance(r.Context(), query.RecipientID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, balance)
}
