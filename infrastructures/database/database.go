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
	Host        string `mapstructure:"db_host"`
	Port        string `mapstructure:"db_port"`
	Username    string `mapstructure:"db_user"`
	Password    string `mapstructure:"db_password"`
	Name        string `mapstructure:"db_name"`
	MaxIdleConn int    `mapstructure:"db_max_idle_conn"`
	MaxOpenConn int    `mapstructure:"db_max_open_conn"`
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

	var connErr error
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Host,
		dbConfig.Port,
	)
	db, connErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connErr != nil {
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

	sqlDB, dbErr := db.DB()
	if dbErr != nil {
		panic(dbErr)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
}

// Migrate is a function to run database migrations
func Migrate() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.Experience{})
	db.AutoMigrate(&models.Education{})
}
