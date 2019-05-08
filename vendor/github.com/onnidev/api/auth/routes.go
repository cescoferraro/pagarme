package auth

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		With(middlewares.ReadCreateJWTRefreshRequestFromBody).
		With(middlewares.AttachCustomerCollection).
		Post("/auth", Auth)
	r.
		With(middlewares.ReadCreateJWTRefreshRequestFromBody).
		With(middlewares.AttachCustomerCollection).
		Post("/auth/customer/login", CustomerLogin)
	r.
		With(middlewares.ReadCreatePagarMeReAuthRequestFromBody).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachClubsCollection).
		Post("/auth/userClub", UserClubReAuth)
	r.
		With(middlewares.ReadCreateJWTRefreshRequestFromBody).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachClubsCollection).
		Post("/auth/leitor", LeitorReAuth)
	r.Get("/auth/sms/{phone}", Sms)
	r.Get("/auth/sms/again/{phone}", Sms)
	r.Get("/auth/sms/totalvoice/{phone}", TotalVoice)
	r.Get("/auth/sms/twilio/{phone}", Twilio)
}
