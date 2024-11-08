package main

import (
	"api-rest/db"
	"api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //Configura um servidor HTTP

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") //Executando o servidor na porta localhost:8080
}

func getEvents(context *gin.Context) {
	events, error := models.GetAllEvents()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events) //Retorno da solicitação GET
}

func createEvent(context *gin.Context) {
	var event models.Event
	error := context.ShouldBindJSON(&event)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.Id = 1
	event.UserId = 1

	error = event.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
