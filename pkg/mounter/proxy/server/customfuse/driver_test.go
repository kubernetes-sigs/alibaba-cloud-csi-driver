package customfuse

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/mount-utils"
)

func TestBuildEnvVars(t *testing.T) {
	tests := []struct {
		name string
		req  *proxy.MountRequest
		want map[string]string
	}{
		{
			name: "full request with bucket",
			req: &proxy.MountRequest{
				Source:  "redis://host:6379/1",
				Target:  "/mnt/data",
				Options: []string{"bucket=my-jfs-data", "url=oss-cn-hangzhou-internal.aliyuncs.com", "otherOpts=--cache-size=1024 --buffer-size=300"},
				Secrets: map[string]string{"accessKeyId": "ak123", "accessKeySecret": "sk456"},
			},
			want: map[string]string{
				"source":          "redis://host:6379/1",
				"mountpoint":      "/mnt/data",
				"bucket":          "my-jfs-data",
				"url":             "oss-cn-hangzhou-internal.aliyuncs.com",
				"otherOpts":       "--cache-size=1024 --buffer-size=300",
				"accessKeyId":     "ak123",
				"accessKeySecret": "sk456",
			},
		},
		{
			name: "url without bucket",
			req: &proxy.MountRequest{
				Source:  "mybucket:/data",
				Target:  "/mnt/oss",
				Options: []string{"url=oss-cn-hangzhou.aliyuncs.com"},
			},
			want: map[string]string{
				"source":     "mybucket:/data",
				"mountpoint": "/mnt/oss",
				"url":        "oss-cn-hangzhou.aliyuncs.com",
			},
		},
		{
			name: "no options no secrets",
			req: &proxy.MountRequest{
				Source: "vol",
				Target: "/mnt/vol",
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
			},
		},
		{
			name: "otherOpts only",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"otherOpts=--buffer-size=300"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"otherOpts":  "--buffer-size=300",
			},
		},
		{
			name: "empty source omitted from env",
			req: &proxy.MountRequest{
				Source: "",
				Target: "/mnt/vol",
			},
			want: map[string]string{
				"mountpoint": "/mnt/vol",
			},
		},
		{
			name: "a source option does not reach the environment",
			req: &proxy.MountRequest{
				Source:  "",
				Target:  "/mnt/vol",
				Options: []string{"source=redis://host:6379/1", "formatOptions=storage=oss"},
			},
			want: map[string]string{
				"mountpoint":    "/mnt/vol",
				"formatOptions": "storage=oss",
			},
		},
		{
			name: "mountOptions passthrough (key=value)",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"cache-size=1024", "debug"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"cache-size": "1024",
				"debug":      "",
			},
		},
		{
			name: "mountOptions mixed with structured options",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"bucket=mybucket", "url=ep.com", "cache-size=1024", "debug"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"bucket":     "mybucket",
				"url":        "ep.com",
				"cache-size": "1024",
				"debug":      "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildEnvVars(tt.req.Source, tt.req.Target, tt.req.Options, tt.req.Secrets, nil)
			gotMap := envSliceToMap(got)
			assert.Equal(t, tt.want, gotMap)
		})
	}
}

func TestBuildEnvVars_SecretsAsDirectEnvVars(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Secrets: map[string]string{"MY_TOKEN": "secret123", "DB_PASS": "p@ss"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)
	assert.Equal(t, "secret123", gotMap["MY_TOKEN"])
	assert.Equal(t, "p@ss", gotMap["DB_PASS"])
}

func TestBuildEnvVars_OtherOptsPassthrough(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Options: []string{"otherOpts=--flag1 --flag2=value"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)
	assert.Equal(t, "--flag1 --flag2=value", gotMap["otherOpts"])
}

