package encrypt

import (
	"encoding/base64"

	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/interfaces/services"
	"paisleypark/kms/util"
)

type EncryptHandler struct {
	repo interfaces.SymmetricKeyRepository
}

func NewEncryptHandler(r interfaces.SymmetricKeyRepository) *EncryptHandler {
	return &EncryptHandler{repo: r}
}

func (h *EncryptHandler) Execute(r *EncryptRequest) (s string, e *util.HttpError) {
	key, err := h.repo.GetKeyById(r.KeyID)
	if err != nil {
		e = util.ErrKeyNotFound
		return
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(key.Ciphertext)

	masterKey := services.MasterKey()

	keyMaterial, err := util.Decrypt(ciphertext, masterKey)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	plaintext := []byte(r.Plaintext)

	ciphertext, err = util.Encrypt(plaintext, keyMaterial)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	ciphertextBlob := key.PPRN() + "." + encodedCiphertext

	s = base64.StdEncoding.EncodeToString([]byte(ciphertextBlob))
	return
}
