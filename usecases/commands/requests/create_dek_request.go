package requests

type CreateDataEncryptionKeyRequest struct {
	AccountID      string `json:"account_id" binding:"required"`
	Algorithm      string `json:"algorithm" binding:"required"`
	Name           string `json:"name_identifier" binding:"required"`
	RotationPeriod int    `json:"rotation_period"`
}
