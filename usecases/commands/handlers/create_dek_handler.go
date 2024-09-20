package handlers

import (
	"os"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io"

	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/commands/requests"
	"paisleypark/kms/util"
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
        panic(err.Error)
    }

	gcm, err := cipher.NewGCM(c)
    if err != nil {
        panic(err.Error)
    }

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error)
    }

	gcm.Seal(nonce, nonce, *key, nil)

	return nil
}
