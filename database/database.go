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
	Host     string `mapstructure:"database_host"`
	Port     string `mapstructure:"postgres_port"`
	Username string `mapstructure:"postgres_user"`
	Password string `mapstructure:"postgres_password"`
	Name     string `mapstructure:"postgres_db"`
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
		msg := fmt.Sprintf(
			"Error when connecting to %s on %s:%s\n",
			dbConfig.Name,
			dbConfig.Host,
			dbConfig.Port,
		)
		panic(msg)
	} else {
		fmt.Printf(
			"Connected to %s on %s:%s\n",
			dbConfig.Name,
			dbConfig.Host,
			dbConfig.Port,
		)
	}
}

// Migrate is a function to run database migrations
func Migrate() {
	db.AutoMigrate(&models.User{})
}
