package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(bytesToEncrypt []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, bytesToEncrypt, nil), nil
}

func Decrypt(bytesToDecrypt []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	// if len(ciphertext) < nonceSize {
	//     return "", err
	// }

	nonce, bytesToDecrypt := bytesToDecrypt[:nonceSize], bytesToDecrypt[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, bytesToDecrypt, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
