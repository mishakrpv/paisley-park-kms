package requests

type CreateDataEncryptionKeyRequest struct {
	AccountID      string `json:"account_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Region         string `json:"region" binding:"required"`
	Algorithm      string `json:"algorithm" binding:"required"`
	RotationPeriod int    `json:"rotation_period"`
}
