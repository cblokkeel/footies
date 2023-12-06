package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPubSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{
		client,
	}
}

func (p *RedisPubSub) Publish(ctx context.Context, channel string, msg interface{}) error {
	return p.client.Publish(ctx, channel, msg).Err()
}

func (p *RedisPubSub) Subscribe(ctx context.Context, channel string) (<-chan *redis.Message, error) {
	sub := p.client.Subscribe(ctx, channel)
	if _, err := sub.Receive(ctx); err != nil {
		return nil, err
	}
	return sub.Channel(), nil
}
