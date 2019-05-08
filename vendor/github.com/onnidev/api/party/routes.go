package party

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes galedira[]
func Routes(r chi.Router) {
	r.Route("/party", func(n chi.Router) {
		n.Use(middlewares.AttachPartiesCollection)

		n.
			With(shared.JWTAuth.Handler).
			Get("/", ListParties)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.AttachClubsCollection).
			Get("/userClub", UserClubListParties)

		n.
			With(middlewares.AttachGridFSCollection).
			Get("/image/{partyId}", Image)
		n.
			With(middlewares.AttachInvoicesCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachCardsCollection).
			With(middlewares.ReadAntiTheftModelRequestFromBody).
			Post("/antitheft/{partyId}", AntiTheft)

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/{mode}/{partyID}", Activate)
		// ONNILOCK
		n.
			With(middlewares.ReadPartyListFilterRequestFromBody).
			Post("/", AppListPartiesFiltered)
		n.
			With(middlewares.ReadSoftPartyPostRequestRequestFromBody).
			With(middlewares.AttachMusicStylesCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachClubMenuTicketCollection).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachClubMenuProductCollection).
			Put("/", Create)
		n.
			With(middlewares.ReadSoftPartyPostRequestRequestFromBody).
			With(middlewares.AttachMusicStylesCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachClubMenuTicketCollection).
			With(middlewares.AttachClubMenuProductCollection).
			Patch("/{partyID}", Patch)

		n.
			With(middlewares.ReadPartyListFilterRequestFromBody).
			Post("/next", ListPartiesFiltered)

		n.
			With(middlewares.AttachUserClubCollection).
			With(interfaces.IsRequestFromApp).
			Get("/club/{clubId}", ReadClubParties)

		n.
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachClubsCollection).
			Get("/soft/club/{clubId}", SoftClubInfo)
		n.
			With(middlewares.AttachUserClubCollection).
			Get("/club/site/{clubId}", ReadClubPartiesSite)

		n.
			// With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachClubMenuTicketCollection).
			With(middlewares.AttachClubMenuProductCollection).
			Get("/soft/{partyId}", ReadPartySoft)
		n.
			// With(shared.JWTAuth.Handler).
			Get("/{partyId}", ReadParty)

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachGridFSCollection).
			Post("/soft/image/{partyId}", UpdateImage)
	})
}
