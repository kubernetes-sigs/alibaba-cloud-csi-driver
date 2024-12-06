package waitstatus

import (
	"context"
	"errors"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"k8s.io/utils/clock"
)

type Simple[T any] struct {
	client desc.Client[T]
	clk    clock.WithTicker

	ClientFactory func(context.Context) (desc.Client[T], error)
}

func NewSimple[T any](client desc.Client[T], clk clock.WithTicker) *Simple[T] {
	return &Simple[T]{
		client: client,
		clk:    clk,
	}
}

func (w *Simple[T]) WaitFor(ctx context.Context, id string, pred StatusPredicate[*T]) (*T, error) {
	client := w.client
	if w.ClientFactory != nil {
		var err error
		client, err = w.ClientFactory(ctx)
		if err != nil {
			return nil, err
		}
	}
	ticker := w.clk.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		resp, err := client.Describe([]string{id})
		if err != nil {
			return nil, err
		}
		if len(resp.Resources) == 0 {
			return nil, nil
		}
		if len(resp.Resources) > 1 {
			return nil, errors.New("too many resources returned")
		}
		res := &resp.Resources[0]
		if pred(res) {
			return res, nil
		}
		select {
		case <-ticker.C():
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
