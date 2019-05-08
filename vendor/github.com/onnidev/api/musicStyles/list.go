package musicStyles

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// List is dskj
func List(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.MusicStylesRepoKey).(interfaces.MusicStylesRepo)
	log.Println(repo, ok)
	all, err := repo.List()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, all)
}
