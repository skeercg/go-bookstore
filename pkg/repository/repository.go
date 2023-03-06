package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	CreateUser() (string, error)
	GetUser() (string, error)
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
