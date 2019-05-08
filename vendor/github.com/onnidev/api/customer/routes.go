package customer

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is the thing
func Routes(j chi.Router) {
	j.Route("/customer", func(r chi.Router) {
		r.Use(middlewares.AttachCustomerCollection)
		r.
			With(shared.JWTAuth.Handler).
			Get("/", ListCustomers)
		r.
			Get("/image/{customerId}", Photo)
		r.
			Get("/weblogin/{location}", FB)
		r.
			Get("/redirect", Redirect)

		r.
			With(middlewares.ReadCustomerPostFromBody).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			Post("/", CreateEndpoint)
		r.
			With(middlewares.ReadCustomerPostFromBody).
			With(middlewares.AttachInvitedCustomerCollection).
			With(middlewares.AttachPromotionalCustomerCollection).
			With(middlewares.AttachVoucherCollection).
			Post("/newcomer", NewComerEndpoint)

		r.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/trust/{mode}/{customerId}", Activate)

		r.
			With(middlewares.AttachClubsCollection).
			With(middlewares.ReadCustomerPatchFromBody).
			With(middlewares.AttachClubsCollection).
			With(middlewares.GetCustomerFromToken).
			Patch("/favorite", PatchFavorite)

		r.
			With(middlewares.AttachClubsCollection).
			With(middlewares.ReadCustomerPatchFromBody).
			With(middlewares.AttachClubsCollection).
			With(middlewares.GetCustomerFromToken).
			Patch("/", Patch)
		r.
			With(middlewares.AttachClubsCollection).
			With(middlewares.ReadCustomerPatchFromBody).
			With(middlewares.AttachClubsCollection).
			With(middlewares.GetCustomerFromToken).
			Patch("/fixup", FixUpPatch)
		r.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachBanCollection).
			With(middlewares.AttachInvoicesCollection).
			With(middlewares.AttachCardsCollection).
			Get("/intel/{customerId}", ReadFullCustomers)

		r.
			With(middlewares.ReadLoginRequestFromBody).
			Post("/login", Login)

		r.Route("/password", func(n chi.Router) {
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.ReadChangePasswordRequestFromBody).
				With(middlewares.GetCustomerFromToken).
				Post("/update", UpdatePassword)
			n.
				With(middlewares.ReadResetFromBody).
				Post("/reset", Reset)
		})

		r.
			With(middlewares.ReadFacebookLoginRequest).
			Post("/login/facebook", FacebookLogin)

		r.
			With(middlewares.ReadFacebookSignUpRequest).
			Post("/signup/facebook", FacebookSignUp)
		r.
			With(middlewares.ReadFacebookSignUpRequest).
			Post("/profile/facebook", FacebookProfile)

		r.
			// With(shared.JWTAuth.Handler).
			With(middlewares.AttachClubsCollection).
			Get("/{customerId}", Read)

		r.
			With(shared.JWTAuth.Handler).
			With(middlewares.ReadCustomerQueryFromBody).
			Post("/query", Query)
		r.
			With(middlewares.ReadCustomerCheckFromBody).
			Post("/check", Check)
	})

}
