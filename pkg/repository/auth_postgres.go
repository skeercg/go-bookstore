package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser() (string, error) {
	return "", nil
}

func (r *AuthPostgres) GetUser() (string, error) {
	return "", nil
}
