package club

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// NextRead sdkjfn
func NextRead(w http.ResponseWriter, r *http.Request) {
	repo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	id := chi.URLParam(r, "clubId")
	card, err := repo.GetByID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, 22)
		return
	}

	recirepo := r.Context().Value(middlewares.RecipientCollectionKey).(interfaces.RecipientsRepo)
	user, err := recirepo.GetByClubID(card.ID.Hex())
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	card.Recipients = &user
	render.Status(r, http.StatusOK)
	render.JSON(w, r, card)
}
