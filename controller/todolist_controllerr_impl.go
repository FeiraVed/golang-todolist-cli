package controller

import (
	"FeiraVed/todolist/entity"
	"FeiraVed/todolist/entity/app"
	"FeiraVed/todolist/helper"
	"FeiraVed/todolist/service"
	"fmt"
	"io"
	"strings"
)

type TodolistControllerImpl struct {
	service service.TodolistService
}

func NewTodolistController(service service.TodolistService) *TodolistControllerImpl {
	return &TodolistControllerImpl{
		service: service,
	}
}

func (controller *TodolistControllerImpl) Create(w io.Writer, name string) {

	if helper.Validate(name) {
		request := app.TodolistCreateRequest{
			Name: name,
		}

		response := controller.service.Create(request)
		todolist := entity.Todolist{
			Id:   response.Id,
			Name: response.Name,
		}
		fmt.Fprintln(w, todolist)

	} else {
		panic("[+] Name cannot be empty")
	}
}

func (controller *TodolistControllerImpl) Update(w io.Writer, id int, name string) {
	if helper.ValidateId(id) && helper.Validate(name) {

		request := app.TodolistUpdateRequest{
			Id:   id,
			Name: name,
		}

		response := controller.service.Update(request)
		todolist := entity.Todolist{
			Id:   response.Id,
			Name: response.Name,
		}
		fmt.Fprintln(w, todolist)

	} else {
		panic("[+] Id or name invalid")
	}

}

func (controller *TodolistControllerImpl) Delete(id int) {
	if helper.ValidateId(id) {
		controller.service.Delete(id)
	} else {
		panic("[+] Id not found")
	}
}
func (controller *TodolistControllerImpl) FindById(w io.Writer, id int) {
	if helper.ValidateId(id) {
		response := controller.service.FindById(id)

		todolist := entity.Todolist{
			Id:   response.Id,
			Name: response.Name,
		}
		fmt.Fprintln(w, todolist)
	} else {
		panic("[+] Id not found")
	}
}
func (controller *TodolistControllerImpl) FindAll(w io.Writer) {
	var data []string

	responses := controller.service.FindAll()
	for _, response := range responses {
		name := strings.Replace(response.Name, " ", "+", -1)
		data = append(data, name)
	}
	fmt.Fprint(w, data)
}
