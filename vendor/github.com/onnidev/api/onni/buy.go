package onni

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"

	"gopkg.in/mgo.v2/bson"
)

// CreatePendingDrinksInvoiceAndVouchersBuy TODO: NEEDS COMMENT INFO
func CreatePendingDrinksInvoiceAndVouchersBuy(ctx context.Context,
	club types.Club,
	party types.Party,
	customer types.Customer,
	products types.BuyPostList,
	logg types.Log,
) (types.Invoice, []types.Voucher, error) {
	horario := types.Timestamp(time.Now())
	vouchers := []types.Voucher{}
	drinkInvoice := types.Invoice{}
	invoiceID := bson.NewObjectId()
	if products.Drinks().Size() != 0 {
		log.Println(">>>>>> vou criar drinks invoice")
		drinkInvoice = types.Invoice{
			ID:            invoiceID,
			CreationDate:  &horario,
			CustomerID:    customer.ID,
			PartyID:       party.ID,
			ClubID:        club.ID,
			Status:        "PENDING",
			OperationType: "ONNI_APP",
			Itens:         products.Drinks().InvoiceItens(),
			Total:         types.Price{Value: products.SumDrinks(party, club), CurrentIsoCode: "BRL"},
			Log: types.Log{
				AppVersion: "4",
				DeviceID:   logg.DeviceID,
				DeviceSO:   strings.ToUpper(logg.DeviceSO),
				Latitude:   logg.Latitude,
				Longitude:  logg.Longitude,
			},
		}
		invoiceRepo, ok := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
		if !ok {
			err := errors.New("assert")
			return drinkInvoice, vouchers, err
		}
		err := invoiceRepo.Collection.Insert(drinkInvoice)
		if err != nil {
			return drinkInvoice, vouchers, err
		}
		for _, item := range products.Drinks().HasPromotion() {
			repo, ok := ctx.Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
			if !ok {
				return drinkInvoice, vouchers, err
			}
			promo, err := repo.PromotionalCustomersFromPromotion(item.Promotion.ID.Hex(), customer.ID.Hex())
			if err != nil {
				return drinkInvoice, vouchers, err
			}
			for i := 1; i <= int(item.Quantity); i++ {
				fmt.Println("criando um voucher promotion", i)
				voucher := types.Voucher{
					ID:                    bson.NewObjectId(),
					CreationDate:          &horario,
					StartDate:             party.StartDate,
					EndDate:               party.EndDate,
					CustomerID:            customer.ID,
					PartyID:               party.ID,
					ClubID:                club.ID,
					PartyProductID:        item.PartyProduct.ID,
					PromotionID:           &item.Promotion.ID,
					InvoiceID:             &invoiceID,
					ClubName:              club.Name,
					PartyName:             party.Name,
					CustomerName:          customer.Name(),
					Status:                "PROCESSING",
					ResponsableUserClubID: &promo.PromoterID,
					Price: types.Price{
						Value:          types.PartyPPrice(item.PartyProduct, item.Promotion, party, club),
						CurrentIsoCode: "BRL",
					},
					Product: item.VoucherProduct(),
					Type:    "PROMOTION",
				}
				vouchers = append(vouchers, voucher)
			}
		}
		for _, item := range products.Drinks().DoesNotHavePromotion() {
			for i := 1; i <= int(item.Quantity); i++ {
				fmt.Println("criando um regular voucher ", i)
				voucher := types.Voucher{
					ID:             bson.NewObjectId(),
					CreationDate:   &horario,
					StartDate:      party.StartDate,
					EndDate:        party.EndDate,
					CustomerID:     customer.ID,
					PartyID:        party.ID,
					ClubID:         club.ID,
					PartyProductID: item.PartyProduct.ID,
					InvoiceID:      &invoiceID,
					ClubName:       club.Name,
					PartyName:      party.Name,
					CustomerName:   customer.Name(),
					Status:         "PROCESSING",
					Price: types.Price{
						Value:          types.PartyPPrice(item.PartyProduct, &types.Promotion{}, party, club),
						CurrentIsoCode: "BRL",
					},
					Product: item.VoucherProduct(),
					Type:    "NORMAL",
				}
				vouchers = append(vouchers, voucher)
			}
		}
		vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
		if !ok {
			err := errors.New("assert")
			return drinkInvoice, vouchers, err
		}
		for _, voucher := range vouchers {
			err = vouchersCollection.Collection.Insert(voucher)
			if err != nil {
				return drinkInvoice, vouchers, err
			}
		}
		return drinkInvoice, vouchers, nil
	}
	log.Println(">>> returrn empty drinkinvoice")
	return drinkInvoice, vouchers, nil
}

