package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/beingmohit/go-secretes/src/api"
	"github.com/beingmohit/go-secretes/src/database"
)

func main() {
	router := gin.Default()

	databaseClient := database.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	
	apiHandler := api.NewAPIHandler(router.Group("api"), databaseClient)
	apiHandler.RegisterRoutes()
	
	router.Run()
}