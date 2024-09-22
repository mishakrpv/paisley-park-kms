package createkey

import (
	"encoding/base64"
	"net/http"

	"paisleypark/kms/domain/entities/keys/symmetric"
	config "paisleypark/kms/interfaces/configuration"
	interfaces "paisleypark/kms/interfaces/repositories"
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

func (h *CreateKeyHandler) Execute(r *CreateKeyRequest) (*dto.KeyDTO, *util.HttpErr) {
	accountId, err := uuid.Parse(r.AccountID)
	if err != nil {
		return nil, util.HandleErr(http.StatusBadRequest, "incorrect account id")
	}

	size, exists := symmetric.MapKeySize[r.KeySpec]
	if !exists {
		return nil, util.HandleErr(http.StatusBadRequest, "unsupported key specification")
	}

	keyMaterial, err := util.RandomBytes(size)
	if err != nil {
		return nil, util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Key material generated", zap.ByteString("key_material", keyMaterial))

	masterKey := config.Config.Get("MASTER_KEY")
	masterKeyMaterial, err := base64.StdEncoding.DecodeString(masterKey)
	if err != nil {
		return nil, util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Master key retrieved",
		zap.String("master_key", masterKey),
		zap.String("size", string(rune(len(masterKeyMaterial)))))

	encryptedKeyMaterial, err := util.Encrypt(keyMaterial, masterKeyMaterial)
	if err != nil {
		return nil, util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	ciphertext := base64.StdEncoding.EncodeToString(encryptedKeyMaterial)

	// validate region
	sk := symmetric.NewKey(accountId, r.Region, r.Description, r.KeySpec, ciphertext)

	err = h.repo.Create(sk)
	if err != nil {
		return nil, util.HandleErr(http.StatusInternalServerError, err.Error())
	}

	zap.L().Debug("Key created in db", zap.String("key_id", sk.KeyID.String()))

	return dto.MapKeyToDTO(sk), nil
}
