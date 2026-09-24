package common

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestIsSubstrateVolumeContext(t *testing.T) {
	tests := []struct {
		name string
		ctx  map[string]string
		want bool
	}{
		{"nil", nil, false},
		{"empty", map[string]string{}, false},
		{"missing key", map[string]string{"foo": "bar"}, false},
		{"wrong value", map[string]string{SubstrateModeKey: "false"}, false},
		{"true", map[string]string{SubstrateModeKey: "true"}, true},
		{"with other keys", map[string]string{SubstrateModeKey: "true", "server": "x"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsSubstrateVolumeContext(tt.ctx))
		})
	}
}

type fakeNodeServer struct {
	csi.UnimplementedNodeServer
	stageErr   error
	publishErr error
}

func (f *fakeNodeServer) NodeStageVolume(_ context.Context, _ *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, f.stageErr
}

func (f *fakeNodeServer) NodePublishVolume(_ context.Context, _ *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return &csi.NodePublishVolumeResponse{}, f.publishErr
}

func TestNodeStageVolumeSubstrateBypassesStagingPathCheck(t *testing.T) {
	validator := NodeServerWithValidator{NodeServer: &fakeNodeServer{}}

	resp, err := validator.NodeStageVolume(context.Background(), &csi.NodeStageVolumeRequest{
		VolumeId:          "vol-1",
		StagingTargetPath: "/var/lib/ateom-gvisor/staging/vol-1",
		VolumeCapability:  &csi.VolumeCapability{},
		VolumeContext:     map[string]string{SubstrateModeKey: "true"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestNodeStageVolumeNonSubstrateRejectsBadPath(t *testing.T) {
	validator := NodeServerWithValidator{NodeServer: &fakeNodeServer{}}

	_, err := validator.NodeStageVolume(context.Background(), &csi.NodeStageVolumeRequest{
		VolumeId:          "vol-1",
		StagingTargetPath: "/var/lib/ateom-gvisor/staging/vol-1",
		VolumeCapability:  &csi.VolumeCapability{},
	})
	assert.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestNodePublishVolumeSubstrateBypassesTargetPathCheck(t *testing.T) {
	validator := NodeServerWithValidator{NodeServer: &fakeNodeServer{}}

	resp, err := validator.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
		VolumeId:         "vol-1",
		TargetPath:       "/var/lib/ateom-gvisor/actors/uid/volumes/data",
		VolumeCapability: &csi.VolumeCapability{},
		VolumeContext:    map[string]string{SubstrateModeKey: "true"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestNodePublishVolumeNonSubstrateRejectsBadPath(t *testing.T) {
	validator := NodeServerWithValidator{NodeServer: &fakeNodeServer{}}

	_, err := validator.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
		VolumeId:         "vol-1",
		TargetPath:       "/var/lib/ateom-gvisor/actors/uid/volumes/data",
		VolumeCapability: &csi.VolumeCapability{},
	})
	assert.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestMetricRecorderRetainsKubeletPodLookupRetry(t *testing.T) {
	client := fake.NewClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "example-pod", Namespace: "example"},
		Status:     v1.PodStatus{Phase: v1.PodRunning},
	})
	failed := false
	client.PrependReactor("get", "pods", func(k8stesting.Action) (bool, runtime.Object, error) {
		if !failed {
			failed = true
			return true, nil, apierrors.NewTimeoutError("temporary pod lookup failure", 1)
		}
		return false, nil, nil
	})
	server := &NodeServerWithMetricRecorder{NodeServer: &fakeNodeServer{}, client: client}
	_, err := server.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{utils.PodNameKey: "example-pod", utils.PodNamespaceKey: "example"},
	})
	assert.NoError(t, err)
	assert.Len(t, client.Actions(), 2)
}

func TestMetricRecorderSkipsPodLookupOnlyForSubstrate(t *testing.T) {
	client := fake.NewClientset()
	server := &NodeServerWithMetricRecorder{NodeServer: &fakeNodeServer{}, client: client}
	_, err := server.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			SubstrateModeKey: "true", utils.PodNameKey: "actor", utils.PodNamespaceKey: "example",
		},
	})
	assert.NoError(t, err)
	assert.Empty(t, client.Actions())
}
