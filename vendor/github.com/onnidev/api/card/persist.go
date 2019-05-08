package card

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PersistToDB is a comented function
// swagger:route POST /cards app postCard
//
// Create a card for a customer.
//
// This will enables you to create a card.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Security:
//       X-AUTH-APPLICATION-TOKEN:
//       X-AUTH-ID:
//       X-CLIENT-ID:
//
//
//     Responses:
//       200: dbCard
//       404: error
func PersistToDB(w http.ResponseWriter, r *http.Request) {
	response := r.Context().Value(middlewares.PagarmeResponseRepoKey).(types.PagarmeCardResponse)
	cardsRepo := r.Context().Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	customer := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	err := makeCustomerCardsDull(cardsRepo, customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	id := bson.NewObjectId()
	horario := types.Timestamp(time.Now())
	mongocard := types.Card{
		ID:           id,
		CardToken:    response.ID,
		CreationDate: &horario,
		UpdateDate:   &horario,
		CustomerID:   bson.ObjectIdHex(customer.ID.Hex()),
		Last4:        response.LastDigits,
		Brand:        strings.ToUpper(response.Brand),
		GatewayType:  "PAGARME",
		Default:      true,
	}
	afterDB, err := cardsRepo.Create(mongocard)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(afterDB, "", "    ")
		log.Println("Card created on MongoDB")
		log.Println(string(j))
	}
	card, err := cardsRepo.GetByID(id.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.JSON(w, r, card)
	return
}

func makeCustomerCardsDull(cardsRepo interfaces.CardsRepo, customer types.Customer) error {
	cards, err := cardsRepo.GetAllByCustomerID(customer.ID.Hex())
	if err != nil {
		return err
	}
	for _, card := range cards {
		if card.Default {
			result := types.Card{}
			change := mgo.Change{
				Update:    bson.M{"$set": bson.M{"defaultCard": false}},
				ReturnNew: true,
			}
			_, err := cardsRepo.
				Collection.
				Find(bson.M{"_id": card.ID}).Apply(change, &result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
