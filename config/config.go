package config

import (
	"github.com/spf13/viper"
)

// Instance ...
type Instance struct {
	Environment string `mapstructure:"app_env"`
	*MainDatabase
	*TestDatabase
	*Storage
}

var appConfig *Instance

// Get is a function to get application config
func Get() *Instance {
	if appConfig == nil {
		Load()
	}
	return appConfig
}

// Clear is a function clear application configs
func Clear() {
	appConfig = nil
}

// Load is a function to load application configs
func Load() {
	provider := viper.New()
	provider.AddConfigPath("../../..")
	provider.AddConfigPath("../..")
	provider.AddConfigPath(".")
	provider.SetConfigName("app")
	provider.SetConfigType("env")
	provider.AutomaticEnv()

	provider.SetDefault("APP_ENV", "development")
	LoadDefaultMainDatabaseConfig(provider)
	LoadDefaultTestDatabaseConfig(provider)
	LoadDefaultStorageConfig(provider)

	err := provider.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Do nothing because we already set the default value
		} else {
			panic(err)
		}
	}

	err = provider.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	var mainDBConfig *MainDatabase
	err = provider.Unmarshal(&mainDBConfig)
	if err != nil {
		panic(err)
	}
	appConfig.MainDatabase = mainDBConfig

	var testDBConfig *TestDatabase
	err = provider.Unmarshal(&testDBConfig)
	if err != nil {
		panic(err)
	}
	appConfig.TestDatabase = testDBConfig

	var storageConfig *Storage
	err = provider.Unmarshal(&storageConfig)
	if err != nil {
		panic(err)
	}
	appConfig.Storage = storageConfig
}

// IsTesting ...
func IsTesting() bool {
	return Get().Environment == "test"
}

// IsDevelopment ...
func IsDevelopment() bool {
	return Get().Environment == "development"
}

// IsProduction ...
func IsProduction() bool {
	return Get().Environment == "production"
}
