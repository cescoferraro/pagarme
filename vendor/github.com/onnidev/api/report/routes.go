package report

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/report/dates"
	"github.com/onnidev/api/report/list"
	"github.com/onnidev/api/report/summary"
	"github.com/onnidev/api/report/users"
	"github.com/onnidev/api/shared"
	"github.com/xuri/excelize"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.With(shared.JWTAuth.Handler).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.GetVouchers).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.GetPartyByID).
		With(middlewares.AttachClubsCollection).
		With(middlewares.GetClubFromPartyMiddleware).
		With(middlewares.AttachUserClubCollection).
		With(middlewares.GetUserClubFromToken).
		Post("/report/mail/{id}", SendMail)

	r.With(shared.JWTAuth.Handler).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.GetVouchers).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.GetPartyByID).
		With(middlewares.AttachClubsCollection).
		With(middlewares.GetClubFromPartyMiddleware).
		With(middlewares.AttachUserClubCollection).
		Post("/report/party/{id}", CloseParty)

	r.With(shared.JWTAuth.Handler).
		With(middlewares.AttachVoucherCollection).
		With(middlewares.GetVouchers).
		With(middlewares.AttachPartiesCollection).
		With(middlewares.GetPartyByID).
		With(middlewares.AttachClubsCollection).
		With(middlewares.GetClubFromPartyMiddleware).
		Get("/report/download/{id}", GenerateExcel)
}

//Excel report
func Excel(data common.Data) *excelize.File {
	xlsx := excelize.NewFile()
	list.Report(xlsx, data)
	summary.Report(xlsx, data)
	dates.Report(xlsx, data)
	users.Report(xlsx, data)
	return xlsx
}
