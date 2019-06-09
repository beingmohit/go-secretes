package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/beingmohit/go-secretes/src/database"
	"github.com/beingmohit/go-secretes/src/api/secrets"
)

type APIHandler struct {
	router *gin.RouterGroup
	databaseClient *database.Client
}

func CreateAPIHandler(router *gin.RouterGroup, databaseClient *database.Client) APIHandler {
	apiHandler := APIHandler{router: router, databaseClient: databaseClient}
	return apiHandler
}

func (apiHandler *APIHandler) RegisterRoutes() {
	fmt.Println("Registering api routes")

	secretsHandler := secrets.CreateAPIHandler(apiHandler.databaseClient)
	
	secretRoutes := apiHandler.router.Group("/secret")
	secretRoutes.GET("/", secretsHandler.Index)
	secretRoutes.POST("/", secretsHandler.Store)
	secretRoutes.GET("/:hash", secretsHandler.Show)
	secretRoutes.DELETE("/:hash", secretsHandler.Delete)
}

func Validate(context *gin.Context) {
	
}