package oplog

import (
	"github.com/go-chi/chi"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Get("/", Info)
}