func TestBuildEnvVars_JSONRoundTrip(t *testing.T) {
	original := &proxy.MountRequest{
		Source:  "mybucket:/data",
		Target:  "/mnt/data",
		Options: []string{"url=oss-cn-hangzhou.aliyuncs.com", "otherOpts=-o a=b --c=d --e f"},
		Secrets: map[string]string{"AK": "id", "SK": "secret=with=equals"},
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded proxy.MountRequest
	require.NoError(t, json.Unmarshal(data, &decoded))

	got := buildEnvVars(decoded.Source, decoded.Target, decoded.Options, decoded.Secrets, nil)
	gotMap := envSliceToMap(got)

	assert.Equal(t, "mybucket:/data", gotMap["source"])
	assert.Equal(t, "/mnt/data", gotMap["mountpoint"])
	assert.Equal(t, "oss-cn-hangzhou.aliyuncs.com", gotMap["url"])
	assert.Equal(t, "-o a=b --c=d --e f", gotMap["otherOpts"])
	assert.Equal(t, "id", gotMap["AK"])
	assert.Equal(t, "secret=with=equals", gotMap["SK"])
}

// Which credential spellings a client accepts is the client's business, so a
// Secret key arrives as that variable name and nothing else: no prefix, no second
// spelling derived from it. An adapter that takes more than one resolves them
// itself, where the alternatives are known.
func TestBuildEnvVars_SecretKeysArriveUnaliased(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Secrets: map[string]string{"legacyKeyId": "myak", "legacyKeySecret": "mysk"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)

	assert.Equal(t, "myak", gotMap["legacyKeyId"])
	assert.Equal(t, "mysk", gotMap["legacyKeySecret"])
	assert.Len(t, got, 4, "source, mountpoint and the two Secret keys, with nothing derived from them")
}

// The driver sets mountpoint and source itself, and then waits for a mount point
// to appear at mountpoint. exec.Cmd keeps the last value for a duplicated name,
// so a volume parameter listed after them would win and the client would mount
// somewhere nothing is watching.
func TestBuildEnvVars_DriverOwnedNamesAreNotOverridable(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt/vol",
		Options: []string{"mountpoint=/tmp/elsewhere", "source=other", "bucket=b"},
		Secrets: map[string]string{"mountpoint": "/tmp/from-secret", "accessKeyId": "ak"},
	}
	gotMap := envSliceToMap(buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil))
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "vol", gotMap["source"])
	assert.Equal(t, "b", gotMap["bucket"], "a sibling option is unaffected")
	assert.Equal(t, "ak", gotMap["accessKeyId"])
	assert.Len(t, gotMap, 4)
}

// The client runs inside the mount-proxy's own process environment, which carries
// the settings that process runs on. A volume parameter must not be able to
// redefine them, or whoever can edit the volume decides how the proxy behaves.
func TestBuildEnvVars_InheritedNamesAreNotShadowed(t *testing.T) {
	inherited := []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"HOME=/root",
		"AGENT_IDENTITY_ENDPOINT=https://identity.invalid",
		"AGENT_IDENTITY_TOKEN_DIR=/var/run/secrets/tokens",
	}
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt/vol",
		Options: []string{"PATH=/tmp/attacker/bin", "cache-size=1024"},
		Secrets: map[string]string{"AGENT_IDENTITY_TOKEN_DIR": "/tmp/elsewhere", "accessKeyId": "ak"},
	}
	gotMap := envSliceToMap(buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, inherited))
	assert.NotContains(t, gotMap, "PATH")
	assert.NotContains(t, gotMap, "HOME")
	assert.NotContains(t, gotMap, "AGENT_IDENTITY_ENDPOINT")
	assert.NotContains(t, gotMap, "AGENT_IDENTITY_TOKEN_DIR")
	assert.Equal(t, "1024", gotMap["cache-size"], "a name nothing else claims still passes through")
	assert.Equal(t, "ak", gotMap["accessKeyId"])
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "vol", gotMap["source"])
}

// readOnly is the one driver-set name that arrives through the options leg rather
// than from the request itself, so it has to be claimed before that leg is read.
// The direction that matters is a Secret clearing a read-only request: with
// last-wins environment semantics that publishes a read-only volume read-write.
func TestBuildEnvVars_ReadOnlyIsNotSettableFromVolumeData(t *testing.T) {
	t.Run("a secret cannot clear a read-only request", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"readOnly=true", "bucket=b"},
			map[string]string{"readOnly": "false"}, nil))
		assert.Equal(t, "true", gotMap["readOnly"])
		assert.Equal(t, "b", gotMap["bucket"], "a sibling option is unaffected")
	})

	t.Run("a secret cannot make a read-write volume read-only", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"bucket=b"},
			map[string]string{"readOnly": "true"}, nil))
		assert.NotContains(t, gotMap, "readOnly")
		assert.Equal(t, "b", gotMap["bucket"])
	})

	t.Run("an option spelling the driver does not emit is refused", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"readOnly=false", "bucket=b"}, nil, nil))
		assert.NotContains(t, gotMap, "readOnly")
		assert.Equal(t, "b", gotMap["bucket"])
	})

	t.Run("position does not decide it", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"bucket=b", "readOnly=false", "readOnly=true"}, nil, nil))
		assert.Equal(t, "true", gotMap["readOnly"])
		assert.Equal(t, "b", gotMap["bucket"])
	})
}

