package promotion

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ListPartyPromotions carinho
func ListPartyPromotions(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "partyId")
	allUser, err := productsCollection.GetByPartyID(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	partiesCollection, ok := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	party, err := partiesCollection.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	repo, ok := r.Context().
		Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	promotions := []SoftPromotion{}
	for _, product := range allUser {
		if product.PromotionalPrices != nil {
			for _, promotion := range *product.PromotionalPrices {
				customers, err := repo.PromotionalCustomersAttachedToPromotion(promotion.ID.Hex())
				if err != nil {
					shared.MakeONNiError(w, r, 400, err)
					return
				}
				timex := types.Timestamp(time.Time{})
				promo := SoftPromotion{
					PromotionID:       promotion.ID.Hex(),
					MakePublic:        promotion.MakePublic,
					PartyName:         party.Name,
					ProductName:       product.Name,
					PromotionName:     promotion.Name,
					PromotionPrice:    promotion.Price.Value,
					QuantitySold:      int(promotion.QuantityPurchased),
					PromotionStart:    &timex,
					PromotionEnd:      &timex,
					QuantityCustomers: len(customers),
				}
				if promotion.StartDate != nil {
					promo.PromotionStart = promotion.StartDate
				}
				if promotion.EndDate != nil {
					promo.PromotionEnd = promotion.EndDate
				}
				if promotion.Promoters != nil {
					promo.QuantityPromoters = len(*promotion.Promoters)
				}
				if promotion.QuantityTotal != nil {
					next := int(*promotion.QuantityTotal)
					promo.QuantityTotal = &next
				}
				if userClub.Profile == "PROMOTER" {
					if promotion.Promoters != nil {
						for _, promoter := range *promotion.Promoters {
							if userClub.ID.Hex() == promoter.PromoterID.Hex() {
								promotions = append(promotions, promo)
							}
							continue
						}
					}
					continue
				}
				promotions = append(promotions, promo)
			}
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, promotions)
}

// SoftPromotion TODO: NEEDS COMMENT INFO
type SoftPromotion struct {
	PromotionID       string  `json:"promotionId" bson:"promotionId"`
	MakePublic        bool    `json:"makePublic" bson:"makePublic"`
	PartyName         string  `json:"partyName" bson:"partyName"`
	ProductName       string  `json:"productName" bson:"productName"`
	PromotionName     string  `json:"promotionName" bson:"promotionName"`
	PromotionPrice    float64 `json:"promotionPrice" bson:"promotionPrice"`
	QuantitySold      int     `json:"quantitySold" bson:"quantitySold"`
	QuantityPromoters int     `json:"quantityPromoters" bson:"quantityPromoters"`

	PromotionEnd   *types.Timestamp `json:"promotionEnd" bson:"promotionId"`
	PromotionStart *types.Timestamp `json:"promotionStart" bson:"promotionStart"`

	QuantityCustomers int  `json:"quantityCustomers" bson:"quantityCustomers"`
	QuantityTotal     *int `json:"quantityTotal" bson:"quantityTotal"`
}
