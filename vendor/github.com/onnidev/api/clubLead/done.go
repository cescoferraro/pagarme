package clubLead

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Done TODO: NEEDS COMMENT INFO
func Done(w http.ResponseWriter, r *http.Request) {
	patch, ok := r.Context().Value(middlewares.ClubLeadPatchRequestKey).(types.ClubLeadPatch)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	lead, err := onni.PatchLead(r.Context(), chi.URLParam(r, "clubLeadID"), patch)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if lead.Stage != "4" {
		err := errors.New("you are not ready daniel san")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	_, onniRecipient, err := onni.CreateALLRecipients(r.Context(), lead)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	password := shared.RangeIn(100000, 999999)
	userClub := lead.UserClub(password)
	err = onni.PersistUserClub(r.Context(), userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	img, err := onni.ImageONNi(r.Context(), lead.Image)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	bimg, err := onni.ImageONNi(r.Context(), lead.BackgroundImage)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := onni.ClubFromLead(r.Context(), lead, onniRecipient.ID, img, bimg)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	go onni.MailNewUserClub(club, userClub, password)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, club)
}
