package party

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// SoftClubInfo is the shit
func SoftClubInfo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clubId")
	clubsRepo, ok := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := clubsRepo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, struct {
		ClubID      string        `json:"clubId" bson:"clubId"`
		MusicStyles []types.Style `json:"musicStyles" bson:"musicStyles"`
		Address     types.Address `json:"address" bson:"address"`
	}{

		Address:     club.Address,
		MusicStyles: club.MusicStyles,
		ClubID:      club.ID.Hex(),
	})
}
