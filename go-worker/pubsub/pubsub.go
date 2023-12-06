package pubsub

import (
	"context"
)

type PubSub[T any] interface {
	Publish(context.Context, string, interface{}) error
	Subscribe(context.Context, string) (<-chan T, error)
}
