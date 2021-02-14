package database

import (
	"fmt"

	"github.com/fikrimohammad/ficree-api/config"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

// Load is a function to load database connection
func Load() *pg.DB {
	if db == nil {
		Connect()
	}
	return db
}

// Close is a function to close database connection
func Close() {
	db.Close()
	db = nil
}

// Connect is a function to connect to a database
func Connect() {
	db = pg.Connect(&pg.Options{
		User:     config.GetDBUsername(),
		Password: config.GetDBPassword(),
		Database: config.GetDBName(),
		Addr:     fmt.Sprintf("%v:%v", config.GetDBHost(), config.GetDBPort()),
	})
}
