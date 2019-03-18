package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Document ..
type Document struct {
	ID    int
	Count int
}

// TableName ..
func (Document) TableName() string {
	return "documents"
}

// DocumentStorage ..
type DocumentStorage interface {
	ByID(id int) (obj Document)
}

type documentImpl struct {
	db *gorm.DB
}

// NewDocumentStorage ..
func NewDocumentStorage() DocumentStorage {
	return &documentImpl{db: db}
}

func (s *documentImpl) ByID(id int) (obj Document) {
	dbc := s.db.First(&obj, id)
	if dbc.Error != nil {
		log.Fatal(dbc.Error)
	}
	return obj
}

func (s *documentImpl) Create() {

}
