package requests

type EncryptRequest struct {
	KeyID     string `json:"key_id" binding:"required"`
	Plaintext string `json:"plaintext"`
}
