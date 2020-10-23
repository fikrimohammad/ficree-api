package database

import (
	"fmt"

	"github.com/fikrimohammad/ficree-api/app/models"
	"github.com/fikrimohammad/ficree-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig is a struct to store database configuration
type DBConfig struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	Username string `mapstructure:"db_username"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
}

var db *gorm.DB

// Instance is a function to fetch database connection
func Instance() *gorm.DB {
	return db
}

// Connect is a function to connect to a database
func Connect() {
	var dbConfig DBConfig
	configErr := config.Instance().Unmarshal(&dbConfig)
	if configErr != nil {
		panic(configErr.Error())
	}

	var err error
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Host,
		dbConfig.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error when connecting to database")
	}
}

// Migrate is a function to run database migrations
func Migrate() {
	db.AutoMigrate(&models.User{})
}
