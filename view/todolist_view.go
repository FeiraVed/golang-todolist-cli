package view

import (
	"io"
	"os"
)

type TodolistView interface {
	ShowTodolist(w io.Writer)
	AddTodolist(w io.Writer, input *os.File)
	UpdateTodolist(w io.Writer, inputId *os.File, inputName *os.File)
	Delete(w io.Writer, input *os.File)
}
