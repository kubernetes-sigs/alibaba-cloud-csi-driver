package nas

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
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

func startListeningFor(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		if conn != nil {
			conn.Close()
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
		expected bool
	}{
		{
			name:     "Nil CNFS",
			cnfs:     nil,
			expected: false,
		},
		{
			name:     "Always fallback strategy",
			cnfs:     fakeCNFS("", "", "", "", v1beta1.FallbackStrategyAlways),
			expected: true,
		},
		{
			name:     "IfConnectFailed fallback strategy-server reachable",
			cnfs:     fakeCNFS("", "", "localhost", "", v1beta1.FallbackStrategyIfConnectFailed),
			expected: false,
		},
		{
			name:     "IfMountTargetUnhealthy fallback strategy-status Available",
			cnfs:     fakeCNFS("", v1beta1.StatusAvailable, "", "", v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			expected: false,
		},
		{
			name:     "IfMountTargetUnhealthy fallback strategy-status Unavailable",
			cnfs:     fakeCNFS("", v1beta1.StatusUnavailable, "", "", v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			expected: true,
		},
	}

	go startListeningFor("localhost:2049")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := cnfsNeedsFallback(tt.cnfs)
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
		expectFallback bool
		expectErr      bool
	}{
		{
			name:           "Always fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, "", "", fallbackCNFSName, v1beta1.FallbackStrategyAlways),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectFallback: true,
		},
		{
			name:           "IfConnectFailed fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, "", "", fallbackCNFSName, v1beta1.FallbackStrategyIfConnectFailed),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
			expectFallback: true,
		},
		{
			name:           "IfMountTargetUnhealthy fallback strategy",
			primaryCNFS:    fakeCNFS(primaryCNFSName, v1beta1.StatusUnavailable, "", fallbackCNFSName, v1beta1.FallbackStrategyIfMountTargetUnhealthy),
			fallbackCNFS:   fakeCNFS(fallbackCNFSName, "", "", "", ""),
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
			} else {
				if tt.expectFallback {
					assert.Equal(t, *tt.fallbackCNFS, *actual)
					assert.Len(t, eventRecorder.Events, 1)

					oldCNFSName, newCNFSName := tt.primaryCNFS.Name, tt.fallbackCNFS.Name
					msg := <-eventRecorder.Events
					switch tt.primaryCNFS.Spec.Fallback.Strategy {
					case v1beta1.FallbackStrategyAlways:
						assert.Contains(t, msg, fmt.Sprintf(cnfsAlwaysFallbackEventTmpl, oldCNFSName, newCNFSName))
					case v1beta1.FallbackStrategyIfConnectFailed:
						assert.Contains(t, msg, fmt.Sprintf(cnfsIfConnectFailedFallbackEventTmpl, oldCNFSName, newCNFSName))
					case v1beta1.FallbackStrategyIfMountTargetUnhealthy:
						assert.Contains(t, msg, fmt.Sprintf(cnfsIfMountTargetUnhealthyFallbackEventTmpl, oldCNFSName, newCNFSName))
					}
				} else {
					assert.Equal(t, *tt.primaryCNFS, *actual)
				}
			}
		})
	}
}
