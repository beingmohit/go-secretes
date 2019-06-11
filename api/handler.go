package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/beingmohit/go-secretes/database"
	"github.com/beingmohit/go-secretes/api/secrets"
)

type APIHandler struct {
	router *gin.RouterGroup
	databaseClient *database.Client
}

func NewAPIHandler(router *gin.RouterGroup, databaseClient *database.Client) APIHandler {
	apiHandler := APIHandler{router: router, databaseClient: databaseClient}
	return apiHandler
}

func (apiHandler *APIHandler) RegisterRoutes() {
	fmt.Println("Registering api routes")

	secretsController := secrets.NewController(apiHandler.databaseClient)
	
	secretRoutes := apiHandler.router.Group("/secret")
	secretRoutes.POST("/", secretsController.Store)
	secretRoutes.GET("/:hash", secretsController.Show)
}