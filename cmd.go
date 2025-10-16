package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

type Command struct {
	Add    string
	Del    int
	Edit   int
	Set    string
	Toggle int
	List   bool
}

func NewCommand() *Command {
	cf := Command{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo item")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo item by index")
	flag.IntVar(&cf.Edit, "edit", -1, "Edit a todo item by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the done status of a todo item by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")

	flag.Parse()
	return &cf

}

func (c *Command) Execute(todos *Todos) error {
	switch {
	case c.Add != "":
		todos.add(c.Add)
	case c.Del != -1:
		return todos.delete(c.Del)
	case c.Edit != -1:
		if c.Set == "" {
			return fmt.Errorf("missing -set\"new title\" for --edit")
		}
		return todos.edit(c.Edit, c.Set)
	case c.Toggle != -1:
		return todos.toggleComplete(c.Toggle)
	case c.List:
		todos.listOfTodos()
	default:
		color.Blue("Invalid Commands")
	}
	return nil

}
