package queries

import (
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/usecases/dto"
)

type DeksByAccountIdQuery struct {
	Repository interfaces.DataEncryptionKeyRepository
}

func NewDeksByAccountIdQuery(r interfaces.DataEncryptionKeyRepository) *DeksByAccountIdQuery {
	return &DeksByAccountIdQuery{Repository: r}
}

func (handler *DeksByAccountIdQuery) Execute(accountId string) *[]dto.DekDto {
	keys := handler.Repository.FindDeksByAccountId(accountId)

	dtos := []dto.DekDto{}

	for _, key := range keys {
		dtos = append(dtos, *dto.MapDekToDto(&key))
	}

	return &dtos
}
