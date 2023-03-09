package repository

import (
	"github.com/jmoiron/sqlx"
	"go-bookstore/pkg/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) error {
	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	return model.User{}, nil
}
