package onni

import (
	"context"
	"errors"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// BanRoutine TODO: NEEDS COMMENT INFO
func BanRoutine(ctx context.Context, customer types.Customer, buylog *types.Log, party types.Party, club types.Club) (*[]types.Ban, error) {
	response := []types.Ban{}
	if customer.Trusted != nil {
		if *customer.Trusted == "ACTIVE" {
			return &response, nil
		}
	}
	bans, err := WhayShouldIBan(ctx, customer)
	if err != nil {
		return &response, err
	}
	log.Println(bans)
	match, err := NewBanQuery(ctx, bans)
	if err != nil {
		return &response, err
	}
	log.Println(match)
	if len(match) != 0 {
		var emptyuser types.UserClub
		BanCustomerAndRefund(ctx, customer, party, club, buylog, &emptyuser)
		return &match, errors.New("banned")
	}
	return &response, nil
}

// CustomerUUIDFromPushRegistry TODO: NEEDS COMMENT INFO
func CustomerUUIDFromPushRegistry(ctx context.Context, id string) ([]string, error) {
	response := []string{}
	pushRegRepo, ok := ctx.Value(middlewares.PushRegistryRepoKey).(interfaces.PushRegistryRepo)
	if !ok {
		log.Println("pushhh")
		err := errors.New("context bug")
		return response, err
	}
	publishdPushRegistrys, err := pushRegRepo.GetCustomersUUID(id)
	if err != nil {
		return response, err
	}
	return publishdPushRegistrys, nil
}

// CustomerUUIDFromInvoices TODO: NEEDS COMMENT INFO
func CustomerUUIDFromInvoices(ctx context.Context, id string) ([]string, error) {
	response := []string{}
	invoiceRepo, ok := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
	if !ok {
		log.Println("invoices")
		err := errors.New("context bug")
		return response, err
	}
	publishdPushRegistrys, err := invoiceRepo.GetUniqueDeviceIDFromCustomer(id)
	if err != nil {
		return response, err
	}
	return publishdPushRegistrys, nil
}

// NewBanQuery TODO: NEEDS COMMENT INFO
func NewBanQuery(ctx context.Context, bans []types.Ban) ([]types.Ban, error) {
	bansRepo, ok := ctx.Value(middlewares.BanRepoKey).(interfaces.BansRepo)
	if !ok {
		err := errors.New("bug assert")
		return bans, err
	}
	all, err := bansRepo.ByPayload(bans)
	if err != nil {
		return bans, err
	}
	return all, nil
}

// WhayShouldIBan TODO: NEEDS COMMENT INFO
func WhayShouldIBan(ctx context.Context, customer types.Customer) ([]types.Ban, error) {
	bans := []types.Ban{}
	customerID := customer.ID.Hex()
	bans = append(bans, customerIDBan(customerID))
	if customer.Phone != "" {
		bans = append(bans, customerPhoneBan(customer))
	}
	allCards, err := CustomerCards(ctx, customerID)
	if err != nil {
		return bans, err
	}
	for _, card := range allCards {
		if card.CardToken != "" {
			bans = append(bans, cardHashBan(card))
		}
	}
	publishdPushRegistrys, err := CustomerUUIDFromPushRegistry(ctx, customer.ID.Hex())
	if err != nil {
		return bans, err
	}
	invoiceTokens, err := CustomerUUIDFromInvoices(ctx, customer.ID.Hex())
	if err != nil {
		return bans, err
	}
	for _, deviceid := range append(publishdPushRegistrys, invoiceTokens...) {
		if !ContainsBan(bans, deviceid) {
			bans = append(bans, deviceHashBan(deviceid))
		}
	}
	return bans, nil
}

// ContainsBan kjsdnf
func ContainsBan(s []types.Ban, e string) bool {
	for _, a := range s {
		if a.Payload == e {
			return true
		}
	}
	return false
}

// BanCustomer TODO: NEEDS COMMENT INFO
func BanCustomer(ctx context.Context, customer types.Customer, buylog *types.Log, userClub *types.UserClub) ([]types.Ban, error) {
	bans := []types.Ban{}
	all, err := WhayShouldIBan(ctx, customer)
	if err != nil {
		return bans, err
	}
	log.Println(all)
	for _, ban := range all {
		err := BanCreate(ctx, ban, buylog, userClub, customer)
		if err != nil {
			return bans, err
		}
	}
	return bans, nil
}

// BanCreate Insert the card on a database
func BanCreate(ctx context.Context, og types.Ban, buylog *types.Log, userClub *types.UserClub, customer types.Customer) error {
	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	bansRepo, err := interfaces.NewBansCollection(db)
	if err != nil {
		defer db.Session.Close()
		return err
	}
	ban, err := bansRepo.ByTypeAndPayload(og)
	if err != nil {
		log.Println("errro")
		if err == mgo.ErrNotFound {
			err := bansRepo.Collection.Insert(og.GenerateOccurences(customer, buylog, userClub))
			if err != nil {
				log.Println(err.Error())
				return err
			}
			log.Println("inserido")
			return nil
		}
		log.Println("doenerrro")
		return err
	}
	match := bson.M{"_id": ban.ID}
	now := types.Timestamp(time.Now())
	change := bson.M{
		"$set":  bson.M{"updateDate": &now},
		"$push": bson.M{"occurrences": ban.GetInfo(buylog, userClub, customer)},
	}
	err = bansRepo.Collection.Update(match, change)
	if err != nil {
		return err
	}
	log.Println(222)
	return nil
}

// BanCustomerAndRefund TODO: NEEDS COMMENT INFO
func BanCustomerAndRefund(ctx context.Context, customer types.Customer, party types.Party, club types.Club, buylog *types.Log, userClub *types.UserClub) ([]types.Ban, error) {
	bans, err := BanCustomer(ctx, customer, buylog, userClub)
	if err != nil {
		return bans, err
	}
	// TODO Refund from Vouchers
	// err = BanRefundVoucher(ctx, customer, party, club)
	// if err != nil {
	// 	return bans, err
	// }
	return bans, nil
}

// BanRefundVoucher TODO: NEEDS COMMENT INFO
func BanRefundVoucher(ctx context.Context, customer types.Customer, party types.Party, club types.Club) error {
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug assert")
		return err
	}
	vouchers, err := vouchersCollection.GetSimpleActualVoucherByCustomer(customer.ID.Hex())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, voucher := range vouchers {
		log.Println(">>>>>> separar os produtos")
		products, _ := BuyPostProductsTyped(ctx, voucher.BuyPartyProductsItemRequest())
		_ = RefundVoucher(ctx, voucher, products)
		_ = CancelVoucher(ctx, voucher, types.UserClub{ID: bson.NewObjectId(), Name: "BANGOD"})
	}
	return nil
}

