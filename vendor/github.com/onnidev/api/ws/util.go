package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/onnidev/api/infra"
)

//  TODO: NEEDS COMMENT INFO
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Logger TODO: NEEDS COMMENT INFO
var Logger = infra.NewLogger("WEBSOCKET")

// WebSocketMsg sdfkjnsdf
type WebSocketMsg struct {
	Type  string      `json:"type" bson:"type"`
	Kind  *string     `json:"kind" bson:"kind"`
	Token *string     `json:"token,omitempty" bson:"token,omitempty"`
	Data  interface{} `json:"data,omitempty" bson:"data,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) ReadPump(channel string) {
	log.Printf("Listening to Websocket messages at %s", channel)
	defer func() {
		DashboardHub.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.ws.ReadMessage()
		Logger.Print("just received : ", string(message))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				Logger.Printf("error: %v", err)
			}
			break
		}
		var result WebSocketMsg
		_ = json.Unmarshal(message, &result)
		if result.Kind != nil {
			if *result.Kind == "BROADCAST" {
				Logger.Print("before publishing TO REDIS " + channel)
				Publish(channel, string(message))
				continue
			}
		}
		sendWebSocketMsg(c, channel, message)
		log.Println(result)
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) WritePump(channelname string) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			sendWebSocketMsg(c, channelname, message)

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func sendWebSocketMsg(c *connection, channelname string, message []byte) {
	Logger.Printf("responding websocket mensage %s", string(message))
	var result WebSocketMsg
	err := json.Unmarshal(message, &result)
	if err != nil {
		c.ws.Close()
		return
	}
	switch {
	case strings.Contains(channelname, "staff"):
		HandleStaff(c, channelname, result)
		break
	case strings.Contains(channelname, "dashboard"):
		HandleDashboard(c, channelname, result)
		break
	case strings.Contains(channelname, "app"):
		HandleApp(c, channelname, result)
		break
	}
}
