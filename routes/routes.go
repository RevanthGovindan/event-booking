package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", createEvent)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
}
