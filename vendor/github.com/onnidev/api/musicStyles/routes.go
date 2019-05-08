package musicStyles

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
			With(middlewares.AttachMusicStylesCollection).
			Get("/", List)
	}
	r.Route("/musicstyles", endpoints)
	r.Route("/musicStyles", endpoints)
}
