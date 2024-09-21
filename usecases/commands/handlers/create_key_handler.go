package handlers

import (
	"encoding/base64"

	"paisleypark/kms/domain/entities"
	config "paisleypark/kms/interfaces/configuration"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/commands/requests"
	"paisleypark/kms/util"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CreateDekHandler struct {
	Repository interfaces.DataEncryptionKeyRepository
}

func NewCreateDekHandler(r interfaces.DataEncryptionKeyRepository) *CreateDekHandler {
	return &CreateDekHandler{Repository: r}
}

func (handler *CreateDekHandler) Execute(request *requests.CreateDataEncryptionKeyRequest) error {

	accountId, err := uuid.Parse(request.AccountID)
	if err != nil {
		zap.L().Debug("Failed to parse \"account_id\"", zap.Error(err))
		return err
	}

	var size int
	switch request.Algorithm {
	case "AES-256":
		size = 32
	case "AES-192":
		size = 24
	case "AES-128":
		size = 16
	case "AES-256 HSM":
		size = 32
	}

	key, err := util.RandomBytes(size)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return err
	}

	masterKey, _ := base64.StdEncoding.DecodeString(config.Config.Get("MASTER_KEY"))

	ciphertext, err := util.Encrypt(key, masterKey)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return err
	}

	dek := entities.NewDataEncryptionKey(
		accountId,
		request.Name,
		request.Region,
		request.Algorithm,
		request.RotationPeriod,
		base64.RawStdEncoding.EncodeToString(ciphertext))

	err = handler.Repository.Create(dek)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return err
	}

	return nil
}
