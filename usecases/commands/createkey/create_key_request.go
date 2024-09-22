package createkey

type CreateKeyRequest struct {
	AccountID   string `json:"account_id" binding:"required"`
	Region      string `json:"region" binding:"required"`
	Description string `json:"description"`
	KeySpec     string `json:"key_spec" binding:"required"`
}
