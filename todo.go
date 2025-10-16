package main

import (
	"errors"
	"fmt"
	"time"
)

// Todo represents a single todo item with its details structured as a Go struct.
type Todo struct {
	Title      string
	IsComplete bool
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

// Todos is a slice of Todo items, allowing for easy management of multiple todos.
type Todos []Todo // act like class name

// add is a method on Todos that allows adding a new todo item.
func (todos *Todos) add(title string) {

	todo := Todo{
		Title:      title,
		CreatedAt:  time.Now(),
		IsComplete: false,
		UpdatedAt:  nil,
	}
	*todos = append(*todos, todo)
}

// validateIndex checks if the provided index is valid for the Todos slice.
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

// deleting the todo on the basis of index

func (todos *Todos) delete(index int) error {
	t := *todos
	err := t.validateIndex(index)
	if err != nil {
		return err
	}
	*todos = append((t)[:index], (t)[index+1:]...)
	return nil
}

// printing all the todos

func (todos *Todos) listOfTodos() {
	for _, todo := range *todos {
		fmt.Printf("%+v\n", todo)
	}
}

// markComplete marks a todo item as complete based on its index.
func (todos *Todos) toggleComplete(index int) error {
	t := *todos
	err := t.validateIndex(index)
	if err != nil {
		return err
	}
	foundTodo := &t[index]
	foundTodo.IsComplete = !foundTodo.IsComplete
	currentTime := time.Now()
	foundTodo.UpdatedAt = &currentTime
	return nil
}
