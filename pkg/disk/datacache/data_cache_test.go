package datacache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/synctest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2/ktesting"
)

func TestGetOpts(t *testing.T) {
	tests := []struct {
		name     string
		opts     map[string]string
		wantMode Mode
		wantSize string // empty => zero
		wantErr  string // substring; empty => no error
	}{
		{
			name: "not configured",
			opts: map[string]string{},
		},
		{
			name:     "size only defaults to writethrough",
			opts:     map[string]string{sizeKey: "10Gi"},
			wantMode: Writethrough,
			wantSize: "10Gi",
		},
		{
			name:     "explicit writeback",
			opts:     map[string]string{sizeKey: "10Gi", modeKey: "writeback"},
			wantMode: Writeback,
			wantSize: "10Gi",
		},
		{
			name:     "explicit writethrough",
			opts:     map[string]string{sizeKey: "5Gi", modeKey: "writethrough"},
			wantMode: Writethrough,
			wantSize: "5Gi",
		},
		{
			name:    "mode without size is an error",
			opts:    map[string]string{modeKey: "writeback"},
			wantErr: "must specify non-zero",
		},
		{
			name:    "zero size with mode is an error",
			opts:    map[string]string{modeKey: "writeback", sizeKey: "0"},
			wantErr: "must specify non-zero",
		},
		{
			name:    "invalid size",
			opts:    map[string]string{sizeKey: "not-a-size"},
			wantErr: "invalid " + sizeKey,
		},
		{
			name:    "unrecognized mode",
			opts:    map[string]string{sizeKey: "10Gi", modeKey: "bogus"},
			wantErr: "unrecognized " + modeKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Opts
			err := GetOpts(tt.opts, &d)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantMode, d.Mode)
			wantSize := resource.Quantity{}
			if tt.wantSize != "" {
				wantSize = resource.MustParse(tt.wantSize)
			}
			assert.Truef(t, d.Size.Cmp(wantSize) == 0, "size = %v, want %v", d.Size, wantSize)
		})
	}
}

func TestParseCacheStatus(t *testing.T) {
	// Real status lines captured from kernel 6.6 (dmsetup status, minus the
	// "0 <len> cache" device-mapper prefix that queryTable strips).
	const wbMq = "8 9/4096 512 0/256 0 48 0 0 0 0 0 2 metadata2 writeback 2 migration_threshold 4096 mq 10 random_threshold 0 sequential_threshold 0 discard_promote_adjustment 0 read_promote_adjustment 0 write_promote_adjustment 0 rw -"
	// no_discard_passdown adds a third feature, shifting the policy-name index.
	// Real capture, with the dirty field (index 10) edited to 9 to also exercise
	// non-zero dirty parsing.
	const wbMqNdp = "8 9/4096 512 9/256 14 34 0 0 0 9 9 3 metadata2 writeback no_discard_passdown 2 migration_threshold 4096 mq 10 random_threshold 0 sequential_threshold 0 discard_promote_adjustment 0 read_promote_adjustment 0 write_promote_adjustment 0 rw -"
	const cleaner = "8 12/4096 512 0/256 0 96 0 0 0 0 0 2 metadata2 writethrough 2 migration_threshold 2048 cleaner 0 rw -"

	tests := []struct {
		name          string
		status        string
		wantDirty     uint64
		wantWriteback bool
		wantPolicy    string
		wantErr       bool
	}{
		{name: "writeback mq clean", status: wbMq, wantDirty: 0, wantWriteback: true, wantPolicy: "mq"},
		{name: "writeback mq dirty + 3 features", status: wbMqNdp, wantDirty: 9, wantWriteback: true, wantPolicy: "mq"},
		{name: "cleaner writethrough", status: cleaner, wantDirty: 0, wantWriteback: false, wantPolicy: cleanerPolicy},
		{name: "empty", status: "", wantErr: true},
		{name: "too short", status: "8 9/4096 512 0/256 0 48 0 0", wantErr: true},
		{name: "non-numeric dirty", status: "8 9/4096 512 0/256 0 48 0 0 0 0 X 2 metadata2 writeback 2 migration_threshold 4096 mq 0", wantErr: true},
		{name: "bad feature count", status: "8 9/4096 512 0/256 0 48 0 0 0 0 0 Z metadata2 writeback", wantErr: true},
		{name: "feature count overflows fields", status: "8 9/4096 512 0/256 0 48 0 0 0 0 0 9 metadata2 writeback", wantErr: true},
		// Features are the last fields, no <#core args> after: 12+nFeat == len(f).
		{name: "no field after features", status: "8 9/4096 512 0/256 0 48 0 0 0 0 0 2 metadata2 writeback", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := parseCacheStatus(tt.status)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, cacheStatus{dirty: tt.wantDirty, writeback: tt.wantWriteback, policy: tt.wantPolicy}, s)
		})
	}
}

