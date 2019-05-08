package interfaces

import (
	"fmt"
	"log"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// VouchersRepo is a struc that hold a mongo collection
type VouchersRepo struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// NewVoucherCollection creates a new types.VoucherDAO
func NewVoucherCollection(store *infra.MongoStore) (VouchersRepo, error) {
	repo := VouchersRepo{
		Session:    store.Session,
		Collection: store.Session.DB(store.Database).C("voucher"),
	}
	return repo, nil
}

// PromotionPurchasedbyCustomer Insert the user on a database
func (c *VouchersRepo) PromotionPurchasedbyCustomer(promotionID, customerID string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"promotionId": bson.ObjectIdHex(promotionID),
			"customerId":  bson.ObjectIdHex(customerID),
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return result, err
	}
	vouchers := []types.CompleteVoucher{}
	for _, voucher := range result {
		log.Println("###############", voucher.Status)
		if voucher.Status == "AVAILABLE" || voucher.Status == "USED" {
			vouchers = append(vouchers, voucher)
		}
	}
	return vouchers, nil
}

// PromotionCountPurchasedbyCustomer Insert the user on a database
func (c *VouchersRepo) PromotionCountPurchasedbyCustomer(promotionID, customerID string) (int, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"promotionId": bson.ObjectIdHex(promotionID),
			"customerId":  bson.ObjectIdHex(customerID),
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return 0, err
	}
	vouchers := []types.CompleteVoucher{}
	for _, voucher := range result {
		log.Println("###############", voucher.Status)
		if voucher.Status == "AVAILABLE" || voucher.Status == "USED" {
			vouchers = append(vouchers, voucher)
		}
	}
	return len(vouchers), nil
}

// GetActualAppVoucherByCustomer Insert the user on a database
func (c *VouchersRepo) GetActualAppVoucherByCustomer(token string) ([]types.AppCompleteVoucher, error) {
	result := []types.AppCompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(token),
			"endDate":    bson.M{"$gt": time.Now()},
			"status":     "AVAILABLE",
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.AppCompleteVoucher{}, err
	}
	return result, nil
}

// GetAccountablesVouchersByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) GetErroredVouchersByPartyAndCustomer(party, customer string) ([]types.Voucher, error) {
	result := []types.Voucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(customer),
			"partyId":    bson.ObjectIdHex(party),
			// "type":       bson.M{"$or": []string{"NORMAL", "PROMOTION"}},
			"$or": []interface{}{
				bson.M{"type": "NORMAL"},
				bson.M{"type": "PROMOTION"},
			},
			"status": "ERROR",
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.Voucher{}, err
	}
	return result, nil
}

// GetAccountablesVouchersByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) GetCompleteVouchersByPartyAndCustomer(party, customer string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(customer),
			"partyId":    bson.ObjectIdHex(party),
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	log.Println("###### ", len(result))
	return result, nil
}

// GetAccountablesVouchersByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) GetAccountablesVouchersByPartyAndCustomer(party, customer string) ([]types.Voucher, error) {
	result := []types.Voucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(customer),
			"partyId":    bson.ObjectIdHex(party),
			"$and": []bson.M{
				{
					"$or": []interface{}{
						bson.M{"status": "AVAILABLE"},
						bson.M{"status": "USED"},
					},
				},
				{
					"$or": []interface{}{
						bson.M{"type": "NORMAL"},
						bson.M{"type": "PROMOTION"},
					},
				},
			},
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.Voucher{}, err
	}
	log.Println("###### ", len(result))
	return result, nil
}

// GetSimpleActualVoucherByCustomer Insert the user on a database
func (c *VouchersRepo) GetSimpleActualVoucherByCustomer(token string) ([]types.Voucher, error) {
	result := []types.Voucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(token),
			"endDate":    bson.M{"$gt": time.Now()},
			"status":     "AVAILABLE",
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.Voucher{}, err
	}
	return result, nil
}

// GetActualVoucherByCustomer Insert the user on a database
func (c *VouchersRepo) GetActualVoucherByCustomer(token string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(token),
			"endDate":    bson.M{"$gt": time.Now()},
			"status":     "AVAILABLE",
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// GetActualVoucherByParty Insert the user on a database
func (c *VouchersRepo) GetActualVoucherByParty(token string) ([]types.Voucher, error) {
	result := []types.Voucher{}
	err := c.Collection.Pipe(
		[]bson.M{
			bson.M{"$match": bson.M{
				"partyId": bson.ObjectIdHex(token),
			}}}).All(&result)
	if err != nil {
		return []types.Voucher{}, err
	}
	return result, nil
}

// ReadByUserClubID Insert the user on a database
func (c *VouchersRepo) ReadByUserClubID(userClubID string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"status":               "USED",
			"startDate":            bson.M{"$gt": time.Now().Add(-time.Duration(24*10) * time.Hour)},
			"voucherUseUserClubId": bson.ObjectIdHex(userClubID),
		},
		bson.M{"voucherUseDate": -1}, true)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// ReadByUserClubIDandClubID Insert the user on a database
