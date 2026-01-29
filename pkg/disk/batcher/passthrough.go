package batcher

import (
	"context"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
)

type Passthrough[T any] struct {
	client desc.Client[T]
}

func NewPassthrough[T any](client desc.Client[T]) Passthrough[T] {
	return Passthrough[T]{
		client: client,
	}
}

func (w Passthrough[T]) Describe(ctx context.Context, id string) (*T, error) {
	resp, err := w.client.Describe([]string{id})
	if err != nil {
		return nil, err
	}
	return &resp.Resources[0], nil
}
