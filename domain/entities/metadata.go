package entities

type Metadata struct {
	KeyIdentifier     string            `json:"key_identifier"`
	Algorithm         string            `json:"algorithm"`
	EncryptionContext map[string]string `json:"encryption_context"`
}
