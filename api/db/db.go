package db

import (
	"curry/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init(dataSourceName string) {
	db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err := db.AutoMigrate(&models.Cache{}, &models.Store{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func GetDB() *gorm.DB {
	return db
}
