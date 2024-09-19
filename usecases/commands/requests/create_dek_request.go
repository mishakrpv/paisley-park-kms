package requests

type CreateDataEncryptionKeyRequest struct {
	AccountID      string `json:"account_id"`
	Algorithm      string `json:"algorithm"`
	Name           string `json:"name_identifier"`
	RotationPeriod int    `json:"rotation_period"`
}
