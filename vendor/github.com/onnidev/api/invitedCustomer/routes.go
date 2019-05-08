package invitedCustomer

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	result := func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachTokenCollection).
			Get("/", ListCards)
		j.
			Get("/fb/{customerId}", FB)
		j.
			With(middlewares.ReadInvitedLinkCustomerPostRequestFromBody).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			Post("/link/{inviteId}", Link)
		j.
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachCustomerCollection).
			Get("/return", Return)
	}
	r.Route("/invitedcustomer", result)
	r.Route("/invitedCustomer", result)
}
