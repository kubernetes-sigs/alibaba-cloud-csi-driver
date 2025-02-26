// batcher package allows for batching of requests to the cloud provider
package batcher

import "context"

type Batcher[T any] interface {
	Describe(ctx context.Context, id string) (*T, error)
}
