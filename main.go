package main

import (
	dbconnection "OrderTracker/dbConnection"
	"OrderTracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbconnection.InitDB()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
