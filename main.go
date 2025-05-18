package main

import (
	"github.com/RevanthGovindan/event-booking/db"
	"github.com/RevanthGovindan/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
