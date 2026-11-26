package cloud

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

const (
	// slowStartInterval is how often the rate doubles before the first
	// Throttling response. Starting at the published floor and doubling takes
	// an action from 2 QPS to 64 QPS in about five seconds, so a fresh process
	// finds the account quota before the first batch is done.
	slowStartInterval = time.Second
	// steadyIncreaseInterval is the additive-increase cadence after the first
	// throttle. It is deliberately slower than slow start: from here on every
	// increase is a probe that may cost a rejected request.
	steadyIncreaseInterval = 5 * time.Second
	// steadyIncreaseStep is the additive increase in QPS per interval.
	steadyIncreaseStep = 1.0
	// decreaseFactor multiplies the rate on a Throttling response. Halving is
	// too coarse: an action that overshoots a 40 QPS quota from 64 lands at 20
	// and needs tens of seconds of additive increase to get back.
	decreaseFactor = 0.75
	// throttleDedupeWindow collapses a burst of concurrent throttles into one
	// event. Without it 100 concurrent responses would apply 100 decreases and
	// put the action on the floor before any of them recovers.
	throttleDedupeWindow = 5 * time.Second
)

// actionLimiter paces one OpenAPI action. The NAS quota is enforced per
// (region, API, account), so each action discovers its own quota instead of
// sharing whatever a deployment guessed:
//
//   - slow start: from the published floor, double the rate every
//     slowStartInterval while calls succeed;
//   - the first Throttling response ends slow start and multiplies the rate by
//     decreaseFactor; further throttles within throttleDedupeWindow change
//     nothing;
//   - from then on the rate grows by steadyIncreaseStep every
//     steadyIncreaseInterval, so a raised account quota is picked up without a
//     restart, and a quota shared with another consumer is tracked without
//     hammering it.
//
// max bounds every phase, so a generous account never sends more than the
// configured ceiling.
type actionLimiter struct {
	name    string
	mu      sync.Mutex
	limiter *rate.Limiter
	limit   float64
	min     float64
	max     float64
	// slowStart is true until the first Throttling response.
	slowStart bool
	// lastChange is when limit last moved; it gates both directions.
	lastChange time.Time
	// lastDecrease gates the decrease separately from lastChange: a burst of
	// concurrent throttles must count as one event.
	lastDecrease time.Time
	now          func() time.Time
}

func newActionLimiter(name string, qps, maxQPS float64) *actionLimiter {
	if qps < 1 {
		qps = 1
	}
	if maxQPS < qps {
		maxQPS = qps
	}
	return &actionLimiter{
		name:       name,
		limiter:    rate.NewLimiter(rate.Limit(qps), burstFor(qps)),
		limit:      qps,
		min:        1,
		max:        maxQPS,
		slowStart:  true,
		lastChange: time.Now(),
		now:        time.Now,
	}
}

// burstFor keeps a burst within one second of quota: a larger burst would put
// more calls on the wire in that second than the service accepts, even when
// the sustained rate is within quota.
func burstFor(qps float64) int {
	burst := int(qps)
	if burst < 1 {
		burst = 1
	}
	if burst > 10 {
		burst = 10
	}
	return burst
}

func (a *actionLimiter) Wait(ctx context.Context) error {
	return a.limiter.Wait(ctx)
}

// OnSuccess grows the rate: multiplicatively during slow start, additively
// afterwards, in both cases up to max.
func (a *actionLimiter) OnSuccess() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.limit >= a.max {
		return
	}
	interval := steadyIncreaseInterval
	if a.slowStart {
		interval = slowStartInterval
	}
	if a.now().Sub(a.lastChange) < interval {
		return
	}
	next := a.limit + steadyIncreaseStep
	if a.slowStart {
		next = a.limit * 2
	}
	if next > a.max {
		next = a.max
	}
	a.limit = next
	a.apply()
}

// OnThrottle lowers the rate. The service rejected the call because a quota was
// exceeded, so the client was sending too fast for this action.
func (a *actionLimiter) OnThrottle() {
	a.mu.Lock()
	defer a.mu.Unlock()
	now := a.now()
	if !a.lastDecrease.IsZero() && now.Sub(a.lastDecrease) < throttleDedupeWindow {
		return
	}
	a.slowStart = false
	next := a.limit * decreaseFactor
	if next < a.min {
		next = a.min
	}
	a.lastDecrease = now
	a.lastChange = now
	if next != a.limit {
		a.limit = next
		a.apply()
	}
}

func (a *actionLimiter) apply() {
	a.lastChange = a.now()
	a.limiter.SetLimit(rate.Limit(a.limit))
	a.limiter.SetBurst(burstFor(a.limit))
	klog.InfoS("nas rate limit changed", "action", a.name, "qps", a.limit, "slowStart", a.slowStart)
}
