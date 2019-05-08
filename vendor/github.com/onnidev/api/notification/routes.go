package notification

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazingskdjfn
func Routes(r chi.Router) {
	r.Route("/notification", func(j chi.Router) {
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachNotificationCollection).
			Get("/", ListNotifications)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachNotificationCollection).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachCustomerCollection).
			With(middlewares.GetCustomerFromToken).
			Get("/customer/{customerID}", ByCustomer)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachNotificationCollection).
			With(middlewares.AttachPushRegistryCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/publish/{notificationID}", Publish)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachVoucherCollection).
			With(middlewares.AttachNotificationCollection).
			With(middlewares.AttachPushRegistryCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			Post("/publish/{notificationID}/{partyProductID}", Publish)
		j.
			With(shared.JWTAuth.Handler).
			With(middlewares.AttachBannerCollection).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachNotificationCollection).
			With(middlewares.AttachPartiesCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.GetUserClubFromToken).
			With(middlewares.ReadCreateNotificationRequestFromBody).
			Post("/", PersistToDB)

		j.Route("/{notificationID}", func(n chi.Router) {
			n.
				// With(shared.JWTAuth.Handler).
				With(middlewares.AttachNotificationCollection).
				Get("/", Read)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachNotificationCollection).
				Delete("/", Delete)
			n.
				With(shared.JWTAuth.Handler).
				With(middlewares.AttachClubsCollection).
				With(middlewares.AttachNotificationCollection).
				With(middlewares.AttachPartiesCollection).
				With(middlewares.AttachUserClubCollection).
				With(middlewares.GetUserClubFromToken).
				With(middlewares.ReadCreateNotificationPatchRequestFromBody).
				Patch("/", Patch)
		})
	})
}
