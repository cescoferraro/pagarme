package bans

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
		With(middlewares.AttachPushRegistryCollection).
		With(middlewares.AttachBanCollection).
		With(middlewares.AttachInvoicesCollection).
		With(middlewares.AttachCardsCollection).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.GetUserClubFromToken).
		Post("/ban/{customerId}", BanEndpoint)
}
