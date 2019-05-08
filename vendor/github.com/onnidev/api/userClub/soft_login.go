package userClub

import (
	"net/http"

	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// SoftLoginUser TODO: NEEDS COMMENT INFO
func SoftLoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := onni.SoftLoginUserClub(r.Context())
	if err != nil {
		shared.MakeONNiError(w, r, 520, err)
		return
	}
	clubs, err := onni.UserClubClubsIDS(ctx, user)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	response, err := onni.SoftResponse(r.Context(), user, clubs)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
