package card_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/card"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
	// . "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	viper.SetDefault("verbose", true)
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	customer.Routes(r)
	card.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	var cardResponse types.Card
	headers := map[string]string{
		"JWT_TOKEN":    loginHelper.Token,
		"Content-Type": "application/json"}
	t.Run("Create Card", func(t *testing.T) {
		card := types.CardRequest{
			Number:         "3411206056921859",
			HolderName:     "Francesco Ferraro",
			Cvv:            "696",
			ExpirationDate: "1025"}
		j, _ := json.MarshalIndent(card, "", "    ")
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/card",
			Headers: headers,
			Body:    bytes.NewBuffer(j)}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
		json.Unmarshal(body, &cardResponse)
		log.Println(cardResponse)
		assert.True(t, cardResponse.Default)
	})
	t.Run("Get All Cards from customer", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/card",
			Headers: headers,
		}
		_, body := infra.NewTestRequest(t, ajax)
		var allCards []types.Card
		json.Unmarshal(body, &allCards)
		log.Println(allCards)
		var created bool
		for _, card := range allCards {
			if card.ID.Hex() == cardResponse.ID.Hex() {
				created = true
			}
		}
		assert.True(t, created)
	})
	t.Run("Update a Card with specific ID", func(t *testing.T) {
		log.Println(cardResponse.ID.Hex())
		card := types.CardPatch{Default: true}
		j, _ := json.MarshalIndent(card, "", "    ")
		ajax := infra.Ajax{
			Method:  "PATCH",
			Path:    infra.TestServer.URL + "/card/" + cardResponse.ID.Hex(),
			Headers: headers,
			Body:    bytes.NewBuffer(j),
		}
		_, body := infra.NewTestRequest(t, ajax)

		var updatedCardResponse types.Card
		err := json.Unmarshal(body, &updatedCardResponse)
		assert.NoError(t, err)
		assert.True(t, updatedCardResponse.Default)
	})
	t.Run("Get a Card with specific ID", func(t *testing.T) {
		_, body := infra.NewTestRequest(t, infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/card/" + cardResponse.ID.Hex(),
			Headers: headers,
		})
		var readCardResponse types.Card
		err := json.Unmarshal(body, &readCardResponse)
		assert.NoError(t, err)
		assert.Equal(t, readCardResponse.ID.Hex(), cardResponse.ID.Hex())
	})
	t.Run("Delete a Card with specific ID", func(t *testing.T) {
		_, body := infra.NewTestRequest(t,
			infra.Ajax{
				Method:  "DELETE",
				Path:    infra.TestServer.URL + "/card/" + cardResponse.ID.Hex(),
				Headers: headers,
			})
		var deleteCardResponse string
		err := json.Unmarshal(body, &deleteCardResponse)
		assert.NoError(t, err)
		assert.Equal(t, deleteCardResponse, cardResponse.ID.Hex())
	})
	// t.Run("Fail to Create Card", func(t *testing.T) {
	// 	card := types.MistakeCardRequest{
	// 		Number:         "dsf",
	// 		HolderName:     fake.FirstName(),
	// 		Cvv:            "345",
	// 		ExpirationDate: "1125"}
	// 	j, _ := json.MarshalIndent(card, "", "    ")
	// 	ajax := infra.Ajax{
	// 		Method:  "POST",
	// 		Path:    infra.TestServer.URL + "/card",
	// 		Headers: headers,
	// 		Body:    bytes.NewBuffer(j)}
	// 	resp, _ := infra.NewTestRequest(t, ajax)
	// 	assert.Equal(t, resp.StatusCode, 400)
	// })
}
