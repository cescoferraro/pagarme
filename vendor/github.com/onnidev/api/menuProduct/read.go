package menuProduct

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Read is dskj
func Read(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	all, err := repo.GetByID(chi.URLParam(r, "menuId"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, all)
}
