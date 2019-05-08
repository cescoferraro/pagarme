package tester

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// LogInGmailCescoUserClub sdfjkn
func LogInGmailCescoUserClub(t *testing.T) *types.UserClubLoginResponse {
	return LoginTestingHelperFactory(t, GmailCescoUserClubLogin)
}

// LogInCescoUserClub sdfjkn
func LogInCescoUserClub(t *testing.T) *types.UserClubLoginResponse {
	return LoginTestingHelperFactory(t, CescoUserClubLogin)
}

// GmailCescoUserClubLogin sdkjfn
var GmailCescoUserClubLogin = types.LoginRequest{
	Email:    "francescoaferraro@gmail.com",
	Password: "111008"}

// CescoUserClubLogin sdkjfn
var CescoUserClubLogin = types.LoginRequest{
	Email:    "cescoferraro@onni.live",
	Password: "descriptor8"}

// LoginTestingHelperFactory TODO: NEEDS COMMENT INFO
func LoginTestingHelperFactory(t *testing.T, user types.LoginRequest) *types.UserClubLoginResponse {
	var User types.UserClubLoginResponse
	function := func(t *testing.T) {
		bolB, _ := json.Marshal(user)
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/userClub/login",
			Body:   bytes.NewBuffer(bolB),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		response, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &User)
		assert.Equal(t, response.StatusCode, 200)
	}
	t.Run("Login to the existing ClubApp API", function)
	return &User
}

// LeitorLoginTestingHelperFactory TODO: NEEDS COMMENT INFO
func LeitorLoginTestingHelperFactory(t *testing.T, user types.LoginRequest) *types.UserClubLoginResponse {
	var User types.UserClubLoginResponse
	function := func(t *testing.T) {
		bolB, _ := json.Marshal(user)
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/userClub/login/leitor",
			Body:   bytes.NewBuffer(bolB),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		_, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &User)
		log.Println(string(body))
	}
	t.Run("Login Leitor QR COde Staff", function)
	return &User
}

// GetPartyProducts TODO: NEEDS COMMENT INFO
func GetClubParties(t *testing.T, helper *types.UserClubLoginResponse) *[]types.Party {
	var User []types.Party
	club := helper.Clubs[rand.Intn(len(helper.Clubs))]
	log.Println(helper.Clubs)
	log.Println(club)
	function := func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/" + club.ID.Hex(),
			Headers: map[string]string{
				"Content-Type": "application/json",
				"JWT_TOKEN":    helper.Token},
		}
		_, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &User)
	}
	t.Run("Get PartyProducts", function)
	return &User
}

// GetPartyProducts TODO: NEEDS COMMENT INFO
func GetPartyProducts(t *testing.T, helper *types.UserClubLoginResponse, id string) *[]types.PartyProduct {
	var User []types.PartyProduct
	function := func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/partyProduct/" + id,
			Headers: map[string]string{
				"Content-Type": "application/json",
				"JWT_TOKEN":    helper.Token},
		}
		_, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &User)
	}
	t.Run("Get PartyProducts", function)
	return &User
}
