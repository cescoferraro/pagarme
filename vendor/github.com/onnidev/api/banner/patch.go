package banner

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Patch skdjfn
func Patch(w http.ResponseWriter, r *http.Request) {
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	bannerRepo := r.Context().Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	bannerPatch := r.Context().Value(middlewares.BannerPatchRequestKey).(types.BannerPatch)
	id := chi.URLParam(r, "bannerID")
	banner, err := bannerRepo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	image := banner.BannerImage
	if bannerPatch.Image != "" {
		grid := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
		file, err := grid.FS.OpenId(bson.ObjectIdHex(bannerPatch.Image))
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		fileDate := types.Timestamp(file.UploadDate())
		image = types.Image{
			FileID:       bson.ObjectIdHex(bannerPatch.Image),
			MimeType:     file.ContentType(),
			CreationDate: &fileDate,
		}
	}
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":  types.Timestamp(time.Now()),
			"updatedBy":   userClub.ID,
			"type":        orBlank(banner.Type, bannerPatch.Type),
			"action":      orBlankObjectID(banner.Action, bannerPatch.Action),
			"name":        orBlank(banner.Name, bannerPatch.Name),
			"description": orBlank(banner.Description, bannerPatch.Description),
			"status":      orBlank(banner.Status, bannerPatch.Status),
			"bannerImage": image,
		}},
		ReturnNew: true,
	}

	var patchedBanner types.Banner
	_, err = bannerRepo.Collection.Find(bson.M{"_id": banner.ID}).Apply(change, &patchedBanner)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(patchedBanner.Name)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedBanner)
}

func orBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}

func orBlankObjectID(og bson.ObjectId, sent string) bson.ObjectId {
	if sent == "" || !bson.ObjectIdHex(sent).Valid() {
		return og
	}
	return bson.ObjectIdHex(sent)
}
