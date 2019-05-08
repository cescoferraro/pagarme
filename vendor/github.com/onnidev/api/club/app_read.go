package club

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// AppRead sdkjfn
func AppRead(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	card, err := repo.AppGetByID(chi.URLParam(r, "clubId"))
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, 22)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, card)
}
