package voucher

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/voucher", func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.ReadVoucherPostRequestFromBody).
			With(middlewares.GetUserClubFromToken).
			Post("/", Create)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.ReadVoucherPostRequestFromBody).
			With(middlewares.GetUserClubFromToken).
			Post("/invite", Invite)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.ReadCustomerTranferableEmailFromBody).
			With(middlewares.GetUserClubFromToken).
			Post("/transfer/{id}", Transfer)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachInvoicesCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/refund/{voucherId}", Refund)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			Get("/party/{partyId}", ByParty)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/party/soft/{partyId}", ByPartySoft)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachClubsCollection).
			Get("/party/graph/soft/{partyId}", DashSoft)

		// ONNILOCK
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/customer", AppByCustomer)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/customer/next", ByCustomer)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/{voucherId}", Read)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.CheckToken).
			With(middlewares.HeaderScan).
			Get("/soft/{voucherId}", Read)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Delete("/{voucherId}", Delete)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.ReadVoucherUsingConstrains).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/use/soft/{voucherId}", UseDash)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.ReadVoucherUsingConstrains).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/use/{voucherId}", Use)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/use/android/{voucherId}", UseAndroid)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/undo/{voucherId}", Undo)
	})
}
