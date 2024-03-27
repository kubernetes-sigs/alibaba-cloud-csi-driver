package waitstatus

import (
	"context"
	"errors"

	"k8s.io/utils/clock"
)

type Simple[T any] struct {
	client ECSDescribeResources[T]
	clk    clock.WithTicker
}

func NewSimple[T any](client ECSDescribeResources[T], clk clock.WithTicker) *Simple[T] {
	return &Simple[T]{
		client: client,
		clk:    clk,
	}
}

func (w *Simple[T]) WaitFor(ctx context.Context, id string, pred StatusPredicate[*T]) (*T, error) {
	ticker := w.clk.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		resp, err := w.client.Describe([]string{id})
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
