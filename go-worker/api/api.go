package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cblokkeel/footies/client"
)

type Mock struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type Api struct {
	jsonPlaceholderClient *client.JsonPlaceholderClient
}

func NewApi(jsonPlaceholderClient *client.JsonPlaceholderClient) *Api {
	return &Api{
		jsonPlaceholderClient,
	}
}

func (a *Api) Start() {
	http.HandleFunc("/todo", a.handleGetTodo)
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// TODO: separate the logic from this endpoint handler, in a service for exemple
func (a *Api) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := a.jsonPlaceholderClient.GetTodo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	json.NewEncoder(w).Encode(todo)
}
