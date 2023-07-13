package service

import (
	"FeiraVed/todolist/db"
	"FeiraVed/todolist/entity"
	"FeiraVed/todolist/entity/app"
	"FeiraVed/todolist/helper"
	"FeiraVed/todolist/repository"
)

type TodolistServiceImpl struct {
	repository repository.TodolistRepository
}

// INFO: Simulasi basis data key-value
var data = db.TodolistDB

func NewTodolistService(todolistRepository repository.TodolistRepository) TodolistService {

	return &TodolistServiceImpl{
		repository: todolistRepository,
	}
}

func (service *TodolistServiceImpl) Create(
	request app.TodolistCreateRequest,
) (_ *app.TodolistResponse) {

	todolist := entity.Todolist{
		Name: request.Name,
	}

	response := service.repository.Save(data, &todolist)
	appResponse := app.TodolistResponse{
		Id:   response.Id,
		Name: response.Name,
	}

	return &appResponse

}

func (service *TodolistServiceImpl) Update(
	request app.TodolistUpdateRequest,
) (_ *app.TodolistResponse) {

	keys := helper.SortTodolistKey(data)

	if request.Id > len(keys) {
		panic("[+] Id not found")
	} else if request.Id < 1 {
		panic("[+] Id not found")
	}

	todolist := entity.Todolist{
		Id:   keys[request.Id-1],
		Name: request.Name,
	}

	response := service.repository.Update(data, &todolist)

	return &app.TodolistResponse{
		Id:   response.Id,
		Name: response.Name,
	}
}

func (service *TodolistServiceImpl) Delete(id int) {

	keys := helper.SortTodolistKey(data)
	if id > len(keys) {
		panic("[+] Id not found")
	} else if id < 1 {
		panic("[+] Id not found")
	}
	todolist := entity.Todolist{
		Id: keys[id-1],
	}
	service.repository.Delete(data, &todolist)
}

func (service *TodolistServiceImpl) FindById(id int) (_ *app.TodolistResponse) {
	response, err := service.repository.FindById(data, id)
	helper.PanicIfError(err)

	return &app.TodolistResponse{
		Id:   response.Id,
		Name: response.Name,
	}

}

func (service *TodolistServiceImpl) FindAll() (_ []app.TodolistResponse) {
	responses := []app.TodolistResponse{}

	keys := helper.SortTodolistKey(data)
	for i := 0; i < len(keys); i++ {
		todolist := app.TodolistResponse{
			Id:   keys[i],
			Name: data[keys[i]],
		}
		responses = append(responses, todolist)
	}
	return responses
}
