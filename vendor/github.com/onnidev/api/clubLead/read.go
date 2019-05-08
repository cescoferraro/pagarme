package clubLead

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// ReadEndpoint TODO: NEEDS COMMENT INFO
func ReadEndpoint(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.ClubLeadKey).(interfaces.ClubLeadRepo)
	leads, err := repo.GetByID(chi.URLParam(r, "clubLeadID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, leads)
}
