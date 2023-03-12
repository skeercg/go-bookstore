package model

type Book struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cost        int    `json:"cost" binding:"required"`
}
