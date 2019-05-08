package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

// HandleApp TODO: NEEDS COMMENT INFO
func HandleApp(c *connection, channelname string, message WebSocketMsg) {
	err := c.write(websocket.TextMessage, []byte(message.Type))
	if err != nil {
		log.Println(err.Error())
		return
	}
}
