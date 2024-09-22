package repositories

import (
	"paisleypark/kms/domain/entities/keys/symmetric"
)

type SymmetricKeyRepository interface {
	Create(sk *symmetric.Key) error
	GetKeyById(keyId string) (*symmetric.Key, error)
	GetKeysByAccountId(accountId string) ([]symmetric.Key, error)
	Delete(keyId string) error
}
