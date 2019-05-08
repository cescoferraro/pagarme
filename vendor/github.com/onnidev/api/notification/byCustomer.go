package notification

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/bradfitz/slice"
	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// ByCustomer sdkfjn
func ByCustomer(w http.ResponseWriter, r *http.Request) {
	notificationRepo := r.Context().Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	notifications, err := notificationRepo.GetByCustomer(chi.URLParam(r, "customerID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	voucherRepo, ok := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	customer, ok := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	vouchers, err := voucherRepo.GetActualVoucherByCustomer(customer.ID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(len(vouchers))
	parties := []string{}
	for _, vouch := range vouchers {
		if !shared.Contains(parties, vouch.Party.ID.Hex()) {
			parties = append(parties, vouch.Party.ID.Hex())
		}
	}
	log.Println(len(parties))
	for _, id := range parties {
		partyNot, err := notificationRepo.GetByParty(id)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		log.Println("party", len(partyNot))
		notifications = append(notifications, partyNot...)
	}

	log.Println(len(notifications))
	filtered := []types.Notification{}
	for _, notificate := range notifications {
		log.Println(notificate.Type)
		if notificate.Party != nil {
			log.Println("has party s")
			if notificate.Party.Status == "ACTIVE" && !notificate.Party.EndDate.Time().Before(time.Now()) {
				filtered = append(filtered, notificate)
			}
		}
	}
	log.Println(len(filtered))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, SortNotifications(filtered))
}

// SortNotifications TODO: NEEDS COMMENT INFO
func SortNotifications(parties []types.Notification) []types.Notification {
	slice.Sort(parties[:], func(i, j int) bool {
		return (parties[j].CreationDate.Time().Before(parties[i].CreationDate.Time()))
	})
	return parties
}
