package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ritesh-15/websocket-advanced/utils"
)

type WsHandler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *WsHandler {
	return &WsHandler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type CreateRoomReq struct {
	Name string `json:"name"`
}

func (h *WsHandler) CreateRoom(ctx *gin.Context) {
	var req CreateRoomReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.NewApiError(false, "unprocessable entity"))
		return
	}

	roomId := uuid.New().String()

	// create room in database

	// TODO: instead of storing into memory store the room in redis like datase
	h.hub.Rooms[roomId] = &Room{
		ID:      roomId,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	ctx.JSON(http.StatusOK, utils.NewApiResponse(true, "room created successfully", h.hub.Rooms[roomId]))
}

func (h *WsHandler) JoinRoom(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Fatal("Error while upgrading conn: ", err)
		return
	}

	// get the room details
	roomId := ctx.Param("roomId")
	username := ctx.Query("username")

	// we get this id from database or authenticated service but for now just generate random client id
	clientId := uuid.New().String()

	// create client
	client := &Client{
		ID:             clientId,
		Conn:           conn,
		RoomID:         roomId,
		Username:       username,
		MessageChannel: make(chan *Message, 10),
	}

	// register client with hub
	h.hub.Register <- client

	go client.SendMessage(h.hub)
	go client.ReadMessage(h.hub)
}
