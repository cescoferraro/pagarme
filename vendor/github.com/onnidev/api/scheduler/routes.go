package scheduler

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		With(shared.JWTAuth.Handler).
		Post("/task", CreateTask)
}
