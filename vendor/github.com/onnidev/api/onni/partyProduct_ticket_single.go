package onni

import (
	"log"
	"time"

	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// CreatePartyProductsTicketSingle TODO: NEEDS COMMENT INFO
func CreatePartyProductsTicketSingle(ticket types.SoftTicketPostPartyProduct, partyID bson.ObjectId) ([]types.PartyProduct, error) {
	horario := types.Timestamp(time.Now())
	log.Println("**************************")
	log.Println(ticket.ExhibitionName)
	if len(ticket.Batches) == 0 {
		instance := types.PartyProduct{
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
			AcceptsEvaluation: false,
			Featured:          false,
			Status:            PartyProductStatus(ticket.Active),
		}
		log.Println("adding a regular ticket")
		return []types.PartyProduct{instance}, nil
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
		AcceptsEvaluation: false,
		Featured:          false,
		Status:            PartyProductStatus(ticket.Active),
		Batches:           all,
	})
	return instances, nil
}
