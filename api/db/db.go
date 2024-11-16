package db

import (
	"github.com/bitebait/curry/api/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func Init(dataSourceName string) {
	sqliteDialector := sqlite.Open(dataSourceName)
	Database, err = gorm.Open(sqliteDialector, &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	migrateDatabase()
}

func migrateDatabase() {
	err = Database.AutoMigrate(&models.Cache{}, &models.Store{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
