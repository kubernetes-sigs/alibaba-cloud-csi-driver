//go:build !windows

package nas

import (
	"context"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
)

type recordingAgenticController struct {
	internal.MockController
	attributes map[string]string
	calls      int
}

func (c *recordingAgenticController) DeleteVolume(_ context.Context, _ *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	c.calls++
	c.attributes = pv.Spec.CSI.VolumeAttributes
	return &csi.DeleteVolumeResponse{}, nil
}

func (c *recordingAgenticController) ControllerExpandVolume(_ context.Context, _ *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	c.calls++
	c.attributes = pv.Spec.CSI.VolumeAttributes
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: 20 << 30}, nil
}

func TestSubstrateLegacyDeleteDoesNotReportFalseSuccess(t *testing.T) {
	server := newMockControllerServer()
	response, err := server.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: "substrate-old-actor-data"})
	require.Error(t, err)
	assert.Equal(t, codes.FailedPrecondition, status.Code(err))
	assert.Nil(t, response)
}

func TestSubstrateHandleRoutesWithoutPV(t *testing.T) {
	for _, operation := range []string{"delete", "expand"} {
		t.Run(operation, func(t *testing.T) {
			server := newMockControllerServer()
			backend := &recordingAgenticController{}
			server.Modes[agenticFsVolumeAs] = backend
			id := "substrate-agenticfs:v1:fs-1:agentic-1"
			if operation == "delete" {
				_, err := server.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: id})
				require.NoError(t, err)
			} else {
				_, err := server.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{VolumeId: id})
				require.NoError(t, err)
			}
			assert.Equal(t, 1, backend.calls)
			assert.Equal(t, "fs-1", backend.attributes[filesystemIDKey])
			assert.Equal(t, "agentic-1", backend.attributes[vcKeyAgenticSpaceId])
			assert.Equal(t, agenticFsVolumeAs, backend.attributes["volumeAs"])
			assert.Empty(t, server.kubeClient.(*fake.Clientset).Actions())
		})
	}
}

func TestSubstrateHandleRejectsMalformedIDs(t *testing.T) {
	for _, id := range []string{
		"substrate-agenticfs:v2:fs-1:agentic-1",
		"substrate-agenticfs:v1::agentic-1",
		"substrate-agenticfs:v1:fs-1:",
		"substrate-agenticfs:v1:../fs:agentic-1",
		"substrate-agenticfs:v1:fs-1:agentic-1:extra",
	} {
		t.Run(id, func(t *testing.T) {
			server := newMockControllerServer()
			backend := &recordingAgenticController{}
			server.Modes[agenticFsVolumeAs] = backend
			_, err := server.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: id})
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Zero(t, backend.calls)
			assert.Empty(t, server.kubeClient.(*fake.Clientset).Actions())
		})
	}
}

func TestSubstrateCreateUsesStableBackendHandleOnlyWhenExplicit(t *testing.T) {
	controller := newAgenticfsDirectCtrl(t, &fakeNasClientV2{})
	for _, mode := range []string{"", "false", "true"} {
		t.Run(mode, func(t *testing.T) {
			req := agenticfsDirectCreateReq("substrate-test-data", 10<<30, map[string]string{common.SubstrateModeKey: mode})
			args, err := controller.getAgenticfsVolumeOptions(context.Background(), req)
			require.NoError(t, err)
			volume := args.volume("agentic-1", "ap-1", "server")
			if mode == "true" {
				assert.Equal(t, "substrate-agenticfs:v1:"+testAgenticFsDirectFilesystemID+":agentic-1", volume.VolumeId)
			} else {
				assert.Equal(t, req.Name, volume.VolumeId)
			}
			assert.Equal(t, volume.VolumeId, args.volume("agentic-1", "ap-1", "server").VolumeId)
		})
	}
}

func TestSubstrateCreateDeleteReclaimsBackendWithoutPV(t *testing.T) {
	backend := newFakeNasClientV2()
	controller := newAgenticfsDirectCtrl(t, backend)
	server := newMockControllerServer()
	server.Modes[agenticFsVolumeAs] = controller
	request := agenticfsDirectCreateReq("substrate-test-data", 10<<30, map[string]string{common.SubstrateModeKey: "true"})
	created, err := server.CreateVolume(context.Background(), request)
	require.NoError(t, err)
	retried, err := server.CreateVolume(context.Background(), request)
	require.NoError(t, err)
	require.Equal(t, created.Volume.VolumeId, retried.Volume.VolumeId)
	require.NotEqual(t, request.Name, created.Volume.VolumeId)
	backend.apStatuses = []string{fakeStatusNotFound}
	backend.describeIdx = 0
	backend.callOrder = nil
	for range 2 {
		freshServer := newMockControllerServer()
		freshServer.Modes[agenticFsVolumeAs] = controller
		_, err := freshServer.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: created.Volume.VolumeId})
		require.NoError(t, err)
		assert.Empty(t, freshServer.kubeClient.(*fake.Clientset).Actions())
	}
	require.Len(t, backend.deleteAgenticSpaceReqs, 2)
	for _, req := range backend.deleteAgenticSpaceReqs {
		assert.Equal(t, testAgenticFsDirectFilesystemID, tea.StringValue(req.FileSystemId))
		assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(req.AgenticSpaceId))
		assert.Len(t, tea.StringValue(req.ClientToken), 64)
	}
	assert.Equal(t, tea.StringValue(backend.deleteAgenticSpaceReqs[0].ClientToken), tea.StringValue(backend.deleteAgenticSpaceReqs[1].ClientToken))
	assert.Equal(t, []string{
		"ListAccesspoints", "DeleteAccesspoint", "DescribeAccesspoint", "DeleteAgenticSpace",
		"ListAccesspoints", "DeleteAccesspoint", "DescribeAccesspoint", "DeleteAgenticSpace",
	}, backend.callOrder)
}

func TestSubstrateDeletePropagatesBackendFailure(t *testing.T) {
	backend := newDeleteFakeNasClientV2()
	backend.createdAccessPoints = append(backend.createdAccessPoints, apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain))
	backend.deleteAgenticSpaceErr = aliErr("Forbidden")
	server := newMockControllerServer()
	server.Modes[agenticFsVolumeAs] = newAgenticfsCtrl(t, backend)
	_, err := server.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{
		VolumeId: "substrate-agenticfs:v1:" + testAgenticFsFilesystemID + ":" + testAgenticFsAgenticSpaceID,
	})
	require.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "nas:DeleteAgenticSpace")
	assert.Contains(t, err.Error(), "Forbidden")
	require.Len(t, backend.deleteAgenticSpaceReqs, 1)
}

func TestSubstrateDeleteAlreadyGoneIsIdempotent(t *testing.T) {
	backend := newDeleteFakeNasClientV2()
	backend.listAccessPointsErr = aliErr("InvalidAgenticSpaceId.NotFound")
	server := newMockControllerServer()
	server.Modes[agenticFsVolumeAs] = newAgenticfsCtrl(t, backend)
	_, err := server.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{
		VolumeId: "substrate-agenticfs:v1:" + testAgenticFsFilesystemID + ":" + testAgenticFsAgenticSpaceID,
	})
	require.NoError(t, err)
	assert.Empty(t, backend.deleteAgenticSpaceReqs)
}
