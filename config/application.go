package config

import "github.com/spf13/viper"

// AppConfig is a variable to store application config
var appConfig *viper.Viper

// Instance is a function to get application config
func Instance() *viper.Viper {
	return appConfig
}

// LoadAppConfig is a function to load application configuration
func LoadAppConfig() {
	provider := viper.New()
	provider.SetConfigName("application")
	provider.SetConfigType("yml")
	provider.AddConfigPath("./config")
	provider.AddConfigPath(".")

	setDefaultDatabaseConfig(provider)

	err := provider.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Do nothing because we already set the default value
		} else {
			panic(err.Error())
		}
	}
	appConfig = provider
}

func setDefaultDatabaseConfig(provider *viper.Viper) {
	provider.SetDefault("DB_HOST", "localhost")
	provider.SetDefault("DB_PORT", "5432")
	provider.SetDefault("DB_USERNAME", "postgres")
	provider.SetDefault("DB_PASSWORD", "postgres")
	provider.SetDefault("DB_NAME", "ficree_development")
}
