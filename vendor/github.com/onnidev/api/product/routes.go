package product

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes galedira[]
func Routes(r chi.Router) {
	r.Route("/product", func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachGridFSCollection).
			Put("/", Add)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			Get("/", Products)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			Get("/{productId}", Read)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.ReadProductPostRequestFromBody).
			Post("/", Post)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.ReadProductPatchRequestFromBody).
			Patch("/{productId}", Patch)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.ReadProductSoftPatchRequestFromBody).
			Patch("/soft/{productId}", SoftPatch)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachGridFSCollection).
			Post("/soft/image/{productId}", UpdateImage)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachProductsCollection).
			With(middlewares.AttachGridFSCollection).
			Delete("/{id}", Delete)
	})
}
