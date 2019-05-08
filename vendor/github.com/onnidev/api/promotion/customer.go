package promotion

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// CustomersSummary carinho
func CustomersSummary(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "promotionID")
	_, promotion, err := productsCollection.GetPromotion(id)
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	repo, ok := r.Context().Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	customerPromotions, err := repo.PromotionalCustomersAttachedToPromotion(promotion.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchersCollection, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	userClub := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	minicustomers := []MiniCustomerPromotion{}
	for _, customerPromotion := range customerPromotions {
		vouchers, err := vouchersCollection.PromotionPurchasedbyCustomer(promotion.ID.Hex(), customerPromotion.CustomerID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		promo := MiniCustomerPromotion{
			CustomerName:      customerPromotion.CustomerName,
			QuantityPurchased: len(vouchers),
			IndicatedDate:     customerPromotion.CreationDate,
			PromoterName:      customerPromotion.PromoterName,
		}
		if len(vouchers) != 0 {
			promo.PurchasedDate = vouchers[len(vouchers)-1].CreationDate
		}
		if userClub.Profile == "PROMOTER" {
			log.Println("é promoter")
			if userClub.ID.Hex() == customerPromotion.PromoterID.Hex() {
				minicustomers = append(minicustomers, promo)
			}
			continue
		}
		minicustomers = append(minicustomers, promo)
	}
	// for _, hehe := range minicustomers {

	// }
	render.Status(r, http.StatusOK)
	render.JSON(w, r, PromotionCustomerSoft{
		Customers: minicustomers,
		Promotion: MiniPromotion{
			PromotionID:            id,
			MakePublic:             promotion.MakePublic,
			IgnoreActiveBatchRules: promotion.IgnoreActiveBatchRules,
			AvailableToFollowers:   promotion.AvailableToFollowers,
		},
	})
}

// MiniCustomerPromotion TODO: NEEDS COMMENT INFO
type MiniCustomerPromotion struct {
	CustomerName      string           `json:"customerName"`
	PromoterName      string           `json:"promoterName"`
	QuantityPurchased int              `json:"quantityPurchased"`
	IndicatedDate     *types.Timestamp `json:"indicatedDate"`
	PurchasedDate     *types.Timestamp `json:"purchasedDate"`
}

// PromotionCustomerSoft TODO: NEEDS COMMENT INFO
type PromotionCustomerSoft struct {
	Customers []MiniCustomerPromotion `json:"customers"`
	Promotion MiniPromotion           `json:"promotion"`
}

// MiniPromotion TODO: NEEDS COMMENT INFO
type MiniPromotion struct {
	AvailableToFollowers   bool
	IgnoreActiveBatchRules bool
	MakePublic             bool
	PromotionID            string
}
