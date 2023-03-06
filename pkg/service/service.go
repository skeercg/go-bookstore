package service

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
