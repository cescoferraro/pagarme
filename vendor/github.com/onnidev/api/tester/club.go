package tester

import (
	"bytes"
	"encoding/json"
	"image"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// ListAllParties dkjns
func ListAllParties(t *testing.T, helper *types.UserClubLoginResponse) []types.Party {
	headers := map[string]string{
		"JWT_TOKEN":    helper.Token,
		"Content-Type": "application/json"}
	var party []types.Party
	t.Run("List Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/party",
			Headers: headers}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &party)
		assert.NoError(t, err)
	})
	return party
}

// ListAllClubs dkjns
func ListAllClubs(t *testing.T, helper *types.UserClubLoginResponse) []types.Club {
	var clubs []types.Club
	t.Run("List Clubs", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/club",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &clubs)
		assert.NoError(t, err)
	})
	return clubs
}

// RandomClubID TODO: NEEDS COMMENT INFO
func RandomClubID(clubs []types.Club) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return clubs[rand.Intn(len(clubs))].ID.Hex()
}

// ReadClubImage dkjns
func ReadClubImage(t *testing.T, helper *types.UserClubLoginResponse, id string) image.Image {
	var cimage image.Image
	t.Run("Read Club Image", func(t *testing.T) {
		// var err error
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/club/image/" + id,
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		_ = bytes.NewReader(body)
		log.Println("")
		// cimage, _, err = image.Decode(buf)
		// assert.NoError(t, err)
	})
	return cimage
}

// ReadClubDashboard dkjns
func ReadClubDashboard(t *testing.T, helper *types.UserClubLoginResponse, id string) types.ClubDashboard {
	var club types.ClubDashboard
	t.Run("Read ClubDashboard", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/club/dashboard/" + id,
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &club)
		assert.NoError(t, err)
	})
	return club
}

// ReadClub dkjns
func ReadClub(t *testing.T, helper *types.UserClubLoginResponse, id string) types.AppClub {
	var club types.AppClub
	t.Run("Read Club", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/club/" + id,
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &club)
		assert.NoError(t, err)
	})
	return club
}

// ListAllClubsByToken sdkjfn
func ListAllClubsByToken(t *testing.T, token string) []types.Club {
	htat := types.UserClubLoginResponse{Token: token}
	return ListAllClubs(t, &htat)
}
