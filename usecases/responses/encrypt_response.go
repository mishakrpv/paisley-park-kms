package responses

type EncryptResponse struct {
	KeyID          string `json:"key_id"`
	CiphertextBlob string `json:"ciphertext_blob"`
}
