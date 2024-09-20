package repositories

import "paisleypark/kms/domain/entities"

type DataEncryptionKeyRepository interface {
	Create(dek *entities.DataEncryptionKey) error
	FindDeksByAccountId(accountId string) []entities.DataEncryptionKey
	Delete(keyId string) error
}