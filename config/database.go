package config

import "github.com/spf13/viper"

// LoadDefaultDatabaseConfig is a function to load default value for database configs
func LoadDefaultDatabaseConfig(provider *viper.Viper) {
	provider.SetDefault("DB_HOST", "localhost")
	provider.SetDefault("DB_PORT", "5432")
	provider.SetDefault("DB_USER", "postgres")
	provider.SetDefault("DB_PASSWORD", "postgres")
	provider.SetDefault("DB_NAME", "ficree_development")
	provider.SetDefault("DB_MAX_IDLE_CONN", 5)
	provider.SetDefault("DB_MAX_OPEN_CONN", 5)
}
