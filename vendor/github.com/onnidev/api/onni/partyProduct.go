package onni

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SoftTicket FKNJSD  iiii qnj q
func SoftTicket(ctx context.Context, partyP types.PartyProduct) (types.SoftTicketPartyProduct, error) {
	status := false
	if partyP.Status == "ACTIVE" {
		status = true
	}
	empty := types.Timestamp(time.Time{})
	product := types.SoftTicketPartyProduct{
		PartyProductID:    partyP.ID,
		DoNotSellMoreThan: partyP.DoNotSellMoreThan,
		DoNotSellBefore:   &empty,
		DoNotSellAfter:    &empty,
		ExhibitionName:    partyP.Name,
		Batches:           []types.SoftSlaveTicketPartyProduct{},
		Featured:          partyP.Featured,
		MenuTicketID:      partyP.MenuTicketID,
		Price:             partyP.MoneyAmount,
		AdditionalInformations: types.AdditionalInformations{},
		Active:                 status,
	}
	if partyP.DoNotSellAfter != nil {
		product.DoNotSellAfter = partyP.DoNotSellAfter
	}
	if partyP.DoNotSellBefore != nil {
		product.DoNotSellBefore = partyP.DoNotSellBefore
	}
	if partyP.MenuTicketID != nil {
		id := *partyP.MenuTicketID
		pp, err := MenuTicket(ctx, partyP.PartyID.Hex(), id.Hex())
		if err != nil {
			log.Println("erooo pppp")
			return product, nil
		}
		log.Println("###### i FOUND THE INFORMATION")
		product.AdditionalInformations = types.AdditionalInformations{
			General:     pp.GeneralInformation,
			Free:        pp.FreeInformation,
			Anniversary: pp.AnniversaryInformation,
		}
	}
	return product, nil

}

// PartyProduct TODO: NEEDS COMMENT INFO
func PartyProduct(ctx context.Context, id string) (types.PartyProduct, error) {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return types.PartyProduct{}, err
	}
	return partyProductCollection.GetByID(id)
}

// Promotion TODO: NEEDS COMMENT INFO
func Promotion(ctx context.Context, id string) (types.PartyProduct, types.Promotion, error) {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return types.PartyProduct{}, types.Promotion{}, err
	}
	return partyProductCollection.GetPromotion(id)
}

// PartyPartyProducts TODO: NEEDS COMMENT INFO
func PartyPartyProducts(ctx context.Context, partyID string) ([]types.PartyProduct, error) {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return []types.PartyProduct{}, err
	}
	return partyProductCollection.GetByPartyID(partyID)
}

// PartyProductStatus TODO: NEEDS COMMENT INFO
func PartyProductStatus(is bool) string {
	if is {
		return "ACTIVE"
	}
	return "INACTIVE"
}

// DeprecatePartyProducts TODO: NEEDS COMMENT INFO
func DeprecatePartyProducts(ctx context.Context, all []types.PartyProduct) error {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	log.Println("partyProduct to deprecate", len(all))
	for _, product := range all {
		horario := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate": &horario,
					"deprecated": true,
				}},
			ReturnNew: true,
		}
		var result types.PartyProduct
		log.Println("deprecating product ", product.Name)
		_, err := partyProductCollection.
			Collection.Find(bson.M{"_id": product.ID}).Apply(change, &result)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeprecatePartyProductsByID TODO: NEEDS COMMENT INFO
func DeprecatePartyProductsByID(ctx context.Context, all []string) error {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	log.Println("partyProduct to deprecate", len(all), "=================")
	for _, product := range all {
		horario := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"updateDate": &horario,
					"deprecated": true,
				}},
			ReturnNew: true,
		}
		var result types.PartyProduct
		_, err := partyProductCollection.
			Collection.Find(bson.M{"_id": bson.ObjectIdHex(product)}).Apply(change, &result)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeprecatePartyProductsByType TODO: NEEDS COMMENT INFO
func DeprecatePartyProductsByType(ctx context.Context, partyID bson.ObjectId, kind string) error {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	log.Println("trying to deprecate partyProducts of type ", kind)
	all, err := partyProductCollection.GetByPartyIDAndType(partyID.Hex(), kind)
	if err != nil {
		return err
	}
	return DeprecatePartyProducts(ctx, all)
}

