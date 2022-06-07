package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "mpg-tracker/api/models"
)

var basePath = "/Users/nick.pisarski/Repos/sqlite-databases"
var databases = []string{"mpgtracker.db", "mpgtracker_test.db"}

func MigrateDatabases() {
	for _, database := range databases {
		fmt.Printf("Attempting to migrate: %s", database)

		dbPath := fmt.Sprintf("%s/%s", basePath, database)
		db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		err = db.AutoMigrate(&FillUpEntity{})
		if err != nil {
			panic("failed to migrate database")
		}

		fmt.Println(" ...Success")
	}
}

func main() {
	MigrateDatabases()
}
