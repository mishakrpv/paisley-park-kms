package entities

import (
	"time"

	"github.com/google/uuid"
)

type DataEncryptionKey struct {
	AccountID      uuid.UUID `json:"account_id"`
	KeyID          uuid.UUID `json:"key_id"`
	Algorithm      string    `json:"algorithm"`
	DateCreated    time.Time `json:"date_created"`
	CiphertextBlob string    `json:"ciphertext_blob"`
}

func NewDataEncryptionKey(
	accountId uuid.UUID,
	algorithm string,
	ciphertextBlob string) *DataEncryptionKey {
	key := DataEncryptionKey{DateCreated: time.Now().UTC()}

	key.AccountID = accountId
	key.KeyID = uuid.New()
	key.Algorithm = algorithm
	key.CiphertextBlob = ciphertextBlob

	return &key
}