func TestSplitCacheTable(t *testing.T) {
	devices, tail, err := splitCacheTable("7:0 7:1 7:2 512 2 metadata2 writeback mq 2 migration_threshold 4096")
	require.NoError(t, err)
	assert.Equal(t, "7:0 7:1 7:2", devices)
	assert.Equal(t, "512 2 metadata2 writeback mq 2 migration_threshold 4096", tail)

	_, _, err = splitCacheTable("7:0 7:1")
	assert.Error(t, err, "expected error for too-few fields")
}

func TestMetaSize(t *testing.T) {
	const gib = 1 << 30
	const mib = 1 << 20
	const kib = 1 << 10
	cases := []struct {
		name     string
		dataSize int64
		want     int64
	}{
		// The 4 MiB transaction overhead dominates a tiny cache; the few
		// per-block bytes then round up to the next 4 KiB.
		{"tiny", 1 * mib, 4*mib + 4096},         // 4 MiB + 4 blocks * 44 -> next 4 KiB
		{"1GiB", 1 * gib, 4*mib + 176*kib},      // 4 MiB + 4096 blocks * 44, already 4 KiB-aligned
		{"100GiB", 100 * gib, 21*mib + 192*kib}, // 4 MiB + 409600 blocks * 44
		// Well past the point where the formula exceeds the kernel cap.
		{"huge", 1 << 50, 17045913600}, // ~16 GiB
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, metaSize(c.dataSize))
		})
	}
}

// fakeDm is an in-memory dmDevice for deterministic tests of the flush/reconcile
// control flow, with no kernel, root, or loop devices involved.
//
// It models the dm-cache states the logic reacts to (policy, io mode, dirty
// count) and produces real status/table strings, so the tests exercise the
// actual parseCacheStatus / splitCacheTable / arg-builder round-trip. tableLoad
// updates the modelled state from the loaded args exactly as the kernel would
// (cleaner forces writethrough); the drain is advanced by an explicit onStatus
// hook rather than a timer, so every transition is under the test's control.
type fakeDm struct {
	dirty     uint64
	writeback bool
	policy    string
	size      uint64
	absent    bool // status returns ENXIO

	loads   []string
	created int
	removed int

	// onStatus, if set, runs before each status read, letting a test advance the
	// modelled drain or inject transitions deterministically.
	onStatus func(f *fakeDm)
}

// tail reconstructs the current table tail from the modelled state.
func (f *fakeDm) tail() string {
	if f.policy == cleanerPolicy {
		return cleanerArgsTail()
	}
	mode := Writethrough
	if f.writeback {
		mode = Writeback
	}
	return cacheArgsTail(mode)
}

func (f *fakeDm) tableStatus(flags uint32) (uint64, string, error) {
	if f.onStatus != nil {
		f.onStatus(f)
	}
	if f.absent {
		return 0, "", fmt.Errorf("get status: %w", unix.ENXIO)
	}
	if flags&dmStatusTableFlag != 0 {
		return f.size, "7:0 7:1 7:2 " + f.tail(), nil
	}
	mode := "writethrough"
	if f.writeback {
		mode = "writeback"
	}
	// Real INFO status layout: dirty is field 10, then <#feat> features, then
	// <#core> core args, then the policy name.
	info := fmt.Sprintf("8 9/4096 512 0/256 0 0 0 0 0 0 %d 2 metadata2 %s 2 migration_threshold 4096 %s 0 rw -",
		f.dirty, mode, f.policy)
	return f.size, info, nil
}

func (f *fakeDm) tableLoad(size uint64, args string) error {
	f.loads = append(f.loads, args)
	f.size = size
	_, tail, err := splitCacheTable(args)
	if err != nil {
		return err
	}
	switch {
	case strings.Contains(tail, " "+cleanerPolicy+" "):
		f.policy = cleanerPolicy
		f.writeback = false // cleaner forces writethrough
	default:
		f.policy = "mq"
		f.writeback = strings.Contains(tail, " "+string(Writeback)+" ")
	}
	return nil
}

func (f *fakeDm) create() error { f.created++; return nil }
func (f *fakeDm) remove() error { f.removed++; return nil }

func (f *fakeDm) switchedToCleaner() int {
	n := 0
	for _, a := range f.loads {
		if strings.Contains(a, " "+cleanerPolicy+" ") {
			n++
		}
	}
	return n
}

// flushToClean waits flushPollInterval between status reads; the tests run it in
// a synctest bubble so that real interval elapses instantly on the fake clock,
// exercising the production constant rather than overriding it.

func TestFlushToClean_CleanWriteback_SwitchesToCleaner(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		// Clean writeback cache: no dirty blocks, but must still switch to cleaner
		// so a late write during teardown cannot land dirty in the cache.
		f := &fakeDm{dirty: 0, writeback: true, policy: "mq"}

		require.NoError(t, flushToClean(ctx, logger, f))
		assert.Equal(t, 1, f.switchedToCleaner(), "expected exactly one switch to cleaner")
		assert.Equal(t, cleanerPolicy, f.policy)
		assert.False(t, f.writeback, "cache must be left writethrough")
	})
}

