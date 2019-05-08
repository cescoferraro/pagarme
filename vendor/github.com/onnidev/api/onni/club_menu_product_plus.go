package onni

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeprecatePartyProductAndUpdateClubMenuProduct TODO: NEEDS COMMENT INFO
func DeprecatePartyProductAndUpdateClubMenuProduct(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuProduct, error) {
	result := types.ClubMenuProduct{}
	err := DeprecatePartyProductsByType(ctx, party.ID, "DRINK")
	if err != nil {
		return result, err
	}
	menu, err := UpdateClubMenuProduct(ctx, req)
	if err != nil {
		return result, err
	}
	_, err = BrandNewProductsFromMenu(ctx, menu, party.ID)
	if err != nil {
		return result, err
	}
	return menu, nil
}

// DeprecatePartyProductAndCreateClubMenuProduct TODO: NEEDS COMMENT INFO
func DeprecatePartyProductAndCreateClubMenuProduct(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuProduct, error) {
	result := types.ClubMenuProduct{}
	err := DeprecatePartyProductsByType(ctx, party.ID, "DRINK")
	if err != nil {
		return result, err
	}
	menu, err := CreateClubMenuProduct(ctx, req)
	if err != nil {
		return result, err
	}
	_, err = BrandNewProductsFromMenu(ctx, menu, party.ID)
	if err != nil {
		return result, err
	}
	return menu, nil
}

// UpdateCurrentClubMenuuProductAndPartyProducts TODO: NEEDS COMMENT INFO
func UpdateCurrentClubMenuProductAndPartyProducts(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) (types.ClubMenuProduct, error) {
	menuid := bson.ObjectIdHex(*req.ClubMenuProduct.ID)
	result := types.ClubMenuProduct{}
	repo, ok := ctx.Value(middlewares.ClubMenuProductRepoKey).(interfaces.ClubMenuProductRepo)
	if !ok {
		err := errors.New("assert")
		return result, err
	}
	menu, err := repo.GetByID(menuid.Hex())
	if err != nil {
		return result, err
	}
	err = DeprecateUnusedProducts(ctx, party, req)
	if err != nil {
		return result, err
	}

	for _, product := range req.Products {
		log.Println("====================================")
		log.Printf("#### We are looping over product name %v\n", product.ExhibitionName)
		if bson.IsObjectIdHex(product.MenuProductID) {
			err := DealWithMenuProduct(ctx, product, party, menu, req.ClubMenuProduct.Name)
			if err != nil {
				return result, err
			}
			continue
		}
		log.Println("##Não tem product menu product")
		err := AddMenuProduct(ctx, product, party, req)
		if err != nil {
			return result, err
		}
	}
	return menu, nil

}

// DealWithMenuProduct TODO: NEEDS COMMENT INFO
func DealWithMenuProduct(ctx context.Context, product types.SoftProductPostPartyProduct, party types.Party, menu types.ClubMenuProduct, newname string) error {
	log.Println("########################################")
	log.Printf("########## RECEIVED MENUPRODUCTID OF %v\n", product.MenuProductID)
	log.Printf("#### We patching this existing menuProduct  %v\n", product.ExhibitionName)
	log.Printf("#### with the image %v\n", product.Image[:12])
	image, err := ImageONNiCombo(ctx, product.Image)
	if err != nil {
		return err
	}
	after := types.Timestamp(time.Time{})
	before := types.Timestamp(time.Time{})
	master := types.MenuProduct{
		ID:                 bson.ObjectIdHex(product.MenuProductID),
		Featured:           product.Featured,
		Price:              types.Price{Value: product.Price, CurrentIsoCode: "BRL"},
		ExhibitionName:     product.ExhibitionName,
		DoNotSellMoreThan:  product.DoNotSellMoreThan,
		Image:              image,
		DoNotSellBefore:    &before,
		DoNotSellAfter:     &after,
		Category:           product.Category,
		Active:             product.Active,
		GeneralInformation: product.GeneralInformation,
		Combo:              product.Combo,
	}
	if bson.IsObjectIdHex(product.ProductID) {
		id := bson.ObjectIdHex(product.ProductID)
		master.ProductID = &id
	}
	err = EditMenuProduct(ctx, master, party, newname)
	if err != nil {
		return err
	}
	if bson.IsObjectIdHex(product.PartyProductID) {
		log.Println("patch this motherfucking product")
		err = PatchProductPartyProduct(ctx, product.PartyProductID, master)
		if err != nil {
			return err
		}
		return nil
	}
	log.Println("supposed to create this motherfucking product")
	all, err := SingleBrandNewProduct(ctx, master, party.ID)
	if err != nil {
		return err
	}
	j, _ := json.MarshalIndent(all, "", "    ")
	log.Println(string(j))
	err = InsertPartyProducts(ctx, []types.PartyProduct{all})
	if err != nil {
		return err
	}
	return nil
}

// DeprecateUnusedProducts TODO: NEEDS COMMENT INFO
func DeprecateUnusedProducts(ctx context.Context, party types.Party, req types.SoftPartyPostRequest) error {
	partyPRepo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	stillused := []string{}
	for _, ticket := range req.Products {
		if ticket.PartyProductID != "" {
			stillused = append(stillused, ticket.PartyProductID)
		}
	}
	current, err := partyPRepo.GetByPartyIDAndType(party.ID.Hex(), "DRINK")
	if err != nil {
		return err
	}
	obsolete := []string{}
	for _, product := range current {
		if !shared.Contains(stillused, product.ID.Hex()) {
			obsolete = append(obsolete, product.ID.Hex())
		}
	}
	err = DeprecatePartyProductsByID(ctx, obsolete)
	if err != nil {
		return err
	}
	return nil
}

// PatchProductPartyProduct TODO: NEEDS COMMENT INFO
func PatchProductPartyProduct(ctx context.Context, id string, instance types.MenuProduct) error {
	repo, ok := ctx.Value(middlewares.PartyProductsRepoKey).(interfaces.PartyProductsRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	partyP, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	log.Println(partyP.Batches)
	// log.Println(pbatches)
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updateDate":        &now,
				"moneyAmount":       instance.Price,
				"featured":          instance.Featured,
				"category":          instance.Category,
				"exhibitionName":    instance.ExhibitionName,
				"combo":             instance.Combo,
				"doNotSellMoreThan": instance.DoNotSellMoreThan,
				"active":            instance.Active,
				"status":            PartyProductStatus(instance.Active),
			}},
		ReturnNew: true,
	}
	patchedProduct := types.PartyProduct{}
	_, err = repo.Collection.Find(bson.M{"_id": partyP.ID}).Apply(change, &patchedProduct)
	if err != nil {
		return err
	}
	return nil
}
