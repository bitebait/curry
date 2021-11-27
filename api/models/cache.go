package models

import "time"

type Cache struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Stores    []Store    `json:"items"`
}
