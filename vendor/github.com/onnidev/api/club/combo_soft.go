package club

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ComboSoft is commented
func ComboSoft(w http.ResponseWriter, r *http.Request) {
	UserRepo := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	allUser, err := UserRepo.MineClubs(userClub)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	combo := []Combo{}
	for _, club := range allUser {
		combo = append(combo, Combo{
			ID:            club.ID,
			Name:          club.Name,
			PercentTicket: club.PercentTicket,
			Tags:          club.Tags,
		})
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, combo)
}

// Combo TODO: NEEDS COMMENT INFO
type Combo struct {
	ID            bson.ObjectId `json:"id" bson:"id"`
	Name          string        `json:"name" bson:"name"`
	Tags          *[]string     `json:"tags" bson:"tags"`
	PercentTicket float64       `json:"percentTicket" bson:"percentTicket"`
}
