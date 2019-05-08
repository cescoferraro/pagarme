package customer

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Login sdkjfn
func Login(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	loginRequest := r.Context().Value(middlewares.UserClubLoginReq).(types.LoginRequest)
	customer, err := collection.Login(loginRequest)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	response, err := customer.LogInCustomer()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
