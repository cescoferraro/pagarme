package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// HandleStaff TODO: NEEDS COMMENT INFO
func HandleStaff(c *connection, channelname string, message WebSocketMsg) {
	Logger.Print("handling staff command")
	if message.Type == "VOUCHERS" {
		err := hey(c, channelname, message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return

	}
	j, _ := json.Marshal(message)
	k, _ := json.MarshalIndent(message, "", "     ")
	log.Println(string(k))
	err := c.write(websocket.TextMessage, j)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