func (c *VouchersRepo) ReadByUserClubIDandClubID(userClubID, clubID string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"status":               "USED",
			"startDate":            bson.M{"$gt": time.Now().Add(-time.Duration(24*10) * time.Hour)},
			"clubId":               bson.ObjectIdHex(clubID),
			"voucherUseUserClubId": bson.ObjectIdHex(userClubID),
		},
		bson.M{"voucherUseDate": -1}, true)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

func makeComplete(match bson.M, sort bson.M, club bool) []bson.M {
	all := []bson.M{
		bson.M{"$match": match},
		bson.M{
			"$lookup": bson.M{
				"from":         "customer",
				"localField":   "customerId",
				"foreignField": "_id",
				"as":           "customer"}},
		bson.M{"$unwind": bson.M{
			"path": "$customer",
			"preserveNullAndEmptyArrays": true}},
		bson.M{
			"$lookup": bson.M{
				"from":         "partyProduct",
				"localField":   "partyProductId",
				"foreignField": "_id",
				"as":           "partyProduct"}},
		bson.M{"$unwind": bson.M{
			"path": "$partyProduct",
			"preserveNullAndEmptyArrays": true}},
		bson.M{
			"$lookup": bson.M{
				"from":         "userClub",
				"localField":   "responsibleUserClubId",
				"foreignField": "_id",
				"as":           "responsable"}},
		bson.M{"$unwind": bson.M{
			"path": "$responsable",
			"preserveNullAndEmptyArrays": true}},
		bson.M{
			"$lookup": bson.M{
				"from":         "party",
				"localField":   "partyId",
				"foreignField": "_id",
				"as":           "party"}},
		bson.M{"$unwind": bson.M{
			"path": "$party",
			"preserveNullAndEmptyArrays": true}},
		bson.M{"$sort": sort},
	}
	if club {
		all = append(all, bson.M{
			"$lookup": bson.M{
				"from":         "club",
				"localField":   "party.clubId",
				"foreignField": "_id",
				"as":           "club"}})
		all = append(all, bson.M{"$unwind": bson.M{
			"path": "$club",
			"preserveNullAndEmptyArrays": true}})

		all = append(all, bson.M{
			"$lookup": bson.M{
				"from":         "partyProduct",
				"localField":   "partyProductId",
				"foreignField": "_id",
				"as":           "partyProduct"}})
		all = append(all, bson.M{"$unwind": bson.M{
			"path": "$partyProduct",
			"preserveNullAndEmptyArrays": true}})

	}
	return all
}

// GetByCustomer Insert the user on a database
func (c *VouchersRepo) GetByCustomer(id string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"customerId": bson.ObjectIdHex(id)},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// GetByClub Insert the user on a database
func (c *VouchersRepo) GetByClub(token string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"clubId": bson.ObjectIdHex(token)},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// GetByParty Insert the user on a database
func (c *VouchersRepo) GetOnGoingByParty(token string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"partyId": bson.ObjectIdHex(token),
			"startDate": bson.M{"$gt": time.Now()}},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// GetByParty Insert the user on a database
func (c *VouchersRepo) GetByParty(token string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"partyId": bson.ObjectIdHex(token)},
		bson.M{"creationDate": 1},
		false)).AllowDiskUse().All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// CountByParty Insert the user on a database
func (c *VouchersRepo) CountByParty(token string) (int, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"partyId": bson.ObjectIdHex(token)},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return len(result), err
	}
	return len(result), nil
}

// FutureByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) FutureByPartyAndCustomer(token, useriD string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"partyId":    bson.ObjectIdHex(token),
			"customerId": bson.ObjectIdHex(useriD),
			"startDate":  bson.M{"$gt": time.Now()},
		},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FutureByCustomer Insert the user on a database
func (c *VouchersRepo) FutureByCustomer(useriD string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(useriD),
			"startDate":  bson.M{"$gt": time.Now()},
		},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// ByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) ByPartyAndCustomer(token, useriD string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"partyId":    bson.ObjectIdHex(token),
			"customerId": bson.ObjectIdHex(useriD),
		},
		bson.M{"creationDate": 1},
		false)).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// CountByPartyAndCustomer Insert the user on a database
