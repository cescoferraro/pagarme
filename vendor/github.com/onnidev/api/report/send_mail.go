package report

import (
	"bytes"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// SendMail is the shit
// swagger:route POST /report/mail/{partyId} backoffice sendMail
//
// Sends a party report for the club email
//
// This will enables club promoters to have a better understanding of
// their gross sales
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: https
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: ok
func SendMail(w http.ResponseWriter, r *http.Request) {
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	all := r.Context().Value(middlewares.PartyVouchers).([]types.CompleteVoucher)
	party := r.Context().Value(middlewares.PartyIDKey).(types.Party)
	club := r.Context().Value(middlewares.ByIDKey).(types.Club)
	buf := new(bytes.Buffer)
	data := common.Data{Club: club, Party: party, Vouchers: types.VouchersList(all)}
	xlsx := Excel(data)
	xlsx.Write(buf)
	go common.MailPartyReport(userClub, party, buf)
	render.JSON(w, r, "success")
}
