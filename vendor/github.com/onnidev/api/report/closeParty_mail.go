package report

import (
	"bytes"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// CloseParty is the shit
// swagger:route POST /report/party/{partyId} backoffice reportCloseParty
//
// Sends a party report for all club admins
//
// This will enables club promoters to have a better understanding of
// their gross sales at right when the party closes.
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
func CloseParty(w http.ResponseWriter, r *http.Request) {
	all := r.Context().Value(middlewares.PartyVouchers).([]types.CompleteVoucher)
	party := r.Context().Value(middlewares.PartyIDKey).(types.Party)
	club := r.Context().Value(middlewares.ByIDKey).(types.Club)
	buf := new(bytes.Buffer)
	data := common.Data{Club: club, Party: party, Vouchers: types.VouchersList(all)}
	xlsx := Excel(data)
	err := xlsx.Write(buf)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	userClubCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	allUsers, err := userClubCollection.ListByClub(club.ID.Hex())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	for _, user := range allUsers {
		if user.Profile == "ADMIN" {
			go common.MailPartyReport(user, party, buf)

		}
	}
	render.JSON(w, r, "success")
}
