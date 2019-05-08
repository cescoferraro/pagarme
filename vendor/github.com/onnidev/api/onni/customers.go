package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// CustomerFromToken TODO: NEEDS COMMENT INFO
func CustomerFromToken(ctx context.Context) (types.Customer, error) {
	customer, ok := ctx.Value(middlewares.CustomersKey).(types.Customer)
	if !ok {
		err := errors.New("bug")
		return customer, err
	}
	return customer, nil
}

// PartyUniqueActiveCustomers asdfnjkas
func PartyUniqueActiveCustomers(ctx context.Context, partyID string) ([]bson.ObjectId, error) {
	var customers []bson.ObjectId
	ref, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return customers, err
	}
	vouchers, err := ref.GetByParty(partyID)
	if err != nil {
		return customers, err
	}
	for _, voucher := range vouchers {
		if voucher.Customer != nil {
			if !shared.ContainsObjectID(customers, voucher.Customer.ID) {
				customers = append(customers, voucher.Customer.ID)
			}
		}
	}
	return customers, nil
}

// PartyUniqueActiveCustomersBySpecificPartyProductID asdfnjkas
func PartyUniqueActiveCustomersBySpecificPartyProductID(ctx context.Context, partyID string, partyProductID string) ([]bson.ObjectId, error) {
	var customers []bson.ObjectId
	ref, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return customers, err
	}
	vouchers, err := ref.GetByParty(partyID)
	if err != nil {
		return customers, err
	}
	for _, voucher := range vouchers {
		if voucher.PartyProductID.Hex() == partyProductID {
			if voucher.Customer != nil {
				if !shared.ContainsObjectID(customers, voucher.Customer.ID) {
					customers = append(customers, voucher.Customer.ID)
				}
			}
		}
	}
	return customers, nil
}

// RemoveInactiveClubs sdkjfnsdf
func RemoveInactiveClubs(ctx context.Context, customer types.Customer) types.Customer {
	repo := ctx.Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	var active []bson.ObjectId
	for _, clubID := range customer.FavoriteClubs {
		club, err := repo.GetByID(clubID.Hex())
		if err == nil {
			if club.Status == "ACTIVE" {
				active = append(active, club.ID)
			}
		}
	}
	customer.FavoriteClubs = active
	return customer
}

// Customer sdkjfnsdf
func Customer(ctx context.Context, id string) (types.Customer, error) {
	repo, ok := ctx.Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	if !ok {
		err := errors.New("bug")
		return types.Customer{}, err
	}
	customer, err := repo.GetByID(id)
	if err == nil {
		return customer, err
	}
	return customer, nil
}
