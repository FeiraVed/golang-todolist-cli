package helper

import "FeiraVed/todolist/db"

func ValidateId(id int) bool {
	data := db.TodolistDB
	keys := SortTodolistKey(data)

	if id > len(keys) {
		return false
	} else if id < 1 {
		return false
	}

	return true
}
