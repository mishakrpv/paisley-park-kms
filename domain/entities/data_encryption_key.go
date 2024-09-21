package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DataEncryptionKey struct {
	KeyID          uuid.UUID `json:"key_id"`
	AccountID      uuid.UUID `json:"account_id"`
	Name           string    `json:"name"`
	Region         string    `json:"region"`
	Algorithm      string    `json:"algorithm"`
	RotationPeriod int       `json:"rotation_period"`

	DateCreated    time.Time `json:"date_created"`
	CiphertextBlob string    `json:"ciphertext_blob"`
}

func NewDataEncryptionKey(
	accountId uuid.UUID,
	name string,
	region string,
	algorithm string,
	rotationPeriod int,
	ciphertextBlob string) *DataEncryptionKey {
	key := DataEncryptionKey{DateCreated: time.Now().UTC()}

	key.AccountID = accountId
	key.Name = name
	key.Region = region
	key.KeyID = uuid.New()
	key.Algorithm = algorithm
	key.RotationPeriod = rotationPeriod
	key.CiphertextBlob = ciphertextBlob

	return &key
}

func (dek *DataEncryptionKey) Identifier() (keyIdentifier string) {
	keyIdentifier = fmt.Sprintf("pprn:ppws:kms:%s:%s:key/%s", dek.Region, dek.AccountID, dek.KeyID)
	return
}

func UUIDFromIdentifier(identifier string) (keyId uuid.UUID) {
	keyId, err := uuid.Parse(strings.Split(identifier, "/")[1])
	if err != nil {
		zap.L().Warn("Wrong key identifier format", zap.String("key_identifier", identifier))
	}
	return
}

func (DataEncryptionKey) TableName() string {
	return "symmetric_keys"
}
