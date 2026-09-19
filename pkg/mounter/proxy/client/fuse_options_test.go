package client

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/textlogger"
)

// captureKlog redirects klog output to a buffer via textlogger and restores it on cleanup.
// Returns a function that returns the captured output.
func captureKlog(t *testing.T) func() string {
	t.Helper()
	var buf bytes.Buffer
	config := textlogger.NewConfig(textlogger.Output(&buf), textlogger.Verbosity(10))
	logger := textlogger.NewLogger(config)
	klog.SetLoggerWithOptions(logger, klog.FlushLogger(func() {}))
	t.Cleanup(func() { klog.ClearLogger() })
	return func() string {
		return buf.String()
	}
}

func TestSplitFuseOptions_EmptyOptions(t *testing.T) {
	fuseOpts, daemonOpts := splitFuseOptions(nil, nil)
	// Should have all default options
	assert.NotEmpty(t, fuseOpts)
	assert.Empty(t, daemonOpts)
	// Verify defaults are present
	assertContainsOption(t, fuseOpts, "nodev")
	assertContainsOption(t, fuseOpts, "nosuid")
	assertContainsOption(t, fuseOpts, "allow_other")
	assertContainsOption(t, fuseOpts, "default_permissions")
}

func TestSplitFuseOptions_DaemonOnlyOptions(t *testing.T) {
	input := []string{"bucket=my-bucket", "endpoint=oss-cn-hangzhou.aliyuncs.com", "passwd_file=/etc/passwd-ossfs"}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	// All input should be daemon options
	assert.ElementsMatch(t, input, daemonOpts)
	// Fuse options should still have defaults
	assert.NotEmpty(t, fuseOpts)
}

func TestSplitFuseOptions_FuseOptionsOverride(t *testing.T) {
	input := []string{"rootmode=40755", "user_id=1000"}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	require.Empty(t, daemonOpts)
	// rootmode should be overridden
	assertContainsKV(t, fuseOpts, "rootmode", "40755")
	// user_id should be overridden
	assertContainsKV(t, fuseOpts, "user_id", "1000")
}

func TestSplitFuseOptions_RoOption(t *testing.T) {
	// New semantics: "ro" in options is NOT a FUSE option unless defaults/mountFlags
	// already have it; it becomes a daemon option.
	input := []string{"ro", "bucket=test"}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	for _, o := range fuseOpts {
		assert.NotEqual(t, "ro", o, "ro from options should not appear in fuseOpts when not in defaults/mountFlags")
	}
	assert.Contains(t, daemonOpts, "ro")
	assert.Contains(t, daemonOpts, "bucket=test")
}

func TestSplitFuseOptions_RwIgnored(t *testing.T) {
	input := []string{"rw", "bucket=test"}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	// rw should be ignored (not in fuse or daemon options)
	for _, o := range fuseOpts {
		assert.NotEqual(t, "rw", o)
	}
	assert.NotContains(t, daemonOpts, "rw")
	// bucket should remain in daemon
	assert.Contains(t, daemonOpts, "bucket=test")
}

func TestSplitFuseOptions_MixedOptions(t *testing.T) {
	// New semantics: keys not present in defaults/mountFlags become daemon options,
	// including "ro".
	input := []string{
		"bucket=test-bucket",
		"ro",
		"endpoint=oss-internal.aliyuncs.com",
		"passwd_file=/etc/creds",
	}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	for _, o := range fuseOpts {
		assert.NotEqual(t, "ro", o)
	}
	assert.Contains(t, daemonOpts, "ro")
	assert.Contains(t, daemonOpts, "bucket=test-bucket")
	assert.Contains(t, daemonOpts, "endpoint=oss-internal.aliyuncs.com")
	assert.Contains(t, daemonOpts, "passwd_file=/etc/creds")
}

func TestSplitFuseOptions_EmptyAndBlankOptions(t *testing.T) {
	input := []string{"", "bucket=test"}
	fuseOpts, daemonOpts := splitFuseOptions(input, nil)
	assert.NotEmpty(t, fuseOpts) // defaults
	assert.Contains(t, daemonOpts, "bucket=test")
}

func TestSplitFuseOptions_OtherOptionsAreDaemonOptions(t *testing.T) {
	// Keys that don't overlap with fuseOptionsMap (defaults ∪ mountFlags) become daemon options.
	input := []string{"noexec", "noatime", "debug", "dev", "suid", "auto_unmount"}
	_, daemonOpts := splitFuseOptions(input, nil)
	assert.Contains(t, daemonOpts, "noexec")
	assert.Contains(t, daemonOpts, "noatime")
	assert.Contains(t, daemonOpts, "debug")
	assert.Contains(t, daemonOpts, "dev")
	assert.Contains(t, daemonOpts, "suid")
	assert.Contains(t, daemonOpts, "auto_unmount")
}

