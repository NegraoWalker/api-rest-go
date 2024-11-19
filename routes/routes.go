package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)       //findAll
	server.GET("/events/:id", getEvent)    //findById
	server.POST("/events", createEvent)    //insert
	server.PUT("/events/:id", updateEvent) //update
}
