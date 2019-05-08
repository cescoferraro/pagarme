package proxy

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/buy"
	"github.com/onnidev/api/cart"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {

	r.
		With(shared.JWTAuth.Handler).
		With(middlewares.AttachPartyProductsCollection).
		With(middlewares.AttachCustomerCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.AttachClubsCollection).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.AttachPromotionalCustomerCollection).
		With(middlewares.GetCustomerFromToken).
		Get("/proxy/app/v4/shopping/cart/customer/{customerId}/party/{partyID}", cart.Cart)
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
		With(middlewares.AttachCardsCollection).
		With(middlewares.AttachPartyProductsCollection).
		With(middlewares.ReadCreateBuyPostFromBody).
		With(middlewares.GetCustomerFromToken).
		With(middlewares.AttachPromotionalCustomerCollection).
		With(middlewares.AttachBanCollection).
		Post("/proxy/app/v4/party/products/buy", buy.Business)
	r.HandleFunc("/proxy/*", Proxy)
}
