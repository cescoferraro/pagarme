package userClub

import (
	"net/http"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// LoginUser TODO: NEEDS COMMENT INFO
func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := onni.LoginUserClub(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 520, err)
		return
	}
	clubs, err := onni.UserClubClubs(ctx, user)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	response, err := onni.Response(r.Context(), user, clubs)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
