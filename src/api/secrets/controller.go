package secrets

import (
	"strings"
	"errors"
	"time"
	"net/http"
	"crypto/md5"
	"encoding/base32"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/beingmohit/go-secretes/src/api/secrets/requests"
	"github.com/beingmohit/go-secretes/src/database"
	"github.com/beingmohit/go-secretes/src/database/models"
)

type Controller struct {
	databaseClient *database.Client
}

func NewController(databaseClient *database.Client) Controller {
	controller := Controller{databaseClient: databaseClient}
	return controller
}

func (controller *Controller) Store(context *gin.Context) {
	var request requests.StoreRequest

	error := context.Bind(&request)
	
	if error != nil {
		controller.error(context, http.StatusBadRequest, error)
		return
	}
	
	if request.ExpireAfterViews < 1 {
		controller.error(context, http.StatusBadRequest, errors.New("Expire after views must be greater than 0"))
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(request.SecretText))
	hash := base32.StdEncoding.EncodeToString(hasher.Sum(nil))
	hash = strings.TrimRight(hash, "=")
	hash = strings.ToLower(hash)
	
	createdAt := time.Now().UTC()	
	expiresAt := time.Time{}

	if request.ExpireAfter > 0 {
		expiresAt = createdAt.Add(time.Minute * time.Duration(request.ExpireAfter))
	}

	if error != nil {
		controller.error(context, http.StatusForbidden, error)
		return
	}
	
	secret := &models.Secret{
		Hash: hash,
		SecretText: request.SecretText,
		CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
		ExpiresAt: expiresAt.Format("2006-01-02 15:04:05"),
		RemainingViews: request.ExpireAfterViews,
	}
	
	error = controller.databaseClient.Save("secrets", secret)
	
	if error != nil {
		controller.error(context, http.StatusBadRequest, error)
		return
	}
	
	controller.success(context, secret)
	
	return
}

func (controller *Controller) Show(context *gin.Context) {
	hash := context.Param("hash")

	secret := &models.Secret{}

	error := controller.databaseClient.Find("secrets", secret, bson.M{"hash":hash})
	
	if error != nil {
		controller.error(context, http.StatusNotFound, error)
		return
	}
	
	if secret.RemainingViews < 1 {
		controller.error(context, http.StatusBadRequest, errors.New("No views remaining"))
		return
	}

	secret.RemainingViews = secret.RemainingViews - 1

	expiresAt, error := time.Parse("2006-01-02 15:04:05", secret.ExpiresAt)

	if error != nil {
		controller.error(context, http.StatusBadRequest, error)
		return
	}
	
	if (!expiresAt.IsZero() && time.Now().UTC().After(expiresAt)) {
		controller.error(context, http.StatusBadRequest, errors.New("Expired"))
		return 
	} 

	error = controller.databaseClient.Save("secrets", secret)
	
	if error != nil {
		controller.error(context, http.StatusBadRequest, error)
		return
	}
	
	controller.success(context, secret)

	return
}

func (controller *Controller) success(context *gin.Context, obj interface{}) {
	controller.response(context, http.StatusOK, obj)
}

func (controller *Controller) error(context *gin.Context, code int, err error) {
	controller.response(context, code, gin.H{"error": err.Error()})
}

func (controller *Controller) response(context *gin.Context, code int, obj interface{}) {
	context.Negotiate(code, gin.Negotiate{
        Offered: []string{gin.MIMEJSON, gin.MIMEXML},
        Data: obj,
	})
}