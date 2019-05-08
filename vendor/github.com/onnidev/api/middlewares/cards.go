package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"unicode"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// CardRequestKey is the shit
var CardRequestKey key = "cards-request-repo"

// ReadCreateCardRequestFromBody skjdfn
func ReadCreateCardRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.CardRequest
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		if !(digits(card.Number) == 16 || digits(card.Number) == 15) {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if digits(card.ExpirationDate) != 4 {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if digits(card.Cvv) == 3 || digits(card.Cvv) == 4 {
			ctx := context.WithValue(r.Context(), CardRequestKey, card)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		http.Error(w, "bad.request", http.StatusBadRequest)
	})
}
func digits(number string) int {
	i := 0
	for _, ch := range number {
		if unicode.IsDigit(ch) {
			i++
		}
	}
	return i
}

// CardUpdateKey is the shit
var CardUpdateKey key = "cards-update-repo"

// ReadUpdateCardRequestFromBody skjdfn
func ReadUpdateCardRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var card types.CardPatch
		err := decoder.Decode(&card)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(card, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), CardUpdateKey, card)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CardsRepoKey is the shit
var CardsRepoKey key = "cards-repo"

// AttachCardsCollection skjdfn
func AttachCardsCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		collection, err := interfaces.NewCardsCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)

			return
		}

		ctx := context.WithValue(r.Context(), CardsRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PagarmeResponseRepoKey is the shit
var PagarmeResponseRepoKey key = "pagarme-cards-request-repo"

// PagarMeCreateCardHash skjdfn
func PagarMeCreateCardHash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cardFromBody := r.Context().Value(CardRequestKey).(types.CardRequest)
		url := "https://api.pagar.me/1/cards"
		card := types.PagarmeCardRequest{
			APIKey:         viper.GetString("PAGARME"),
			Number:         cardFromBody.Number,
			HolderName:     cardFromBody.HolderName,
			Cvv:            cardFromBody.Cvv,
			ExpirationDate: cardFromBody.ExpirationDate,
		}
		j, _ := json.MarshalIndent(card, "", "    ")
		if viper.GetBool("verbose") {
			log.Println("Request sent to PagarMe")
			log.Println(string(j))
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer resp.Body.Close()
		var response types.PagarmeCardResponse
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(response, "", "    ")
			log.Println("Request received from PagarMe")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), PagarmeResponseRepoKey, response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
