package util

import (
	"crypto/rand"

	"go.uber.org/zap"
)

func NewByteSequence(length int) *[]byte {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	return &randomBytes
}