// A name is claimed by whichever leg reaches it first, so the Secret leg goes
// first: a volume parameter must not be able to pass its own value off as a
// credential the volume's Secret carries.
func TestBuildEnvVars_SecretBeatsOptionOfSameName(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"accessKeyId=from-option", "bucket=b"},
		map[string]string{"accessKeyId": "from-secret"}, nil))
	assert.Equal(t, "from-secret", gotMap["accessKeyId"])
	assert.Equal(t, "b", gotMap["bucket"], "a name nothing else claims still passes through")
}

// The plugin refuses these case-insensitively, but it passes Secrets through
// unfiltered, so this is the only half that can catch a differently cased key.
func TestBuildEnvVars_DriverOwnedNamesAreRefusedInAnyCasing(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"SOURCE=from-option", "bucket=b"},
		map[string]string{"MountPoint": "/tmp/from-secret", "READONLY": "false"}, nil))
	assert.NotContains(t, gotMap, "SOURCE")
	assert.NotContains(t, gotMap, "MountPoint")
	assert.NotContains(t, gotMap, "READONLY")
	assert.Equal(t, "vol", gotMap["source"], "the driver's own spelling is what survives")
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "b", gotMap["bucket"])
	assert.Len(t, gotMap, 3)
}

// PATH in the mount-proxy's environment and path= in a volume are different
// variables to the entrypoint, and path is a documented driver field. The
// inherited-name set is exact-match for that reason; making it case-insensitive
// would drop every volume's sub-path.
func TestBuildEnvVars_InheritedNamesMatchExactly(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"path=/sub", "PATH=/tmp/attacker/bin"},
		nil, []string{"PATH=/usr/bin:/bin"}))
	assert.Equal(t, "/sub", gotMap["path"])
	assert.NotContains(t, gotMap, "PATH")
}

// Options and secrets travel through the same list, so the report of what the
// entrypoint will see has to stop at the names.
func TestEnvNamesNeverCarriesValues(t *testing.T) {
	assert.Equal(t, "mountpoint,accessKeySecret",
		envNames([]string{"mountpoint=/mnt/vol", "accessKeySecret=hunter2"}))
}

func envSliceToMap(envs []string) map[string]string {
	sort.Strings(envs)
	m := make(map[string]string, len(envs))
	for _, e := range envs {
		for i := 0; i < len(e); i++ {
			if e[i] == '=' {
				m[e[:i]] = e[i+1:]
				break
			}
		}
	}
	return m
}

type fakeMountChecker struct {
	mount.Interface
	notMnt bool
}

func (f *fakeMountChecker) IsLikelyNotMountPoint(string) (bool, error) {
	return f.notMnt, nil
}

// newTestMounter builds an extendedMounter whose mountpoint check is scripted and
// whose entrypoint is replaced for the duration of the test, so no test needs a
// script at either of the two fixed paths this container looks at.
//
// Cleanup signals whole process groups the same way the driver does. A stand-in
// entrypoint that forks leaves something behind, and killing only the pid the
// driver recorded is exactly the mistake under test here.
func newTestMounter(t *testing.T, notMnt bool, entrypoint func(path string) *exec.Cmd) *extendedMounter {
	t.Helper()
	original := newEntrypointCmd
	newEntrypointCmd = entrypoint
	m := &extendedMounter{
		driver: &Driver{
			monitorManager: server.NewMountMonitorManager(),
		},
		Interface: &fakeMountChecker{notMnt: notMnt},
	}
	t.Cleanup(func() {
		newEntrypointCmd = original
		m.driver.pids.Range(func(key, _ any) bool {
			_ = signalGroup(key.(int), syscall.SIGKILL)
			return true
		})
		// Also an assertion, not just hygiene: killing the group is what releases the
		// stderr pipe cmd.Wait is blocked on, so a wait group that does not drain here
		// means the driver is leaking a goroutine per mount.
		if !waitGroupTimeout(&m.driver.wg, 10*time.Second) {
			t.Error("a goroutine was still blocked in cmd.Wait after its process group was killed")
		}
	})
	return m
}

// processAlive reports whether pid is still running. A zombie counts as dead: it is
// no longer executing, and the kernel has already released its descriptors, which is
// the only thing the driver waits on. Nothing here is obliged to reap a grandchild,
// so insisting the pid vanish entirely would be waiting on something the code under
// test does not control.
func processAlive(pid int) bool {
	data, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		return false
	}
	s := string(data)
	// The state byte follows the command name, which is parenthesised and may itself
	// contain spaces and parentheses, so only the last ')' is a delimiter.
	i := strings.LastIndex(s, ")")
	if i < 0 || i+2 >= len(s) {
		return false
	}
	return s[i+2] != 'Z'
}

