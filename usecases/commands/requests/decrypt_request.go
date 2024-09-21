package requests

type DecryptRequest struct {
	CiphertextBlob string `json:"ciphertext_blob"`
}