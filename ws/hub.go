package ws

import "log"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room `json:"rooms"`
	Register   chan *Client
	UnRegister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan *Message, 10),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			// validate the room id first then join that room
			if _, ok := hub.Rooms[client.RoomID]; ok {
				// join the room
				room := hub.Rooms[client.RoomID]
				room.Clients[client.ID] = client

				// broadcase the message that user joined a the room
				message := &Message{
					RoomID:   client.RoomID,
					ClientID: client.ID,
					Message:  "user joined a room",
					Type:     MessageType.JOINED_ROOM,
					Username: client.Username,
				}

				hub.Broadcast <- message
			} else {
				message := &Message{
					RoomID:   client.RoomID,
					ClientID: client.ID,
					Message:  "room not found",
					Type:     MessageType.ERROR,
					Username: client.Username,
				}

				hub.Broadcast <- message
			}

		case client := <-hub.UnRegister:
			// check if room is exits or not
			if _, ok := hub.Rooms[client.RoomID]; ok {
				// check if user extis or not
				room := hub.Rooms[client.RoomID]

				if _, ok := room.Clients[client.RoomID]; ok {
					// broadcase message that user leaves
					message := &Message{
						RoomID:   client.RoomID,
						ClientID: client.ID,
						Message:  "user leaves a room",
						Type:     MessageType.LEAVE_ROOM,
						Username: client.Username,
					}

					hub.Broadcast <- message

					// close message channel for client
					close(room.Clients[client.ID].MessageChannel)

					// remove user from clients
					delete(room.Clients, client.ID)
				}
			}

		case message := <-hub.Broadcast:
			// check if room is exits or not
			if _, ok := hub.Rooms[message.RoomID]; ok {
				// broadcase message to all connected clients
				room := hub.Rooms[message.RoomID]

				log.Printf("total clients in room %v \n", len(room.Clients))

				for _, client := range room.Clients {
					// send the message
					client.MessageChannel <- message
				}
			}
		}
	}
}
