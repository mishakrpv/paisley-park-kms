package handlers

import (
	"encoding/base64"
	"encoding/json"

	"paisleypark/kms/domain/entities"
	config "paisleypark/kms/interfaces/configuration"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/commands/requests"
	"paisleypark/kms/usecases/responses"
	"paisleypark/kms/util"

	"go.uber.org/zap"
)

type EncryptHandler struct {
	Repository interfaces.DataEncryptionKeyRepository
}

func NewEncryptHandler(r interfaces.DataEncryptionKeyRepository) *EncryptHandler {
	return &EncryptHandler{Repository: r}
}

func (handler *EncryptHandler) Execute(request *requests.EncryptRequest) (*responses.EncryptResponse, error) {

	key := handler.Repository.FindById(request.KeyID)
	if key == nil {
		return nil, nil
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(key.CiphertextBlob)
	masterKey, _ := base64.StdEncoding.DecodeString(config.Config.Get("MASTER_KEY"))

	keyMaterial, err := util.Decrypt(ciphertext, masterKey)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return nil, err
	}

	zap.L().Debug("Key has been decrypted",
		zap.String("from", key.CiphertextBlob),
		zap.ByteString("to", keyMaterial))

	plaintext, _ := base64.StdEncoding.DecodeString(request.Plaintext)

	ciphertext, err = util.Encrypt(plaintext, keyMaterial)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return nil, err
	}

	encryptionContext := map[string]string{
		"": "",
	}

	metadata, _ := json.Marshal(entities.Metadata{
		KeyIdentifier:     key.Identifier(),
		Algorithm:         key.Algorithm,
		EncryptionContext: encryptionContext})

	return &responses.EncryptResponse{KeyID: key.KeyID.String(),
		CiphertextBlob: base64.StdEncoding.EncodeToString(metadata) + "." + base64.StdEncoding.EncodeToString(ciphertext)}, nil
}
