package routes

import (
	"net/http"
	"github.com/spf13/viper"

	"paisleypark/kms/infrastructure/repositories"
	"paisleypark/kms/usecases/queries"

	"github.com/gin-gonic/gin"
)

func GETKeys(c *gin.Context) {
	repository := repositories.NewMySqlDekRepository(viper.GetString("CONNECTION_STRINGS__DB_CONNECTION"))
	query := queries.NewDeksByAccountIdQuery(repository)

	var json map[string]string

	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	keys := query.Execute(json["account_id"])

	c.JSON(http.StatusOK, *keys)
}
