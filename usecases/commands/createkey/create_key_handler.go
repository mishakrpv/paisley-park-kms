package createkey

import (
	"encoding/base64"

	"paisleypark/kms/domain/entities/keys/symmetric"
	interfaces "paisleypark/kms/interfaces/repositories"
	"paisleypark/kms/interfaces/services"
	"paisleypark/kms/usecases/dto"
	"paisleypark/kms/util"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CreateKeyHandler struct {
	repo interfaces.SymmetricKeyRepository
}

func NewCreateKeyHandler(r interfaces.SymmetricKeyRepository) *CreateKeyHandler {
	return &CreateKeyHandler{repo: r}
}

func (h *CreateKeyHandler) Execute(r *CreateKeyRequest) (k *dto.KeyDTO, e *util.HttpError) {
	accountId, err := uuid.Parse(r.AccountID)
	if err != nil {
		e = util.ErrInvalidAccountId
		return
	}

	size, exists := symmetric.MapKeySize[r.KeySpec]
	if !exists {
		e = util.ErrUnsupportedKeySpec
		return
	}

	keyMaterial, err := util.RandomBytes(size)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	masterKey := services.MasterKey()

	ciphertext, err := util.Encrypt(keyMaterial, masterKey)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	// validate region
	sk := symmetric.NewKey(accountId, r.Region, r.Description, r.KeySpec, encodedCiphertext)

	err = h.repo.Create(sk)
	if err != nil {
		e = util.ErrInternalServer(err)
		return
	}

	zap.L().Debug("Key created in db", zap.String("key_id", sk.KeyID.String()))

	k = dto.MapKeyToDTO(sk)
	return
}