// waitForPid blocks until the stand-in entrypoint has recorded the pid of the
// subprocess it forked. That file is how a test learns about a process the driver
// itself never sees, and which is the reason the cleanup has to reach past the pid
// it started.
func waitForPid(t *testing.T, path string) int {
	t.Helper()
	var content string
	require.Eventually(t, func() bool {
		b, err := os.ReadFile(path)
		if err != nil {
			return false
		}
		if strings.TrimSpace(string(b)) == "" {
			return false
		}
		content = string(b)
		return true
	}, 10*time.Second, 20*time.Millisecond, "the entrypoint never recorded its subprocess")
	pid, err := strconv.Atoi(strings.TrimSpace(content))
	require.NoError(t, err)
	return pid
}

func TestExtendedMount(t *testing.T) {
	t.Run("the entrypoint's stderr reaches the returned error", func(t *testing.T) {
		m := newTestMounter(t, true, func(string) *exec.Cmd {
			return exec.Command("sh", "-c", `echo "myfuse: mount failed: connection refused" >&2; exit 1`)
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "myfuse: mount failed: connection refused")
		assert.Contains(t, err.Error(), "exit status 1")
	})

	// The entrypoint is a script the customer wrote, so how much it prints before
	// failing is not knowable from here. What is knowable is that the error carrying
	// it has to stay sendable: the proxy caps a message at proxy.MaxMsgSize and
	// refuses the request outright past that, so an unbounded capture can destroy the
	// very diagnosis it exists to deliver.
	t.Run("an oversized stderr keeps its end and says it was cut", func(t *testing.T) {
		// ExtendedMount mirrors the entrypoint's stderr to os.Stderr, which under a
		// test is the test log. Sending it away keeps this about what was retained
		// rather than about pages of filler in the output.
		devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
		require.NoError(t, err)
		original := os.Stderr
		os.Stderr = devNull
		t.Cleanup(func() {
			os.Stderr = original
			_ = devNull.Close()
		})

		const last = "myfuse: the reason is on the last line"
		m := newTestMounter(t, true, func(string) *exec.Cmd {
			return exec.Command("sh", "-c", fmt.Sprintf(
				`head -c %d /dev/zero | tr '\0' 'x' >&2; echo %q >&2; exit 1`,
				stderrTailLimit*3, last))
		})

		err = m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), last, "what a caller acts on is at the end of the stream")
		assert.Contains(t, err.Error(), "earlier byte(s) dropped",
			"a partial stream must not read as the whole of what the client said")
		assert.Less(t, len(err.Error()), stderrTailLimit+2048,
			"the error stays well inside the %d byte message cap", proxy.MaxMsgSize)
	})

	t.Run("a successful mount reports the pid and a channel to watch", func(t *testing.T) {
		m := newTestMounter(t, false, func(string) *exec.Cmd {
			return exec.Command("sleep", "30")
		})
		op := &mounter.MountOperation{Target: t.TempDir()}

		require.NoError(t, m.ExtendedMount(context.Background(), op))

		res, ok := op.MountResult.(server.FuseMountResult)
		require.True(t, ok, "MountResult must carry FuseMountResult")
		assert.NotZero(t, res.PID)
		assert.NotNil(t, res.ExitChan)
	})

	t.Run("an entrypoint that cannot start is reported as such", func(t *testing.T) {
		missing := filepath.Join(t.TempDir(), "no-such-entrypoint")
		m := newTestMounter(t, true, func(string) *exec.Cmd {
			return exec.Command(missing)
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "start entrypoint failed")
	})

	t.Run("the caller's deadline bounds the attempt", func(t *testing.T) {
		m := newTestMounter(t, true, func(string) *exec.Cmd {
			return exec.Command("sleep", "300")
		})
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})
		elapsed := time.Since(start)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "mount timeout after")
		assert.Less(t, elapsed, 5*time.Second, "the attempt must stop at the deadline it was given")
	})
}

