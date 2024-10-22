package main

import (
    "github.com/gin-gonic/gin"
	"project2/routes"
	"project2/database"
)

func main () {
	r := gin.Default()

	database.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8080")
}