package site

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/site", func(n chi.Router) {

		n.
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			Get("/cart/{partyID}", TicketCart)
		n.
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			Get("/cart/ticket/{partyID}", TicketCart)
		n.
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/cart/promo/{partyID}", PromoCart)

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/customer/next", ByCustomer)
	})

}
