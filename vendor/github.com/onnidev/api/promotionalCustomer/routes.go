package promotionalCustomer

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes galedira[]
func Routes(r chi.Router) {
	endpoint := func(n chi.Router) {
		n.
			With(middlewares.AttachPromotionalCustomerCollection).
			Get("/", List)
		n.
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.ReadPromotionalCustomerQueryFromBody).
			Post("/", Create)
		n.
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.ReadPromotionalCustomerQueryFromBody).
			Post("/invite", Invite)
	}
	r.Route("/promotionalcustomer", endpoint)
	r.Route("/promotionalCustomer", endpoint)
}
