package decrypt

import (
	"encoding/base64"
	"net/http"
	"strings"

	"paisleypark/kms/domain/entities/keys/symmetric"
	config "paisleypark/kms/interfaces/configuration"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/util"

	"go.uber.org/zap"
)

type DecryptHandler struct {
	repo interfaces.SymmetricKeyRepository
}

func NewDecryptHandler(r interfaces.SymmetricKeyRepository) *DecryptHandler {
	return &DecryptHandler{repo: r}
}

func (h *DecryptHandler) Execute(r *DecryptRequest) (string, *util.HttpErr) {
	ciphertextBlob, err := base64.StdEncoding.DecodeString(r.CiphertextBlob)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	substrings := strings.Split(string(ciphertextBlob), ".")
	if len(substrings) != 2 {
		return "", util.HandleErr(http.StatusBadRequest, "")
	}

	zap.L().Debug("",
		zap.String("pprn", substrings[0]),
		zap.String("ciphertext", substrings[1]))

	keyId, err := symmetric.UUIDFromPPRN(substrings[0])
	if err != nil {
		return "", util.HandleErr(http.StatusBadRequest, err.Error())
	}

	key, err := h.repo.GetKeyById(keyId.String())
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	if key == nil {
		return "", util.HandleErr(http.StatusNotFound, "key not found")
	}

	encryptedKeyMaterial, err := base64.StdEncoding.DecodeString(key.Ciphertext)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	masterKey := config.Config.Get("MASTER_KEY")
	masterKeyMaterial, err := base64.StdEncoding.DecodeString(masterKey)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Master key retrieved",
		zap.String("master_key", masterKey),
		zap.String("size", string(rune(len(masterKeyMaterial)))))

	keyMaterial, err := util.Decrypt(encryptedKeyMaterial, masterKeyMaterial)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	ciphertext, err := base64.StdEncoding.DecodeString(substrings[1])
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	plaintext, err := util.Decrypt(ciphertext, keyMaterial)
	if err != nil {
		return "", util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	encoded := base64.StdEncoding.EncodeToString(plaintext)

	return encoded, nil
}
