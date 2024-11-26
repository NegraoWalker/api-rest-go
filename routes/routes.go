package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)          //findAll events
	server.GET("/events/:id", getEvent)       //findById event
	server.POST("/events", createEvent)       //insert event
	server.PUT("/events/:id", updateEvent)    //update event
	server.DELETE("/events/:id", deleteEvent) //delete event
	server.POST("/signup", signup)            //insert user
	server.POST("/login", login)
}
