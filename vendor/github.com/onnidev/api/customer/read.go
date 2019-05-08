package customer

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Read is the shit
func Read(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "customerId")
	customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customer, err := customersRepo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, onni.RemoveInactiveClubs(r.Context(), customer))
}
