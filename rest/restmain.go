package main

import (
	"log"
	"net/http"
	"rest/todo/routes"
)

func main() {
	todoRouter := todo.TodoRouter()
	log.Fatal(http.ListenAndServe(":8080", todoRouter))
}
