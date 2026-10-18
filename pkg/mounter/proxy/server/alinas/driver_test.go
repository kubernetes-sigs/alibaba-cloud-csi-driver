package alinas

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/mount-utils"
)

func TestAddAutoFallbackNFSMountOptions(t *testing.T) {
	tests := []struct {
		name         string
		inputOptions []string
		expected     []string
	}{
		{
			name:         "Empty options",
			inputOptions: []string{},
			expected:     []string{},
		},
		{
			name:         "EFC without VSC should add both no_add_cgroup and auto_fallback_nfs",
			inputOptions: []string{"efc"},
			expected:     []string{"efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "EFC with VSC should only add no_add_cgroup",
			inputOptions: []string{"efc", "net=vsc"},
			expected:     []string{"efc", "net=vsc", "no_add_cgroup"},
		},
		{
			name:         "EFC with non-VSC net option should add both options",
			inputOptions: []string{"efc", "net=other"},
			expected:     []string{"efc", "net=other", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC but no VSC",
			inputOptions: []string{"rw", "efc", "hard"},
			expected:     []string{"rw", "efc", "hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC and VSC",
			inputOptions: []string{"rw", "efc", "net=vsc", "hard"},
			expected:     []string{"rw", "efc", "net=vsc", "hard", "no_add_cgroup"},
		},
		{
			name:         "No EFC should not add any options",
			inputOptions: []string{"rw", "hard"},
			expected:     []string{"rw", "hard"},
		},
		{
			name:         "Comma-separated options with EFC but no VSC",
			inputOptions: []string{"rw,efc,hard"},
			expected:     []string{"rw,efc,hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Comma-separated options with EFC and VSC",
			inputOptions: []string{"rw,efc,net=vsc,hard"},
			expected:     []string{"rw,efc,net=vsc,hard", "no_add_cgroup"},
		},
		{
			name:         "Empty option string in slice",
			inputOptions: []string{"", "efc"},
			expected:     []string{"", "efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addAutoFallbackNFSMountOptions(tt.inputOptions)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtendedMounter_TracksTargets(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	// Mount should track the target
	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share",
		Target: "/mnt/nas1",
		FsType: "alinas",
	})
	assert.NoError(t, err)

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "target should be tracked after successful mount")

	// Mount a second target
	err = m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share2",
		Target: "/mnt/nas2",
		FsType: "cpfs-nfs",
	})
	assert.NoError(t, err)

	_, loaded = driver.targets.Load("/mnt/nas2")
	assert.True(t, loaded, "second target should be tracked")
}

func TestExtendedMounter_UnmountRemovesTarget(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	// Mount first
	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share",
		Target: "/mnt/nas1",
		FsType: "alinas",
	})
	assert.NoError(t, err)

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded)

	// Unmount should remove the target
	err = m.Unmount("/mnt/nas1")
	assert.NoError(t, err)

	_, loaded = driver.targets.Load("/mnt/nas1")
	assert.False(t, loaded, "target should be removed after unmount")
}

func TestExtendedMounterPassesSensitiveOptions(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source:           "192.168.1.1:/share",
		Target:           "/mnt/nas-sens",
		FsType:           "alinas",
		Options:          []string{"tls", "vers=3"},
		SensitiveOptions: []string{"access_key_id=AK", "access_key_secret=SK", "security_token=TOK"},
	})
	require.NoError(t, err)

	require.Len(t, fakeMounter.MountPoints, 1)
	// FakeMounter appends sensitiveOptions to the recorded Opts, proving they
	// were passed through MountSensitive (and thus masked in real logs).
	opts := fakeMounter.MountPoints[0].Opts
	assert.Contains(t, opts, "tls")
	assert.Contains(t, opts, "access_key_secret=SK")
	assert.Contains(t, opts, "security_token=TOK")
}

