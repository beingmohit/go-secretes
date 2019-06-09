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
	fmt.Println("Index")

	secret := &models.Secret{
		Hash: "77788",
		SecretText: "77788",
		CreatedAt: "77788",
		ExpiresAt: "77788",
		RemainingViews: 1000,
	}

	error := apiHandler.databaseClient.Save("secrets", secret)
	
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"status": true})
	return
}	

func (apiHandler *APIHandler) Store(context *gin.Context) {
	fmt.Println("Store")
}

func (apiHandler *APIHandler) Show(context *gin.Context) {
	fmt.Println("Show")
}

func (apiHandler *APIHandler) Delete(context *gin.Context) {
	fmt.Println("Delete")
}