func (c *VouchersRepo) CountByPartyAndCustomer(token, useriD string) (int, error) {
	result, err := c.ByPartyAndCustomer(token, useriD)
	if err != nil {
		return len(result), err
	}
	return len(result), nil
}

// GetSimpleByID Insert the user on a database
func (c *VouchersRepo) GetSimpleByID(id string) (types.Voucher, error) {
	result := types.Voucher{}
	if bson.IsObjectIdHex(id) {
		err := c.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			return types.Voucher{}, err
		}
		return result, nil
	}
	return types.Voucher{}, fmt.Errorf("not a valid object id")
}

// GetByID Insert the user on a database
func (c *VouchersRepo) GetByID(id string) (types.CompleteVoucher, error) {
	result := types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"creationDate": 1},
		false)).One(&result)
	if err != nil {
		return types.CompleteVoucher{}, err
	}
	return result, nil
}

// Insert Insert the user on a database
func (c *VouchersRepo) Insert(voucher types.Voucher) error {
	return c.Collection.Insert(voucher)
}

// CancelVoucher Insert the user on a database
func (c *VouchersRepo) CancelVoucher(voucher types.Voucher, userClub types.UserClub) (types.Voucher, error) {
	result := types.Voucher{}
	if bson.IsObjectIdHex(voucher.ID.Hex()) {
		updateDate := types.Timestamp(time.Now())
		log.Println("using")
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"status":                 "CANCELED",
					"updateDate":             &updateDate,
					"voucherUseDate":         &updateDate,
					"voucherUseUserClubId":   &userClub.ID,
					"voucherUseUserClubName": &userClub.Name,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(voucher.ID.Hex())}).Apply(change, &result)
		if err != nil {
			return types.Voucher{}, err
		}
		return result, nil
	}
	return types.Voucher{}, fmt.Errorf("not a valid object id")
}

// UseVoucher Insert the user on a database
func (c *VouchersRepo) UseVoucher(voucher types.Voucher, userClub types.UserClub) (types.Voucher, error) {
	result := types.Voucher{}
	if bson.IsObjectIdHex(voucher.ID.Hex()) {
		updateDate := types.Timestamp(time.Now())
		log.Println("using")
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"status":                 "USED",
					"updateDate":             &updateDate,
					"voucherUseDate":         &updateDate,
					"voucherUseUserClubId":   &userClub.ID,
					"voucherUseUserClubName": &userClub.Name,
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(voucher.ID.Hex())}).Apply(change, &result)
		if err != nil {
			return types.Voucher{}, err
		}
		return result, nil
	}
	return types.Voucher{}, fmt.Errorf("not a valid object id")
}

// GetByIDAndTransfer Insert the user on a database
func (c *VouchersRepo) GetByIDAndTransfer(id string) (types.Voucher, error) {
	result := types.Voucher{}
	if bson.IsObjectIdHex(id) {
		now := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{"status": "TRANSFERED", "updateDate": &now}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.Voucher{}, err
		}
		return result, nil
	}
	return types.Voucher{}, fmt.Errorf("not a valid object id")
}

// List lists a token collection
func (c *VouchersRepo) List() ([]types.Voucher, error) {
	result := []types.Voucher{}
	err := c.Collection.Find(bson.M{}).Limit(10).All(&result)
	log.Println(result[0])
	return result, err
}

// GetAllCustomerVouchers Insert the user on a database
func (c *VouchersRepo) GetAllCustomerVouchers(token string) ([]types.CompleteVoucher, error) {
	result := []types.CompleteVoucher{}
	err := c.Collection.Pipe(makeComplete(
		bson.M{
			"customerId": bson.ObjectIdHex(token),
		},
		bson.M{"startDate": -1},
		false)).All(&result)
	if err != nil {
		return []types.CompleteVoucher{}, err
	}
	return result, nil
}

// SetInvitedNameOnVoucher Insert the user on a database
func (c *VouchersRepo) SetInvitedNameOnVoucher(id, name string) (types.Voucher, error) {
	result := types.Voucher{}
	if bson.IsObjectIdHex(id) {
		now := types.Timestamp(time.Now())
		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"customerName": name,
					"updateDate":   &now,
					"status":       "AVAILABLE",
				}},
			ReturnNew: true,
		}
		_, err := c.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
		if err != nil {
			return types.Voucher{}, err
		}
		return result, nil
	}
	return types.Voucher{}, fmt.Errorf("not a valid object id")
}
