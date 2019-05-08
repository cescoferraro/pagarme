package promotion

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes galedira[]
func Routes(r chi.Router) {
	endpoints := func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.ReadPromotionPostRequestRequestFromBody).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			Post("/", CreatePromotion)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.ReadPromotionPatchRequestRequestFromBody).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			Patch("/{promotionID}", PatchPromotion)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			Get("/{promotionID}", ReadPromotion)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/party/{partyId}", ListPartyPromotions)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/customer/{promotionID}", CustomersSummary)
	}
	r.Route("/promotion", endpoints)
	r.Route("/promotions", endpoints)
}
