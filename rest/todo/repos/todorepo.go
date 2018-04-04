package repos

import (
	"fmt"
)

var todos Todos
var currentID int

func init() {
	CreateTodo(Todo{Name: " Enhancing repo"})
	CreateTodo(Todo{Name: " Adding CRUD "})

}

/* CreateToDo */
func CreateTodo(t Todo) Todo {
	currentID += 1
	t.Id = currentID
	todos = append(todos, t)
	return t

}
func FindAllTodo() []Todo {
	return todos
}
func FindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}
func DeleteTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with ID: %d", id)
}
