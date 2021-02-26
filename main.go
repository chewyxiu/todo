package main

import (
	"github.com/gorilla/mux"
	"github.com/todo/api"
	"github.com/todo/dto"
	"github.com/todo/storage"
	"github.com/todo/utils"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos/create", utils.Prepare(dto.CreateTodoRequest{}, api.CreateTodo)).Methods("POST")
	router.HandleFunc("/todos/update", utils.Prepare(dto.UpdateTodoRequest{}, api.UpdateTodo)).Methods("POST")
	router.HandleFunc("/todos", utils.Prepare(dto.GetTodosRequest{}, api.GetTodos)).Methods("GET")
	http.Handle("/", router)

	storage.Init()

	log.Println("Server started at port 8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Println(err)
	}
}