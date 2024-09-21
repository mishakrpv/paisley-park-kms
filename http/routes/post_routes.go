package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"paisleypark/kms/infrastructure/repositories"
	"paisleypark/kms/usecases/commands/handlers"
	"paisleypark/kms/usecases/commands/requests"
)

func POSTKeys(c *gin.Context) {
	repository := repositories.NewGormDekRepository(Db)
	handler := handlers.NewCreateDekHandler(repository)

	var json requests.CreateDataEncryptionKeyRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	err := handler.Execute(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are great"})
}

func POSTEncrypt(c *gin.Context) {
	repository := repositories.NewGormDekRepository(Db)
	handler := handlers.NewEncryptHandler(repository)

	var json requests.EncryptRequest
	
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	response, err := handler.Execute(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *response)
}

func POSTDecrypt(c *gin.Context) {
	repository := repositories.NewGormDekRepository(Db)
	handler := handlers.NewDecryptHandler(repository)

	var json requests.DecryptRequest
	
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	response, err := handler.Execute(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *response)
}
