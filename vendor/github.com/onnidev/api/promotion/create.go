package promotion

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreatePromotion TODO: NEEDS COMMENT INFO
func CreatePromotion(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.PromotionPostRequestKey).(types.PromotionPostRequest)
	if !ok {
		err := errors.New("sandro bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	partyP, err := repo.GetByID(req.PartyProductID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	horario := types.Timestamp(time.Now())
	timex := types.Timestamp(time.Time{})
	promotion := types.Promotion{
		ID:                     bson.NewObjectId(),
		QuantityPurchased:      0,
		StartDate:              &timex,
		EndDate:                &timex,
		MakePublic:             req.MakePublic,
		IgnoreActiveBatchRules: req.IgnoreActiveBatchRules,
		AvailableToFollowers:   req.AvailableToFollowers,
		Name:                   req.Name,
		Price:                  types.Price{Value: req.Price, CurrentIsoCode: "BRL"},
	}
	if req.StartDate != nil {
		promotion.StartDate = req.StartDate
	}
	if req.EndDate != nil {
		promotion.EndDate = req.EndDate
	}
	if req.QuantityTotal != nil {
		promotion.QuantityTotal = req.QuantityTotal
	}
	if req.QuantityPerCustomer != nil {
		promotion.QuantityPerCustomer = req.QuantityPerCustomer
	}
	if len(req.Promoters) != 0 {

		userRepo, ok := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
		if !ok {
			err := errors.New("asset bug")
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		ids := []bson.ObjectId{}
		for _, id := range req.Promoters {
			ids = append(ids, bson.ObjectIdHex(id))
		}
		users, err := userRepo.GetByIDS(ids)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		log.Println(users)
		promoters := []types.Promoter{}
		for _, user := range users {
			promoters = append(promoters, types.Promoter{
				PromoterID: user.ID,
				Deleted:    false,
			})
		}
		promotion.Promoters = &promoters

	}
	promotions := []types.Promotion{promotion}
	if partyP.PromotionalPrices != nil {
		promotions = append(promotions, *partyP.PromotionalPrices...)
	}
	partyProduct := types.PartyProduct{}
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":        &horario,
			"promotionalPrices": &promotions,
		}},
		ReturnNew: true,
	}
	log.Println("beofre creating promotion")
	_, err = repo.Collection.Find(bson.M{"_id": partyP.ID}).Apply(change, &partyProduct)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, req)
}
