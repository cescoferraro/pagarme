package clubLead

import (
	"log"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// ListEndpoint TODO: NEEDS COMMENT INFO
func ListEndpoint(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.ClubLeadKey).(interfaces.ClubLeadRepo)
	leads, err := repo.List()
	if err != nil {
		log.Println(err.Error())
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, leads)
}