func TestSplitFuseOptions_MountFlagsPriority(t *testing.T) {
	// Test: mountFlags can override defaults
	fuseOpts, daemonOpts := splitFuseOptions(nil, []string{"rootmode=40755"})
	require.Empty(t, daemonOpts)
	assertContainsKV(t, fuseOpts, "rootmode", "40755")
}

func TestSplitFuseOptions_OptionsOverrideMountFlags(t *testing.T) {
	// Test: options override mountFlags for the same key (with warning)
	flush := captureKlog(t)
	fuseOpts, daemonOpts := splitFuseOptions(
		[]string{"rootmode=41000"}, // options
		[]string{"rootmode=40755"}, // mountFlags
	)
	require.Empty(t, daemonOpts)
	assertContainsKV(t, fuseOpts, "rootmode", "41000")
	assert.Contains(t, flush(), "overrides existing fuseOption")
}

func TestSplitFuseOptions_AllMountFlagsToFuse(t *testing.T) {
	// New semantics: ALL mountFlags entries go into fuseOptions, regardless of whether
	// they are known FUSE keys. Non-FUSE-known entries (e.g. "bucket=test") still appear
	// in fuseOptions.
	fuseOpts, daemonOpts := splitFuseOptions(nil, []string{"ro", "bucket=test", "foo"})
	assertContainsOption(t, fuseOpts, "ro")
	assertContainsKV(t, fuseOpts, "bucket", "test")
	assertContainsOption(t, fuseOpts, "foo")
	assert.Empty(t, daemonOpts, "mountFlags must never go to daemon options")
}

func TestSplitFuseOptions_MountFlagsRwAndEmptyIgnored(t *testing.T) {
	// rw and empty entries are silently dropped from mountFlags as well.
	fuseOpts, _ := splitFuseOptions(nil, []string{"rw", "", "=value"})
	for _, o := range fuseOpts {
		assert.NotEqual(t, "rw", o)
		assert.NotEqual(t, "", o)
	}
}

func TestSplitFuseOptions_OptionsOverrideMountFlagsNonFuseKey(t *testing.T) {
	// New semantics: a non-FUSE key (e.g. "bucket") that came in via mountFlags is now
	// part of fuseOptions; if the same key appears in options, options wins + warning.
	flush := captureKlog(t)
	fuseOpts, daemonOpts := splitFuseOptions(
		[]string{"bucket=from-options"}, // options
		[]string{"bucket=from-flags"},   // mountFlags
	)
	assertContainsKV(t, fuseOpts, "bucket", "from-options")
	assert.Empty(t, daemonOpts)
	log := flush()
	assert.Contains(t, log, "bucket=from-options")
	assert.Contains(t, log, "overrides existing fuseOption")
}

func TestSplitFuseOptions_FullPriorityChain(t *testing.T) {
	// Full test:
	//   defaults: rootmode=40000
	//   mountFlags: rootmode=40755, ro, bucket=from-flags
	//   options: rootmode=41000, bucket=from-opts (overrides mountFlags), endpoint=...
	flush := captureKlog(t)
	fuseOpts, daemonOpts := splitFuseOptions(
		[]string{"rootmode=41000", "bucket=from-opts", "endpoint=oss-internal"},
		[]string{"rootmode=40755", "ro", "bucket=from-flags"},
	)
	// rootmode 41000 (options win)
	assertContainsKV(t, fuseOpts, "rootmode", "41000")
	// ro from mountFlags
	assertContainsOption(t, fuseOpts, "ro")
	// bucket: options override mountFlags-provided value
	assertContainsKV(t, fuseOpts, "bucket", "from-opts")
	// endpoint not in fuseOptions (no overlap) -> daemon
	assert.Contains(t, daemonOpts, "endpoint=oss-internal")
	log := flush()
	assert.Contains(t, log, "rootmode=41000")
	assert.Contains(t, log, "bucket=from-opts")
}

func TestParseOptionKey(t *testing.T) {
	tests := []struct {
		name    string
		opt     string
		wantKey string
		wantOk  bool
	}{
		{name: "empty string", opt: "", wantKey: "", wantOk: false},
		{name: "empty key with value", opt: "=value", wantKey: "", wantOk: false},
		{name: "rw silently ignored", opt: "rw", wantKey: "", wantOk: false},
		{name: "flag-style key", opt: "ro", wantKey: "ro", wantOk: true},
		{name: "kv-style key", opt: "rootmode=40755", wantKey: "rootmode", wantOk: true},
		{name: "non-FUSE key still returns key (no membership check)", opt: "bucket=test", wantKey: "bucket", wantOk: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, ok := parseOptionKey(tc.opt)
			assert.Equal(t, tc.wantKey, key)
			assert.Equal(t, tc.wantOk, ok)
		})
	}
}

