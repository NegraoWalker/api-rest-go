package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() //Configura um servidor HTTP

	server.GET("/events", getEvents)

	server.Run(":8080") //Executando o servidor na porta localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) //Retorno da solicitação GET
}
