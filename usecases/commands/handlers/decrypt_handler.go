package handlers

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"paisleypark/kms/domain/entities"
	config "paisleypark/kms/interfaces/configuration"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/commands/requests"
	"paisleypark/kms/usecases/responses"
	"paisleypark/kms/util"

	"go.uber.org/zap"
)

type DecryptHandler struct {
	Repository interfaces.DataEncryptionKeyRepository
}

func NewDecryptHandler(r interfaces.DataEncryptionKeyRepository) *DecryptHandler {
	return &DecryptHandler{Repository: r}
}

func (handler *DecryptHandler) Execute(request *requests.DecryptRequest) (*responses.DecryptResponse, error) {
	parts := strings.Split(request.CiphertextBlob, ".")

	metadataBytes, _ := base64.StdEncoding.DecodeString(parts[0])

	var metadata entities.Metadata

	err := json.Unmarshal(metadataBytes, &metadata)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
		return nil, err
	}

	key := handler.Repository.FindById(entities.UUIDFromIdentifier(metadata.KeyIdentifier).String())
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

	ciphertextBlob, _ := base64.StdEncoding.DecodeString(parts[1])

	ciphertextBlob, _ = util.Decrypt(ciphertextBlob, keyMaterial)

	return &responses.DecryptResponse{Plaintext: base64.StdEncoding.EncodeToString(ciphertextBlob)}, nil
}
