package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"

	"go.uber.org/zap"
)

func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return string(ciphertext)
}

func Decrypt(encryptedString string, keyString string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}

	return string(plaintext)
}