func TestFlushToClean_Writethrough_NoSwitch(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		// A configured writethrough cache is already clean; nothing to do.
		f := &fakeDm{dirty: 0, writeback: false, policy: "mq"}

		require.NoError(t, flushToClean(ctx, logger, f))
		assert.Equal(t, 0, f.switchedToCleaner(), "writethrough cache should not switch to cleaner")
	})
}

func TestFlushToClean_DirtyWriteback_DrainsThenReturns(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		// Dirty writeback cache. After switching to cleaner, model the drain: each
		// subsequent status read sees one fewer dirty block until zero.
		f := &fakeDm{dirty: 3, writeback: true, policy: "mq"}
		f.onStatus = func(f *fakeDm) {
			if f.policy == cleanerPolicy && f.dirty > 0 {
				f.dirty--
			}
		}

		require.NoError(t, flushToClean(ctx, logger, f))
		assert.Equal(t, 1, f.switchedToCleaner(), "expected one switch to cleaner")
		assert.Zero(t, f.dirty, "expected drain to complete")
	})
}

// TestFlushToClean_RetryAfterAbort_KeepsDraining models a teardown that switched
// to cleaner and then timed out with dirty blocks remaining, and asserts that a
// retry does NOT mistake the cleaner's writethrough status for "already clean"
// and return early: it must keep draining. This is the data-loss guard.
func TestFlushToClean_RetryAfterAbort_KeepsDraining(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		// State left by an aborted flush: cleaner policy (=> writethrough) but still
		// dirty. A buggy `!writeback => done` check would return immediately here.
		f := &fakeDm{dirty: 2, writeback: false, policy: cleanerPolicy}
		f.onStatus = func(f *fakeDm) {
			if f.dirty > 0 {
				f.dirty--
			}
		}

		require.NoError(t, flushToClean(ctx, logger, f))
		assert.Equal(t, 0, f.switchedToCleaner(), "retry should not switch again (already cleaner)")
		assert.Zero(t, f.dirty, "retry returned without draining")
	})
}

func TestFlushToClean_ContextCancelledMidDrain(t *testing.T) {
	logger, _ := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())

		// Never drains; cancel after the switch so the poll observes ctx.Done().
		f := &fakeDm{dirty: 5, writeback: true, policy: "mq"}
		f.onStatus = func(f *fakeDm) {
			if f.policy == cleanerPolicy {
				cancel()
			}
		}

		err := flushToClean(ctx, logger, f)
		assert.ErrorIs(t, err, context.Canceled)
		assert.NotZero(t, f.dirty, "test bug: cache should still be dirty at cancellation")
	})
}

// flushToClean propagates ENXIO rather than swallowing it; teardownDataCache
// decides that an absent device is not fatal (and still cleans up the files).
func TestFlushToClean_DeviceAbsentPropagatesENXIO(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	f := &fakeDm{absent: true}

	assert.ErrorIs(t, flushToClean(ctx, logger, f), unix.ENXIO)
	assert.Equal(t, 0, f.switchedToCleaner(), "absent device should not switch policy")
}

func TestReconcileTable(t *testing.T) {
	logger, _ := ktesting.NewTestContext(t)

	t.Run("cleaner is reconciled to normal", func(t *testing.T) {
		f := &fakeDm{policy: cleanerPolicy, writeback: false}
		require.NoError(t, reconcileTable(logger, f, Writeback))
		assert.Len(t, f.loads, 1, "expected one reload")
		assert.Equal(t, "mq", f.policy)
		assert.True(t, f.writeback, "not reconciled to writeback")
	})

	t.Run("matching table is left alone", func(t *testing.T) {
		f := &fakeDm{policy: "mq", writeback: true}
		require.NoError(t, reconcileTable(logger, f, Writeback))
		assert.Empty(t, f.loads, "should not reload a matching table")
	})

	t.Run("wrong mode is reconciled", func(t *testing.T) {
		f := &fakeDm{policy: "mq", writeback: false} // writethrough on disk
		require.NoError(t, reconcileTable(logger, f, Writeback))
		assert.Len(t, f.loads, 1)
		assert.True(t, f.writeback, "wrong mode not reconciled")
	})
}

func TestWaitForCacheDevice(t *testing.T) {
	t.Run("existing device returns immediately", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "dev")
		require.NoError(t, os.WriteFile(path, nil, 0600))
		assert.NoError(t, waitForCacheDevice(t.Context(), path))
	})

	t.Run("absent device times out", func(t *testing.T) {
		// synctest's fake clock lets the real deviceAppearTimeout elapse
		// instantly instead of mutating the production constants.
		dir := t.TempDir()
		synctest.Test(t, func(t *testing.T) {
			err := waitForCacheDevice(t.Context(), filepath.Join(dir, "nonexistent"))
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		})
	})
}