// PartyProducts TODO: NEEDS COMMENT INFO
func InsertPartyProducts(ctx context.Context, products []types.PartyProduct) error {
	partyProductCollection, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	for _, product := range products {
		err := partyProductCollection.Collection.Insert(product)
		if err != nil {
			return err
		}
	}
	return nil
}

// TicketsFromParty TODO: NEEDS COMMENT INFO
func TicketsFromParty(ctx context.Context, partyID string) ([]types.SoftTicketPartyProduct, error) {
	softticket := []types.SoftTicketPartyProduct{}
	repo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return softticket, err
	}
	tickets, err := repo.GetTicketsByPartyID(partyID)
	if err != nil {
		return softticket, err
	}
	log.Println("-----------------------------")
	log.Printf("got %v ticket partyProducts\n", len(tickets))
	log.Println("-----------------------------")
	for _, ticket := range tickets {
		log.Printf("ticket %v", ticket.ID)
		if !ticket.Deprecated {
			log.Printf("not deprecated")
			if len(ticket.Batches) == 0 && ticket.OwnerBatchID == nil {
				result, err := SoftTicket(ctx, ticket)
				if err != nil {
					return softticket, err
				}
				softticket = append(softticket, result)
			}
		}
	}
	log.Printf("#### numero de partyP  without batches %v", len(softticket))
	for _, ticket := range tickets {
		if !ticket.Deprecated {
			if len(ticket.Batches) != 0 {
				tk, err := SoftMasterTicket(ticket, ctx, tickets)
				if err != nil {
					return softticket, err
				}
				softticket = append(softticket, tk)
			}
		}
	}
	return softticket, nil
}

func getProduct(all []types.PartyProduct, id string) (types.SoftSlaveTicketPartyProduct, error) {
	for _, partyP := range all {
		if partyP.ID.Hex() == id {
			return partyP.SoftSlaveTicket(), nil
		}
	}
	err := errors.New("deu erro")
	return types.SoftSlaveTicketPartyProduct{}, err
}

//  SoftMasterTicket FKNJSD  iiii qnj q
func SoftMasterTicket(partyP types.PartyProduct, ctx context.Context, all []types.PartyProduct) (types.SoftTicketPartyProduct, error) {
	status := false
	if partyP.Status == "ACTIVE" {
		status = true
	}
	ids := []types.SoftSlaveTicketPartyProduct{}
	log.Println("*************")
	log.Println(partyP.Name, partyP.ID.Hex())
	log.Println(partyP.Batches)
	for _, id := range partyP.Batches {
		log.Println(" batches ============")
		log.Printf("trying to find %v", id.PartyProductID.Hex())
		instance, err := getProduct(all, id.PartyProductID.Hex())
		if err != nil {
			log.Println("(((((((((((())))))))))))")
			continue
		}
		log.Println(id.PartyProductID.Hex(), instance.ExhibitionName, instance.Price.Value, status)
		ids = append(ids, instance)
		log.Println("***********")
	}
	empty := types.Timestamp(time.Time{})
	product := types.SoftTicketPartyProduct{
		PartyProductID:    partyP.ID,
		DoNotSellMoreThan: partyP.DoNotSellMoreThan,
		DoNotSellBefore:   &empty,
		DoNotSellAfter:    &empty,
		ExhibitionName:    partyP.Name,
		Batches:           ids,
		Featured:          partyP.Featured,
		MenuTicketID:      partyP.MenuTicketID,
		Price:             partyP.MoneyAmount,
		AdditionalInformations: types.AdditionalInformations{},
		Active:                 status,
	}
	if partyP.DoNotSellAfter != nil {
		product.DoNotSellAfter = partyP.DoNotSellAfter
	}
	if partyP.DoNotSellBefore != nil {
		product.DoNotSellBefore = partyP.DoNotSellBefore
	}
	if partyP.MenuTicketID != nil {
		id := *partyP.MenuTicketID
		pp, err := MenuTicket(ctx, partyP.PartyID.Hex(), id.Hex())
		if err != nil {
			log.Println("erooo menu ticket")
			log.Println(id.Hex())
			return product, err
		}
		log.Println("###### i FOUND THE INFORMATION")
		product.AdditionalInformations = types.AdditionalInformations{
			General:     pp.GeneralInformation,
			Free:        pp.FreeInformation,
			Anniversary: pp.AnniversaryInformation,
		}
	}
	return product, nil
}

