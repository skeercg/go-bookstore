package repository

import (
	"go-bookstore/pkg/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Bookstore interface {
	Create(book model.Book) error
	GetAll(title, sort string) ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Delete(id int) error
	Update(book model.Book, id int) error
}

type Repository struct {
	Authorization
	Bookstore
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewBookstorePostgres(db),
	}
}
