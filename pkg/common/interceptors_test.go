package common

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEarlyTimeout_None(t *testing.T) {
	handler := func(ctx context.Context, _ interface{}) (interface{}, error) {
		_, ok := ctx.Deadline()
		assert.False(t, ok)
		return struct{}{}, nil
	}
	earlyTimeout(context.Background(), nil, nil, handler)
}

func TestEarlyTimeout(t *testing.T) {
	cases := []struct {
		before, after time.Duration
	}{
		{60 * time.Second, 59 * time.Second},
		{1 * time.Second, 900 * time.Millisecond},
	}
	for _, tc := range cases {
		t.Run(tc.before.String(), func(t *testing.T) {
			now := time.Now()
			handler := func(ctx context.Context, _ interface{}) (interface{}, error) {
				deadline, ok := ctx.Deadline()
				assert.True(t, ok)
				assert.WithinDuration(t, now.Add(tc.after), deadline, 10*time.Millisecond)
				return struct{}{}, nil
			}
			ctx, cancel := context.WithDeadline(context.Background(), now.Add(tc.before))
			t.Cleanup(cancel)
			earlyTimeout(ctx, nil, nil, handler)
		})
	}

}
