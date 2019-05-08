package userClub

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes galedira[]
func Routes(r chi.Router) {
	endpoints := func(j chi.Router) {

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			Get("/", ListUsersClub)
		j.
			With(middlewares.AttachUserClubCollection).
			With(middlewares.ReadCustomerCheckFromBody).
			Post("/check", Check)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.ReadUserClubPostRequestFromBody).
			Post("/", Create)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/{mode}/{userClubID}", Activate)

		j.Get("/image/{token}", FBImage)

		j.Route("/{userClubID}", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.ReadUserClubPatchRequestFromBody).
				With(middlewares.AttachUserClubCollection).
				With(middlewares.GetUserClubFromToken).
				Patch("/", Patch)
		})
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/android/vouchers", SimpleVoucherReadByUserClub)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Get("/vouchers", VoucherReadByUserClub)

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.ReadChangePasswordRequestFromBody).
			Patch("/password", ChangePasswordEndpoint)

		j.Route("/club", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachUserClubCollection).
				Get("/{clubId}", ByClub)
		})
		j.Route("/parties", func(n chi.Router) {
			n.
				With(middlewares.AttachUserClubCollection).
				With(middlewares.AttachVoucherCollection).
				With(middlewares.AttachClubsCollection).
				With(middlewares.AttachPartiesCollection).
				With(middlewares.GetUserClubFromToken).
				Get("/", PartiesUserClub)
		})
		j.Route("/password", func(n chi.Router) {
			n.
				With(middlewares.AttachUserClubCollection).
				With(middlewares.ReadResetFromBody).
				Post("/reset", Reset)
		})

		j.Route("/onni", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachUserClubCollection).
				Get("/", ONNi)
		})
		j.Route("/login", func(i chi.Router) {
			i.
				With(middlewares.AttachUserClubCollection).
				With(middlewares.ReadLoginRequestFromBody).
				With(middlewares.AttachClubsCollection).
				Post("/", LoginUser)
			i.
				With(middlewares.AttachUserClubCollection).
				With(middlewares.ReadSoftLoginRequestFromBody).
				With(middlewares.AttachClubsCollection).
				Post("/soft", SoftLoginUser)

			i.Route("/leitor", func(f chi.Router) {

				f.
					With(middlewares.AttachUserClubCollection).
					With(middlewares.ReadLoginRequestFromBody).
					With(middlewares.AttachClubsCollection).
					With(middlewares.AttachPartiesCollection).
					Post("/", LoginLeitor)

				f.
					With(middlewares.AttachUserClubCollection).
					With(middlewares.AttachClubsCollection).
					With(middlewares.ReadOauth2LoginEmail).
					Post("/oauth2", Oauth2)
			})
		})

		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachPartyProductsCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			Get("/promoter/{clubId}", SoftCombo)
	}
	r.Route("/userclub", endpoints)
	r.Route("/userClub", endpoints)
}
