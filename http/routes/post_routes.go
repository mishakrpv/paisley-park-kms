package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"paisleypark/kms/infrastructure/repositories"
	"paisleypark/kms/usecases/commands/createkey"
	"paisleypark/kms/usecases/commands/decrypt"
	"paisleypark/kms/usecases/commands/encrypt"
)

func POSTKeys(c *gin.Context) {
	repository := repositories.NewGormSkRepository(Db)
	handler := createkey.NewCreateKeyHandler(repository)

	var json createkey.CreateKeyRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	dto, err := handler.Execute(&json)

	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &dto)
}

func POSTEncrypt(c *gin.Context) {
	repository := repositories.NewGormSkRepository(Db)
	handler := encrypt.NewEncryptHandler(repository)

	var json encrypt.EncryptRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	ciphertextBlob, err := handler.Execute(&json)
	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ciphertext_blob": &ciphertextBlob})
}

func POSTDecrypt(c *gin.Context) {
	repository := repositories.NewGormSkRepository(Db)
	handler := decrypt.NewDecryptHandler(repository)

	var json decrypt.DecryptRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}

	plaintext, err := handler.Execute(&json)

	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &plaintext)
}
