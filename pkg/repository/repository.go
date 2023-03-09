package repository

import (
	"github.com/jmoiron/sqlx"
	"go-bookstore/pkg/model"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (model.User, error)
}

type Bookstore interface {
	Create() error
	GetAll() error
	GetById() error
	Delete() error
	Update() error
}

type Repository struct {
	Authorization
	Bookstore
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewBookstorePostgres(db),
	}
}
