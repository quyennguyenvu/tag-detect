package storage

import "github.com/jinzhu/gorm"

// OneWord ..
type OneWord struct {
	ID     int
	Word   string
	Appear int
	TFIDF  float64
}

// TableName ..
func (OneWord) TableName() string {
	return "one_words"
}

// OneWordStorage ..
type OneWordStorage interface {
	// ByID(id int) (obj Document)
}

type oneWordImpl struct {
	db *gorm.DB
}

// NewOneWordStorage ..
func NewOneWordStorage() OneWordStorage {
	return &oneWordImpl{db: db}
}
