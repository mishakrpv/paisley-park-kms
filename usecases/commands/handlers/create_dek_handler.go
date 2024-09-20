package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/commands/requests"
	"paisleypark/kms/util"

	"go.uber.org/zap"
)

type CreateDekHandler struct {
	Repository interfaces.DataEncryptionKeyRepository
}

func NewCreateDekHandler(r interfaces.DataEncryptionKeyRepository) *CreateDekHandler {
	return &CreateDekHandler{Repository: r}
}

func (handler *CreateDekHandler) Execute(request *requests.CreateDataEncryptionKeyRequest) error {
	key := util.NewByteSequence(256)
	masterKey := []byte(os.Getenv("MASTER_KEY"))

	c, err := aes.NewCipher(masterKey)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	gcm.Seal(nonce, nonce, *key, nil)

	return nil
}
