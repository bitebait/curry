package cache

import (
	"log"

	"github.com/bitebait/curry/api/db"
	"github.com/bitebait/curry/api/models"
	"github.com/bluele/gcache"
	"gorm.io/gorm"
)

const cacheKey = "cache"

var gc gcache.Cache

func init() {
	gc = gcache.New(1).Build()
}

// SetCache sets the cache with the provided data
func SetCache(cacheData *models.Cache) {
	log.Println("Running cache update process...")
	if err := createCacheEntry(cacheData); err != nil {
		log.Printf("Failed to create cache: %v\n", err)
		return
	}

	if err := preloadCacheData(cacheData); err != nil {
		log.Printf("Failed to load cache: %v\n", err)
		return
	}

	gc.Purge()
	gc.Set(cacheKey, cacheData)
}

func createCacheEntry(cacheData *models.Cache) error {
	return db.Database.Create(cacheData).Error
}

func preloadCacheData(cacheData *models.Cache) error {
	return db.Database.Preload("Stores", func(db *gorm.DB) *gorm.DB {
		return db.Order("stores.name ASC")
	}).Last(cacheData).Error
}

// GetCache retrieves the cache from memory
func GetCache() (interface{}, error) {
	cache, err := gc.Get(cacheKey)
	if err != nil {
		log.Printf("Failed to retrieve cache: %v\n", err)
		return nil, err
	}
	return cache, nil
}
