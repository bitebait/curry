package models

import "time"

type Cache struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	Stores    []Store   `json:"items"`
}
