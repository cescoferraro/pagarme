package site

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ByCustomer dskfjnsd
func ByCustomer(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	vouchers, err := vouchersCollection.GetActualVoucherByCustomer(customer.ID.Hex())
	if err != nil {
		log.Println(err.Error())
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	// j, _ := json.MarshalIndent(result, "", "    ")
	// log.Println("Request received from body")
	render.Status(r, http.StatusOK)
	render.JSON(w, r, vouchers)
}
