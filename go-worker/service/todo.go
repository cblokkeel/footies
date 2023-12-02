package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
)

type TodoService struct {
	cache  cache.Cacher
	client *client.JsonPlaceholderClient
}

func NewTodoService(cache cache.Cacher, client *client.JsonPlaceholderClient) *TodoService {
	return &TodoService{
		cache,
		client,
	}
}

func (s *TodoService) GetTodo(ctx context.Context) (*client.Todo, error) {
	cacheKey := s.getCacheKeyForTodo(1) // Mocking id here, in real scenario we will get the id from the req
	var cachedTodo client.Todo
	if exists := s.cache.Get(ctx, cacheKey, &cachedTodo); exists {
		fmt.Println("Get from cache")
		return &cachedTodo, nil
	}

	todo, err := s.client.GetTodo()
	if err != nil {
		return nil, err
	}
	s.cache.Set(ctx, cacheKey, todo, time.Minute) // ttl should be in env
	return todo, nil
}

func (s *TodoService) getCacheKeyForTodo(id int) string {
	return fmt.Sprintf("todo_%d", id)
}
