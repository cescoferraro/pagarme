package buy

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// Business is a comented function
func Business(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>> retrieving request from body")
	buy, ok := r.Context().Value(middlewares.BuyPostKey).(types.BuyPost)
	if !ok {
		err := errors.New("context bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>> retrieving customer from token")
	customer, ok := r.Context().Value(middlewares.CustomersKey).(types.Customer)
	if !ok {
		err := errors.New("context bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>> retrieving party and club")
	party, club, err := onni.PartyAndClub(r.Context(), buy.PartyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>>>> validate product request")
	err = onni.ValidateBuyPostRequest(r.Context(), buy, club, party, customer)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>>>> separar os produtos")
	products, err := onni.BuyPostProductsTyped(r.Context(), buy.Products)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	req, err := onni.CreateBuyRequest(r.Context(), club, party, customer, products)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>> antifraude ")
	antitheftResult, err := onni.CalculateAntiFraud(r.Context(), club, party, customer, products)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	err = onni.AntiTheftRoutine(r.Context(), club, party, customer, products, antitheftResult)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>> ban routine")
	ban, err := onni.BanRoutine(r.Context(), customer, &buy.Log, party, club)
	if err != nil {
		onni.CreateAntiTheftRecordBan(r.Context(), antitheftResult, ban)
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>>>> criar invoices e vouchers")
	invoices, vouchers, err := onni.CreateInvoicesAndVouchersBuy(r.Context(), club, party, customer, products, buy.Log)
	if err != nil {
		onni.CreateAntiTheftRecordInvoiceVoucherCreation(r.Context(), antitheftResult, invoices, vouchers)
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(">>>>>> indo ao pagarme")
	pg, err := onni.BuyPagarme(r.Context(), req)
	if err != nil {
		log.Println(">>>>>> buy error")
		onni.CreateAntiTheftRecordAfterPagarME(r.Context(), antitheftResult, invoices, vouchers, pg)
		ierr := onni.BuyError(invoices, vouchers)
		if ierr != nil {
			log.Println("ierr", ierr)
		}
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	tid := strconv.Itoa(pg.TID)
	log.Println(">>>>>> buy success")
	log.Println(antitheftResult, req)
	onni.CreateAntiTheftRecordAfterPagarME(r.Context(), antitheftResult, invoices, vouchers, pg)
	err = onni.BuySuccess(club, party, invoices, vouchers, products, customer, tid)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, buy)
}
