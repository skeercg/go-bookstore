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

func (r *BookstorePostgres) GetById(id int) (model.Book, error) {
	var book model.Book

	result := r.db.First(&book, id)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

func (r *BookstorePostgres) Delete(id int) error {
	result := r.db.Delete(&model.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookstorePostgres) Update(book model.Book, id int) error {
	var oldBook model.Book

	result := r.db.First(&oldBook, id)

	if result.Error != nil {
		return result.Error
	}

	oldBook.Title = book.Title
	oldBook.Description = book.Description
	oldBook.Cost = book.Cost

	result = r.db.Save(&oldBook)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