func TestExtendedMounterUnmountStopsJWTAuthRefresher(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	const target = "/mnt/nas-jwt"
	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share",
		Target: target,
		FsType: "alinas",
	})
	require.NoError(t, err)

	// Simulate the refresher registered by AlinasJWTAuthInterceptor. The sink
	// is irrelevant here: this test only covers the unmount -> stop wiring.
	refresher := jwtauth.NewRefresher(jwtauth.Opts{}, noopCredentialSink{})
	require.NoError(t, refresher.StartWith(&jwtauth.STSToken{
		Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}))
	jwtauth.DefaultManager.Add(target, refresher)

	require.NoError(t, m.Unmount(target))
	assert.False(t, jwtauth.DefaultManager.HasTarget(target),
		"jwtauth refresher should be stopped when the mount is unmounted")
}

type noopCredentialSink struct{}

func (noopCredentialSink) Apply(*jwtauth.STSToken) error { return nil }

func (noopCredentialSink) Cleanup() {}

func TestShouldCleanup(t *testing.T) {
	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")

	driver := &Driver{ResetFlagPath: flagFile}

	assert.False(t, driver.shouldCleanup(), "should return false when flag file does not exist")

	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	assert.True(t, driver.shouldCleanup(), "should return true when flag file exists")
}

func TestTerminate_CleansUpMounts(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")
	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	fakeMounter := mount.NewFakeMounter(nil)
	driver := &Driver{ResetFlagPath: flagFile}
	driver.Mounter = mounter.NewForMounter(
		&extendedMounter{driver: driver, Interface: fakeMounter},
	)
	driver.targets.Store("/mnt/nas1", struct{}{})
	driver.targets.Store("/mnt/nas2", struct{}{})

	driver.Terminate()

	// After Terminate, targets should have been cleaned up via Unmount
	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.False(t, loaded, "target /mnt/nas1 should be removed after terminate")
	_, loaded = driver.targets.Load("/mnt/nas2")
	assert.False(t, loaded, "target /mnt/nas2 should be removed after terminate")
}

func TestTerminate_CleanupDisabled(t *testing.T) {
	server.SetCleanupNASMountsOnExit(false)
	defer server.SetCleanupNASMountsOnExit(false)

	driver := &Driver{}
	driver.targets.Store("/mnt/nas1", struct{}{})

	driver.Terminate()

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "targets should remain when cleanup is disabled")
}

func TestTerminate_CleanupEnabledButNoResetFlag(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	driver := &Driver{ResetFlagPath: filepath.Join(tmpDir, "nonexistent")}
	driver.targets.Store("/mnt/nas1", struct{}{})

	driver.Terminate()

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "targets should remain when reset flag is absent")
}

// stubAlinasKeygen installs a fake `alinas-keygen` on PATH that writes a key to
// keyPath, so tests can exercise initRSAPrivateKey without the real
// aliyun-alinas-utils package or the openssl RSA-3072 generation cost.
func stubAlinasKeygen(t *testing.T, keyPath string) {
	t.Helper()
	installStubKeygen(t, "#!/bin/sh\n"+
		"set -eu\n"+
		"KEY_FILE=\""+keyPath+"\"\n"+
		"mkdir -p \"$(dirname \"$KEY_FILE\")\"\n"+
		"echo stub-key > \"$KEY_FILE\"\n"+
		"chmod 400 \"$KEY_FILE\"\n"+
		"echo \"Private key generated at $KEY_FILE\"\n")
}

// installStubKeygen writes the given script as `alinas-keygen` and prepends it
// to PATH for the duration of the test.
func installStubKeygen(t *testing.T, script string) {
	t.Helper()
	binDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(binDir, alinasKeygenCommand), []byte(script), 0755))
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func TestInitRSAPrivateKey(t *testing.T) {
	driver := &Driver{ConfigDir: t.TempDir()}
	keyPath := driver.rsaPrivateKeyPath()
	stubAlinasKeygen(t, keyPath)

	// First run generates the key at the ConfigDir-derived path.
	require.NoError(t, driver.initRSAPrivateKey())
	fi, err := os.Stat(keyPath)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0400), fi.Mode().Perm())

	data, err := os.ReadFile(keyPath)
	require.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestInitRSAPrivateKey_Idempotent(t *testing.T) {
	driver := &Driver{ConfigDir: t.TempDir()}
	keyPath := driver.rsaPrivateKeyPath()

	// Pre-create the key, then install a keygen stub that fails if invoked.
	require.NoError(t, os.MkdirAll(filepath.Dir(keyPath), 0755))
	require.NoError(t, os.WriteFile(keyPath, []byte("existing-key"), 0400))
	installStubKeygen(t, "#!/bin/sh\necho 'alinas-keygen must not run when key exists' >&2\nexit 1\n")

	// initRSAPrivateKey must skip the script (Go-side stat guard) and preserve
	// the existing key, so it is never rotated on restart.
	require.NoError(t, driver.initRSAPrivateKey())
	data, err := os.ReadFile(keyPath)
	require.NoError(t, err)
	assert.Equal(t, "existing-key", string(data), "existing key should be preserved (idempotent)")
}

