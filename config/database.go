package config

import "github.com/spf13/viper"

// MainDatabase ...
type MainDatabase struct {
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUsername string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
}

// TestDatabase ...
type TestDatabase struct {
	TestDBHost     string `mapstructure:"test_db_host"`
	TestDBPort     string `mapstructure:"test_db_port"`
	TestDBUsername string `mapstructure:"test_db_user"`
	TestDBPassword string `mapstructure:"test_db_password"`
	TestDBName     string `mapstructure:"test_db_name"`
}

// LoadDefaultMainDatabaseConfig is a function to load default value for main database configs
func LoadDefaultMainDatabaseConfig(provider *viper.Viper) {
	provider.SetDefault("DB_HOST", "localhost")
	provider.SetDefault("DB_PORT", "5432")
	provider.SetDefault("DB_USER", "postgres")
	provider.SetDefault("DB_PASSWORD", "postgres")
	provider.SetDefault("DB_NAME", "ficree_development")
	provider.SetDefault("DB_MAX_IDLE_CONN", 5)
	provider.SetDefault("DB_MAX_OPEN_CONN", 5)
}

// LoadDefaultTestDatabaseConfig is a function to load default value for test database configs
func LoadDefaultTestDatabaseConfig(provider *viper.Viper) {
	provider.SetDefault("TEST_DB_HOST", "localhost")
	provider.SetDefault("TEST_DB_PORT", "5432")
	provider.SetDefault("TEST_DB_USER", "postgres")
	provider.SetDefault("TEST_DB_PASSWORD", "postgres")
	provider.SetDefault("TEST_DB_NAME", "ficree_development")
	provider.SetDefault("TEST_DB_MAX_IDLE_CONN", 5)
	provider.SetDefault("TEST_DB_MAX_OPEN_CONN", 5)
}

// GetDBHost ...
func GetDBHost() string {
	if IsTesting() {
		return appConfig.TestDBHost
	}
	return appConfig.DBHost
}

// GetDBPort ...
func GetDBPort() string {
	if IsTesting() {
		return appConfig.TestDBPort
	}
	return appConfig.DBPort
}

// GetDBUsername ...
func GetDBUsername() string {
	if IsTesting() {
		return appConfig.TestDBUsername
	}
	return appConfig.DBUsername
}

// GetDBPassword ...
func GetDBPassword() string {
	if IsTesting() {
		return appConfig.TestDBPassword
	}
	return appConfig.DBPassword
}

// GetDBName ...
func GetDBName() string {
	if IsTesting() {
		return appConfig.TestDBName
	}
	return appConfig.DBName
}
