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

// SetCache sets the cache with the provided data
func SetCache(cacheData *models.Cache) {
	log.Println("Running cache update process...")
	err := db.Database.Create(cacheData).Error
	if err != nil {
		log.Printf("Failed to create cache: %v\n", err)
		return
	}

	err = db.Database.Preload("Stores",
		func(db *gorm.DB) *gorm.DB {
			return db.Order("stores.name ASC")
		}).Last(cacheData).Error
	if err != nil {
		log.Printf("Failed to load cache: %v\n", err)
		return
	}

	gc.Purge()
	gc.Set("cache", cacheData)
}

// GetCache retrieves the cache from memory
func GetCache() (interface{}, error) {
	cache, err := gc.Get("cache")
	if err != nil {
		log.Printf("Failed to retrieve cache: %v\n", err)
		return nil, err
	}

	return cache, nil
}
