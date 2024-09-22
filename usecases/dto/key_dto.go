package dto

import (
	"paisleypark/kms/domain/entities/keys/symmetric"
	"time"
)

type KeyDTO struct {
	PPRN         string    `json:"pprn"`
	CreationDate time.Time `json:"creation_date"`
}

func MapKeyToDTO(sk *symmetric.Key) *KeyDTO {
	return &KeyDTO{PPRN: sk.PPRN(), CreationDate: sk.CreationDate}
}
