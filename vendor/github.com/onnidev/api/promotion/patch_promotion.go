package promotion

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PatchPromotion carinho
func PatchPromotion(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(middlewares.PromotionPatchRequestKey).(types.PromotionPatchRequest)
	if !ok {
		err := errors.New("sandro bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "promotionID")
	partyP, promo, err := repo.GetPromotion(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(req)
	timex := types.Timestamp(time.Time{})
	next := types.Promotion{
		ID:                     bson.ObjectIdHex(id),
		Name:                   shared.OrBlank(promo.Name, req.Name),
		Price:                  promo.Price,
		MakePublic:             req.MakePublic,
		IgnoreActiveBatchRules: req.IgnoreActiveBatchRules,
		AvailableToFollowers:   req.AvailableToFollowers,
		QuantityPurchased:      promo.QuantityPurchased,
		// StartDate:              promo.StartDate,
		// EndDate:                promo.EndDate,
		StartDate: promo.StartDate,
		EndDate:   promo.EndDate,
		// QuantityTotal          *int          `json:"quantityTotal" bson:"quantityTotal"`
		// QuantityPerCustomer    *int          `json:"quantityPerCustomer" bson:"quantityPerCustomer"`
		QuantityTotal:       req.QuantityTotal,
		QuantityPerCustomer: req.QuantityPerCustomer,
		// Promoters              *[]Promoter   `json:"promoters" bson:"promoters"`
		Promoters: promo.Promoters,
	}
	if req.StartDate != nil && req.EndDate != nil {
		next.StartDate = req.StartDate
		next.EndDate = req.EndDate
	}
	if req.StartDate == nil && req.EndDate == nil {
		next.StartDate = &timex
		next.EndDate = &timex
	}

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
	promoters := []types.Promoter{}
	for _, user := range users {
		promoter := types.Promoter{
			PromoterID: user.ID,
			Deleted:    false,
		}
		promoters = append(promoters, promoter)
	}
	next.Promoters = &promoters
	result := []types.Promotion{}
	if partyP.PromotionalPrices != nil {
		for _, promotion := range *partyP.PromotionalPrices {
			if promotion.ID.Hex() == promo.ID.Hex() {
				result = append(result, next)
				continue
			}
			result = append(result, promotion)
		}

	}
	horario := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":        &horario,
				"promotionalPrices": &result,
			}},
		ReturnNew: true,
	}
	newPartyP := types.PartyProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": partyP.ID}).Apply(change, &newPartyP)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, next)
}
