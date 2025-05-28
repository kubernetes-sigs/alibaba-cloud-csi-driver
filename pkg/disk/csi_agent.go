package disk

import (
	"context"
	"errors"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

type CSIAgent struct {
	csi.UnimplementedNodeServer

	ns *nodeServer
}

func NewCSIAgent() *CSIAgent {
	podCgroup, err := utils.NewPodCGroup()
	if err != nil {
		klog.Fatalf("Failed to initialize pod cgroup: %v", err)
	}

	return &CSIAgent{
		ns: &nodeServer{
			mounter:    utils.NewMounter(),
			k8smounter: k8smount.NewWithoutSystemd(""),
			podCGroup:  podCgroup,
			locks:      utils.NewVolumeLocks(),
		},
	}
}

func (a *CSIAgent) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return a.ns.NodeGetCapabilities(ctx, req)
}

func (a *CSIAgent) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return localExpandVolume(ctx, req)
}

func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	serial := req.VolumeContext["serialNumber"]
	if serial == "" {
		return nil, status.Error(codes.InvalidArgument, "serialNumber is empty")
	}

	if !a.ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer a.ns.locks.Release(req.VolumeId)

	targetPath := req.TargetPath
	notMounted, err := a.ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, status.Errorf(codes.Internal, "failed to check for mounted: %v", err)
	}
	if err == nil && !notMounted {
		return &csi.NodePublishVolumeResponse{}, nil
	}

	device, err := DefaultDeviceManager.WaitDevice(ctx, serial)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "device not found: %v", err)
	}

	err = a.ns.setupDisk(ctx, device, targetPath, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &csi.NodePublishVolumeResponse{}, nil
}

func (a *CSIAgent) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	err := utils.CleanupSimpleMount(req.TargetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cleanup %s: %v", req.TargetPath, err)
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}
