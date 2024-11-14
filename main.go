package main

import (
	"api-rest/db"
	"api-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //Configura um servidor HTTP
	routes.RegisterRoutes(server)
	server.Run(":8080") //Executando o servidor na porta localhost:8080
}
