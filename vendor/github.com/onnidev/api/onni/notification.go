package onni

import (
	"context"
	"errors"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetNotification TODO: NEEDS COMMENT INFO
func GetNotification(ctx context.Context, id string) (types.Notification, error) {
	repo, ok := ctx.Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	if !ok {
		err := errors.New("bug")
		return types.Notification{}, err
	}
	notification, err := repo.GetByID(id)
	if err != nil {
		return types.Notification{}, err
	}
	return notification, nil
}

// PublishAndroidNotification TODO: NEEDS COMMENT INFO
func PublishAndroidNotification(ctx context.Context, notification types.Notification, ids []bson.ObjectId) error {
	pushRegistryRepo, ok := ctx.Value(middlewares.PushRegistryRepoKey).(interfaces.PushRegistryRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	androidIDS, err := pushRegistryRepo.GetMobileTokenFromCustomers(ids, "ANDROID")
	if err != nil {
		return err
	}
	androidIDS = append(androidIDS, "sdjfnsdkfn")
	go AndroidPushDevelopmentPushNotification(notification, androidIDS)
	return nil
}

// PublishIOSNotification TODO: NEEDS COMMENT INFO
func PublishIOSNotification(ctx context.Context, notification types.Notification, ids []bson.ObjectId) error {
	pushRegistryRepo, ok := ctx.Value(middlewares.PushRegistryRepoKey).(interfaces.PushRegistryRepo)
	if !ok {
		err := errors.New("bug")
		return err
	}
	iosIDS, err := pushRegistryRepo.GetMobileTokenFromCustomers(ids, "IOS")
	if err != nil {
		return err
	}
	if viper.GetString("env") == "production" || viper.GetString("env") == "dev" {
		go IosPushProductionPushNotification(notification, iosIDS)
		return nil
	}
	go IosPushDevelopmentPushNotification(ctx, notification, iosIDS)
	return nil
}

// PublishNotificationForPartyProductBuyers TODO: NEEDS COMMENT INFO
func PublishNotificationForPartyProductBuyers(ctx context.Context, notification types.Notification, partyProductID string) (types.Notification, error) {
	if notification.Status != "DRAFT" {
		err := errors.New("only draft notification can be published")
		return types.Notification{}, err
	}
	ids, err := PartyUniqueActiveCustomersBySpecificPartyProductID(ctx, notification.PartyID.Hex(), bson.ObjectIdHex(partyProductID).Hex())
	if err != nil {
		return types.Notification{}, err
	}
	err = PublishAndroidNotification(ctx, notification, ids)
	if err != nil {
		return types.Notification{}, err
	}
	err = PublishIOSNotification(ctx, notification, ids)
	if err != nil {
		return types.Notification{}, err
	}
	patchedNotification, err := MarkNotifcationAsSentBy(ctx, notification)
	if err != nil {
		return types.Notification{}, err
	}
	return patchedNotification, nil
}

// PublishNotification TODO: NEEDS COMMENT INFO
func PublishNotification(ctx context.Context, notification types.Notification) (types.Notification, error) {
	if notification.Status != "DRAFT" {
		err := errors.New("only draft notification can be published")
		return types.Notification{}, err
	}
	ids, err := PartyUniqueActiveCustomers(ctx, notification.PartyID.Hex())
	if err != nil {
		return types.Notification{}, err
	}
	err = PublishAndroidNotification(ctx, notification, ids)
	if err != nil {
		return types.Notification{}, err
	}
	err = PublishIOSNotification(ctx, notification, ids)
	if err != nil {
		return types.Notification{}, err
	}
	patchedNotification, err := MarkNotifcationAsSentBy(ctx, notification)
	if err != nil {
		return types.Notification{}, err
	}
	return patchedNotification, nil
}

// MarkNotifcationAsSentBy TODO: NEEDS COMMENT INFO
func MarkNotifcationAsSentBy(ctx context.Context, notification types.Notification) (types.Notification, error) {
	userClub, ok := ctx.Value(middlewares.UserClubKey).(types.UserClub)
	if !ok {
		err := errors.New("bug")
		return types.Notification{}, err
	}
	now := types.Timestamp(time.Now())
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"updateDate":    &now,
			"updatedBy":     &userClub.ID,
			"updatedByName": &userClub.Name,
			"status":        "ACTIVE",
		}},
		ReturnNew: true,
	}
	repo, ok := ctx.Value(middlewares.NotificationRepoKey).(interfaces.NotificationRepo)
	if !ok {
		err := errors.New("bug")
		return types.Notification{}, err
	}
	var patchedNotification types.Notification
	_, err := repo.Collection.Find(bson.M{"_id": notification.ID}).Apply(change, &patchedNotification)
	if err != nil {
		return types.Notification{}, err
	}
	return patchedNotification, nil
}
