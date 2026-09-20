package nas

import (
	"errors"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const credTarget = "/var/lib/kubelet/pods/uid/volumes/kubernetes.io~csi/pv/mount"

// brokerConfig is a node whose NAS mounts go through the mount broker, which is
// what makes installing a credential on a live mount possible at all.
func brokerConfig() *internal.NodeConfig {
	return &internal.NodeConfig{MountProxySocket: "/run/cnfs/alinas-mounter.sock"}
}

func longTermOptions(akID string) *Options {
	opt := apOptions(akID)
	opt.SecurityToken = ""
	opt.Expiration = time.Time{}
	return opt
}

func apOptions(akID string) *Options {
	return &Options{
		Accesspoint: "ap-xxx.nas.aliyuncs.com",
		Path:        "/",
		stsCredential: stsCredential{
			AkID: akID, AkSecret: "sk", SecurityToken: "token",
			Expiration: time.Now().Add(time.Hour),
		},
	}
}

func TestSyncMountCredentials(t *testing.T) {
	t.Run("installs a credential the mount does not have yet", func(t *testing.T) {
		m := &fakeRefreshMounter{can: true}
		ns := &nodeServer{mounter: m, config: brokerConfig()}
		ns.installed.record(credTarget, "STS.old")

		opt := apOptions("STS.new")
		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, opt))

		op := m.installed[credTarget]
		require.NotNil(t, op)
		assert.Equal(t, map[string]string{
			interceptors.SecretKeyAccessKeyID:     "STS.new",
			interceptors.SecretKeyAccessKeySecret: "sk",
			interceptors.SecretKeySecurityToken:   "token",
		}, op.Secrets)
		// The fstype the mount was made with, so the broker routes the refresh
		// without having to remember the mount.
		assert.Equal(t, "alinas", op.FsType)

		akID, known := ns.installed.get(credTarget)
		assert.True(t, known)
		assert.Equal(t, "STS.new", akID)
	})

	t.Run("an unchanged credential costs no refresh", func(t *testing.T) {
		// kubelet republishes about once a minute per volume; a mount already
		// running on this credential must not pay an install each time.
		m := &fakeRefreshMounter{can: true}
		ns := &nodeServer{mounter: m, config: brokerConfig()}
		ns.installed.record(credTarget, "STS.same")

		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, apOptions("STS.same")))
		assert.Empty(t, m.installed)
	})

	t.Run("a failed install is retried, not recorded", func(t *testing.T) {
		m := &fakeRefreshMounter{can: true, refreshErr: errors.New("broker restarted")}
		ns := &nodeServer{mounter: m, config: brokerConfig()}
		ns.installed.record(credTarget, "STS.old")

		err := ns.syncMountCredentials(t.Context(), "vol-123", credTarget, apOptions("STS.new"))
		require.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		akID, _ := ns.installed.get(credTarget)
		assert.Equal(t, "STS.old", akID, "the mount is still on the old credential")

		// The next republish tries again rather than believing the rotation happened.
		m.refreshErr = nil
		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, apOptions("STS.new")))
		akID, _ = ns.installed.get(credTarget)
		assert.Equal(t, "STS.new", akID)
	})

	t.Run("a mount without credentials is left alone", func(t *testing.T) {
		// Plain NFS, or jwtauth: there the broker holds the token and rotates on
		// its own schedule.
		m := &fakeRefreshMounter{can: true}
		ns := &nodeServer{mounter: m, config: brokerConfig()}

		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget,
			&Options{Accesspoint: "ap-xxx.nas.aliyuncs.com", Path: "/"}))
		assert.Empty(t, m.installed)
		_, known := ns.installed.get(credTarget)
		assert.False(t, known)
	})

	t.Run("losetup volumes are left alone", func(t *testing.T) {
		// The share is mounted elsewhere and the target is a loop device, so there
		// is nothing at this path to install a credential on.
		m := &fakeRefreshMounter{can: true}
		ns := &nodeServer{mounter: m, config: brokerConfig()}

		opt := apOptions("STS.new")
		opt.MountType = LosetupType
		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, opt))
		assert.Empty(t, m.installed)
	})

	t.Run("a long-term key a node cannot install is best effort", func(t *testing.T) {
		// Connector or agent mode: the change has nowhere to go, and the key does not
		// expire, so the mount keeps working on what it was mounted with. This is the
		// behaviour such a node always had.
		ns := &nodeServer{mounter: &recordingMounter{}, config: &internal.NodeConfig{}}
		ns.installed.record(credTarget, "LTAI.old")

		require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, longTermOptions("LTAI.new")))
		akID, _ := ns.installed.get(credTarget)
		assert.Equal(t, "LTAI.old", akID, "the mount is still on the credential it was mounted with")
	})

	t.Run("an STS credential a node cannot install fails the republish", func(t *testing.T) {
		// Silence would leave the mount to die when the token expires, and it would
		// die long after anyone connected it to this Secret.
		ns := &nodeServer{mounter: &recordingMounter{}, config: &internal.NodeConfig{}}
		ns.installed.record(credTarget, "STS.old")

		err := ns.syncMountCredentials(t.Context(), "vol-123", credTarget, apOptions("STS.new"))
		require.Error(t, err)
		assert.Equal(t, codes.FailedPrecondition, status.Code(err))
		assert.Contains(t, err.Error(), "expires")
		assert.Contains(t, err.Error(), "AlinasMountProxy")
	})
}

