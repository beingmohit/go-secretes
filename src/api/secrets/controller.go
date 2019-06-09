package secrets

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/beingmohit/go-secretes/src/database"
	"github.com/beingmohit/go-secretes/src/database/models"
)

type APIHandler struct {
	databaseClient *database.Client
}

func CreateAPIHandler(databaseClient *database.Client) APIHandler {
	apiHandler := APIHandler{databaseClient: databaseClient}
	return apiHandler
}

func (apiHandler *APIHandler) Index(context *gin.Context) {
	
}	

func (apiHandler *APIHandler) Store(context *gin.Context) {
	var request StoreRequest

	if context.BindJSON(&request) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bind Error"})
		return
	}

	secret := &models.Secret{
		Hash: request.Hash,
		SecretText: request.SecretText,
		CreatedAt: request.CreatedAt,
		ExpiresAt: request.ExpiresAt,
		RemainingViews: request.RemainingViews,
	}
	
	error := apiHandler.databaseClient.Save("secrets", secret)
	
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	
	context.JSON(http.StatusOK, secret)
	return
}

func (apiHandler *APIHandler) Show(context *gin.Context) {
	fmt.Println("Show")
}

func (apiHandler *APIHandler) Delete(context *gin.Context) {
	fmt.Println("Delete")
}

