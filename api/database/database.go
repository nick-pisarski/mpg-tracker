package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New Creates a new DB client from a connection string
func New(connStr string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}
	return db
}
