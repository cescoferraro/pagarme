package customer

import (
	"errors"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Query is the shit
func Query(w http.ResponseWriter, r *http.Request) {
	customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	query := r.Context().Value(middlewares.CustomersQueryKey).(types.CustomerQuery)
	allCustomers, err := customersRepo.Query(query)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if len(allCustomers) > 50 {
		err := errors.New("QUERY TOO LARGE")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allCustomers)
}
