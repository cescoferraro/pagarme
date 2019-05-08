package ws

import (
	"encoding/json"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// HandleDashboard TODO: NEEDS COMMENT INFO
func HandleDashboard(c *connection, channelname string, message WebSocketMsg) {

	Logger.Print("handling dashboard command")
	Logger.Print("handling dashboard command")
	Logger.Print("handling dashboard command")
	if message.Type == "VOUCHERS" {
		err := hey(c, channelname, message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}
	j, _ := json.Marshal(message)
	err := c.write(websocket.TextMessage, j)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func hey(c *connection, channelname string, message WebSocketMsg) error {
	var hey types.MyClaimsType
	_, err := jwt.ParseWithClaims(
		*message.Token,
		&hey,
		shared.JWTAuth.Options.ValidationKeyGetter)
	if err != nil {
		c.ws.Close()
		return err
	}

	db, err := infra.Cloner()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	collection, err := interfaces.NewUserClubCollection(db)
	if err != nil {
		return err
	}
	partyCollection, err := interfaces.NewPartiesCollection(db)
	if err != nil {
		return err
	}
	clubsRepo, err := interfaces.NewClubsCollection(db)
	if err != nil {
		return err
	}
	voucherCollection, err := interfaces.NewVoucherCollection(db)
	if err != nil {
		return err
	}
	reports, err := onni.WSGetPartyReport(hey.ClientID,
		voucherCollection,
		partyCollection,
		clubsRepo,
		collection, true)
	if err != nil {
		return err
	}
	msg := WebSocketMsg{Type: "VOUCHERS_RESPONSE", Data: &reports}
	j, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	k, _ := json.MarshalIndent(msg, "", "     ")
	log.Println(string(k))
	err = c.write(websocket.TextMessage, j)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil

}
