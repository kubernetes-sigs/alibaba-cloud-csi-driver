//go:build !windows

package nas

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/record"
)

type fakeCNFSGetter struct {
	cnfsMap map[string]*v1beta1.ContainerNetworkFileSystem
}

func newFakeCNFSGetter(cnfsList ...*v1beta1.ContainerNetworkFileSystem) *fakeCNFSGetter {
	cnfsMap := make(map[string]*v1beta1.ContainerNetworkFileSystem)
	for _, cnfs := range cnfsList {
		if cnfs != nil {
			cnfsMap[cnfs.Name] = cnfs
		}
	}
	return &fakeCNFSGetter{cnfsMap}
}

func (f *fakeCNFSGetter) GetCNFS(_ context.Context, name string) (*v1beta1.ContainerNetworkFileSystem, error) {
	if cnfs, ok := f.cnfsMap[name]; ok {
		return cnfs, nil
	}
	return nil, fmt.Errorf("CNFS %s not found", name)
}

func startListeningFor(ctx context.Context, wg *sync.WaitGroup, listener net.Listener) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			conn, _ := listener.Accept()
			if conn != nil {
				conn.Close()
			}
		}
	}
}

func fakeCNFS(name, status, server, fallbackName string, fallbackStrategy v1beta1.FallbackStrategy) *v1beta1.ContainerNetworkFileSystem {
	return &v1beta1.ContainerNetworkFileSystem{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: v1beta1.ContainerNetworkFileSystemStatus{
			Status: status,
		},
		Spec: v1beta1.ContainerNetworkFileSystemSpec{
			Fallback: v1beta1.Fallback{
				Name:     fallbackName,
				Strategy: fallbackStrategy,
			},
			Parameters: v1beta1.Parameters{
				Server: server,
			},
		},
	}
}

func TestCNFSNeedsFallback(t *testing.T) {
	tests := []struct {
		name     string
		cnfs     *v1beta1.ContainerNetworkFileSystem
		listen   bool
		expected bool
	}{
		{
			name:     "Nil CNFS",
			cnfs:     nil,
			expected: false,
		},
		{
			name:     "No fallback strategy",
			cnfs:     &v1beta1.ContainerNetworkFileSystem{},
			expected: false,
		},
		{
			name:     "Always fallback strategy",
			cnfs:     fakeCNFS("", "", "", "", v1beta1.FallbackStrategyAlways),
			expected: true,
		},
		{
			name:     "IfConnectFailed fallback strategy - server reachable",
			cnfs:     fakeCNFS("", "", "localhost", "", v1beta1.FallbackStrategyIfConnectFailed),
			listen:   true,
			expected: false,
		},
		{
			name:     "IfConnectFailed fallback strategy - server unreachable",
			cnfs:     fakeCNFS("", "", "localhost", "", v1beta1.FallbackStrategyIfConnectFailed),
			expected: true,
		},
		{
			name:     "IfMountTargetUnhealthy fallback strategy - status Available",
			cnfs:     fakeCNFS("", v1beta1.StatusAvailable, "", "", v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			expected: false,
		},
		{
			name:     "IfMountTargetUnhealthy fallback strategy - status Unavailable",
			cnfs:     fakeCNFS("", v1beta1.StatusUnavailable, "", "", v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if tt.listen {
				var wg sync.WaitGroup
				addr := "localhost:2049"
				listener, err := net.Listen("tcp", addr)
				require.NoError(t, err)

				wg.Add(1)
				go startListeningFor(ctx, &wg, listener)

				t.Cleanup(func() {
					cancel()
					_ = listener.Close()
					wg.Wait()
				})
			}
			actual := cnfsNeedsFallback(ctx, tt.cnfs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFallbackCNFSAndRecord(t *testing.T) {
	client := fake.NewSimpleClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: "default",
		},
	})
	req := &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			utils.PodNameKey:      "pod1",
			utils.PodNamespaceKey: "default",
		},
	}
	eventRecorder := record.NewFakeRecorder(5)
	ctx, _ := utils.WithPodInfo(context.Background(), client, req)

	primaryCNFSName, fallbackCNFSName := "primary", "fallback"
	tests := []struct {
		name           string
		primaryCNFS    *v1beta1.ContainerNetworkFileSystem
		fallbackCNFS   *v1beta1.ContainerNetworkFileSystem
		expectedEvent  string
		expectFallback bool
		expectErr      bool
	}{
		{
			name:           "Always fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, "", "", fallbackCNFSName, v1beta1.FallbackStrategyAlways),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectedEvent:  fmt.Sprintf(cnfsAlwaysFallbackEventTmpl, primaryCNFSName, fallbackCNFSName),
			expectFallback: true,
		},
		{
			name:           "IfConnectFailed fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, "", "", fallbackCNFSName, v1beta1.FallbackStrategyIfConnectFailed),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectedEvent:  fmt.Sprintf(cnfsIfConnectFailedFallbackEventTmpl, primaryCNFSName, fallbackCNFSName),
			expectFallback: true,
		},
		{
			name:           "IfMountTargetUnhealthy fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, v1beta1.StatusUnavailable, "", fallbackCNFSName, v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectedEvent:  fmt.Sprintf(cnfsIfMountTargetUnhealthyFallbackEventTmpl, primaryCNFSName, fallbackCNFSName),
			expectFallback: true,
		},
		{
			name:         "Non-existent fallback CNFS",
			primaryCNFS:  fakeCNFS("primary", "", "", "non-existent-cnfs", ""),
			fallbackCNFS: fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnfsGetter := newFakeCNFSGetter(tt.fallbackCNFS)
			server := nodeServer{
				config: &internal.NodeConfig{
					KubeClient: client,
					CNFSGetter: cnfsGetter,
				},
				recorder: eventRecorder,
			}
			actual, err := server.fallbackCNFSAndRecord(ctx, req, tt.primaryCNFS)
			if tt.expectErr {
				assert.Error(t, err)
			} else if tt.expectFallback {
				assert.Equal(t, *tt.fallbackCNFS, *actual)
				assert.Len(t, eventRecorder.Events, 1)
				msg := <-eventRecorder.Events
				assert.Contains(t, msg, tt.expectedEvent)
			} else {
				assert.Equal(t, *tt.primaryCNFS, *actual)
			}
		})
	}
}

