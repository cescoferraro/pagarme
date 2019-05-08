package clubLead

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	patch, ok := r.Context().Value(middlewares.ClubLeadPatchRequestKey).(types.ClubLeadPatch)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	j, _ := json.MarshalIndent(patch, "", "    ")
	log.Println("ClubLead Patch Request")
	log.Println(string(j))
	patchedClubLead, err := onni.PatchLead(r.Context(), chi.URLParam(r, "clubLeadID"), patch)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedClubLead)
}
