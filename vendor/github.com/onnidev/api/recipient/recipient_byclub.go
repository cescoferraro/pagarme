package recipient

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ListRecipients is commented
func ListClubRecipients(w http.ResponseWriter, r *http.Request) {
	recipientsRepo := r.Context().
		Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	all, err := recipientsRepo.GetByClubID(chi.URLParam(r, "clubId"))
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, all)
}
