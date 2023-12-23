package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn           *websocket.Conn
	RoomID         string `json:"room_id"`
	ID             string `json:"id"`
	Username       string `json:"username"`
	MessageChannel chan *Message
}

type ClientMessageTypes struct {
	JOINED_ROOM string
	LEAVE_ROOM  string
	MESSAGE     string
	ERROR       string
}

var MessageType = ClientMessageTypes{
	JOINED_ROOM: "JOINED_ROOM",
	LEAVE_ROOM:  "LEAVE_ROOM",
	MESSAGE:     "MESSAGE",
	ERROR:       "ERROR",
}

type Message struct {
	RoomID   string `json:"roomId"`
	ClientID string `json:"clientId"`
	Message  string `json:"message"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

func (client *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.UnRegister <- client
		client.Conn.Close()
	}()

	for {

		_, data, err := client.Conn.ReadMessage()

		if err != nil {
			log.Println("socket error: ", err)
			return
		}

		message := &Message{
			RoomID:   client.RoomID,
			ClientID: client.ID,
			Message:  string(data),
			Type:     MessageType.MESSAGE,
			Username: client.Username,
		}

		hub.Broadcast <- message
	}
}

func (client *Client) SendMessage(hub *Hub) {
	defer func() {
		hub.UnRegister <- client
		client.Conn.Close()
	}()

	for {
		message, ok := <-client.MessageChannel
		if !ok {
			log.Println("unable to send message")
			return
		}

		if err := client.Conn.WriteJSON(message); err != nil {
			log.Printf("writting message error: %v", err)
			return
		}
	}
}
