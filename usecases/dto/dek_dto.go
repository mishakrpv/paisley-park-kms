package dto

import (
	"paisleypark/kms/domain/entities"
)

type DekDto struct {
	AccountID   string `json:"account_id"`
	KeyID       string `json:"key_id"`
	DateCreated string `json:"date_created"`
}

func MapDekToDto(key *entities.DataEncryptionKey) *DekDto {
	dto := DekDto{
		AccountID:   key.AccountID.String(),
		KeyID:       key.KeyID.String(),
		DateCreated: key.DateCreated.String(),
	}

	return &dto
}
