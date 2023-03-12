package repository

import (
	"errors"
	"go-bookstore/pkg/model"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) error {
	r.db.Create(&user)

	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (*model.User, error) {
	var user *model.User

	if result := r.db.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, result.Error
	}

	if user.Password != password {
		return nil, errors.New("no such user")
	}

	return user, nil
}
