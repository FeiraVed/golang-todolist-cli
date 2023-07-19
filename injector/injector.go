//go:build wireinject
// +build wireinject

package injector

import (
	"FeiraVed/todolist/controller"
	"FeiraVed/todolist/repository"
	"FeiraVed/todolist/service"
	"FeiraVed/todolist/view"

	"github.com/google/wire"
)

var todolistSet = wire.NewSet(
	repository.NewTodolistRepository,
	wire.Bind(new(repository.TodolistRepository), new(*repository.TodolistRepositoryImpl)),
	service.NewTodolistService,
	wire.Bind(new(service.TodolistService), new(*service.TodolistServiceImpl)),
	controller.NewTodolistController,
	wire.Bind(new(controller.TodolistController), new(*controller.TodolistControllerImpl)),
)

func InitializedView() *view.TodolistViewImpl {
	wire.Build(todolistSet, view.NewTodolistView)
	return nil
}