// TestSyncMountCredentialsWithoutARecord covers what a csi-plugin restart leaves
// behind, since the record does not survive one. A node that can install treats the
// unknown state as "install it", whatever kind of credential it is. A node that
// cannot install leaves the mount exactly as it found it, and remembers nothing, so
// what the user sees does not change with a restart.
func TestSyncMountCredentialsWithoutARecord(t *testing.T) {
	tests := []struct {
		name        string
		opt         *Options
		mounter     mounter.Mounter
		noBroker    bool
		wantInstall bool
	}{
		{
			name:        "STS token is installed",
			opt:         apOptions("STS.freshly-exchanged"),
			mounter:     &fakeRefreshMounter{can: true},
			wantInstall: true,
		},
		{
			name:        "long-term key on a node that can install is installed too",
			opt:         longTermOptions("LTAI.user"),
			mounter:     &fakeRefreshMounter{can: true},
			wantInstall: true,
		},
		{
			// Mounted before this bookkeeping existed, on a node that never could
			// rotate anything.
			name:     "node that cannot install is left alone",
			opt:      longTermOptions("LTAI.user"),
			mounter:  &recordingMounter{},
			noBroker: true,
		},
		{
			name:    "long-term key with an unreachable broker is left alone",
			opt:     longTermOptions("LTAI.user"),
			mounter: &fakeRefreshMounter{canErr: errors.New("dial unix: no such file")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := brokerConfig()
			if tt.noBroker {
				cfg = &internal.NodeConfig{}
			}
			ns := &nodeServer{mounter: tt.mounter, config: cfg}
			require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, tt.opt))

			akID, known := ns.installed.get(credTarget)
			assert.Equal(t, tt.wantInstall, known, "only an installed credential is remembered")
			if tt.wantInstall {
				assert.Equal(t, tt.opt.AkID, akID)
			}

			if m, ok := tt.mounter.(*fakeRefreshMounter); ok {
				if tt.wantInstall {
					require.NotNil(t, m.installed[credTarget])
					assert.Equal(t, tt.opt.AkID, m.installed[credTarget].Secrets[interceptors.SecretKeyAccessKeyID])
				} else {
					assert.Empty(t, m.installed)
				}
			}
		})
	}
}

