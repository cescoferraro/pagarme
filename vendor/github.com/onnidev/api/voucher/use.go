package voucher

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/ws"
	"github.com/pressly/chi/render"
)

// Use sdkjf
func Use(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "voucherId")
	userClub, ok := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("bug assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	voucher, err := onni.VoucherUseComplete(r.Context(), id, true, onni.AllKindsConstrain, userClub)
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
