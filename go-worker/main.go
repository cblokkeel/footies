package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cblokkeel/footies/api"
	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
	"github.com/cblokkeel/footies/pubsub"
	"github.com/cblokkeel/footies/service"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No env file found")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := redisClient.Set(context.Background(), "key", "val", time.Second).Err(); err != nil {
		panic(fmt.Errorf("Couldnt connect to redis"))
	}
	redisCache := cache.NewRedisCache(redisClient)
	pubsub := pubsub.NewRedisPubSub(redisClient)

	baseClient := client.NewClient(&http.Client{Timeout: time.Second * 10})
	footballAPIClient := client.NewFootballAPIClient(baseClient, "https://api-football-v1.p.rapidapi.com/v3/fixtures")
	matchService := service.NewMatchService(redisCache, pubsub, footballAPIClient)

	api := api.NewApi(matchService)
	go func() {
		api.Start()
	}()

	ch, err := pubsub.Subscribe(context.Background(), "monitoring")
	if err != nil {
		panic(err)
	}
	for msg := range ch {
		matchID := msg.Payload
		go matchService.MonitorMatch(context.Background(), matchID)
	}
}
