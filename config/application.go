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
	provider.SetConfigFile(".env")

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
	provider.SetDefault("DATABASE_HOST", "localhost")
	provider.SetDefault("POSTGRES_PORT", "5432")
	provider.SetDefault("POSTGRES_USER", "postgres")
	provider.SetDefault("POSTGRES_PASSWORD", "postgres")
	provider.SetDefault("POSTGRES_DB", "ficree_development")
}
