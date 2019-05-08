package file

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Get TODO: NEEDS COMMENT INFO
func Get(w http.ResponseWriter, r *http.Request) {
	fs := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	id := chi.URLParam(r, "id")
	log.Println(id)
	file, err := fs.FS.OpenId(bson.ObjectIdHex(id))
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	defer file.Close()
	log.Println(file.ContentType())
	w.Header().Set("Content-Disposition", "attachment; filename="+id)
	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
