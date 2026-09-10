package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
)

const stuckMountError = "mount timeout, the reason a caller needs to see"

// stuckDriver never finishes a mount, then winds down the way the FUSE drivers
// do: it waits MountShutdownGrace after its context fires before returning the
// error. That delay is what a caller's deadline has to be kept clear of, so a
// driver without it would let this suite pass while the real ones still lose
// their errors.
type stuckDriver struct {
	fstype string
	// returned reports when the error was ready, to tell an error that raced the
	// caller apart from one that arrived in time.
	returned chan time.Time
}

func (d *stuckDriver) Name() string      { return d.fstype }
func (d *stuckDriver) Fstypes() []string { return []string{d.fstype} }
func (d *stuckDriver) Init()             {}
func (d *stuckDriver) Terminate()        {}
func (d *stuckDriver) ApplyOptionDefaults(options []string) []string {
	return options
}

func (d *stuckDriver) Mount(ctx context.Context, _ *proxy.MountRequest, _ int) error {
	<-ctx.Done()
	time.Sleep(proxy.MountShutdownGrace)
	d.returned <- time.Now()
	return fmt.Errorf("%s", stuckMountError)
}

func registerStuckDriver(t *testing.T) *stuckDriver {
	t.Helper()
	d := &stuckDriver{fstype: "stuck-" + t.Name(), returned: make(chan time.Time, 1)}
	server.RegisterDriver(d)
	server.Init([]string{d.Name()})
	return d
}

// A caller that runs out of time has to receive what the driver said, not its
// own deadline: that error is the whole reason the reserve exists.
func TestTimedOutMountReportsItsOwnError(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	driver := registerStuckDriver(t)
	socketPath := newTestServer(t)

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	callerDeadline, ok := ctx.Deadline()
	require.True(t, ok)

	resp, err := client.NewClient(socketPath).Mount(ctx, &proxy.MountRequest{
		Fstype: driver.fstype,
		Source: "stuck://bucket",
		Target: t.TempDir(),
	})

	require.NoError(t, err, "the request itself must complete; only the mount fails")
	require.NotNil(t, resp)
	assert.Contains(t, resp.Error, stuckMountError,
		"the caller got %q instead of what the driver reported", resp.Error)

	select {
	case returnedAt := <-driver.returned:
		assert.True(t, returnedAt.Before(callerDeadline),
			"the driver finished at %v, after the caller stopped waiting at %v",
			returnedAt, callerDeadline)
	case <-time.After(time.Second):
		t.Fatal("driver never returned")
	}
}

// Without a deadline the connection timeout takes over, and the error still has
// to come back rather than being cut off with the connection.
func TestTimedOutMountReportsWithoutCallerDeadline(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	driver := registerStuckDriver(t)
	socketPath := newTestServer(t)

	resp, err := client.NewClient(socketPath).Mount(ctx, &proxy.MountRequest{
		Fstype: driver.fstype,
		Source: "stuck://bucket",
		Target: t.TempDir(),
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Contains(t, resp.Error, stuckMountError)
}
