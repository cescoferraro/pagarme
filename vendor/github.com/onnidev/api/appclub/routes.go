package appclub

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachTokenCollection).
		With(middlewares.ReadSoftLoginRequestFromBody).
		Post("/appclub/v1/userclub/login", Login)
	r.
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachTokenCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.CheckToken).
		With(middlewares.HeaderScan).
		Get("/appclub/v1/vouchers/read/club/{clubId}/userClub/{userClubId}", History)
	r.
		With(middlewares.ReadVoucherSoftValidateReqFromBody).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachTokenCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.HeaderScan).
		With(middlewares.CheckToken).
		Post("/appclub/v1/voucher/validate", Validate)
	r.
		With(middlewares.ReadVoucherSoftReadReqFromBody).
		With(middlewares.AttachPartyProductsCollection).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.AttachTokenCollection).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.HeaderScan).
		With(middlewares.CheckToken).
		Post("/appclub/v1/voucher/read", Read)

}
