package voucher

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/ws"
	"github.com/pressly/chi/render"
)

// UseDash sdkjf
func UseDash(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "voucherId")
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	voucher, err := onni.VoucherUseComplete(r.Context(), id, false, onni.AllKindsConstrain, userClub)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	partyID := voucher.PartyID.Hex()
	result, err := json.Marshal(ws.WebSocketMsg{
		Type: "VOUCHERS_CHANGES",
		Data: struct {
			PartyID string `json:"partyID"`
		}{
			PartyID: partyID,
		}},
	)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	ws.Publish(shared.RedisPrefixer("dashboard"), string(result))
	ws.Publish(shared.RedisPrefixer("staff"), string(result))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, voucher)
}

// UseAndroid sdkjf
func UseAndroid(w http.ResponseWriter, r *http.Request) {
	log.Println("contenstando")
	id := chi.URLParam(r, "voucherId")
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("bug assert")
		androidResponse(w, r, 520, err.Error(), types.Voucher{})
		return
	}
	voucher, err := onni.VoucherUse(r.Context(), id, true, onni.AllKindsConstrain, userClub)
	if err != nil {
		androidResponse(w, r, 520, err.Error(), voucher)
		return
	}
	partyID := voucher.PartyID.Hex()
	result, err := json.Marshal(ws.WebSocketMsg{
		Type: "VOUCHERS_CHANGES",
		Data: struct {
			PartyID string `json:"partyID"`
		}{
			PartyID: partyID,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		androidResponse(w, r, 520, err.Error(), voucher)
		return
	}
	ws.Publish(shared.RedisPrefixer("dashboard"), string(result))
	ws.Publish(shared.RedisPrefixer("staff"), string(result))
	androidResponse(w, r, 200, "sucess", voucher)
}

func androidResponse(w http.ResponseWriter, r *http.Request, code int, reason string, voucher types.Voucher) {
	render.Status(r, code)
	render.JSON(w, r, AndroidResponse{
		Code:   code,
		Reason: reason,
		Voucher: MiniVoucher{
			CustomerName: voucher.CustomerName,
			ProductName:  voucher.Product.Name,
			PartyName:    voucher.PartyName,
			Type:         voucher.Type,
		},
	})
}

// AndroidResponse TODO: NEEDS COMMENT INFO
type AndroidResponse struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Voucher MiniVoucher `json:"voucher"`
}

// MiniVoucher TODO: NEEDS COMMENT INFO
type MiniVoucher struct {
	CustomerName string `json:"customerName"`
	Type         string `json:"type"`
	PartyName    string `json:"partyName"`
	ProductName  string `json:"productName"`
}
