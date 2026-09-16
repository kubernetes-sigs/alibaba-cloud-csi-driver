//go:build !windows

package nas

import (
	"context"
	"errors"
	"testing"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/utils/clock"
	clocktesting "k8s.io/utils/clock/testing"
)

// An unexpected non-Describe call fails instead of silently returning a mock
// default. The polling policy is testable without constructing a full controller.
type agenticfsDescribeClient struct {
	interfaces.NasClientV2Interface
	describe func(context.Context, string, string) (*sdk.DescribeAccessPointResponse, error)
}

func (c agenticfsDescribeClient) DescribeAccesspoint(ctx context.Context, fs, ap string) (*sdk.DescribeAccessPointResponse, error) {
	return c.describe(ctx, fs, ap)
}

// Notify after Reset so tests advance virtual time only when the polling timer
// is armed, rather than racing its initial construction or an in-progress call.
type agenticfsPollClock struct {
	*clocktesting.FakeClock
	reset chan struct{}
}

func newAgenticfsPollClock() *agenticfsPollClock {
	return &agenticfsPollClock{
		FakeClock: clocktesting.NewFakeClock(time.Unix(1000, 0)),
		reset:     make(chan struct{}, 1),
	}
}

func (c *agenticfsPollClock) NewTimer(d time.Duration) clock.Timer {
	return &agenticfsPollTimer{Timer: c.FakeClock.NewTimer(d), reset: c.reset}
}

type agenticfsPollTimer struct {
	clock.Timer
	reset chan struct{}
}

func (t *agenticfsPollTimer) Reset(d time.Duration) bool {
	active := t.Timer.Reset(d)
	select {
	case t.reset <- struct{}{}:
	default:
	}
	return active
}

func awaitAgenticfsPollTimer(t *testing.T, c *agenticfsPollClock) {
	t.Helper()
	select {
	case <-c.reset:
	case <-time.After(5 * time.Second):
		t.Fatal("poll did not arm its timer")
	}
}

func awaitAgenticfsPollResult(t *testing.T, done <-chan error) error {
	t.Helper()
	select {
	case err := <-done:
		return err
	case <-time.After(5 * time.Second):
		t.Fatal("poll did not return")
		return nil
	}
}

func agenticfsDescribeResponse(state, domain string) *sdk.DescribeAccessPointResponse {
	return &sdk.DescribeAccessPointResponse{Body: &sdk.DescribeAccessPointResponseBody{
		AccessPoint: &sdk.DescribeAccessPointResponseBodyAccessPoint{
			Status: tea.String(state), DomainName: tea.String(domain),
		},
	}}
}

func TestAgenticfsPollRecoversReadAfterWriteAndTransientErrorsWithFakeClock(t *testing.T) {
	clk := newAgenticfsPollClock()
	calls := 0
	ctrl := &agenticfsController{
		clock: clk, apPollInterval: 3 * time.Second, apPollTimeout: 45 * time.Second,
		nasClient: agenticfsDescribeClient{describe: func(_ context.Context, fs, ap string) (*sdk.DescribeAccessPointResponse, error) {
			assert.Equal(t, "fs", fs)
			assert.Equal(t, "ap", ap)
			calls++
			switch calls {
			case 1:
				return nil, aliErr("InvalidAccessPoint.NotFound")
			case 2:
				return nil, errors.New("connection reset")
			case 3:
				return nil, nil // malformed response is retried within the budget
			case 4:
				return agenticfsDescribeResponse("Pending", ""), nil
			default:
				return agenticfsDescribeResponse(accessPointStatusActive, "ap.example.com"), nil
			}
		}},
	}
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	server := ""
	done := make(chan error, 1)
	go func() { done <- ctrl.waitAccessPointActive(ctx, "fs", "ap", &server) }()
	for range 4 {
		awaitAgenticfsPollTimer(t, clk)
		clk.Step(ctrl.apPollInterval)
	}
	require.NoError(t, awaitAgenticfsPollResult(t, done))
	assert.Equal(t, 5, calls)
	assert.Equal(t, "ap.example.com", server)
	assert.False(t, clk.HasWaiters(), "timers must be stopped on return")
}

