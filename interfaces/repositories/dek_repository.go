package repositories

import "paisleypark/kms/domain/entities"

type DataEncryptionKeyRepository interface {
	Create(dek *entities.DataEncryptionKey) error
	Delete(keyId string) error
}