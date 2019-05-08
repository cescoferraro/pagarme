package cart

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/cart", func(n chi.Router) {

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/{partyID}", Cart)
	})

}