// The examples all fork a preparation subprocess before exec'ing the client, to
// format a volume or set a quota. Signalling the pid the driver started leaves that
// subprocess behind holding the inherited stderr pipe open, so cmd.Wait never
// returns, so the goroutine owning the pid map entry never finishes, and the wait
// group shutdown blocks on never drains. Killing the group ends all three.
func TestExtendedMountTimeoutKillsTheWholeProcessGroup(t *testing.T) {
	pidFile := filepath.Join(t.TempDir(), "child.pid")
	m := newTestMounter(t, true, func(string) *exec.Cmd {
		// `wait` keeps the entrypoint alive and holding the pipe open, which is the
		// arrangement that deadlocks a pid-only cleanup. The subprocess outliving it
		// is the point, not an accident of the script.
		return exec.Command("sh", "-c", "sleep 300 & echo $! > "+pidFile+"; wait")
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "process group terminated with SIGTERM")
	child := waitForPid(t, pidFile)
	require.Eventually(t, func() bool { return !processAlive(child) },
		10*time.Second, 50*time.Millisecond,
		"the subprocess the entrypoint forked must not outlive the mount attempt")
	assert.True(t, waitGroupTimeout(&m.driver.wg, 10*time.Second),
		"cmd.Wait is still blocked on a stderr pipe somebody left open")

	var tracked int
	m.driver.pids.Range(func(_, _ any) bool { tracked++; return true })
	assert.Zero(t, tracked, "a stale entry gets signalled again on the next shutdown")
}

// A driver that ignores SIGTERM must still return within the grace the deadline
// arithmetic reserves for it, or a timed-out mount reports into a connection the
// caller has already abandoned.
func TestExtendedMountTimeoutHonoursShutdownGrace(t *testing.T) {
	m := newTestMounter(t, true, func(string) *exec.Cmd {
		// The shell itself ignores SIGTERM and is the process being waited on, so
		// only SIGKILL ends it. Its sleeps are short so it never blocks on a child.
		return exec.Command("sh", "-c", "trap '' TERM; while :; do sleep 0.1; done")
	})
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})
	elapsed := time.Since(start)

	require.Error(t, err)
	t.Logf("returned after %v", elapsed)
	assert.Less(t, elapsed, 200*time.Millisecond+proxy.MountShutdownGrace+time.Second,
		"winding down took longer than the reserved grace")
	assert.Greater(t, elapsed, proxy.MountShutdownGrace,
		"SIGKILL came early, so the grace is not the one the deadline reserves")
}

// Shutdown has to end on the driver's terms rather than the entrypoint's. An
// entrypoint that ignores SIGTERM keeps its subprocess alive too, since SIG_IGN
// survives both fork and exec, so nothing here would end without escalating.
func TestTerminateEscalatesPastAnEntrypointThatIgnoresSIGTERM(t *testing.T) {
	pidFile := filepath.Join(t.TempDir(), "child.pid")
	m := newTestMounter(t, true, func(string) *exec.Cmd {
		return exec.Command("sh", "-c", "trap '' TERM; sleep 300 & echo $! > "+pidFile+"; wait")
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mountErr := make(chan error, 1)
	go func() {
		mountErr <- m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})
	}()

	child := waitForPid(t, pidFile)
	require.True(t, processAlive(child), "the subprocess has to be running for this to mean anything")
	require.Eventually(t, func() bool {
		tracked := false
		m.driver.pids.Range(func(_, _ any) bool {
			tracked = true
			return false
		})
		return tracked
	}, 10*time.Second, 10*time.Millisecond,
		"the driver has to be tracking the mount before shutdown means anything")

	start := time.Now()
	m.driver.Terminate()
	elapsed := time.Since(start)

	t.Logf("Terminate returned after %v", elapsed)
	assert.Greater(t, elapsed, proxy.MountShutdownGrace,
		"it escalated before waiting out the grace it is supposed to give")
	assert.Less(t, elapsed, proxy.MountShutdownGrace+3*time.Second,
		"shutdown must end on its own terms, not on the entrypoint's")
	require.Eventually(t, func() bool { return !processAlive(child) },
		10*time.Second, 50*time.Millisecond,
		"the subprocess the entrypoint forked must not outlive shutdown")
	assert.Error(t, <-mountErr, "the mount attempt it interrupted has to report that")
}

// A group that has already exited is, from the driver's side, indistinguishable from
// one that never started. Reporting either as a failure would put a spurious error in
// the log of every shutdown that races a normal unmount.
func TestSignalGroupToleratesAProcessAlreadyGone(t *testing.T) {
	// Well above any pid the kernel hands out, so this cannot reach a process that
	// exists and cannot land on one whose pid was recycled.
	const absent = 1 << 30
	assert.NoError(t, signalGroup(absent, syscall.SIGTERM))
	assert.NoError(t, signalGroup(absent, syscall.SIGKILL))
}