func TestAgenticfsPollBudgetAndCompletionOrderWithFakeClock(t *testing.T) {
	for _, tt := range []struct {
		name       string
		deleting   bool
		complete   bool
		wantCode   codes.Code
		wantDetail string
	}{
		{"create timeout", false, false, codes.DeadlineExceeded, "last status: Pending"},
		{"delete timeout", true, false, codes.Aborted, "still being deleted"},
		// Existing policy: a call reaching the target wins even if it completes
		// after the polling budget. Refactoring must not reorder these checks.
		{"create completed after budget", false, true, codes.OK, ""},
		{"delete completed after budget", true, true, codes.OK, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			clk := newAgenticfsPollClock()
			calls := 0
			ctrl := &agenticfsController{
				clock: clk, apPollInterval: 3 * time.Second, apPollTimeout: 10 * time.Second,
				nasClient: agenticfsDescribeClient{describe: func(context.Context, string, string) (*sdk.DescribeAccessPointResponse, error) {
					calls++
					if tt.complete && calls == 2 {
						if tt.deleting {
							return nil, aliErr("InvalidAccessPoint.NotFound")
						}
						return agenticfsDescribeResponse(accessPointStatusActive, "new.example.com"), nil
					}
					return agenticfsDescribeResponse("Pending", ""), nil
				}},
			}
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()
			done := make(chan error, 1)
			server := "original.example.com"
			go func() {
				if tt.deleting {
					done <- ctrl.waitAccessPointGone(ctx, "fs", "ap")
				} else {
					done <- ctrl.waitAccessPointActive(ctx, "fs", "ap", &server)
				}
			}()
			awaitAgenticfsPollTimer(t, clk)
			clk.Step(ctrl.apPollTimeout + time.Second)
			err := awaitAgenticfsPollResult(t, done)
			assert.Equal(t, tt.wantCode, status.Code(err), "%v", err)
			if tt.wantDetail != "" {
				require.ErrorContains(t, err, tt.wantDetail)
			}
			assert.Equal(t, 2, calls)
			assert.Equal(t, "original.example.com", server, "do not replace an already known endpoint")
			assert.False(t, clk.HasWaiters())
		})
	}
}

func TestAgenticfsPollCancellationAndPermanentErrors(t *testing.T) {
	for _, tt := range []struct {
		name         string
		cancelBefore bool
		apiError     error
		wantCode     codes.Code
		wantCalls    int
	}{
		{"already cancelled", true, nil, codes.DeadlineExceeded, 0},
		{"cancel during wait", false, nil, codes.DeadlineExceeded, 1},
		{"permanent API rejection", false, aliErr("InvalidParameter.AccessPointId"), codes.InvalidArgument, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			clk := newAgenticfsPollClock()
			calls := 0
			ctrl := &agenticfsController{
				clock: clk, apPollInterval: 3 * time.Second, apPollTimeout: 45 * time.Second,
				nasClient: agenticfsDescribeClient{describe: func(context.Context, string, string) (*sdk.DescribeAccessPointResponse, error) {
					calls++
					return agenticfsDescribeResponse("Pending", ""), tt.apiError
				}},
			}
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()
			if tt.cancelBefore {
				cancel()
			}
			done := make(chan error, 1)
			go func() { done <- ctrl.waitAccessPointActive(ctx, "fs", "ap", nil) }()
			if !tt.cancelBefore && tt.apiError == nil {
				awaitAgenticfsPollTimer(t, clk)
				cancel()
			}
			err := awaitAgenticfsPollResult(t, done)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Equal(t, tt.wantCalls, calls)
			assert.False(t, clk.HasWaiters())
		})
	}
}

func TestAgenticfsAccessPointDeletionOrder(t *testing.T) {
	listed := []*sdk.ListAccessPointsResponseBodyAccessPoints{
		apItem("ap-second", accessPointStatusActive, ""),
		apItem("ap-pv", accessPointStatusActive, ""),
		apItem("", accessPointStatusActive, ""),
		apItem("ap-second", accessPointStatusActive, ""),
		apItem("ap-third", accessPointStatusActive, ""),
	}
	assert.Equal(t, []string{"ap-pv", "ap-second", "ap-third"}, accessPointDeletionOrder("ap-pv", listed))
	assert.Equal(t, []string{"ap-second", "ap-pv", "ap-third"}, accessPointDeletionOrder("", listed))
	assert.Equal(t, "ap-second", tea.StringValue(listed[0].AccessPointId), "leave the cloud snapshot unchanged")
	assert.Empty(t, accessPointDeletionOrder("", nil))
}
