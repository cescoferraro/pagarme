package ws

// Hub TODO: NEEDS COMMENT INFO
type Hub struct {
	// Registered connections.
	connections map[*connection]bool
	// Inbound messages from the connections.
	Broadcast chan []byte
	// Register requests from the connections.
	register chan *connection
	// Unregister requests from connections.
	unregister chan *connection
}

// DashboardHub TODO: NEEDS COMMENT INFO
var DashboardHub = Hub{
	Broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

// AppHub TODO: NEEDS COMMENT INFO
var AppHub = Hub{
	Broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

// StaffHub TODO: NEEDS COMMENT INFO
var StaffHub = Hub{
	Broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

// Run TODO: NEEDS COMMENT INFO
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.Broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}
