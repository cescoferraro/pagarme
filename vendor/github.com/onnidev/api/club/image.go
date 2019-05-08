package club

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// Image is the shit
func Image(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clubId")
	repo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	club, err := repo.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	fs := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	file, err := fs.FS.OpenId(club.Image.FileID)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+id+file.ContentType())
	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
