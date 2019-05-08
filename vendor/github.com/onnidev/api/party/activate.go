package party

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Activate TODO: NEEDS COMMENT INFO
func Activate(w http.ResponseWriter, r *http.Request) {
	log.Println("caminho mais facil")
	log.Println("caminho mais facil")
	ctx := r.Context()
	id := chi.URLParam(r, "partyID")
	party, err := onni.Party(ctx, id)
	if err != nil {
		log.Println("caminho mais difickl")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	mode := chi.URLParam(r, "mode")
	if !shared.Contains([]string{"draft", "active", "inactive"}, mode) {
		err := errors.New("not a plausible mode")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClubADMIN := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate": &now,
			"updatedBy":  userClubADMIN.ID,
			"status":     strings.ToUpper(mode),
		}},
		ReturnNew: true,
	}
	log.Println("before patching ")
	repo, ok := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	var patchedParty types.Party
	_, err = repo.Collection.Find(bson.M{"_id": party.ID}).Apply(change, &patchedParty)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, true)
}
