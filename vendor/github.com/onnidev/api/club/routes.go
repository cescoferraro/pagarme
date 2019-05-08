package club

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/club", func(n chi.Router) {
		n.
			With(middlewares.ReadClubPostRequestFromBody).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachRecipientsCollection).
			Post("/", PersistToDB)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/", ListClubs)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/soft", ComboSoft)

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/{mode}/{clubID}", Activate)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.ReadClubPatchFromBody).
			Patch("/{clubId}", Patch)
		n.
			With(middlewares.AttachClubsCollection).
			Get("/{clubId}", AppRead)
		n.
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachRecipientsCollection).
			Get("/next/{clubId}", NextRead)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachVoucherCollection).
			Get("/dashboard/{clubId}", Dashboard)

		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachGridFSCollection).
			Post("/soft/image/main/{clubId}", UpdateImage)
		n.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachGridFSCollection).
			Post("/soft/image/background/{clubId}", BCKUpdateImage)
		n.
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachGridFSCollection).
			Get("/image/{clubId}", Image)
	})

}