func TestSetFuseOption(t *testing.T) {
	tests := []struct {
		name string
		opt  string
		// expectKey is the key to check in the resulting map ("" means do not check)
		expectKey   string
		expectValue string
		// expectAbsentKey, when non-empty, asserts the key is NOT present in the map
		expectAbsentKey string
	}{
		{name: "empty string is ignored", opt: "", expectAbsentKey: ""},
		{name: "empty key with value is ignored", opt: "=value", expectAbsentKey: ""},
		{name: "rw is silently ignored", opt: "rw", expectAbsentKey: "rw"},
		{name: "flag-style entry", opt: "ro", expectKey: "ro", expectValue: nullVal},
		{name: "kv-style entry", opt: "rootmode=40755", expectKey: "rootmode", expectValue: "40755"},
		{name: "non-FUSE-known key still inserted (kv)", opt: "bucket=test", expectKey: "bucket", expectValue: "test"},
		{name: "non-FUSE-known key still inserted (flag)", opt: "foo", expectKey: "foo", expectValue: nullVal},
		{name: "kv on flag default still uses provided value", opt: "nodev=1", expectKey: "nodev", expectValue: "1"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := make(map[string]string)
			setFuseOption(m, tc.opt)
			if tc.expectKey != "" {
				v, ok := m[tc.expectKey]
				require.True(t, ok, "expected key %q to be present in map", tc.expectKey)
				assert.Equal(t, tc.expectValue, v)
			}
			if tc.expectAbsentKey != "" {
				_, ok := m[tc.expectAbsentKey]
				assert.False(t, ok, "expected key %q to be absent from map", tc.expectAbsentKey)
			}
		})
	}
}

// TestSplitFuseOptions_OptionsOverrideWarn verifies that when options overrides a key
// already present in fuseOptionsMap (defaults ∪ mountFlags), a warning is logged.
func TestSplitFuseOptions_OptionsOverrideWarn(t *testing.T) {
	t.Run("options override default", func(t *testing.T) {
		flush := captureKlog(t)
		fuseOpts, _ := splitFuseOptions([]string{"rootmode=41000"}, nil)
		assertContainsKV(t, fuseOpts, "rootmode", "41000")
		log := flush()
		assert.Contains(t, log, "rootmode=41000")
		assert.Contains(t, log, "overrides existing fuseOption")
	})

	t.Run("flag-style default key shows <flag> placeholder", func(t *testing.T) {
		// allow_other default is nullVal (a flag); when options override it, the
		// warning should fire using the <flag> placeholder for the prior value.
		flush := captureKlog(t)
		fuseOpts, _ := splitFuseOptions([]string{"allow_other"}, nil)
		assertContainsOption(t, fuseOpts, "allow_other")
		log := flush()
		assert.Contains(t, log, "allow_other")
		assert.Contains(t, log, "<flag>")
	})

	t.Run("non-overlapping key in options does not warn", func(t *testing.T) {
		flush := captureKlog(t)
		_, daemonOpts := splitFuseOptions([]string{"bucket=test"}, nil)
		assert.Contains(t, daemonOpts, "bucket=test")
		log := flush()
		assert.NotContains(t, log, "overrides existing fuseOption")
	})

	t.Run("rw in options does not warn and does not appear", func(t *testing.T) {
		flush := captureKlog(t)
		fuseOpts, daemonOpts := splitFuseOptions([]string{"rw"}, nil)
		for _, o := range fuseOpts {
			assert.NotEqual(t, "rw", o)
		}
		assert.NotContains(t, daemonOpts, "rw")
		log := flush()
		assert.NotContains(t, log, "overrides existing fuseOption")
	})
}

// assertContainsOption checks that opts contains the exact option string.
func assertContainsOption(t *testing.T, opts []string, option string) {
	t.Helper()
	for _, o := range opts {
		if o == option {
			return
		}
	}
	t.Errorf("expected options %v to contain %q", opts, option)
}

// assertContainsKV checks that opts contains "key=value".
func assertContainsKV(t *testing.T, opts []string, key, value string) {
	t.Helper()
	expected := key + "=" + value
	for _, o := range opts {
		if o == expected {
			return
		}
	}
	t.Errorf("expected options %v to contain %q", opts, expected)
}
