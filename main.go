package main

import "github.com/fatih/color"

func main() {
	color.New(color.FgRed, color.Bold).Println("======================================TODO CLI APPLICATION=========================================")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	cf := NewCommand()
	cf.Execute(&todos)
	todos.listOfTodos()
	if err != nil {
		color.Red("Error loading todos: %v", err)
	}

}
