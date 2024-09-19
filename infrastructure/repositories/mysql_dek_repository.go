package repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"paisleypark/kms/domain/entities"
	interfaces "paisleypark/kms/interfaces/repositories"
)

type MySqlDekRepository struct {
	db *gorm.DB
}

func NewMySqlDekRepository(dsn string) interfaces.DataEncryptionKeyRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	repository := new(MySqlDekRepository)
	repository.db = db

	return repository
}

func (r *MySqlDekRepository) Create(dek *entities.DataEncryptionKey) error {
	return r.db.Create(dek).Error
}

func (r *MySqlDekRepository) Delete(keyId string) error {
	return r.db.Delete(&entities.DataEncryptionKey{}, keyId).Error
}
