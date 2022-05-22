package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		log.Fatal("Failed to load database")
	}

	db.AutoMigrate(&Quote{})
	DB = db
}
