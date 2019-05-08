package menuProduct

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

// ByClub is dskj
func ByClub(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(repo, ok)
	all, err := repo.GetByClubActive(chi.URLParam(r, "clubId"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, all)
}
