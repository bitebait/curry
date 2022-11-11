package db

import (
	"github.com/bitebait/curry/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func Init(dataSourceName string) {
	Database, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err := Database.AutoMigrate(&models.Cache{}, &models.Store{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
