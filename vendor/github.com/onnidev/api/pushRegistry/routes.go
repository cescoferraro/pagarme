package pushRegistry

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/pushRegistry", func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPushRegistryCollection).
			Get("/", ListPushRegistry)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPushRegistryCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			With(middlewares.ReadCreatePushRegistryRequestFromBody).
			Post("/", Create)
	})
}
