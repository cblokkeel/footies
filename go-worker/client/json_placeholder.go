package client

import (
	"fmt"
)

type JsonPlaceholderClient struct {
	*Client
}

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var url = "https://jsonplaceholder.typicode.com/todos/1"

func NewJsonPlaceholderClient(baseClient *Client) *JsonPlaceholderClient {
	return &JsonPlaceholderClient{
		Client: baseClient,
	}
}

func (c *JsonPlaceholderClient) GetTodo() (*Todo, error) {
	var todo Todo
	if err := c.GetReq(url, &todo); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error while fetching todo")
	}
	return &todo, nil
}
