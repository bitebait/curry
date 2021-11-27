package models

type Store struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Currency string  `json:"currency"`
	Value    float32 `gorm:"type:decimal(10,2);" json:"value"`
	URL      string  `json:"url"`
	CacheID  uint    `json:"-"`
}
