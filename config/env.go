package config

import (
	"github.com/spf13/viper"
)

// Env is a variable to store environment variables
var Env *viper.Viper

// Instance is a function to get application config
func Instance() *viper.Viper {
	return Env
}

// LoadAppEnv is a function to load application environment variables
func LoadAppEnv() {
	provider := viper.New()
	provider.SetConfigFile(".env")

	LoadDefaultDatabaseConfig(provider)
	LoadDefaultStorageConfig(provider)

	err := provider.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Do nothing because we already set the default value
		} else {
			panic(err.Error())
		}
	}
	Env = provider
}
