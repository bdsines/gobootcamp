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

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo repos.Todo
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		respondWithError(w, http.StatusBadRequest, " Bad Payload")
		return
	}

	// if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
	// 	panic(err)
	// 	respondWithError(w, http.StatusBadRequest, "Bad Request")
	// 	return
	// }
	// log.Printf("todo unmarshalled ", todo)
	t := repos.CreateTodo(todo)
	respondWithJson(w, http.StatusOK, t)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
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
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo repos.Todo
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
		respondWithError(w, http.StatusBadRequest, " Bad Payload")
		return
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		panic(err)
		respondWithError(w, http.StatusBadRequest, " Bad Payload")
		return
	}
	err = repos.UpdateTodo(todo)
	if err != nil {
		panic(err)
		respondWithError(w, http.StatusBadRequest, " Failed to Update")
		return
	}
	respondWithJson(w, http.StatusOK, err)
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoIdStr := vars["todoId"]
	todoId, _ := strconv.Atoi(todoIdStr)
	err := repos.DeleteTodo(todoId)

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(repos.DeleteTodo(todoId)); err != nil {
	if err != nil {
		panic(err)
		respondWithError(w, http.StatusBadRequest, " Failed to Delete")
	}
	respondWithJson(w, http.StatusOK, err)
	// fmt.Fprintln(w, "Todo show:", todoId)
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}
