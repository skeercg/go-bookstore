package repository

import (
	"fmt"
	"go-bookstore/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
	booksTable = "books"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Book{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
