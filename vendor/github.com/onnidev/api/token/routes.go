package token

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// TokenRoutes is amazing
func TokenRoutes(r chi.Router) {
	r.
		With(shared.JWTAuth.Handler).
		With(middlewares.AttachTokenCollection).
		Get("/token", ListTokens)
}
