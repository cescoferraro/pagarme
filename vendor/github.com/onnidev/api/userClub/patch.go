package userClub

import (
	"errors"
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
	userClubADMIN := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	userClubRepo := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	userClub, err := userClubRepo.GetByID(chi.URLParam(r, "userClubID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	userClubPatch := r.Context().Value(middlewares.UserClubPatchRequestKey).(types.UserClubPatch)

	var clubs []bson.ObjectId
	for _, club := range userClubPatch.Clubs {
		clubs = append(clubs, bson.ObjectIdHex(club))
	}

	exists, err := userClubRepo.ExistsByKey("mail", userClubPatch.Mail)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if exists && userClub.Mail != userClubPatch.Mail {
		err := errors.New("email already exists")
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate": &now,
			"updatedBy":  userClubADMIN.ID,
			"name":       orBlank(userClub.Name, userClubPatch.Name),
			"mail":       orBlank(userClub.Mail, userClubPatch.Mail),
			"status":     orBlank(userClub.Status, userClubPatch.Status),
			"password":   userClub.Password,
			"clubs":      clubs,
			"profile":    orBlank(userClub.Profile, userClubPatch.Profile),
		}},
		ReturnNew: true,
	}

	var patchedUserClub types.UserClub
	_, err = userClubRepo.Collection.Find(bson.M{"_id": userClub.ID}).Apply(change, &patchedUserClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(patchedUserClub.Name)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, patchedUserClub)
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
