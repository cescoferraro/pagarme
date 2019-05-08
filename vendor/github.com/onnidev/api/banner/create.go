package banner

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// PersistToDB is a comented function
func PersistToDB(w http.ResponseWriter, r *http.Request) {
	bannerRequest := r.Context().Value(middlewares.BannerRequestKey).(types.BannerPostRequest)
	grid := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	file, err := grid.FS.OpenId(bson.ObjectIdHex(bannerRequest.Image))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	bannerRepo := r.Context().Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	horario := types.Timestamp(file.UploadDate())
	now := types.Timestamp(time.Now())
	id := bson.ObjectIdHex(bannerRequest.ID)
	if !id.Valid() {
		id = bson.NewObjectId()
	}
	banner := types.Banner{
		ID:          id,
		Description: bannerRequest.Description,
		Type:        bannerRequest.Type,
		Name:        bannerRequest.Name,
		Status:      "NOT_PUBLISHED",
		BannerImage: types.Image{
			FileID:       bson.ObjectIdHex(bannerRequest.Image),
			MimeType:     file.ContentType(),
			CreationDate: &horario,
		},
		Action:       bson.ObjectIdHex(bannerRequest.Action),
		CreatedByID:  userClub.ID,
		CreationDate: &now,
	}
	err = bannerRepo.Create(banner)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(banner, "", "    ")
		log.Println("Banner created on MongoDB")
		log.Println(string(j))
	}
	render.JSON(w, r, bannerRequest)
}
