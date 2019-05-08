package location

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes galedira[]
func Routes(r chi.Router) {

	r.Route("/locate", func(n chi.Router) {
		n.
			With(middlewares.ReadGeoLocalizationPostRequestFromBody).
			Post("/", Locate)
	})
}
