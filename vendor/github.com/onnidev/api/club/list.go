package club

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ListClubs is commented
func ListClubs(w http.ResponseWriter, r *http.Request) {
	UserRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	allUser, err := UserRepo.MineClubs(userClub)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
