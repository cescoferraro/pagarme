package file

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		// With(shared.JWTAuth.Handler).
		With(middlewares.AttachGridFSCollection).
		With(middlewares.ReadFileFromBody).
		Put("/file", Add)
	r.
		// With(shared.JWTAuth.Handler).
		With(middlewares.AttachGridFSCollection).
		With(middlewares.ReadFileFromBody).
		Put("/file/{min}", AddMin)

	r.
		With(middlewares.AttachGridFSCollection).
		Get("/file/{id}", Get)
	r.
		With(middlewares.AttachGridFSCollection).
		Get("/file/s3/{id}", Aws)
}