func TestInitRSAPrivateKey_CommandMissing(t *testing.T) {
	// With alinas-keygen absent from PATH, initRSAPrivateKey must return an error
	// rather than silently succeeding.
	t.Setenv("PATH", t.TempDir())
	driver := &Driver{ConfigDir: t.TempDir()}
	assert.Error(t, driver.initRSAPrivateKey())
}

func TestConfigDir(t *testing.T) {
	assert.Equal(t, configDir, (&Driver{}).configDir())
	assert.Equal(t, "/custom/aliyun", (&Driver{ConfigDir: "/custom/aliyun"}).configDir())
}

func TestRSAPrivateKeyPath(t *testing.T) {
	assert.Equal(t, filepath.Join(configDir, "alinas", rsaPrivateKeyFile), (&Driver{}).rsaPrivateKeyPath())
	assert.Equal(t, "/custom/aliyun/alinas/"+rsaPrivateKeyFile, (&Driver{ConfigDir: "/custom/aliyun"}).rsaPrivateKeyPath())
}

func TestInitNASRSAPEMFlag(t *testing.T) {
	defer func() { server.InitNASRSAPEM = false }()
	assert.False(t, server.InitNASRSAPEM)
	server.InitNASRSAPEM = true
	assert.True(t, server.InitNASRSAPEM)
}

func TestResetFlagPath_Default(t *testing.T) {
	driver := &Driver{}
	assert.Equal(t, defaultResetFlagPath, driver.resetFlagPath())
}

func TestResetFlagPath_Override(t *testing.T) {
	driver := &Driver{ResetFlagPath: "/custom/path/reset"}
	assert.Equal(t, "/custom/path/reset", driver.resetFlagPath())
}

func TestTerminate_UnmountError(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")
	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	// FakeMounter with a registered mount point and UnmountFunc that returns error
	fakeMounter := mount.NewFakeMounter([]mount.MountPoint{
		{Device: "192.168.1.1:/share", Path: "/mnt/nas1", Type: "nfs"},
	})
	fakeMounter.UnmountFunc = func(path string) error {
		return fmt.Errorf("device busy")
	}
	driver := &Driver{ResetFlagPath: flagFile}
	driver.Mounter = mounter.NewForMounter(
		&extendedMounter{driver: driver, Interface: fakeMounter},
	)
	driver.targets.Store("/mnt/nas1", struct{}{})

	// Should not panic even when unmount fails
	driver.Terminate()

	// Target should still be in the map since Unmount failed (Delete only on success)
	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "target should remain when unmount fails")
}

// TestRefreshCredentialCompleteness pins which credentials may reach the vendor
// tool: the access key signs the certificate, while a security token is only part
// of an STS credential and a long-term access key has none.
func TestRefreshCredentialCompleteness(t *testing.T) {
	tests := []struct {
		name    string
		secrets map[string]string
		errLike string
	}{
		{
			name: "long-term access key needs no token",
			secrets: map[string]string{
				interceptors.SecretKeyAccessKeyID:     "LTAI5tlongterm",
				interceptors.SecretKeyAccessKeySecret: "sk",
			},
			// Reaches the vendor tool, which is absent in a unit test.
			errLike: "alinas-tls-cert-refresh",
		},
		{
			name:    "no access key at all",
			secrets: map[string]string{interceptors.SecretKeySecurityToken: "token"},
			errLike: "need akId and akSecret",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewDriver().Refresh(context.Background(), "/mnt/x", tt.secrets)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errLike)
		})
	}
}
