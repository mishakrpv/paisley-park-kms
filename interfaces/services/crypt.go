package services

import (
	"encoding/base64"
	config "paisleypark/kms/interfaces/configuration"

	"go.uber.org/zap"
)

var masterKey []byte

func MasterKey() []byte {
	if masterKey != nil {
		return masterKey
	}

	encodedMasterKey := config.Config.Get("MASTER_KEY")
	masterKey, err := base64.StdEncoding.DecodeString(encodedMasterKey)
	if err != nil {
		zap.L().Fatal("Configuration missing value for MASTER_KEY")
	}

	zap.L().Debug("Master key retrieved",
		zap.String("master_key", encodedMasterKey))

	return masterKey
}
