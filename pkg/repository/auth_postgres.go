package repository

import (
	"fmt"
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

type userNotFoundError struct{}

func (er *userNotFoundError) Error() string {
	return "user not found"
}

func (r *AuthPostgres) GetUser(username, password string) (*model.User, error) {
	var user *model.User

	if result := r.db.Where("username = ?", username).First(&user); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	if user.Password != password {
		return nil, &userNotFoundError{}
	}

	return user, nil
}
