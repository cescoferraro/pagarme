package ws

import (
	"log"
	"net/http"
)

// Handler TODO: NEEDS COMMENT INFO
func Handler(channel string, hub Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c := &connection{send: make(chan []byte, 256), ws: ws}
		log.Printf("Connecting %s Websocket", channel)
		hub.register <- c
		go c.WritePump(channel)
		c.ReadPump(channel)
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, true)
	}
}
