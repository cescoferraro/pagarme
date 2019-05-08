package onni

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// CreatePartyProductsTicket TODO: NEEDS COMMENT INFO
func CreatePartyProductsTicket(ctx context.Context, tickets []types.SoftTicketPostPartyProduct, menuticketid, partyID bson.ObjectId) error {
	productsCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("assert")
		return err
	}
	horario := types.Timestamp(time.Now())
	log.Println(len(tickets))
	for _, ticket := range tickets {
		log.Println("**************************")
		log.Println(ticket.ExhibitionName)
		if len(ticket.Batches) == 0 {
			instance := types.PartyProduct{
				ID:                bson.NewObjectId(),
				Name:              ticket.ExhibitionName,
				CreationDate:      &horario,
				PartyID:           partyID,
				MoneyAmount:       types.Price{Value: ticket.Price, CurrentIsoCode: "BRL"},
				DoNotSellMoreThan: ticket.DoNotSellMoreThan,
				DoNotSellBefore:   ticket.DoNotSellBefore,
				DoNotSellAfter:    ticket.DoNotSellAfter,
				QuantityPurchased: 0.0,
				Type:              "TICKET",
				QuantityFree:      0,
				MenuTicketID:      &menuticketid,
				AcceptsEvaluation: false,
				Featured:          false,
				Status:            PartyProductStatus(ticket.Active),
				// Product
			}
			log.Println("adding a regular ticket")
			log.Println(ticket.DoNotSellAfter)
			log.Println(ticket.DoNotSellBefore)
			err := productsCollection.Collection.Insert(instance)
			if err != nil {
				return err
			}
			continue
		}
		masterID := bson.NewObjectId()
		instances := []types.PartyProduct{}
		for _, batch := range ticket.Batches {
			log.Println("inserting batch")
			instance := types.PartyProduct{
				ID:                bson.NewObjectId(),
				Name:              batch.ExhibitionName,
				CreationDate:      &horario,
				PartyID:           partyID,
				MoneyAmount:       types.Price{Value: batch.Price, CurrentIsoCode: "BRL"},
				DoNotSellMoreThan: batch.DoNotSellMoreThan,
				DoNotSellBefore:   batch.DoNotSellBefore,
				DoNotSellAfter:    batch.DoNotSellAfter,
				QuantityPurchased: 0.0,
				Type:              "TICKET",
				QuantityFree:      0,
				MenuTicketID:      &menuticketid,
				AcceptsEvaluation: false,
				Featured:          false,
				Status:            PartyProductStatus(batch.Active),
				OwnerBatchID:      &masterID,
			}
			instances = append(instances, instance)
		}
		all := []types.ObjectPartyProduct{}
		for _, pro := range instances {
			all = append(all, types.ObjectPartyProduct{PartyProductID: pro.ID})
		}
		log.Println("inserting master")
		instances = append(instances, types.PartyProduct{
			ID:                masterID,
			Name:              ticket.ExhibitionName,
			CreationDate:      &horario,
			PartyID:           partyID,
			MoneyAmount:       types.Price{Value: ticket.Price, CurrentIsoCode: "BRL"},
			DoNotSellMoreThan: ticket.DoNotSellMoreThan,
			DoNotSellBefore:   ticket.DoNotSellBefore,
			DoNotSellAfter:    ticket.DoNotSellAfter,
			QuantityPurchased: 0.0,
			Type:              "TICKET",
			QuantityFree:      0,
			MenuTicketID:      &menuticketid,
			AcceptsEvaluation: false,
			Featured:          false,
			Status:            PartyProductStatus(ticket.Active),
			Batches:           all,
			// Product
		})
		for _, instance := range instances {
			err := productsCollection.Collection.Insert(instance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
