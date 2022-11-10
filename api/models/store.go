package models

type Store struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Value    string `json:"value"`
	URL      string `json:"url"`
	CacheID  uint   `json:"-"`
}
