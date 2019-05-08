package clubLead

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// CreateEndpoint TODO: NEEDS COMMENT INFO
func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	repo, ok := r.Context().Value(middlewares.ClubLeadKey).(interfaces.ClubLeadRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	req, ok := r.Context().Value(middlewares.ClubLeadRequestKey).(types.ClubLeadPostRequest)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	now := types.Timestamp(time.Now())
	lead := types.ClubLead{
		ID:              bson.NewObjectId(),
		CreationDate:    &now,
		AdminName:       req.AdminName,
		AdminMail:       req.AdminMail,
		Image:           req.Image,
		BackgroundImage: req.BackgroundImage,
		AdminPhone:      req.AdminPhone,
		Stage:           "1",
	}
	err := repo.Create(lead)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(lead, "", "    ")
		log.Println("ClubLead created on MongoDB")
		log.Println(string(j))
	}
	render.Status(r, 200)
	render.JSON(w, r, lead)
}