func customerIDBan(id string) types.Ban {
	horario := types.Timestamp(time.Now())
	return types.Ban{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Type:         "CUSTOMERID",
		Payload:      id,
	}
}
func customerPhoneBan(customer types.Customer) types.Ban {
	horario := types.Timestamp(time.Now())
	return types.Ban{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Type:         "CUSTOMERPHONE",
		Payload:      customer.Phone,
	}
}

func deviceHashBan(deviceid string) types.Ban {
	horario := types.Timestamp(time.Now())
	return types.Ban{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Type:         "DEVICEID",
		Payload:      deviceid,
	}
}
func cardHashBan(card types.Card) types.Ban {
	horario := types.Timestamp(time.Now())
	return types.Ban{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Type:         "CARDHASH",
		Payload:      card.CardToken,
	}
}

// CustomerCards TODO: NEEDS COMMENT INFO
func CustomerCards(ctx context.Context, id string) ([]types.Card, error) {
	cardsRepo, ok := ctx.Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	if !ok {
		err := errors.New("bug assert")
		return []types.Card{}, err
	}
	allCards, err := cardsRepo.GetAllByCustomerID(id)
	if err != nil {
		return allCards, err
	}
	return allCards, nil
}

// CustomerDefaultCard TODO: NEEDS COMMENT INFO
func CustomerDefaultCard(ctx context.Context, id string) (types.Card, error) {
	cardsRepo, ok := ctx.Value(middlewares.CardsRepoKey).(interfaces.CardsRepo)
	if !ok {
		err := errors.New("bug assert")
		return types.Card{}, err
	}
	allCards, err := cardsRepo.GetByCustomerDefaultCard(id)
	if err != nil {
		return allCards, err
	}
	return allCards, nil
}
