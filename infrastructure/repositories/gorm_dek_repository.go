package repositories

import (
	"gorm.io/gorm"

	"paisleypark/kms/domain/entities"
	interfaces "paisleypark/kms/interfaces/repositories"
)

type GormDekRepository struct {
	db *gorm.DB
}

func NewGormDekRepository(db *gorm.DB) interfaces.DataEncryptionKeyRepository {
	repository := new(GormDekRepository)
	repository.db = db

	return repository
}

func (r *GormDekRepository) Create(dek *entities.DataEncryptionKey) error {
	return r.db.Create(dek).Error
}

func (r *GormDekRepository) FindDeksByAccountId(accountId string) []entities.DataEncryptionKey {
	var keys []entities.DataEncryptionKey
	r.db.Find(&keys, "account_id = ?", accountId)
	return keys
}

func (r *GormDekRepository) Delete(keyId string) error {
	return r.db.Delete(&entities.DataEncryptionKey{}, keyId).Error
}
