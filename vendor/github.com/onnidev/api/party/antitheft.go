package party

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// AntiTheft TODO: NEEDS COMMENT INFO
func AntiTheft(w http.ResponseWriter, r *http.Request) {
	partyID := chi.URLParam(r, "partyId")
	model, ok := r.Context().Value(middlewares.AntiTheftModelRequestKey).(types.AntiTheftModel)
	if !ok {
		err := errors.New("assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	repo, ok := r.Context().Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
	if !ok {
		err := errors.New("assert")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	invoices, err := repo.GetByPartyID(partyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	uniqueCustomers := []bson.ObjectId{}
	for _, invoice := range invoices {
		if !shared.ContainsObjectID(uniqueCustomers, invoice.CustomerID) {
			uniqueCustomers = append(uniqueCustomers, invoice.CustomerID)
		}
	}
	customers := []types.Customer{}
	for _, customerID := range uniqueCustomers {
		customer, err := onni.Customer(r.Context(), customerID.Hex())
		if err != nil {
			shared.MakeONNiError(w, r, 400, err)
			return
		}
		customers = append(customers, customer)
	}

	party, err := onni.Party(r.Context(), partyID)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := onni.Club(r.Context(), party.ClubID.Hex())
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	result := []types.AntiTheftResult{}
	for _, customer := range customers {
		for _, invoice := range CustomerInvoices(invoices, customer) {
			totalDrink := float64(0)
			totalTicket := float64(0)
			for _, item := range invoice.Itens {
				if item.Product.Type == "DRINK" {
					totalDrink = totalDrink + (float64(item.Quantity) * item.UnitPrice.Value)
				}
				if item.Product.Type == "TICKET" {
					totalTicket = totalTicket + (float64(item.Quantity) * item.UnitPrice.Value)
				}
			}
			scoreCard, err := onni.AntiTheftScoreCard(r.Context(), club, party, customer, totalTicket, totalDrink, model)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			vouchers, err := onni.PartyCustomerVoucher(r.Context(), party, customer)
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			cards, err := onni.CustomerCards(r.Context(), customer.ID.Hex())
			if err != nil {
				shared.MakeONNiError(w, r, 400, err)
				return
			}
			all := types.VouchersList(vouchers)
			idate := (invoice.CreationDate.Time())
			tickets := onni.TicketScore(party, club, all.Before(idate).Revert(), totalTicket, model)
			drinks := onni.DrinkScore(party, club, all.Before(idate).Revert(), totalDrink, model)
			result = append(result, types.AntiTheftResult{
				CustomerID:         customer.ID,
				Cards:              len(cards),
				CustomerDate:       customer.CreationDate.Time().String(),
				CustomerName:       customer.Name(),
				CustomerMail:       customer.Mail,
				Total:              invoice.Total.Value,
				ScoreCard:          scoreCard,
				ScoreDrinks:        drinks,
				ScoreTickets:       tickets,
				Score:              onni.AntiTheftArithmetics(totalTicket, totalDrink, tickets, drinks, scoreCard),
				AccoutableDrinks:   all.Before(idate).Accountable().FilterDrinks().Sum(),
				AccoutableTickets:  all.Before(idate).Accountable().FilterTickets().Sum(),
				ClubTicketsAverage: onni.ClubTicketsAverage(club),
				ClubDrinksAverage:  onni.ClubDrinksAverage(club),
			})
			// }
		}
	}

	var b bytes.Buffer
	file := io.Writer(&b)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	all := [][]string{}
	for _, value := range result {
		all = append(all, []string{
			value.CustomerID.Hex(),
			value.CustomerName,
			value.CustomerMail,
			strconv.Itoa(value.ScoreCard),
			strconv.Itoa(value.ScoreTickets),
			strconv.Itoa(value.ScoreDrinks),
			strconv.Itoa(value.Score),
			strconv.FormatFloat(value.AccoutableTickets, 'f', -1, 64),
			strconv.FormatFloat(value.AccoutableDrinks, 'f', -1, 64),
			strconv.FormatFloat(value.Total, 'f', -1, 64),
			strconv.FormatFloat(value.ClubTicketsAverage, 'f', -1, 64),
			strconv.FormatFloat(value.ClubDrinksAverage, 'f', -1, 64),
			strconv.Itoa(value.Cards)})
	}
	err = writer.WriteAll(all)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	b.WriteTo(w)
}

// CustomerInvoices TODO: NEEDS COMMENT INFO
func CustomerInvoices(invoices []types.Invoice, customer types.Customer) []types.Invoice {
	result := []types.Invoice{}
	for _, invoice := range invoices {
		if invoice.CustomerID.Hex() == customer.ID.Hex() {
			result = append(result, invoice)
		}
	}
	return result
}
