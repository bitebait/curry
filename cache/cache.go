package cache

import (
	"log"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/models"
	"github.com/bluele/gcache"
	"gorm.io/gorm"
)

var gc gcache.Cache

func init() {
	gc = gcache.New(1).Build()
}

func SetCache(c *models.Cache) {
	log.Println("Running crawler...")
	db.Database.Create(c)
	db.Database.Preload("Stores",
		func(db *gorm.DB) *gorm.DB {
			return db.Order("stores.name ASC")
		}).Last(&c)
	gc.Purge()
	gc.Set("cache", c)
}

func GetCache() (interface{}, error) {
	cache, err := gc.Get("cache")
	if err != nil {
		return nil, err
	}
	return cache, nil
}
