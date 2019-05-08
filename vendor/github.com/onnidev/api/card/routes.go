package card

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/card", func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachTokenCollection).
			With(middlewares.AttachCardsCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/", ListCards)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachCardsCollection).
			With(middlewares.AttachTokenCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			With(middlewares.ReadCreateCardRequestFromBody).
			With(middlewares.PagarMeCreateCardHash).
			Post("/", PersistToDB)

		j.Route("/{id}", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachCardsCollection).
				Get("/", Read)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachCardsCollection).
				With(middlewares.AttachCustomerCollection).
				With(middlewares.GetCustomerFromToken).
				Delete("/", Delete)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachCardsCollection).
				With(middlewares.ReadUpdateCardRequestFromBody).
				With(middlewares.AttachCustomerCollection).
				With(middlewares.GetCustomerFromToken).
				Patch("/", Patch)
		})
	})
}
