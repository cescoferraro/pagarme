package party

import (
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// UserClubListParties sdjknfdskjf
func UserClubListParties(w http.ResponseWriter, r *http.Request) {
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	clubsRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	clubs, err := clubsRepo.MineClubs(userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	var clubIDS []bson.ObjectId
	for _, club := range clubs {
		clubIDS = append(clubIDS, club.ID)
	}
	partiesCollection := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	allParties, err := partiesCollection.GetByClubIDS(clubIDS)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	if viper.GetBool("verbose") {
		// j, _ := json.MarshalIndent(allParties, "", "    ")
		// log.Println("Parties UserClubHasAccess")
		// log.Println(string(j))
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, allParties)
}
