package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./internal/db/short-url.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
