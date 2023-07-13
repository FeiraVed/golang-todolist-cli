package view

import (
	"FeiraVed/todolist/controller"
	"FeiraVed/todolist/helper"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TodolistViewImpl struct {
	controller controller.TodolistController
}

func NewTodolistView(controller controller.TodolistController) TodolistView {
	return &TodolistViewImpl{
		controller: controller,
	}
}

var InputName = helper.InputName
var InputId = helper.InputId

func (view *TodolistViewImpl) ShowTodolist(w io.Writer) {
	fmt.Println("")
	fmt.Println("======= Todolist =======")
	buffer := bytes.Buffer{}
	view.controller.FindAll(&buffer)

	data := regexp.MustCompile(`[^a-zA-Z+]+`)

	text := data.ReplaceAllString(buffer.String(), " ")

	newData := strings.Trim(text, " ")
	s := strings.Split(newData, " ")

	for i, v := range s {
		fmt.Fprintln(w, strconv.Itoa(i+1)+".", strings.Replace(v, "+", " ", -1))
	}
}

func (view *TodolistViewImpl) AddTodolist(w io.Writer, input *os.File) {
	buffer := bytes.Buffer{}
	fmt.Println("===== Add Todolist =====")
	fmt.Print("[+] name?: ")
	name := InputName(input)

	view.controller.Create(&buffer, name)
	fmt.Fprintln(w, "[+] Success add todolist")
}
func (view *TodolistViewImpl) UpdateTodolist(w io.Writer, inputId *os.File, inputName *os.File) {
	fmt.Println("===== Update Todolist =====")
	buffer := bytes.Buffer{}
	fmt.Print("[+] Id?: ")
	id := InputId(inputId)
	fmt.Print("[+] Name?: ")
	name := InputName(inputName)

	view.controller.Update(&buffer, id, name)
	fmt.Fprintln(w, "[+] Success update todolist,id :", id)
}
func (view *TodolistViewImpl) Delete(w io.Writer, input *os.File) {

	fmt.Println("===== Delete Todolist =====")
	fmt.Print("[+] Id?: ")
	id := InputId(input)
	view.controller.Delete(id)
	fmt.Fprintln(w, "[+] Success delete todolist,id :", id)
}
