package report

import (
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
)

// GenerateExcel is the shit
// swagger:route GET /report/download/{partyId} backoffice generateExcel
//
// Downloads an excel Sheet with all vouchers consumed at the party
// with the given id
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
//
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: Binary

func GenerateExcel(w http.ResponseWriter, r *http.Request) {
	all := r.Context().Value(middlewares.PartyVouchers).([]types.CompleteVoucher)
	party := r.Context().Value(middlewares.PartyIDKey).(types.Party)
	club := r.Context().Value(middlewares.ByIDKey).(types.Club)
	data := common.Data{
		Club:     club,
		Party:    party,
		Vouchers: types.VouchersList(all),
	}
	xlsx := Excel(data)
	common.SendThroughHTTP(xlsx, w, party.Name)
	return
}
