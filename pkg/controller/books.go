package controller

import (
	"net/http"
)

func (c *Controller) getBooks(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.GetAll()
	if err != nil {
		return
	}
}

func (c *Controller) getBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.GetById()
	if err != nil {
		return
	}
}

func (c *Controller) deleteBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.Delete()
	if err != nil {
		return
	}
}

func (c *Controller) createBook(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.Create()
	if err != nil {
		return
	}
}

func (c *Controller) updateBookById(w http.ResponseWriter, r *http.Request) {
	err := c.services.Bookstore.Update()
	if err != nil {
		return
	}
}
