package repository

import "github.com/jmoiron/sqlx"

type BookstorePostgres struct {
	db *sqlx.DB
}

func NewBookstorePostgres(db *sqlx.DB) *BookstorePostgres {
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
