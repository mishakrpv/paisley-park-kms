package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ConfigurationManager struct{}

func NewConfigurationManager() *ConfigurationManager {
	cm := &ConfigurationManager{}
	
	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
		os.Setenv("ENV", env)
	}

	viper.SetConfigFile("appsettings.json")

	err := viper.ReadInConfig()

	if err != nil {
		zap.L().Warn("Failed to read configuration file in config", zap.Error(err))
	}

	cm.AddFile(fmt.Sprintf("appsettings.%s.json", env))
	cm.AddFile(".env")

	viper.AutomaticEnv()

	return cm
}

func (cm *ConfigurationManager) AddFile(path string) {
	viper.SetConfigFile(path)

	err := viper.MergeInConfig()
	if err != nil {
		zap.L().Warn("Failed to merge configuration file in config",
			zap.String("filepath", path),
			zap.Error(err))
	}
}

func (cm *ConfigurationManager) Get(key string) string {
	return viper.GetString(key)
}
