package service

import "FeiraVed/todolist/entity/app"

type TodolistService interface {
	Create(request app.TodolistCreateRequest) *app.TodolistResponse
	Update(request app.TodolistUpdateRequest) *app.TodolistResponse
	Delete(id int)
	FindById(id int) *app.TodolistResponse
	FindAll() []app.TodolistResponse
}
