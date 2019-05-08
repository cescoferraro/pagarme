package onni

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/onnidev/api/pagarme"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// BuyPostProductsTyped TODO: NEEDS COMMENT INFO
func BuyPostProductsTyped(ctx context.Context, all []types.BuyPartyProductsItemRequest) (types.BuyPostList, error) {
	result := []types.BuyPartyProductsItem{}
	for _, item := range all {
		partyP, err := PartyProduct(ctx, item.PartyProductID)
		if err != nil {
			return result, err
		}
		log.Println(">>>> retrieving party and club")
		party, club, err := PartyAndClub(ctx, partyP.PartyID.Hex())
		if err != nil {
			return result, err
		}
		instance := types.BuyPartyProductsItem{
			PartyProductID: item.PartyProductID,
			PromotionID:    item.PromotionID,
			InvoiceItem: types.InvoiceItem{
				PartyProductID: bson.ObjectIdHex(item.PartyProductID),
				UnitPrice: types.Price{
					Value:          types.PartyPPrice(partyP, &types.Promotion{}, party, club),
					CurrentIsoCode: "BRL",
				},
				Quantity: item.Quantity,
				Product: types.InvoiceItemProduct{
					Name: partyP.Name,
					Type: partyP.Type,
				}},
			PartyProduct: partyP,
			Quantity:     item.Quantity,
		}

		if bson.IsObjectIdHex(item.PromotionID) {
			id := bson.ObjectIdHex(item.PromotionID)
			instance.InvoiceItem.PromotionID = &id
			_, promotion, err := Promotion(ctx, item.PromotionID)
			if err != nil {
				return types.BuyPostList(result), err
			}
			instance.InvoiceItem.UnitPrice = types.Price{
				Value:          types.PartyPPrice(partyP, &promotion, party, club),
				CurrentIsoCode: "BRL",
			}
			instance.Promotion = &promotion
		}
		result = append(result, instance)
	}
	log.Printf("separating %v products\n", len(result))
	log.Printf("separating %v products\n", len(result))
	return types.BuyPostList(result), nil
}

// RefundVoucher TODO: NEEDS COMMENT INFO
func RefundVoucher(ctx context.Context, voucher types.Voucher, all types.BuyPostList) error {
	party, err := Party(ctx, voucher.PartyID.Hex())
	if err != nil {
		return err
	}
	club, err := Club(ctx, voucher.ClubID.Hex())
	if err != nil {
		return err
	}
	if voucher.InvoiceID == nil {
		return errors.New("impossible")
	}
	id := *voucher.InvoiceID
	invoice, err := Invoice(ctx, id.Hex())
	if err != nil {
		return err
	}
	token := viper.GetString("PAGARME")
	name := club.Name
	if len(name) > 12 {
		name = name[12:]
	}
	api := pagarme.New(token)
	if invoice.TransactionID == nil {
		return errors.New("impossible")
	}
	tid := *invoice.TransactionID
	log.Println("tid")
	log.Println(tid)
	log.Println(tid)
	log.Println(tid)
	log.Println(tid)
	t, err := api.TransactionRead(ctx, tid)
	if err != nil {
		return err
	}
	splits := []types.SplitRule{}
	count := ClubRecipientsCount(t.SplitRules)
	if count != 1 {
		return errors.New("impossible")
	}
	for _, split := range t.SplitRules {
		if IsONNiSplit(split) {
			splits = append(splits, types.SplitRule{
				ID:                  split.ID,
				Amount:              all.SumONNiString(party, club),
				RecipientID:         split.RecipientID,
				Liable:              "true",
				ChargeProcessingFee: "true",
				ChargeRemainderFee:  "true",
			})
			continue
		}
		splits = append(splits, types.SplitRule{
			ID:                  split.ID,
			Amount:              all.SumClubString(party, club),
			RecipientID:         split.RecipientID,
			Liable:              "false",
			ChargeProcessingFee: "false",
			ChargeRemainderFee:  "false",
		})
	}
	req := types.PagarMeTransactionRefundRequest{
		Amount:     all.SumBuy(party, club),
		APIKey:     token,
		Async:      "false",
		SplitRules: splits,
	}
	_, err = api.SplitRefund(ctx, tid, req)
	if err != nil {
		return err
	}
	return nil
}

// ClubRecipientsCount TODO: NEEDS COMMENT INFO
func ClubRecipientsCount(splits []types.SplitRuleResponse) int {
	count := 0
	for _, split := range splits {
		if !IsONNiSplit(split) {
			count = count + 1
		}
	}
	return count
}

// IsONNiSplit TODO: NEEDS COMMENT INFO
func IsONNiSplit(split types.SplitRuleResponse) bool {
	return shared.Contains([]string{
		"re_ciw98fu3s07vrcv5xm0r77hnp",
		"re_cjip2e4if01vzmb6d86hqh9ju",
	}, split.RecipientID)
}

// RefundInvoices TODO: NEEDS COMMENT INFO
func RefundInvoices(ctx context.Context, invoice types.Invoice) error {
	if invoice.TransactionID == nil {
		return errors.New("impossible")
	}
	token := viper.GetString("PAGARME")
	api := pagarme.New(token)
	tid := *invoice.TransactionID
	club, err := Club(ctx, invoice.ClubID.Hex())
	if err != nil {
		return err
	}
	name := club.Name
	if len(name) > 12 {
		name = name[12:]
	}
	t, err := api.TransactionRead(ctx, tid)
	if err != nil {
		return err
	}
	splits := []types.SplitRule{}
	count := ClubRecipientsCount(t.SplitRules)
	if count != 1 {
		return errors.New("impossible")
	}
	for _, split := range t.SplitRules {
		if IsONNiSplit(split) {
			splits = append(splits, types.SplitRule{
				ID:                  split.ID,
				Amount:              shared.AddCents(strconv.FormatFloat(invoice.ValueToOnni, 'f', -1, 64)),
				RecipientID:         split.RecipientID,
				Liable:              "true",
				ChargeProcessingFee: "true",
				ChargeRemainderFee:  "true",
			})
			continue
		}
		splits = append(splits, types.SplitRule{
			ID:                  split.ID,
			Amount:              shared.AddCents(strconv.FormatFloat(invoice.ValueToClub, 'f', -1, 64)),
			RecipientID:         split.RecipientID,
			Liable:              "false",
			ChargeProcessingFee: "false",
			ChargeRemainderFee:  "false",
		})
	}
	req := types.PagarMeTransactionRefundRequest{
		Amount:     invoice.Total.Value,
		APIKey:     token,
		Async:      "false",
		SplitRules: splits,
	}
	_, err = api.SplitRefund(ctx, tid, req)
	if err != nil {
		return err
	}
	return nil
}
