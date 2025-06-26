package main

import (
	"library-api/config"
	"library-api/controllers"
	"library-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	bookCollection := config.DB.Collection("books")
	controllers.SetCollection(bookCollection)

	r := gin.Default()
	routes.BookRoutes(r)
	r.Run(":8080")
}
