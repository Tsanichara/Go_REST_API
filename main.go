package main

import (
	"github.com/Tsanichara/Go_REST_API/db"
	"github.com/Tsanichara/Go_REST_API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
