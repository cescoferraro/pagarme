package userClub

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
)

// ChangePassword TODO: NEEDS COMMENT INFO
func ChangePassword(helper *types.UserClubLoginResponse, req types.ChangePasswordRequest) func(t *testing.T) {
	return func(t *testing.T) {
		bolB, _ := json.Marshal(req)
		var user types.UserClub
		log.Println("trying to change a password using the token ", helper.Token)
		ajax := infra.Ajax{
			Method:  "PATCH",
			Path:    infra.TestServer.URL + "/userClub/password",
			Body:    bytes.NewBuffer(bolB),
			Headers: map[string]string{"JWT_TOKEN": helper.Token},
		}
		reponse, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &user)
		log.Println(reponse.Status)
		log.Println(string(body))
	}
}
