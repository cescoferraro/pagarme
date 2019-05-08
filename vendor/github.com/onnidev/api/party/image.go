package party

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Image is the shit
func Image(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "partyId")
	partyCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	party, err := partyCollection.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	fs := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	file, err := fs.FS.OpenId(party.BackgroundImage.FileID)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+id+file.ContentType())
	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
