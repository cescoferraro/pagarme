package voucher

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Read sdkjf
func Read(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "voucherId")
	// TODO: Verificar se o leitor tem permissão para ler esse voucher
	// if err != nil {
	// 	http.Error(w, err.Error(), http.Status)
	// 	return
	// }
	// http.StatusForbidden
	voucher, err := vouchersCollection.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, voucher)
}

// SoftRead sdkjf
func SoftRead(w http.ResponseWriter, r *http.Request) {
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "voucherId")
	voucher, err := vouchersCollection.GetSimpleByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, voucher)
}
