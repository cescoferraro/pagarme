package appclub

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/twinj/uuid"
)

// Login TODO: NEEDS COMMENT INFO
func Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, ok := r.Context().Value(middlewares.SoftUserClubLoginReq).(types.SoftLoginRequest)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	collection, ok := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	email := strings.ToLower(loginRequest.Email)
	user, err := collection.Login(email, loginRequest.Password)
	if err != nil {
		if err.Error() == "not found" {
			err := errors.New("user.club.login.error.login.failed")
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	// if !shared.Contains([]string{"ATTENDANT", "ADMIN"}, user.Profile) {
	// 	err := errors.New("usuário não faz leitura")
	// 	shared.MakeONNiError(w, r, 400, err)
	// 	return
	// }
	tokenrepo, ok := r.Context().Value(middlewares.TokensRepoKey).(interfaces.TokensRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	token := uuid.NewV4().String()
	horario := types.Timestamp(time.Now())
	err = tokenrepo.Collection.Insert(types.Token{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		UserID:       user.ID,
		Token:        token,
	})
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	response := struct {
		ID     string `json:"id" bson:"id"`
		Token  string `json:"token" bson:"token"`
		Name   string `json:"name" bson:"name"`
		Mail   string `json:"mail" bson:"mail"`
		ClubID string `json:"clubId" bson:"clubId"`
	}{
		ID:     user.ID.Hex(),
		ClubID: user.Clubs[0].Hex(),
		Name:   user.Name,
		Mail:   user.Mail,
		Token:  token,
	}
	render.JSON(w, r, response)
}
