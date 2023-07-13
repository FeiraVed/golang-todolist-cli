package controller

import "io"

type TodolistController interface {
	Create(w io.Writer, name string)
	Update(w io.Writer, id int, name string)
	Delete(id int)
	FindById(w io.Writer, id int)
	FindAll(w io.Writer)
}
