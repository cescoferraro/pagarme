package appclub

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// History TODO: NEEDS COMMENT INFO
func History(w http.ResponseWriter, r *http.Request) {
	clubID := chi.URLParam(r, "clubId")
	userClubID := chi.URLParam(r, "userClubId")
	log.Println(clubID, userClubID)
	vouchersRepo, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchers, err := vouchersRepo.ReadByUserClubIDandClubID(userClubID, clubID)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, vouchers)
}
