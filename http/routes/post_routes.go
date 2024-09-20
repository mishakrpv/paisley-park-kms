package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"

	"paisleypark/kms/infrastructure/repositories"
	"paisleypark/kms/usecases/commands/handlers"
	"paisleypark/kms/usecases/commands/requests"
)

func POSTKeys(c *gin.Context) {
	repository := repositories.NewMySqlDekRepository(viper.GetString("CONNECTION_STRINGS__DB_CONNECTION"))
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
