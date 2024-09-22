package symmetric

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Key struct {
	AccountID    uuid.UUID `json:"account_id"`
	KeyID        uuid.UUID `json:"key_id"`
	Region       string    `json:"region"`
	Description  string    `json:"description"`
	KeySpec      string    `json:"key_spec"`
	Ciphertext   string    `json:"ciphertext"`
	CreationDate time.Time `json:"creation_date"`
}

func NewKey(
	accountId uuid.UUID,
	region string,
	description string,
	keySpec string,
	ciphertext string,
) *Key {
	return &Key{
		AccountID:    accountId,
		KeyID:        uuid.New(),
		Region:       region,
		Description:  description,
		KeySpec:      keySpec,
		Ciphertext:   ciphertext,
		CreationDate: time.Now().UTC(),
	}
}

func UUIDFromPPRN(pprn string) (keyId uuid.UUID, err error) {
	substrings := strings.Split(pprn, "/")
	if len(substrings) != 2 {
		err = errors.New("incorrect pprn")
		return
	}
	keyId, err = uuid.Parse(substrings[1])
	return
}

func (m *Key) PPRN() string {
	return fmt.Sprintf("pprn:ppws:kms:%s:%s:key/%s", m.Region, m.AccountID, m.KeyID)
}

func (Key) TableName() string {
	return "symmetric_keys"
}
