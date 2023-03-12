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
	GetById() error
	Delete() error
	Update() error
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
