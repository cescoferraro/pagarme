package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// VouchersRepoKey is a context key
var VouchersRepoKey VoucherKey = "vouchers-repo"

// VoucherKey is a context key
type VoucherKey string

// AttachVoucherCollection mongo wise
func AttachVoucherCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := infra.Cloner()
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		collection, err := interfaces.NewVoucherCollection(db)
		if err != nil {
			defer db.Session.Close()
			shared.MakeONNiError(w, r, 400, err)
			return
		}

		ctx := context.WithValue(r.Context(), VouchersRepoKey, collection)
		go func() {
			select {
			case <-ctx.Done():
				collection.Session.Close()
			}
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PartyVouchers  is the shit
var PartyVouchers key = "report-vouchers"

// GetVouchers skjdfn
func GetVouchers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := r.Context().Value(VouchersRepoKey).(interfaces.VouchersRepo)
		id := chi.URLParam(r, "id")
		vouchers, err := repo.GetByParty(id)
		if err != nil {
			render.Status(r, http.StatusExpectationFailed)
			render.JSON(w, r, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), PartyVouchers, vouchers)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReadUseConstrains  is the shit
var ReadUseConstrains key = "vouchers-constrains"

// ReadVoucherUsingConstrains skjdfn
func ReadVoucherUsingConstrains(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var login types.VoucherUseConstrain
		err := decoder.Decode(&login)
		if err != nil {
			http.Error(w, "bad.request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if viper.GetBool("verbose") {
			j, _ := json.MarshalIndent(login, "", "    ")
			log.Println("Request received from body")
			log.Println(string(j))
		}
		ctx := context.WithValue(r.Context(), ReadUseConstrains, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// VouchersPostRequestKey is a context key
var VouchersPostRequestKey VoucherKey = "post-vouchers-req"

// ReadVoucherPostRequestFromBody skjdfn
func ReadVoucherPostRequestFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.VoucherPostRequest
		err := decoder.Decode(&productReq)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		// if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(productReq, "", "    ")
		log.Println("Product Post Request received from body")
		log.Println(string(j))
		// }
		ctx := context.WithValue(r.Context(), VouchersPostRequestKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// VoucherSoftReadReqKey is a context key
var VoucherSoftReadReqKey VoucherKey = "soft-read-vouchers-req"

// ReadVoucherSoftReadReqFromBody skjdfn
func ReadVoucherSoftReadReqFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.VoucherSoftReadReq
		err := decoder.Decode(&productReq)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		// if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(productReq, "", "    ")
		log.Println("Product Post Request received from body")
		log.Println(string(j))
		// }
		ctx := context.WithValue(r.Context(), VoucherSoftReadReqKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// VoucherSoftValidateReqKey is a context key
var VoucherSoftValidateReqKey VoucherKey = "soft-validate-vouchers-req"

// ReadVoucherSoftValidateReqFromBody skjdfn
func ReadVoucherSoftValidateReqFromBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var productReq types.VoucherSoftValidateReq
		err := decoder.Decode(&productReq)
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)

			return
		}
		defer r.Body.Close()
		// if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(productReq, "", "    ")
		log.Println("Product Post Request received from body")
		log.Println(string(j))
		// }
		ctx := context.WithValue(r.Context(), VoucherSoftValidateReqKey, productReq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
