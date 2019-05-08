package userClub

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Create is awesome
func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	password := shared.RangeIn(100000, 999999)
	userClub, err := onni.CreateUserClub(ctx, password)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	err = onni.PersistUserClub(ctx, userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	req, ok := ctx.Value(middlewares.ReadUserClubPostKey).(types.UserClubPostRequest)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo, ok := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := repo.GetByID(req.Clubs[rand.Intn(len(req.Clubs))])
	if err != nil {
		err := errors.New("assert bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	go onni.MailNewUserClub(club, userClub, password)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, userClub)
}
