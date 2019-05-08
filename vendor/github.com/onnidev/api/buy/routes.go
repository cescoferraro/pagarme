package buy

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		With(shared.JWTAuth.Handler).
		With(middlewares.AttachCustomerCollection).
		With(middlewares.AttachRecipientsCollection).
		With(middlewares.AttachCardsCollection).
		With(middlewares.AttachBanCollection).
		With(middlewares.AttachPushRegistryCollection).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.AttachClubsCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.AttachInvoicesCollection).
		With(middlewares.AttachPartyProductsCollection).
		With(middlewares.AttachPromotionalCustomerCollection).
		With(middlewares.GetCustomerFromToken).
		With(middlewares.ReadCreateBuyPostFromBody).
		Post("/buy", Business)
	r.
		With(shared.JWTAuth.Handler).
		With(middlewares.AttachCustomerCollection).
		With(middlewares.AttachRecipientsCollection).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachCardsCollection).
		With(middlewares.AttachBanCollection).
		With(middlewares.AttachAntiTheftCollection).
		With(middlewares.AttachPushRegistryCollection).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.AttachClubsCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.AttachInvoicesCollection).
		With(middlewares.AttachPartyProductsCollection).
		With(middlewares.AttachPromotionalCustomerCollection).
		With(middlewares.GetUserClubFromToken).
		Post("/refund/{voucherId}", Refund)
}
