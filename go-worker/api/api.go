package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cblokkeel/footies/service"
)

type Api struct {
	todoService *service.TodoService
}

func NewApi(todoService *service.TodoService) *Api {
	return &Api{
		todoService,
	}
}

func (a *Api) Start() {
	http.HandleFunc("/todo", a.handleGetTodo)
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func (a *Api) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := a.todoService.GetTodo(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
