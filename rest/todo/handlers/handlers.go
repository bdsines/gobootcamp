package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"rest/todo/repos"

	"github.com/gorilla/mux"
)

// var todos repos.Todos

// func Start() {

// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", Index)
// 	router.HandleFunc("/todos", TodoIndex)
// 	router.HandleFunc("/todos/{todoId}", ShowTodo)
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo repos.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := repos.CreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// todos = append(todos, repos.Todo{Name: "Getting Started w/ Go"})
	// todos = append(todos, repos.Todo{Name: "Installation Basics"})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repos.FindAllTodo()); err != nil {
		panic(err)
	}
	// fmt.Fprintln(w, "Show All")

}
func ShowTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoIdStr := vars["todoId"]
	todoId, _ := strconv.Atoi(todoIdStr)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repos.FindTodo(todoId)); err != nil {
		panic(err)
	}
	// fmt.Fprintln(w, "Todo show:", todoId)
}
