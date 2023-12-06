package job

import (
	"context"

	"github.com/cblokkeel/footies/pubsub"
	"github.com/cblokkeel/footies/service"
	"github.com/redis/go-redis/v9"
)

type Job struct {
	pubsub       pubsub.PubSub[*redis.Message]
	matchService *service.MatchService
}

func NewJob(pubsub pubsub.PubSub[*redis.Message], matchService *service.MatchService) *Job {
	return &Job{
		pubsub,
		matchService,
	}
}

func (j *Job) Start() {
	go j.startMonitoringJob()
}

func (j *Job) startMonitoringJob() {
	ch, err := j.pubsub.Subscribe(context.Background(), "monitoring")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case msg := <-ch:
			matchID := msg.Payload
			go j.matchService.MonitorMatch(context.Background(), matchID)
		}
	}
}
