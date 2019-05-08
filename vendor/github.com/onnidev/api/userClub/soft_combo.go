package userClub

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

// SoftCombo dskjfn
func SoftCombo(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("asset bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	users, err := repo.ListPromotersByClub(chi.URLParam(r, "clubId"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	promoters := []types.SoftPromoter{}
	for _, user := range users {
		promoter := types.SoftPromoter{
			PromoterName: user.Name,
			PromoterID:   user.ID.Hex(),
			Selected:     false,
		}
		promoters = append(promoters, promoter)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, promoters)
}
