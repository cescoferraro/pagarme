package partyProduct

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/promotion"
	"github.com/onnidev/api/shared"
)

// Routes galedira[]
func Routes(r chi.Router) {
	endpoints := func(j chi.Router) {
		j.
			// With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			Get("/{partyId}", Products)
		j.
			// With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			Get("/tickets/{partyId}", TicketProducts)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/party/soft/combo/{partyId}", SoftCombo)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/promotion/party/{partyId}", promotion.ListPartyPromotions)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			Get("/promotion/{promotionID}", promotion.ReadPromotion)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/promotion/customer/{promotionID}", promotion.CustomersSummary)
	}
	r.Route("/partyProduct", endpoints)
	r.Route("/partyProducts", endpoints)
	r.Route("/partyproduct", endpoints)
	r.Route("/partyproducts", endpoints)
}
