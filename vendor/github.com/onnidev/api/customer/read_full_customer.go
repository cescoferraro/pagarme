package customer

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

// ListCustomers is the shit
func ReadFullCustomers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "customerId")
	cardsRepo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	customerCards, err := cardsRepo.GetByCustomerID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}

	allInvoices, err := onni.InvoicesFromUser(r.Context(), id)
	alldevice := []string{}
	for _, invoice := range allInvoices {
		if invoice.Log.DeviceID != "" {
			alldevice = shared.SamePayload(alldevice, invoice.Log.DeviceID)
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, IntelResponse{
		ID:      id,
		Devices: alldevice,
		Cards:   customerCards,
	})
}

// IntelResponse sdkjfn
type IntelResponse struct {
	ID      string       `json:"id"`
	Devices []string     `json:"devices"`
	Cards   []types.Card `json:"cards"`
}
