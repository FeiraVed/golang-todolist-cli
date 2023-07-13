package repository

import (
	"FeiraVed/todolist/entity"
	"FeiraVed/todolist/helper"
	"errors"
	"strconv"
)

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() TodolistRepository {

	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Save(
	db DB,
	todolist *entity.Todolist,
) (_ *entity.Todolist) {
	db["todo-"+strconv.Itoa(len(db)+1)] = todolist.Name

	keys := helper.SortTodolistKey(db)
	lastId := keys[len(keys)-1]
	todolist.Id = lastId

	return todolist
}

func (repository *TodolistRepositoryImpl) Update(
	db DB,
	todolist *entity.Todolist,
) (_ *entity.Todolist) {

	db[todolist.Id] = todolist.Name

	keys := helper.SortTodolistKey(db)
	lastId := keys[len(keys)-1]
	todolist.Id = lastId

	return todolist

}
func (repository *TodolistRepositoryImpl) Delete(db DB, todolist *entity.Todolist) {
	delete(db, todolist.Id)
}

func (repository *TodolistRepositoryImpl) FindById(
	db DB,
	todolistId int,
) (_ *entity.Todolist, _ error) {

	keys := helper.SortTodolistKey(db)
	if todolistId > len(keys) {
		return nil, errors.New("[+] Id not found")
	} else if todolistId == 0 {
		return nil, errors.New("[+] Id not found")
	}

	todolist := entity.Todolist{
		Id:   keys[todolistId-1],
		Name: db[keys[todolistId-1]],
	}

	return &todolist, nil
}

func (repository *TodolistRepositoryImpl) FindAll(db DB) (_ []entity.Todolist) {
	data := []entity.Todolist{}
	keys := helper.SortTodolistKey(db)
	for i := 0; i < len(keys); i++ {
		todolist := entity.Todolist{
			Id:   keys[i],
			Name: db[keys[i]],
		}
		data = append(data, todolist)
	}
	return data
}
