package cache

import (
	"context"
	"time"
)

type Cacher interface {
	Get(context.Context, string, interface{}) bool
	Set(context.Context, string, interface{}, time.Duration) error
	HasSet(context.Context, string, string) bool
	AddSet(context.Context, string, string) error
	RemoveSet(context.Context, string, string) error
}
