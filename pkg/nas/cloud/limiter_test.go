package cloud

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestActionLimiterSlowStart(t *testing.T) {
	t.Parallel()
	a := newActionLimiter("CreateAccessPoint", 2, 20)
	base := time.Now()
	cur := base
	a.now = func() time.Time { return cur }
	a.lastChange = base

	cur = base.Add(slowStartInterval / 2)
	a.OnSuccess()
	assert.Equal(t, 2.0, a.limit, "must not grow before the slow start interval")

	for _, want := range []float64{4, 8, 16, 20} {
		cur = cur.Add(slowStartInterval + time.Millisecond)
		a.OnSuccess()
		assert.Equal(t, want, a.limit)
	}
	assert.Equal(t, rate.Limit(20), a.limiter.Limit(), "must stop at max")
}

func TestActionLimiterThrottleEndsSlowStart(t *testing.T) {
	t.Parallel()
	a := newActionLimiter("CreateAccessPoint", 2, 100)
	base := time.Now()
	cur := base
	a.now = func() time.Time { return cur }
	a.lastChange = base

	cur = cur.Add(slowStartInterval + time.Millisecond)
	a.OnSuccess()
	cur = cur.Add(slowStartInterval + time.Millisecond)
	a.OnSuccess()
	assert.Equal(t, 8.0, a.limit)

	a.OnThrottle()
	assert.Equal(t, 6.0, a.limit, "8 * 0.75")
	assert.False(t, a.slowStart, "slow start ends at the first throttle")

	// Recovery is additive from here, so a raise is picked up without
	// producing another overshoot of the same size.
	cur = cur.Add(steadyIncreaseInterval + time.Millisecond)
	a.OnSuccess()
	assert.Equal(t, 7.0, a.limit)
}

func TestActionLimiterThrottleDeduped(t *testing.T) {
	t.Parallel()
	// A burst of concurrent throttles is one event. Decreasing once per
	// response would put a 100-worker burst on the floor immediately.
	a := newActionLimiter("CreateAccessPoint", 40, 40)
	cur := time.Now()
	a.now = func() time.Time { return cur }

	for i := 0; i < 100; i++ {
		a.OnThrottle()
	}
	assert.Equal(t, 30.0, a.limit, "one decrease per dedupe window")
}

func TestActionLimiterFloor(t *testing.T) {
	t.Parallel()
	a := newActionLimiter("CreateAccessPoint", 2, 2)
	cur := time.Now()
	a.now = func() time.Time { return cur }

	for _, want := range []float64{1.5, 1.125, 1} {
		cur = cur.Add(throttleDedupeWindow + time.Millisecond)
		a.OnThrottle()
		assert.Equal(t, want, a.limit)
	}
	cur = cur.Add(throttleDedupeWindow + time.Millisecond)
	a.OnThrottle()
	assert.Equal(t, 1.0, a.limit, "must stop at the floor")
	assert.Equal(t, rate.Limit(1), a.limiter.Limit())
}
