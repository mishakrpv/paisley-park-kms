package encrypt

import (
	"encoding/base64"
	"fmt"
	"net/http"

	config "paisleypark/kms/interfaces/configuration"

	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/util"

	"go.uber.org/zap"
)

type EncryptHandler struct {
	repo interfaces.SymmetricKeyRepository
}

func NewEncryptHandler(r interfaces.SymmetricKeyRepository) *EncryptHandler {
	return &EncryptHandler{repo: r}
}

func (h *EncryptHandler) Execute(r *EncryptRequest) (string, *util.HttpErr) {
	key, err := h.repo.GetKeyById(r.KeyID)
	if err != nil {
		return "", util.HandleErr(http.StatusNotFound, err.Error())
	}

	if key == nil {
		return "", util.HandleErr(http.StatusNotFound, "key not found")
	}

	encryptedKeyMaterial, err := base64.StdEncoding.DecodeString(key.Ciphertext)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Encrypted key material decoded", zap.String("key_id", r.KeyID))

	masterKey := config.Config.Get("MASTER_KEY")
	masterKeyMaterial, err := base64.StdEncoding.DecodeString(masterKey)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Master key retrieved",
		zap.String("master_key", masterKey),
		zap.String("size", fmt.Sprint(len(masterKeyMaterial))))

	keyMaterial, err := util.Decrypt(encryptedKeyMaterial, masterKeyMaterial)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	decodedPlaintext, err := base64.StdEncoding.DecodeString(r.Plaintext)
	if err != nil {
		return "", util.HandleErr(http.StatusBadRequest, err.Error())
	}

	encryptedPlaintext, err := util.Encrypt(decodedPlaintext, keyMaterial)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	ciphertext := base64.StdEncoding.EncodeToString(encryptedPlaintext)

	payload := key.PPRN() + "." + ciphertext

	return base64.StdEncoding.EncodeToString([]byte(payload)), nil
}
