package club

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Dashboard is the shit
// swagger:route GET /clubs/dashboard/{clubId} backoffice getDashboardByClubID
//
// Get all the latest parties.
//
// By latest we meand parties with startDate grater than
// 30 days before today.
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
//       200: partiesList
//       404: error
func Dashboard(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clubId")
	clubCollection := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	club, err := clubCollection.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.ClubDashboard{
		Address: club.Address,
		Total:   33,
	})
}
