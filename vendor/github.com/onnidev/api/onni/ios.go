package onni

import (
	"context"
	"log"

	fcm "github.com/NaySoftware/go-fcm"
	"github.com/fatih/structs"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

// AndroidPushDevelopmentPushNotification sdfkjn
func AndroidPushDevelopmentPushNotification(notification types.Notification, ids []string) (*fcm.FcmResponseStatus, error) {
	var response *fcm.FcmResponseStatus
	if len(ids) != 0 {
		data := structs.Map(&notification)
		//token 14/05 nearly
		c := fcm.NewFcmClient("AAAAUi0c6qo:APA91bHujDnAJD1nH5S1paWZUXV6HQz-Y8p1C-AcLaMnFRGC2cNDr80RKEpglAszkWGFDXrWZ0OYdjoIH92NPoGiTlZufJ4mmA87BIwlxt5l9-vhOaj27s_2aNzPC1QFm7IMppE8dUcp")
		c.NewFcmRegIdsMsg([]string{ids[0]}, data)
		if len(ids) >= 2 {
			c.AppendDevices(ids[1:])
		}
		response, err := c.Send()
		if err != nil {
			return response, err
		}
		log.Println("*****Android*****")
		log.Println(response.StatusCode)
		log.Println(response.Err)
	}
	return response, nil
}

// IosPushDevelopmentPushNotification sdfkjn
func IosPushDevelopmentPushNotification(ctx context.Context, notification types.Notification, ids []string) error {
	client, err := iosPushNotificationClient()
	if err != nil {
		return err
	}
	for _, id := range ids {
		not, err := notification.GetIOSNotification(id)
		if err != nil {
			return err
		}
		resp, err := client.Development().Push(not)
		if err != nil {
			return err
		}
		log.Println("*******IOS DEVELOPMENT*********")
		log.Println(resp)
	}
	return nil
}

// IosPushProductionPushNotification sdfkjn
func IosPushProductionPushNotification(notification types.Notification, ids []string) error {
	client, err := iosPushNotificationClient()
	if err != nil {
		return err
	}
	for _, id := range ids {
		not, err := notification.GetIOSNotification(id)
		if err != nil {
			return err
		}
		resp, err := client.Production().Push(not)
		if err != nil {
			return err
		}
		log.Println("*******IOS PRODUCTION*********")
		log.Println("sending for token ID", id)
		log.Println(resp)
	}
	return nil
}

func iosPushNotificationClient() (*apns2.Client, error) {
	var client apns2.Client
	css, err := shared.Asset("public/key.p8")
	if err != nil {
		return &client, err
	}
	//
	authKey, err := token.AuthKeyFromBytes(css)
	if err != nil {
		return &client, err
	}
	token := &token.Token{
		AuthKey: authKey,
		KeyID:   "JHUG7GDZ9W",
		TeamID:  "LLC5JE67PU",
	}

	return apns2.NewTokenClient(token), nil
}
