package voucher

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ByParty is commented
// swagger:route POST /vouchers/transfer/{voucherId} backoffice transferVoucher
//
// ByParty a voucher
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: https
//
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: vouchersType
func ByParty(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "partyId")
	vouchers, err := vouchersCollection.GetByParty(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, vouchers)
}
