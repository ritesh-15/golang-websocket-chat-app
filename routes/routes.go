package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/websocket-advanced/ws"
)

func InitRoutes(app *gin.Engine, wsHandler *ws.WsHandler) {
	app.POST("/ws/create-room", wsHandler.CreateRoom)

	app.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
}
