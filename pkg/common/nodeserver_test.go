package common

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		{"wrong value", map[string]string{"csi.alibabacloud.com/substrate-mode": "false"}, false},
		{"true", map[string]string{"csi.alibabacloud.com/substrate-mode": "true"}, true},
		{"with other keys", map[string]string{"csi.alibabacloud.com/substrate-mode": "true", "server": "x"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSubstrateVolumeContext(tt.ctx))
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
		VolumeContext:     map[string]string{"csi.alibabacloud.com/substrate-mode": "true"},
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
		VolumeContext:    map[string]string{"csi.alibabacloud.com/substrate-mode": "true"},
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
