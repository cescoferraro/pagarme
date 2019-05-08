package onni

import (
	"log"
	"time"

	"github.com/onnidev/api/types"

	"gopkg.in/mgo.v2/bson"
)

// SoftTicketPartyProductConverter TODO: NEEDS COMMENT INFO
func SoftTicketPartyProductConverter(tickets []types.SoftTicketPostPartyProduct) []types.MenuTicket {
	newTickets := []types.MenuTicket{}
	log.Println("==============================")
	for _, ticket := range tickets {
		newTickets = append(newTickets, SoftTicketPartyProductSingleConverter(ticket))
	}
	return newTickets
}

// SoftTicketPartyProductSingleConverter TODO: NEEDS COMMENT INFO
func SoftTicketPartyProductSingleConverter(ticket types.SoftTicketPostPartyProduct) types.MenuTicket {
	id := bson.NewObjectId()
	if len(ticket.MenuTicketID) != 0 {
		id = bson.ObjectIdHex(ticket.MenuTicketID)
	}
	batches := []types.MenuTicket{}
	credit := new(float64)
	for _, batch := range ticket.Batches {
		id := bson.NewObjectId()
		after := types.Timestamp(time.Time{})
		before := types.Timestamp(time.Time{})
		if len(batch.MenuTicketID) != 0 {
			id = bson.ObjectIdHex(batch.MenuTicketID)
		}
		if batch.DoNotSellBefore != nil {
			before = *batch.DoNotSellBefore
		}
		if batch.DoNotSellAfter != nil {
			after = *batch.DoNotSellAfter
		}
		batches = append(batches, types.MenuTicket{
			ID:                     id,
			Price:                  types.Price{Value: batch.Price, CurrentIsoCode: "BRL"},
			ExhibitionName:         batch.ExhibitionName,
			DoNotSellMoreThan:      batch.DoNotSellMoreThan,
			DoNotSellBefore:        &before,
			CreditConsumable:       credit,
			DoNotSellAfter:         &after,
			Active:                 batch.Active,
			GeneralInformation:     batch.AdditionalInformations.General,
			FreeInformation:        batch.AdditionalInformations.Free,
			AnniversaryInformation: batch.AdditionalInformations.Anniversary,
			Batches:                []types.MenuTicket{},
		})

	}
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	if ticket.DoNotSellBefore != nil {
		before = *ticket.DoNotSellBefore
	}
	if ticket.DoNotSellAfter != nil {
		after = *ticket.DoNotSellAfter
	}
	return types.MenuTicket{
		ID:                     id,
		Price:                  types.Price{Value: ticket.Price, CurrentIsoCode: "BRL"},
		ExhibitionName:         ticket.ExhibitionName,
		DoNotSellMoreThan:      ticket.DoNotSellMoreThan,
		DoNotSellBefore:        &before,
		DoNotSellAfter:         &after,
		Active:                 ticket.Active,
		CreditConsumable:       credit,
		GeneralInformation:     ticket.AdditionalInformations.General,
		FreeInformation:        ticket.AdditionalInformations.Free,
		AnniversaryInformation: ticket.AdditionalInformations.Anniversary,
		Batches:                batches,
	}
}
