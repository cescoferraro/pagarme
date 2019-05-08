package promotion

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ReadPromotion carinho
func ReadPromotion(w http.ResponseWriter, r *http.Request) {
	productsCollection := r.Context().Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	id := chi.URLParam(r, "promotionID")
	partyP, promotion, err := productsCollection.GetPromotion(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	party, err := onni.Party(r.Context(), partyP.PartyID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	timex := types.Timestamp(time.Time{})
	softpromotion := types.SoftPromotion{
		ID:                     promotion.ID,
		Name:                   promotion.Name,
		StartDate:              &timex,
		EndDate:                &timex,
		PartyName:              party.Name,
		Price:                  promotion.Price,
		ProductName:            partyP.Name,
		QuantityTotal:          promotion.QuantityTotal,
		QuantityPurchased:      promotion.QuantityPurchased,
		QuantityPerCustomer:    promotion.QuantityPerCustomer,
		MakePublic:             promotion.MakePublic,
		IgnoreActiveBatchRules: promotion.IgnoreActiveBatchRules,
		AvailableToFollowers:   promotion.AvailableToFollowers,
		// Promoters:              promotion.Promoters,
	}
	if promotion.StartDate != nil {
		softpromotion.StartDate = promotion.StartDate
	}
	if promotion.EndDate != nil {
		softpromotion.EndDate = promotion.EndDate
	}
	if promotion.Promoters != nil {
		ids := []string{}
		for _, user := range *promotion.Promoters {
			ids = append(ids, user.PromoterID.Hex())
		}
		repo, ok := r.Context().Value(middlewares.UserClubRepoKey).(interfaces.UserClubRepo)
		if !ok {
			err := errors.New("asset bug")
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		log.Println(ids)
		users, err := repo.ListPromotersByClub(party.ClubID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		promoters := []types.SoftPromoter{}
		for _, user := range users {
			promoter := types.SoftPromoter{
				PromoterName: user.Name,
				PromoterID:   user.ID.Hex(),
			}
			if shared.Contains(ids, user.ID.Hex()) {
				promoter.Selected = true
			}
			promoters = append(promoters, promoter)
		}
		softpromotion.Promoters = &promoters
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, softpromotion)
}
