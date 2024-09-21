package dto

import (
	"paisleypark/kms/domain/entities"
)

type DekDto struct {
	KeyIdentifier  string `json:"key_identifier"`
	DateCreated    string `json:"date_created"`
	RotationPeriod int    `json:"rotation_period"`
}

func MapDekToDto(key *entities.DataEncryptionKey) *DekDto {
	dto := DekDto{
		KeyIdentifier:  key.Identifier(),
		DateCreated:    key.DateCreated.String(),
		RotationPeriod: key.RotationPeriod,
	}

	return &dto
}
