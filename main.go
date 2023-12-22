package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ritesh-15/websocket-advanced/config"
	"github.com/ritesh-15/websocket-advanced/database"
	"github.com/ritesh-15/websocket-advanced/routes"
	"github.com/ritesh-15/websocket-advanced/ws"
)

func init() {
	config.LoadEnv()
	database.InitDatabase()
}

func main() {
	app := gin.New()

	app.Use(gin.Logger())

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	routes.InitRoutes(app, wsHandler)

	app.Run(":9000")
}
