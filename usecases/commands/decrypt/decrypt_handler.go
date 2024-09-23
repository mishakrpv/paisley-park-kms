package decrypt

import (
	"encoding/base64"
	"strings"

	"paisleypark/kms/domain/entities/keys/symmetric"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/interfaces/services"
	"paisleypark/kms/util"
)

type DecryptHandler struct {
	repo interfaces.SymmetricKeyRepository
}

func NewDecryptHandler(r interfaces.SymmetricKeyRepository) *DecryptHandler {
	return &DecryptHandler{repo: r}
}

func (h *DecryptHandler) Execute(r *DecryptRequest) (s string, e *util.HttpError) {
	encodedCiphertextBlob, err := base64.StdEncoding.DecodeString(r.CiphertextBlob)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	substrings := strings.Split(string(encodedCiphertextBlob), ".")
	if len(substrings) != 2 {
		e = util.ErrInvalidCiphertextBlob
		return
	}

	keyId, err := symmetric.UUIDFromPPRN(substrings[0])
	if err != nil {
		e = util.ErrInvalidPPRN
		return
	}

	ch := make(chan *[]byte)

	go func(keyId string, r interfaces.SymmetricKeyRepository) {
		key, err := r.GetKeyById(keyId)

		if err != nil {
			e = util.ErrKeyNotFound
		}

		ciphertext, err := base64.StdEncoding.DecodeString(key.Ciphertext)
		if err != nil {
			e = util.ErrInternalServer(err)
		}

		ch <- &ciphertext
	}(keyId.String(), h.repo)

	ciphertext, err := base64.StdEncoding.DecodeString(substrings[1])
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	kCiphertext := <- ch

	masterKey := services.MasterKey()

	keyMaterial, err := util.Decrypt(*kCiphertext, masterKey)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	plaintext, err := util.Decrypt(ciphertext, keyMaterial)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	s = string(plaintext)
	return
}
