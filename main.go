package main

import "github.com/fatih/color"

func main() {
	color.Red("======================================TODO CLI APPLICATION=========================================")
	todos := Todos{}
	todos.add("Learning advance Golang concepts")
	todos.add("Doing App dev with React Native")
	todos.add("Doing system engineering in Golang")
	todos.listOfTodos()
	todos.delete(1)
	todos.listOfTodos()
	todos.toggleComplete(1)
	todos.listOfTodos()

}
