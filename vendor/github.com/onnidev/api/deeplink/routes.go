package deeplink

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is the thing
func Routes(j chi.Router) {
	j.Route("/deeplink", func(r chi.Router) {
		r.Get("/", MainDeeplink)
		r.With(middlewares.AttachPartiesCollection).Get("/party/{partyID}", PartyDeeplink)
		r.With(middlewares.AttachClubsCollection).Get("/club/{clubID}", CLubDeepLInk)
	})
	j.Route("/tickets", func(r chi.Router) {
		r.Get("/", BudDeeplink)
	})

}