// CreatePendingTicketsInvoiceAndVouchersBuy TODO: NEEDS COMMENT INFO
func CreatePendingTicketsInvoiceAndVouchersBuy(ctx context.Context,
	club types.Club,
	party types.Party,
	customer types.Customer,
	products types.BuyPostList,
	logg types.Log,
) (types.Invoice, []types.Voucher, error) {
	horario := types.Timestamp(time.Now())
	vouchers := []types.Voucher{}
	ticketInvoice := types.Invoice{}
	if products.Tickets().Size() != 0 {
		log.Println(">>>>>> vou criar tickets invoice")
		invoiceID := bson.NewObjectId()
		ticketInvoice = types.Invoice{
			ID:            invoiceID,
			CreationDate:  &horario,
			CustomerID:    customer.ID,
			PartyID:       party.ID,
			ClubID:        club.ID,
			Status:        "PENDING",
			OperationType: "ONNI_APP",
			Itens:         products.Tickets().InvoiceItens(),
			Total:         types.Price{Value: products.SumTickets(party, club), CurrentIsoCode: "BRL"},
			Log: types.Log{
				AppVersion: "4",
				DeviceID:   logg.DeviceID,
				DeviceSO:   strings.ToUpper(logg.DeviceSO),
				Latitude:   logg.Latitude,
				Longitude:  logg.Longitude,
			},
		}
		log.Println(ticketInvoice.ID.Hex())
		for _, item := range products.Tickets().HasPromotion() {
			repo, ok := ctx.Value(middlewares.PromotionalCustomerRepoKey).(interfaces.PromotionalCustomerRepo)
			if !ok {
				err := errors.New("bug")
				return ticketInvoice, vouchers, err
			}
			promo, err := repo.PromotionalCustomersFromPromotion(item.Promotion.ID.Hex(), customer.ID.Hex())
			if err != nil {
				return ticketInvoice, vouchers, err
			}
			for i := 1; i <= int(item.Quantity); i++ {
				fmt.Println("criando um voucher promotion", i)
				voucher := types.Voucher{
					ID:                    bson.NewObjectId(),
					CreationDate:          &horario,
					StartDate:             party.StartDate,
					EndDate:               party.EndDate,
					CustomerID:            customer.ID,
					PartyID:               party.ID,
					ClubID:                club.ID,
					PartyProductID:        item.PartyProduct.ID,
					PromotionID:           &item.Promotion.ID,
					InvoiceID:             &invoiceID,
					ClubName:              club.Name,
					PartyName:             party.Name,
					CustomerName:          customer.Name(),
					Status:                "PROCESSING",
					ResponsableUserClubID: &promo.PromoterID,
					Price: types.Price{
						Value:          types.PartyPPrice(item.PartyProduct, item.Promotion, party, club),
						CurrentIsoCode: "BRL",
					},
					Product: item.VoucherProduct(),
					Type:    "PROMOTION",
				}
				vouchers = append(vouchers, voucher)
			}
		}
		for _, item := range products.Tickets().DoesNotHavePromotion() {
			for i := 1; i <= int(item.Quantity); i++ {
				fmt.Println("criando um regular voucher ", i)
				voucher := types.Voucher{
					ID:             bson.NewObjectId(),
					CreationDate:   &horario,
					StartDate:      party.StartDate,
					EndDate:        party.EndDate,
					CustomerID:     customer.ID,
					PartyID:        party.ID,
					ClubID:         club.ID,
					PartyProductID: item.PartyProduct.ID,
					InvoiceID:      &invoiceID,
					ClubName:       club.Name,
					PartyName:      party.Name,
					CustomerName:   customer.Name(),
					Status:         "PROCESSING",
					Price: types.Price{
						Value:          types.PartyPPrice(item.PartyProduct, &types.Promotion{}, party, club),
						CurrentIsoCode: "BRL",
					},
					Product: item.VoucherProduct(),
					Type:    "NORMAL",
				}
				vouchers = append(vouchers, voucher)
			}
		}
		invoiceRepo, ok := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
		if !ok {
			err := errors.New("assert")
			return ticketInvoice, vouchers, err
		}
		err := invoiceRepo.Collection.Insert(ticketInvoice)
		if err != nil {
			return ticketInvoice, vouchers, err
		}
		vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
		if !ok {
			err := errors.New("assert")
			return ticketInvoice, vouchers, err
		}
		for _, voucher := range vouchers {
			err = vouchersCollection.Collection.Insert(voucher)
			if err != nil {
				return ticketInvoice, vouchers, err
			}
		}
	}
	log.Println(">>>> returning empoty ticket invoice")
	return ticketInvoice, vouchers, nil
}

// CreateInvoicesAndVouchersBuy TODO: NEEDS COMMENT INFO
func CreateInvoicesAndVouchersBuy(ctx context.Context,
	club types.Club,
	party types.Party,
	customer types.Customer,
	products types.BuyPostList,
	logg types.Log,
) (types.Invoices, []types.Voucher, error) {
	invoices := types.Invoices{}
	vouchers := []types.Voucher{}
	drinksInvoice, dvouchers, err := CreatePendingDrinksInvoiceAndVouchersBuy(ctx, club, party, customer, products, logg)
	if err != nil {
		return invoices, vouchers, nil
	}
	if drinksInvoice.Valid() {
		invoices.Drinks = &drinksInvoice
		vouchers = append(vouchers, dvouchers...)
	}
	ticketInvoice, tvouchers, err := CreatePendingTicketsInvoiceAndVouchersBuy(ctx, club, party, customer, products, logg)
	if err != nil {
		return invoices, vouchers, nil
	}
	if ticketInvoice.Valid() {
		invoices.Tickets = &ticketInvoice
		vouchers = append(vouchers, tvouchers...)
	}
	return invoices, vouchers, nil
}
