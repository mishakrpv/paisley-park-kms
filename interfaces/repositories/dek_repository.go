package repositories

import "paisleypark/kms/domain/entities"

type DataEncryptionKeyRepository interface {
	Create(dek *entities.DataEncryptionKey) error
	FindById(keyId string) *entities.DataEncryptionKey
	FindDeksByAccountId(accountId string) []entities.DataEncryptionKey
	Delete(keyId string) error
}
