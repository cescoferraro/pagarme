package buy_test

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/buy"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	// . "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	customer.Routes(r)
	buy.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoCustomer(t)
	t.Run("Buy stuff", func(t *testing.T) {
		// j, _ := json.MarshalIndent(drink(loginHelper.ID), "", "    ")
		// j, _ := json.MarshalIndent(ticketNASF(loginHelper.ID), "", "    ")
		// j, _ := json.MarshalIndent(ticketNASF(loginHelper.ID), "", "    ")
		j, _ := json.MarshalIndent(carpevita(loginHelper.ID), "", "    ")
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/buy",
			Headers: loginHelper.Headers(),
			Body:    bytes.NewBuffer(j)}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}

func carpevita(id string) types.BuyPost {
	return types.BuyPost{
		CustomerID: id,
		PartyID:    "5b7b587814319f000179daa6",
		ClubID:     "5b6b61a3e405ba0001900841",
		Products: []types.BuyPartyProductsItemRequest{
			{
				PartyProductID: "5b7c884114319f000179e0b9",
				PromotionID:    "",
				Quantity:       int64(1),
			},
		},
		Log: types.Log{
			AppVersion: "4",
			DeviceID:   "kj231ne3",
			DeviceSO:   "129i8wu21w",
			Latitude:   float64(0),
			Longitude:  float64(0),
		},
	}
}

func all(id string) types.BuyPost {
	return types.BuyPost{
		CustomerID: id,
		PartyID:    "5b6a0f07f2ed440001a45aaf",
		ClubID:     "592ca3b3cc922d18bdae4cd8",
		Products: []types.BuyPartyProductsItemRequest{
			{
				PartyProductID: "5b6a0f07f2ed440001a45ab2",
				PromotionID:    "",
				Quantity:       int64(1),
			},
		},
		Log: types.Log{
			AppVersion: "4",
			DeviceID:   "kj231ne3",
			DeviceSO:   "129i8wu21w",
			Latitude:   float64(0),
			Longitude:  float64(0),
		},
	}
}

func drink(id string) types.BuyPost {
	return types.BuyPost{
		CustomerID: id,
		PartyID:    "5b6a0f07f2ed440001a45aaf",
		ClubID:     "592ca3b3cc922d18bdae4cd8",
		Products: []types.BuyPartyProductsItemRequest{
			{
				PartyProductID: "5b6a0f07f2ed440001a45ab2",
				PromotionID:    "5b6a12eef2ed440001a45ac8",
				Quantity:       int64(1),
			},
			{
				PartyProductID: "5b6a0f07f2ed440001a45ab2",
				PromotionID:    "",
				Quantity:       int64(1),
			},
		},
		Log: types.Log{
			AppVersion: "4",
			DeviceID:   "kj231ne3",
			DeviceSO:   "ANDROID",
			Latitude:   float64(0),
			Longitude:  float64(0),
		},
	}
}
func ticketNASF(id string) types.BuyPost {
	return types.BuyPost{
		CustomerID: id,
		PartyID:    "5b6a0f07f2ed440001a45aaf",
		ClubID:     "592ca3b3cc922d18bdae4cd8",
		Products: []types.BuyPartyProductsItemRequest{
			{
				PartyProductID: "5b6a0f07f2ed440001a45ab5",
				PromotionID:    "5b6a12dbf2ed440001a45ac6",
				Quantity:       int64(1),
			},
			{
				PartyProductID: "5b6a0f07f2ed440001a45ab5",
				PromotionID:    "",
				Quantity:       int64(1),
			},
		},
		Log: types.Log{
			AppVersion: "4",
			DeviceID:   "kj231ne3",
			DeviceSO:   "ANDROID",
			Latitude:   float64(0),
			Longitude:  float64(0),
		},
	}
}

func ticketASF(id string) types.BuyPost {
	return types.BuyPost{
		CustomerID: id,
		PartyID:    "5b6a0f59f2ed440001a45ab8",
		ClubID:     "592ca3b3cc922d18bdae4cd8",
		Products: []types.BuyPartyProductsItemRequest{
			// {
			// 	PartyProductID: "5b6a0f59f2ed440001a45abb",
			// 	PromotionID:    "5b6a1253f2ed440001a45ac3",
			// 	Quantity:       int64(1),
			// },
			{
				PartyProductID: "5b6a0f59f2ed440001a45abb",
				PromotionID:    "",
				Quantity:       int64(1),
			},
		},
		Log: types.Log{
			AppVersion: "4",
			DeviceID:   "kj231ne3",
			DeviceSO:   "ANDROID",
			Latitude:   float64(0),
			Longitude:  float64(0),
		},
	}
}