// DrinksFromParty TODO: NEEDS COMMENT INFO
func DrinksFromParty(ctx context.Context, partyID string) ([]types.SoftProductPartyProduct, error) {
	softproducts := []types.SoftProductPartyProduct{}
	repo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return softproducts, err
	}
	products, err := repo.GetDrinksByPartyID(partyID)
	if err != nil {
		return softproducts, err
	}
	for _, ticket := range products {
		if !ticket.Deprecated {
			pro, err := SoftProduct(ctx, ticket)
			if err != nil {
				return softproducts, err
			}
			softproducts = append(softproducts, pro)
		}
	}
	return softproducts, nil
}

// BrandNewTicketsFromMenu TODO: NEEDS COMMENT INFO
func BrandNewTicketsFromMenu(ctx context.Context, menu types.ClubMenuTicket, partyID bson.ObjectId) ([]types.PartyProduct, error) {
	products := BrandNew(menu.Tickets, partyID)
	err := InsertPartyProducts(ctx, products)
	if err != nil {
		return products, err
	}
	return products, nil
}

// BrandNew TODO: NEEDS COMMENT INFO
func BrandNew(tickets []types.MenuTicket, partyID bson.ObjectId) []types.PartyProduct {
	result := []types.PartyProduct{}
	for _, ticket := range tickets {
		instances := BrandNewSingle(ticket, partyID)
		result = append(result, instances...)
	}
	return result
}

