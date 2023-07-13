package main

import (
	"FeiraVed/todolist/controller"
	"FeiraVed/todolist/core"
	"FeiraVed/todolist/helper"
	"FeiraVed/todolist/repository"
	"FeiraVed/todolist/service"
	"FeiraVed/todolist/view"
	"bytes"
	"fmt"
	"os"
	"time"
)

var Repository = repository.NewTodolistRepository()
var Service = service.NewTodolistService(Repository)
var Controller = controller.NewTodolistController(Service)
var View = view.NewTodolistView(Controller)

func ShowTodolist() {
	buffer := bytes.Buffer{}
	View.ShowTodolist(&buffer)
	fmt.Println(buffer.String())
}

func AddTodolist() {
	buffer := bytes.Buffer{}
	View.AddTodolist(&buffer, os.Stdin)
	fmt.Println(buffer.String())
}

func UpdateTodolist() {
	buffer := bytes.Buffer{}
	View.UpdateTodolist(&buffer, os.Stdin, os.Stdin)
	fmt.Println(buffer.String())
}

func DeleteTodolist() {
	buffer := bytes.Buffer{}
	View.Delete(&buffer, os.Stdin)
	fmt.Println(buffer.String())
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			main()
		}
	}()

	for {
		core.ClearScreen()
		core.Banner()
		ShowTodolist()
		fmt.Printf("%s[%s•%s]%s Menu Todolist\n", core.Green, core.Red, core.Green, core.Cyan)
		fmt.Printf(
			" %s╠══ %s[%s1%s]%s. Add Todolist\n",
			core.Yellow,
			core.Red,
			core.Yellow,
			core.Red,
			core.Green,
		)
		fmt.Printf(
			" %s╠══ %s[%s2%s]%s. Update Todolist\n",
			core.Yellow,
			core.Red,
			core.Yellow,
			core.Red,
			core.Green,
		)
		fmt.Printf(
			" %s╠══ %s[%s3%s]%s. Delete Todolist\n",
			core.Yellow,
			core.Red,
			core.Yellow,
			core.Red,
			core.Green,
		)
		fmt.Printf(
			" %s╠══ %s[%s4%s]%s. Exit/Close\n",
			core.Yellow,
			core.Red,
			core.Yellow,
			core.Red,
			core.Green,
		)
		fmt.Printf(" %s║", core.Yellow)

		menu := helper.InputMain(
			core.Yellow+" \n ╚══"+core.Green+"["+core.Yellow+"menu"+core.Green+"]: ",
			nil,
		)

		if menu == "1" {
			AddTodolist()
			time.Sleep(1 * time.Second)
		} else if menu == "2" {
			UpdateTodolist()
			time.Sleep(1 * time.Second)
		} else if menu == "3" {
			DeleteTodolist()
			time.Sleep(1 * time.Second)
		} else if menu == "4" {
			fmt.Println("[+] Bye bye...")
			return
		} else {
			fmt.Println("[+] Wrong input")
			time.Sleep(1 * time.Second)
		}
	}
}
