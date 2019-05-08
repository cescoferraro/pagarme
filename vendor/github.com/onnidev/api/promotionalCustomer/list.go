package promotionalCustomer

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// List sdjknfdskjf
func List(w http.ResponseWriter, r *http.Request) {

	repo := r.Context().
		Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	promotions, err := repo.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, promotions)
}