// TestSyncMountCredentialsInsistsOnUnknownSTS is the half of the rule that cannot
// be relaxed: after a csi-plugin restart nothing is known about the live mount, and
// an STS credential that never reaches it takes the mount down at expiry. Staying
// quiet would mean a restart had swallowed the only warning.
func TestSyncMountCredentialsInsistsOnUnknownSTS(t *testing.T) {
	ns := &nodeServer{mounter: &recordingMounter{}, config: &internal.NodeConfig{}}

	err := ns.syncMountCredentials(t.Context(), "vol-123", credTarget, apOptions("STS.token"))
	require.Error(t, err)
	assert.Equal(t, codes.FailedPrecondition, status.Code(err))
	_, known := ns.installed.get(credTarget)
	assert.False(t, known, "nothing was installed, so nothing is remembered")
}

func TestInstalledCredentialsForget(t *testing.T) {
	var c installedCredentials
	c.record(credTarget, "STS.x")
	akID, known := c.get(credTarget)
	require.True(t, known)
	require.Equal(t, "STS.x", akID)

	// A later mount of the same path belongs to another Pod.
	c.forget(credTarget)
	_, known = c.get(credTarget)
	assert.False(t, known)

	// Nothing to compare against later, so nothing is remembered.
	c.record(credTarget, "")
	_, known = c.get(credTarget)
	assert.False(t, known)
}

// TestSyncMountCredentialsSkipsMountsTheBrokerDoesNotHold covers a Secret shared
// with a plain NFS volume: its access key belongs to no mount the broker performed,
// and asking it to install one would fail every republish with "fstype not
// supported".
func TestSyncMountCredentialsSkipsMountsTheBrokerDoesNotHold(t *testing.T) {
	m := &fakeRefreshMounter{can: true}
	ns := &nodeServer{mounter: m, config: brokerConfig()}

	opt := &Options{
		Server:        "xxx.nas.aliyuncs.com",
		Path:          "/",
		Vers:          "3",
		MountProtocol: MountProtocolNFS,
	}
	opt.AkID, opt.AkSecret = "LTAI.stray", "sk"
	require.NoError(t, ns.syncMountCredentials(t.Context(), "vol-123", credTarget, opt))
	assert.Empty(t, m.installed)
	_, known := ns.installed.get(credTarget)
	assert.False(t, known)
}

// TestCheckExpiringCredential pins which volumes a node that cannot install
// credentials may still mount. What decides it is whether the credential expires,
// not where it came from: an STS credential in a publish secret is as doomed on such
// a node as an rrsa one, while a long-term access key never needed rotation and must
// keep working.
func TestCheckExpiringCredential(t *testing.T) {
	tests := []struct {
		name    string
		opt     *Options
		wantErr bool
	}{
		{
			name:    "rrsa exchanges a credential that expires",
			opt:     &Options{AuthType: AuthTypeRRSA},
			wantErr: true,
		},
		{
			name:    "publish secret carrying a security token",
			opt:     apOptions("STS.user-supplied"),
			wantErr: true,
		},
		{
			name: "long-term access key is exempt",
			opt:  longTermOptions("LTAI.user"),
		},
		{
			name: "no credential at all",
			opt:  &Options{Accesspoint: "ap-xxx.nas.aliyuncs.com", Path: "/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// A node in connector mode: nothing here can reach a live mount.
			ns := &nodeServer{mounter: &recordingMounter{}, config: &internal.NodeConfig{}}
			err := ns.checkExpiringCredential(t.Context(), tt.opt)
			if !tt.wantErr {
				assert.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Equal(t, codes.FailedPrecondition, status.Code(err))
			assert.Contains(t, err.Error(), "AlinasMountProxy")
		})
	}

	t.Run("a node that can install accepts either", func(t *testing.T) {
		ns := &nodeServer{mounter: &fakeRefreshMounter{can: true}, config: brokerConfig()}
		assert.NoError(t, ns.checkExpiringCredential(t.Context(), &Options{AuthType: AuthTypeRRSA}))
		assert.NoError(t, ns.checkExpiringCredential(t.Context(), apOptions("STS.x")))
	})
}
