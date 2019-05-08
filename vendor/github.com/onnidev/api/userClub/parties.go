package userClub

import (
	"log"
	"net/http"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// PartiesUserClub is the shit
func PartiesUserClub(w http.ResponseWriter, r *http.Request) {
	log.Println("right route at least")
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	reports, err := onni.HTTPPartyReport(r.Context(), userClub.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		log.Println(err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, reports)
}
