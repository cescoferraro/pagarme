package voucher

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/ws"
	"github.com/pressly/chi/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Undo sdkjf
func Undo(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "voucherId")
	voucher, err := vouchersCollection.GetSimpleByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	log.Println(voucher.Status)
	if voucher.Status != "USED" {
		http.Error(w,
			errors.New("only used vouchers can be unread").Error(),
			480)
		return
	}
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"status":                 "AVAILABLE",
				"updateDate":             time.Now(),
				"voucherUseDate":         nil,
				"voucherUseUserClubId":   nil,
				"voucherUseUserClubName": nil,
			}},
		ReturnNew: true,
	}
	var result types.Voucher
	_, err = vouchersCollection.
		Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).
		Apply(change, &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	partyID := voucher.PartyID.Hex()
	msg, _ := json.Marshal(ws.WebSocketMsg{
		Type: "VOUCHERS_CHANGES",
		Data: struct {
			PartyID string `json:"partyID"`
		}{
			PartyID: partyID,
		}},
	)
	ws.Publish(shared.RedisPrefixer("dashboard"), string(msg))
	ws.Publish(shared.RedisPrefixer("staff"), string(msg))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
