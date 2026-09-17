//go:build !windows

package nas

import (
	"context"
	"maps"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type agenticfsCNFSGetterFunc func(context.Context, string) (*cnfsv1beta1.ContainerNetworkFileSystem, error)

func (f agenticfsCNFSGetterFunc) GetCNFS(ctx context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
	return f(ctx, name)
}

func TestAgenticfsVolumeOptionsAreAnIndependentSnapshot(t *testing.T) {
	// No NAS client or CNFS getter: direct filesystem parsing must be side-effect free.
	ctrl := &agenticfsController{}
	req := agenticfsDirectCreateReq(testAgenticFsPVName, 10*GiB+1, map[string]string{
		vcKeyOptions:                    "tls,vers=4,ram,noresvport",
		paramAgenticSpaceFileCountLimit: "20000",
		"ignored":                       "not-a-volume-attribute",
	})
	before := maps.Clone(req.Parameters)
	args, err := ctrl.getAgenticfsVolumeOptions(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, before, req.Parameters, "parsing must not mutate the request")
	assert.Equal(t, &agenticfsVolumeArgs{
		Name:           testAgenticFsPVName,
		FileSystemID:   testAgenticFsDirectFilesystemID,
		FileSystemPath: "/" + testAgenticFsPVName + "/",
		ZoneID:         testAgenticFsZoneID,
		VpcID:          testAgenticFsVpcID,
		VSwitchID:      testAgenticFsVSwitchID,
		SizeLimit:      11 * GiB,
		FileCountLimit: 20000,
		MountOptions:   "tls,vers=4,ram,noresvport",
	}, args)

	// Neither later request mutations nor changes to a returned volume can alter
	// the parsed options or another response built from them.
	req.Name = "changed"
	req.Parameters[filesystemIDKey] = "changed"
	req.Parameters[vcKeyOptions] = "changed"
	req.CapacityRange.RequiredBytes = 50 * GiB
	volume := args.volume(testAgenticFsAgenticSpaceID, testAgenticFsAccessPointID, testAgenticFsAPDomain)
	assert.Equal(t, &csi.Volume{
		VolumeId:      testAgenticFsPVName,
		CapacityBytes: 11 * GiB,
		VolumeContext: map[string]string{
			vcKeyServer:         testAgenticFsAPDomain,
			vcKeyPath:           "/",
			vcKeyMountProtocol:  mountProtocolAlinas,
			vcKeyOptions:        "tls,vers=4,ram,noresvport",
			filesystemIDKey:     testAgenticFsDirectFilesystemID,
			vcKeyAccesspointId:  testAgenticFsAccessPointID,
			vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
		},
	}, volume)
	volume.VolumeContext[vcKeyOptions] = "changed"
	assert.Equal(t, args.MountOptions, args.volume("space", "ap", "server").VolumeContext[vcKeyOptions])
}

func TestAgenticfsVolumeOptionsValidationOrder(t *testing.T) {
	for _, tt := range []struct {
		name      string
		mutate    func(*csi.CreateVolumeRequest)
		cnfsErr   error
		wantCalls int
		wantError string
	}{
		{
			name: "name before CNFS",
			mutate: func(req *csi.CreateVolumeRequest) {
				req.Name = "invalid/name"
				req.Parameters[ZoneID] = ""
			},
			wantError: "must not contain",
		},
		{
			name:      "CNFS error before placement",
			mutate:    func(req *csi.CreateVolumeRequest) { req.Parameters[ZoneID] = "" },
			cnfsErr:   status.Error(codes.Unavailable, "CNFS unavailable"),
			wantCalls: 1,
			wantError: "failed to get CNFS",
		},
		{
			name: "direct filesystem bypasses CNFS",
			mutate: func(req *csi.CreateVolumeRequest) {
				req.Parameters[filesystemIDKey] = testAgenticFsDirectFilesystemID
				req.Parameters[ZoneID] = ""
			},
			cnfsErr:   status.Error(codes.Unavailable, "must not be reached"),
			wantError: "parameters.zoneId is empty",
		},
		{
			name: "placement before quota",
			mutate: func(req *csi.CreateVolumeRequest) {
				req.Parameters[VpcID] = ""
				req.Parameters[paramAgenticSpaceSizeLimit] = "invalid"
			},
			wantCalls: 1,
			wantError: "parameters.vpcId is empty",
		},
		{
			name: "size parsing before file count parsing",
			mutate: func(req *csi.CreateVolumeRequest) {
				req.Parameters[paramAgenticSpaceSizeLimit] = "invalid"
				req.Parameters[paramAgenticSpaceFileCountLimit] = "invalid"
			},
			wantCalls: 1,
			wantError: "invalid parameters.agenticSpaceSizeLimit",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			ctrl := &agenticfsController{cnfsGetter: agenticfsCNFSGetterFunc(func(_ context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
				calls++
				return agenticfsCNFS(name), tt.cnfsErr
			})}
			req := agenticfsCreateReq(testAgenticFsPVName, 10*GiB, nil)
			tt.mutate(req)
			args, err := ctrl.getAgenticfsVolumeOptions(context.Background(), req)
			require.ErrorContains(t, err, tt.wantError)
			assert.Nil(t, args)
			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}
