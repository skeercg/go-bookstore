package repository

import (
	"go-bookstore/pkg/model"
	"gorm.io/gorm"
)

type BookstorePostgres struct {
	db *gorm.DB
}

func NewBookstorePostgres(db *gorm.DB) *BookstorePostgres {
	return &BookstorePostgres{db: db}
}

func (r *BookstorePostgres) Create(book model.Book) error {
	r.db.Create(&book)

	return nil
}

func (r *BookstorePostgres) GetAll(title, sort string) ([]model.Book, error) {
	orderQuery := "cost " + sort
	titleQuery := "%" + title + "%"

	var books []model.Book
	result := r.db.Where("title LIKE ?", titleQuery).Order(orderQuery).Find(&books)

	if result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

func (r *BookstorePostgres) GetById() error {
	return nil
}

func (r *BookstorePostgres) Delete() error {
	return nil
}

func (r *BookstorePostgres) Update() error {
	return nil
}
