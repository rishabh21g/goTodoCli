package main

import "github.com/fatih/color"

func main() {
	color.New(color.FgRed, color.Bold).Println("======================================TODO CLI APPLICATION=========================================")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		color.Red("Error loading todos: %v", err)
	}
	todos.add("Learning advance Golang concepts")
	todos.add("Doing App dev with React Native")
	todos.add("Doing system engineering in Golang")
	todos.listOfTodos()
	err = storage.Save(todos)
	if err != nil {
		color.Red("Error saving todos: %v", err)
	} else {
		color.Green("Todos saved successfully!")
	}

}