func TestIsValidServer(t *testing.T) {
	tests := []struct {
		arg      string
		expected bool
	}{
		{
			arg:      "valid.nfs.server",
			expected: true,
		},
		{
			arg:      "invalid:nfs:server",
			expected: false,
		},
		{
			arg:      "[valid::ipv6:nfs:server]",
			expected: true,
		},
		{
			arg:      "[invalid:ipv6:nfs:server]:",
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := isValidServer(tt.arg)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestParseVolumeContextAgentIdentity(t *testing.T) {
	t.Run("sandboxCredProviderName, the spelling OSS uses", func(t *testing.T) {
		opt, _, err := parseVolumeContext(map[string]string{
			"authType":                "agent-identity",
			"sandboxId":               "sb-1",
			"sandboxCredProviderName": "cp-1",
		})
		require.NoError(t, err)
		assert.Equal(t, "agent-identity", opt.AuthType)
		assert.Equal(t, "sb-1", opt.SandboxId)
		assert.Equal(t, "cp-1", opt.SandboxCredProviderName)
	})

	t.Run("credentialProviderName is accepted as OSS accepts it", func(t *testing.T) {
		opt, _, err := parseVolumeContext(map[string]string{
			"credentialProviderName": "cp-2",
		})
		require.NoError(t, err)
		assert.Equal(t, "cp-2", opt.SandboxCredProviderName)
	})

	t.Run("both spellings with the same value agree", func(t *testing.T) {
		opt, _, err := parseVolumeContext(map[string]string{
			"sandboxCredProviderName": "cp-3",
			"credentialProviderName":  "cp-3",
		})
		require.NoError(t, err)
		assert.Equal(t, "cp-3", opt.SandboxCredProviderName)
	})

	t.Run("conflicting spellings resolve to one of them without failing", func(t *testing.T) {
		// VolumeContext is a map, so which one wins is not defined; the mount
		// must not be rejected over it.
		opt, _, err := parseVolumeContext(map[string]string{
			"sandboxCredProviderName": "cp-a",
			"credentialProviderName":  "cp-b",
		})
		require.NoError(t, err)
		assert.Contains(t, []string{"cp-a", "cp-b"}, opt.SandboxCredProviderName)
	})

	t.Run("authType is lowercased", func(t *testing.T) {
		opt, _, err := parseVolumeContext(map[string]string{"authType": "Agent-Identity"})
		require.NoError(t, err)
		assert.Equal(t, "agent-identity", opt.AuthType)
	})

	t.Run("values are taken verbatim, as in OSS", func(t *testing.T) {
		// OSS does not trim either, so surrounding whitespace is not tolerated
		// by either driver rather than by only one of them.
		opt, _, err := parseVolumeContext(map[string]string{
			"authType":                " agent-identity ",
			"sandboxId":               " sb-5 ",
			"sandboxCredProviderName": " cp-5 ",
		})
		require.NoError(t, err)
		assert.Equal(t, " agent-identity ", opt.AuthType)
		assert.Equal(t, " sb-5 ", opt.SandboxId)
		assert.Equal(t, " cp-5 ", opt.SandboxCredProviderName)
	})

	t.Run("keys are matched case-insensitively", func(t *testing.T) {
		opt, _, err := parseVolumeContext(map[string]string{
			"SANDBOXCREDPROVIDERNAME": "cp-4",
			"SandboxID":               "sb-4",
		})
		require.NoError(t, err)
		assert.Equal(t, "cp-4", opt.SandboxCredProviderName)
		assert.Equal(t, "sb-4", opt.SandboxId)
	})

	t.Run("absent options stay empty", func(t *testing.T) {
		opt, cnfsName, err := parseVolumeContext(map[string]string{"server": "1.1.1.1"})
		require.NoError(t, err)
		assert.Empty(t, opt.AuthType)
		assert.Empty(t, opt.SandboxId)
		assert.Empty(t, opt.SandboxCredProviderName)
		assert.Empty(t, cnfsName)
	})
}

func TestParseVolumeContextPreexistingKeys(t *testing.T) {
	opt, cnfsName, err := parseVolumeContext(map[string]string{
		"server":                     "1.1.1.1",
		"path":                       "/share",
		"vers":                       "3",
		"options":                    "nolock,tls",
		"loopImageSize":              "2048",
		"containerNetworkFileSystem": "cnfs-1",
	})
	require.NoError(t, err)
	assert.Equal(t, "1.1.1.1", opt.Server)
	assert.Equal(t, "/share", opt.Path)
	assert.Equal(t, "3", opt.Vers)
	assert.Equal(t, []string{"nolock", "tls"}, opt.Options)
	assert.Equal(t, 2048, opt.LoopImageSize)
	assert.Equal(t, "cnfs-1", cnfsName)

	_, _, err = parseVolumeContext(map[string]string{"loopImageSize": "not-a-number"})
	assert.Error(t, err, "an unparsable loopImageSize must be rejected")
}
