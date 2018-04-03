package todo

import (
	"net/http"
	"rest/todo/handlers"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"ToDoIndex",
		"GET",
		"/todos",
		handlers.TodoIndex,
	},
	Route{
		"ShowToDo",
		"GET",
		"/todos/{todoId}",
		handlers.ShowTodo,
	},
}
