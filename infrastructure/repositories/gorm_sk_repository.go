package repositories

import (
	"gorm.io/gorm"

	"paisleypark/kms/domain/entities/keys/symmetric"
	interfaces "paisleypark/kms/interfaces/repositories"
)

type GormSkRepository struct {
	db *gorm.DB
}

func NewGormSkRepository(db *gorm.DB) interfaces.SymmetricKeyRepository {
	repository := new(GormSkRepository)
	repository.db = db

	return repository
}

func (r *GormSkRepository) Create(sk *symmetric.Key) error {
	return r.db.Create(sk).Error
}

func (r *GormSkRepository) GetKeyById(keyId string) (*symmetric.Key, error) {
	var key symmetric.Key
	err := r.db.First(&key, "key_id = ?", keyId).Error
	return &key, err
}

func (r *GormSkRepository) GetKeysByAccountId(accountId string) ([]symmetric.Key, error) {
	var keys []symmetric.Key
	err := r.db.Find(&keys, "account_id = ?", accountId).Error
	return keys, err
}

func (r *GormSkRepository) Delete(keyId string) error {
	return r.db.Delete(&symmetric.Key{}, keyId).Error
}
