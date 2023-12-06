package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cblokkeel/footies/api"
	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
	"github.com/cblokkeel/footies/constants"
	"github.com/cblokkeel/footies/job"
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
	footballAPIClient := client.NewFootballAPIClient(baseClient, os.Getenv(constants.FootballApiUrl), os.Getenv(constants.FootballApiKey))
	matchService := service.NewMatchService(redisCache, pubsub, footballAPIClient)

	api := api.NewApi(matchService)
	go api.Start()

	job := job.NewJob(pubsub, matchService)
	go job.Start()

	select {} // keep running
}
