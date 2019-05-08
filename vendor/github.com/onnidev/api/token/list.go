package token

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/pressly/chi/render"
)

// ListTokens is commented
func ListTokens(w http.ResponseWriter, r *http.Request) {
	UserRepo := r.Context().Value(middlewares.TokensRepoKey).(interfaces.TokensRepo)
	allUser, err := UserRepo.List()
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allUser)
}
