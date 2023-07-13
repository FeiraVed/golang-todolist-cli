package repository

import (
	"FeiraVed/todolist/entity"
)

type DB map[string]string

type TodolistRepository interface {
	Save(db DB, todolist *entity.Todolist) *entity.Todolist

	Update(db DB, todolist *entity.Todolist) *entity.Todolist

	Delete(db DB, todolist *entity.Todolist)

	FindById(db DB, todolistId int) (*entity.Todolist, error)

	FindAll(db DB) []entity.Todolist
}
