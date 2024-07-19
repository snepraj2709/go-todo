package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var todos []Todo

// Get all todos
func getAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Create a new todo
func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getAllTodo", getAllTodoHandler).Methods("GET")
	router.HandleFunc("/postTodo", postTodoHandler).Methods("POST")
	http.ListenAndServe(":8000", router)
}
