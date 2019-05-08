package menuTicket

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazingskdjfn
func Routes(r chi.Router) {
	endpoints := func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuTicketCollection).
			Get("/club/{clubId}", ByClub)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuTicketCollection).
			Get("/{menuId}", Read)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuTicketCollection).
			Delete("/{menuId}", Delete)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuTicketCollection).
			Post("/clone/{menuId}", Clone)
	}
	r.Route("/menuTicket", endpoints)
	r.Route("/menuticket", endpoints)
}
