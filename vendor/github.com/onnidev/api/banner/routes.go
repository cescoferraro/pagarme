package banner

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/banner", func(j chi.Router) {
		j.
			With(middlewares.AttachBannerCollection).
			// With(interfaces.IsRequestFromApp).
			Get("/", PublishedBanners)
		j.
			With(middlewares.AttachBannerCollection).
			// With(interfaces.IsRequestFromApp).
			Get("/all", AllBanners)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachBannerCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.ReadCreateBannerRequestFromBody).
			Post("/", PersistToDB)

		j.Route("/{bannerID}", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachBannerCollection).
				Get("/", Read)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachBannerCollection).
				Delete("/", Delete)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachBannerCollection).
				With(middlewares.AttachUserClubCollection).
				With(middlewares.AttachGridFSCollection).
				With(middlewares.ReadPatchBannerRequestFromBody).
				With(middlewares.GetUserClubFromToken).
				Patch("/", Patch)
		})
	})
}
