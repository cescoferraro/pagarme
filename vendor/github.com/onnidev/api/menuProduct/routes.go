package menuProduct

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
			With(middlewares.AttachClubMenuProductCollection).
			Get("/{menuId}", Read)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuProductCollection).
			Get("/club/{clubId}", ByClub)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuProductCollection).
			Delete("/{menuId}", Delete)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubMenuProductCollection).
			Post("/clone/{menuId}", Clone)
	}
	r.Route("/menuproduct", endpoints)
	r.Route("/menuProduct", endpoints)
}