// BrandNewSingle TODO: NEEDS COMMENT INFO
func BrandNewSingle(ticket types.MenuTicket, partyID bson.ObjectId) []types.PartyProduct {
	log.Println("|||||||||************************")
	log.Println(ticket.ExhibitionName)
	masterID := bson.NewObjectId()
	instances := []types.PartyProduct{}
	horario := types.Timestamp(time.Now())
	zero := types.Timestamp(time.Time{})
	for _, batch := range ticket.Batches {
		after := types.Timestamp(time.Time{})
		before := types.Timestamp(time.Time{})
		if batch.DoNotSellBefore != nil {
			log.Println("batch has before")
			before = *batch.DoNotSellBefore
		}
		if batch.DoNotSellAfter != nil {
			log.Println("batch has after")
			after = *batch.DoNotSellAfter
		}
		instance := types.PartyProduct{
			ID:                bson.NewObjectId(),
			Name:              batch.ExhibitionName,
			CreationDate:      &horario,
			UpdateDate:        &zero,
			PartyID:           partyID,
			MoneyAmount:       batch.Price,
			DoNotSellMoreThan: batch.DoNotSellMoreThan,
			DoNotSellBefore:   &before,
			DoNotSellAfter:    &after,
			QuantityPurchased: 0.0,
			Type:              "TICKET",
			QuantityFree:      0,
			MenuTicketID:      &ticket.ID,
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
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	if ticket.DoNotSellBefore != nil {
		before = *ticket.DoNotSellBefore
	}
	if ticket.DoNotSellAfter != nil {
		after = *ticket.DoNotSellAfter
	}
	log.Println("inserting master")
	instances = append(instances, types.PartyProduct{
		ID:                masterID,
		Name:              ticket.ExhibitionName,
		CreationDate:      &horario,
		PartyID:           partyID,
		MoneyAmount:       ticket.Price,
		DoNotSellMoreThan: ticket.DoNotSellMoreThan,
		DoNotSellBefore:   &before,
		DoNotSellAfter:    &after,
		QuantityPurchased: 0.0,
		Type:              "TICKET",
		QuantityFree:      0,
		MenuTicketID:      &ticket.ID,
		AcceptsEvaluation: false,
		Featured:          false,
		Status:            PartyProductStatus(ticket.Active),
		Batches:           all,
		// Product
	})
	return instances
}

// BrandNewProductsFromMenu TODO: NEEDS COMMENT INFO
func BrandNewProductsFromMenu(ctx context.Context, menu types.ClubMenuProduct, partyID bson.ObjectId) ([]types.PartyProduct, error) {
	products, err := HHH(ctx, menu, partyID)
	if err != nil {
		return []types.PartyProduct{}, err
	}
	err = InsertPartyProducts(ctx, products)
	if err != nil {
		return products, err
	}
	return products, nil
}

// HHH TODO: NEEDS COMMENT INFO
func HHH(ctx context.Context, menu types.ClubMenuProduct, partyID bson.ObjectId) ([]types.PartyProduct, error) {
	products := menu.Products
	newProducts := []types.PartyProduct{}
	for _, product := range products {
		partyP, err := SingleBrandNewProduct(ctx, product, partyID)
		if err != nil {
			return []types.PartyProduct{}, err
		}
		newProducts = append(newProducts, partyP)
	}
	return newProducts, nil
}

// SingleBrandNewProduct TODO: NEEDS COMMENT INFO
func SingleBrandNewProduct(ctx context.Context, product types.MenuProduct, partyID bson.ObjectId) (types.PartyProduct, error) {
	horario := types.Timestamp(time.Now())
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	if product.DoNotSellBefore != nil {
		before = *product.DoNotSellBefore
	}
	if product.DoNotSellAfter != nil {
		after = *product.DoNotSellAfter
	}
	partyP := types.PartyProduct{
		ID:                bson.NewObjectId(),
		Name:              product.ExhibitionName,
		CreationDate:      &horario,
		PartyID:           partyID,
		MoneyAmount:       product.Price,
		DoNotSellMoreThan: product.DoNotSellMoreThan,
		Featured:          product.Featured,
		DoNotSellBefore:   &before,
		DoNotSellAfter:    &after,
		Category:          &product.Category,
		QuantityPurchased: 0.0,
		Type:              "DRINK",
		QuantityFree:      0,
		MenuProductID:     &product.ID,
		AcceptsEvaluation: false,
		Image:             &product.Image,
		Status:            PartyProductStatus(product.Active),
		Combo:             product.Combo,
	}

	if product.Combo == nil {
		if product.ProductID != nil {
			repo := ctx.Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
			id := *product.ProductID
			product, err := repo.GetByID(id.Hex())
			if err != nil {
				return types.PartyProduct{}, err
			}
			partyP.Product = &product
		}
	}
	j, _ := json.MarshalIndent(product.Combo, "", "    ")
	log.Println("#######", string(j))
	return partyP, nil
}

// SoftProduct FKNJSD  iiii qnj q
func SoftProduct(ctx context.Context, partyP types.PartyProduct) (types.SoftProductPartyProduct, error) {
	log.Println("####### PRODUTO")
	status := false
	if partyP.Status == "ACTIVE" {
		status = true
	}
	product := types.SoftProductPartyProduct{
		PartyProductID:    partyP.ID,
		DoNotSellMoreThan: partyP.DoNotSellMoreThan,
		DoNotSellBefore:   partyP.DoNotSellBefore,
		DoNotSellAfter:    partyP.DoNotSellAfter,
		ExhibitionName:    partyP.Name,
		Featured:          partyP.Featured,
		MenuProductID:     partyP.MenuProductID,
		Price:             partyP.MoneyAmount,
		Combo:             partyP.Combo,
		Active:            status,
	}
	if partyP.Product != nil {
		pp := *partyP.Product
		product.ProductID = &pp.ID

	}
	if partyP.Image != nil {
		product.Image = *partyP.Image
	}
	if partyP.Category != nil {
		product.Category = *partyP.Category
	}
	if partyP.Product != nil {
		product.ProductID = &partyP.Product.ID
	}
	if partyP.MenuProductID != nil {
		id := *partyP.MenuProductID
		log.Println(partyP.ID.Hex(), id.Hex())
		pp, err := MenuProduct(ctx, partyP.PartyID.Hex(), id.Hex())
		if err != nil {
			log.Printf("##### fudeu nao achou a info do producto %v \n", partyP.Name)
			return product, nil
		}
		log.Println("###### i FOUND THE INFORMATION")
		product.GeneralInformation = pp.GeneralInformation
	}
	return product, nil

}
