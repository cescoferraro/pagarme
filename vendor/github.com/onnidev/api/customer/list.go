package customer

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ListCustomers is the shit
func ListCustomers(w http.ResponseWriter, r *http.Request) {
	customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	allUser, err := customersRepo.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
