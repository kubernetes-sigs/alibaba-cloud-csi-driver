package server

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// deadlineSpyDriver reports the deadline of the context its mount was given.
// Nothing else exposes what a request's header turned into by the time a driver
// runs, and that translation is the whole point of the code under test.
type deadlineSpyDriver struct {
	fstype string
	got    chan time.Time
}

func (d *deadlineSpyDriver) Name() string      { return d.fstype }
func (d *deadlineSpyDriver) Fstypes() []string { return []string{d.fstype} }
func (d *deadlineSpyDriver) Init()             {}
func (d *deadlineSpyDriver) Terminate()        {}
func (d *deadlineSpyDriver) ApplyOptionDefaults(options []string) []string {
	return options
}

func (d *deadlineSpyDriver) Mount(ctx context.Context, _ *proxy.MountRequest, _ int) error {
	deadline, _ := ctx.Deadline()
	d.got <- deadline
	return nil
}

func registerDeadlineSpy(t *testing.T) *deadlineSpyDriver {
	t.Helper()
	d := &deadlineSpyDriver{fstype: "deadline-spy", got: make(chan time.Time, 1)}
	fstypeToDriver[d.fstype] = d
	t.Cleanup(func() { delete(fstypeToDriver, d.fstype) })
	return d
}

func mountWithHeader(t *testing.T, socketPath string, header proxy.Header, fstype string) {
	t.Helper()
	conn := dialTestServer(t, socketPath)
	header.Method = proxy.Mount
	data, err := json.Marshal(proxy.Request{
		Header: header,
		Body:   proxy.MountRequest{Fstype: fstype, Source: "spy://b", Target: "/tmp/spy"},
	})
	require.NoError(t, err)
	_, err = conn.Write(append(data, proxy.MessageEnd))
	require.NoError(t, err)
}

func awaitDeadline(t *testing.T, d *deadlineSpyDriver) time.Time {
	t.Helper()
	select {
	case got := <-d.got:
		return got
	case <-time.After(5 * time.Second):
		t.Fatal("driver was never reached")
		return time.Time{}
	}
}

func TestStatedInstantReachesDriver(t *testing.T) {
	spy := registerDeadlineSpy(t)
	socketPath := newTestServer(t)

	stated := time.Now().Add(500 * time.Millisecond)
	mountWithHeader(t, socketPath, proxy.Header{DeadlineUnixNano: stated.UnixNano()}, spy.fstype)

	got := awaitDeadline(t, spy)
	assert.True(t, got.Before(stated), "the driver must stop before the client does: %v vs %v", got, stated)
}

// A client that predates the field leaves the connection timeout in charge,
// which is how this behaved before the header existed.
func TestUnstatedFallsBackToConnectionTimeout(t *testing.T) {
	spy := registerDeadlineSpy(t)
	socketPath := newTestServer(t)

	before := time.Now()
	mountWithHeader(t, socketPath, proxy.Header{}, spy.fstype)

	got := awaitDeadline(t, spy)
	// newTestServer runs with a 1s connection timeout.
	assert.WithinDuration(t, before.Add(time.Second), got, 500*time.Millisecond)
}
