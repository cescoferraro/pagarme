package party_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/party"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	party.Routes(r)
	userClub.Routes(r)
	customer.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	var allParties []types.Party
	t.Run("List Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party",
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allParties)
		assert.NoError(t, err)
		for _, party := range allParties {
			club := party.Club
			if club != nil {
			}
		}
	})
	var allUserClubOwnParties []types.Party
	t.Run("List UserClub Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/userClub",
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allUserClubOwnParties)
		assert.NoError(t, err)
	})
	t.Run("Read Random Party", func(t *testing.T) {
		id := allParties[rand.Intn(len(allParties))].ID.Hex()
		// id := "58f67b90cc922d1471b52ad5"
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/" + id,
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		var party types.Party
		err := json.Unmarshal(body, &party)
		party.Description = ""
		assert.NoError(t, err)
	})
	t.Run("Get Party Image", func(t *testing.T) {
		id := allParties[rand.Intn(len(allParties))].ID.Hex()
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/image/" + id,
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		buf := bytes.NewReader(body)
		_, _, err := image.Decode(buf)
		assert.NoError(t, err)
	})
	var allUserClubParties []types.Party
	clubID := allParties[rand.Intn(len(allParties))].ClubID
	t.Run("UserClub get Club Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/club/" + clubID.Hex(),
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allUserClubParties)
		assert.NoError(t, err)
	})
	cesco := tester.LogInCescoCustomer(t)
	var allCustomerParties []types.AppParty
	t.Run("Compare ", func(t *testing.T) {
		assert.True(t, len(allUserClubParties) >= len(allCustomerParties))
	})
	t.Run("Filter a List of Parties", func(t *testing.T) {
		from := types.Timestamp(time.Now())
		till := types.Timestamp(time.Now().Add(1000000 * time.Minute))
		filter := types.PartyFilter{
			From: &from,
			Till: &till,
			Long: -30.0288606,
			Lat:  -51.2026095,
		}
		filters, _ := json.MarshalIndent(filter, "", "    ")
		ajax := infra.Ajax{
			Method: "POST",
			Body:   bytes.NewBuffer(filters),
			Path:   infra.TestServer.URL + "/party",
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json",
			}}
		var allFilterParties []types.Party
		r, body := infra.NewTestRequest(t, ajax)
		if r.StatusCode >= 400 && r.StatusCode < 600 {
		}
		err := json.Unmarshal(body, &allFilterParties)
		assert.NoError(t, err)
	})
	t.Run("Customer get Club Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/party/club/" + clubID.Hex(),
			Headers: map[string]string{
				"JWT_TOKEN":    cesco.Token,
				"Content-Type": "application/json"}}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &allCustomerParties)
		assert.NoError(t, err)
		for _, party := range allCustomerParties {
			if party.Status != "ACTIVE" {
				err := errors.New("customer should only see ACTIVE parties")
				assert.NoError(t, err)
			}
		}
	})
	t.Run("Party AntiTheft", func(t *testing.T) {
		model := types.AntiTheftModel{
			CardL1:   4,
			CardL2:   5,
			CardL3:   6,
			DrinkL1:  float64(2.0),
			DrinkL2:  float64(2.5),
			DrinkL3:  float64(3.0),
			TicketL1: float64(2.0),
			TicketL2: float64(2.5),
			TicketL3: float64(3.0),
		}
		modelByt, _ := json.MarshalIndent(model, "", "    ")
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/party/antitheft/5b62295cc5f1fa00012f5482",
			Body:   bytes.NewBuffer(modelByt),
			Headers: map[string]string{
				"JWT_TOKEN":    loginHelper.Token,
				"Content-Type": "application/json"}}
		_, _ = infra.NewTestRequest(t, ajax)
		// list := []types.AntiTheftResult{}
		// err := json.Unmarshal(body, &list)
		// assert.NoError(t, err)
		// byt, err := json.MarshalIndent(list, "    ", "")
		// log.Println(string(body))
	})
}
