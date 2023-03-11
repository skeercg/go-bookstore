package repository

import (
	"gorm.io/gorm"
)

type BookstorePostgres struct {
	db *gorm.DB
}

func NewBookstorePostgres(db *gorm.DB) *BookstorePostgres {
	return &BookstorePostgres{db: db}
}

func (r *BookstorePostgres) Create() error {
	return nil
}

func (r *BookstorePostgres) GetAll() error {
	return nil
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
