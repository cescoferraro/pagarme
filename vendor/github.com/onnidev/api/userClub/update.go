package userClub

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ChangePassword is the shit
// swagger:route UPDATE /userClub backoffice updateUserClub
//
// UPDATE a user password
//
// Get all backoffice users
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: https
//
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: partiesList
func ChangePasswordEndpoint(w http.ResponseWriter, r *http.Request) {
	userCollection := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
	userDB := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	req := r.Context().Value(middlewares.UserClubReq).(types.ChangePasswordRequest)
	if !compare(req.Password, userDB.Password) {
		render.Status(r, 401)
		render.JSON(w, r, http.StatusText(401))
		return
	}
	newPassword := encryptPassword2(req.NewPassword)
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{"password": newPassword}},
		ReturnNew: true,
	}
	_, err := userCollection.Collection.Find(bson.M{"_id": userDB.ID}).Apply(change, &userDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userDB.Password = newPassword
	render.JSON(w, r, userDB)
}

func compare(password, hash string) bool {
	return encryptPassword2(password) == hash
}

func encryptPassword2(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
