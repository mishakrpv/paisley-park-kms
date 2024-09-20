package util

import (
	"crypto/rand"
)

func NewByteSequence(length int) *[]byte {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err.Error)
	}

	return &randomBytes
}
