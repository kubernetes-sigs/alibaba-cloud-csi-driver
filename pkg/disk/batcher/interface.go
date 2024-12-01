package batcher

import "context"

type Batcher[T any] interface {
	Describe(ctx context.Context, id string) (*T, error)
}
