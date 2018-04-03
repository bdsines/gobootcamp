package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"rest/todo/repos"

	"github.com/gorilla/mux"
)

var todos repos.Todos

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
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos = append(todos, repos.Todo{Name: "Getting Started w/ Go"})
	todos = append(todos, repos.Todo{Name: "Installation Basics"})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
	// fmt.Fprintln(w, "Show All")

}
func ShowTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
