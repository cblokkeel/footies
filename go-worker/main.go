package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cblokkeel/footies/api"
	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
	"github.com/cblokkeel/footies/service"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := redisClient.Set(context.Background(), "key", "val", time.Second).Err(); err != nil {
		panic(fmt.Errorf("Couldnt connect to redis"))
	}
	redisCache := cache.NewRedisCache(redisClient)

	baseClient := client.NewClient(&http.Client{Timeout: time.Second * 10})
	jsonPlaceHolderClient := client.NewJsonPlaceholderClient(baseClient)
	todoService := service.NewTodoService(redisCache, jsonPlaceHolderClient)

	api := api.NewApi(todoService)
	api.Start()
}
