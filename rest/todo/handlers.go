package todo

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", ShowTodo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// var todos = Todos{
	// 	Todo{"name": "Getting Started w/ Go"},
	// 	Todo{"name": "Explore REST Basics"},
	// }
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
	fmt.Fprintln(w, "Todo Index!")

}
func ShowTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
