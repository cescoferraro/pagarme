package userClub

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
)

// FBImage TODO: NEEDS COMMENT INFO
func FBImage(w http.ResponseWriter, r *http.Request) {
	info, err := onni.GetFacebookProfile(chi.URLParam(r, "token"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	uri := `https://graph.facebook.com/` + info.ID + `/picture?type=square&width=300`
	log.Println(uri)
	http.Redirect(w, r, uri, 301)
}
