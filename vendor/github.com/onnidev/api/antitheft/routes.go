package antitheft

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		With(middlewares.AttachBanCollection).
		With(middlewares.AttachAntiTheftCollection).
		Get("/antitheft", List)
	r.
		With(middlewares.AttachBanCollection).
		With(middlewares.AttachAntiTheftCollection).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.GetUserClubFromToken).
		Post("/antitheft/{mode}/{antitheftId}", Activate)

}
