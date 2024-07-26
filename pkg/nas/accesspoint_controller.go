package nas

import (
	"context"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"path"
	"path/filepath"
	"strconv"

	sdk "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func newAccesspointController(config *internal.ControllerConfig) (internal.Controller, error) {
	nasClient, err := config.NasClientFactory.V2(config.Region)
	if err != nil {
		return nil, err
	}
	return &accesspointController{
		config:    config,
		nasClient: nasClient,
	}, nil
}

type accesspointController struct {
	nasClient interfaces.NasClientV2Interface
	config    *internal.ControllerConfig
}

func (c *accesspointController) VolumeAs() string {
	return "accesspoint"
}

func (c *accesspointController) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	parameters := req.Parameters

	// get CNFS object
	cnfsName := parameters["containerNetworkFileSystem"]
	if cnfsName == "" {
		return nil, status.Error(codes.InvalidArgument, "storageclass parameters.containerNetworkFileSystem is empty")
	}
	cnfs, err := c.config.CNFSGetter.GetCNFS(ctx, cnfsName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.InvalidArgument, "CNFS not found: %s", cnfsName)
		}
		return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
	}

	// determine base path
	var basePath string
	if basePath = parameters["path"]; basePath == "" {
		basePath = "/"
	} else {
		basePath = filepath.Clean(basePath)
	}

	// create accesspoint
	result, err := c.createAccesspoint(ctx, req.Name, basePath, cnfs, parameters)
	if err != nil {
		return nil, err
	}

	volumeContext := map[string]string{}
	volumeContext["containerNetworkFileSystem"] = cnfs.Name
	volumeContext["path"] = "/"
	volumeContext["accesspoint"] = tea.StringValue(result.Body.AccessPoint.AccessPointDomain)
	volumeContext["accesspointId"] = tea.StringValue(result.Body.AccessPoint.AccessPointId)
	if enableRam := parameters["accesspointEnableRam"]; enableRam != "" {
		volumeContext["accesspointEnableRam"] = enableRam
	}
	resp := &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      req.Name,
			VolumeContext: volumeContext,
		},
	}

	// set dir quota
	if parameters["volumeCapacity"] == "true" || parameters["allowVolumeExpansion"] == "true" {
		capacity := req.GetCapacityRange().GetRequiredBytes()
		quota := (capacity + GiB - 1) >> 30
		resp.Volume.CapacityBytes = quota << 30
		resp.Volume.VolumeContext["volumeCapacity"] = "true"
		if err := c.nasClient.SetDirQuota(&sdk.SetDirQuotaRequest{
			FileSystemId: &cnfs.Status.FsAttributes.FilesystemID,
			Path:         tea.String(path.Join(basePath, req.Name)),
			SizeLimit:    &quota,
			QuotaType:    tea.String("Enforcement"),
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:SetDirQuota failed: %v", err)
		}
	}
	return resp, nil
}

func (c *accesspointController) createAccesspoint(ctx context.Context, name, basePath string, cnfs *cnfsv1beta1.ContainerNetworkFileSystem, parameters map[string]string) (*sdk.CreateAccessPointResponse, error) {
	filesystemId := cnfs.Status.FsAttributes.FilesystemID
	if filesystemId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing filesystemId in CNFS status")
	}
	// Only standard filesystems support AccessPoint.
	filesystemType := cnfs.Status.FsAttributes.FilesystemType
	if filesystemType != cloud.FilesystemTypeStandard {
		return nil, status.Error(codes.InvalidArgument, "only filesystems of standard type support accesspoint")
	}
	vpcId := cnfs.Status.FsAttributes.VpcID
	if vpcId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing vpc id in CNFS status")
	}
	vswId := cnfs.Status.FsAttributes.VSwitchID
	if vswId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing vSwitch id in CNFS status")
	}

	req := &sdk.CreateAccessPointRequest{
		AccessGroup:     tea.String(cloud.DefaultAccessGroup),
		FileSystemId:    &filesystemId,
		VpcId:           &vpcId,
		VswId:           &vswId,
		AccessPointName: &name,
		RootDirectory:   tea.String(path.Join(basePath, name)),
		OwnerUserId:     tea.Int32(0),
		OwnerGroupId:    tea.Int32(0),
	}

	var invalidParameters []string
	parseParameter := func(key string, target any) {
		value := parameters[key]
		if value == "" {
			return
		}
		switch p := target.(type) {
		case **string:
			*p = &value
		case **int32:
			intVal, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				invalidParameters = append(invalidParameters, key)
				return
			}
			int32Val := int32(intVal)
			*p = &int32Val
		case **bool:
			boolVal, err := strconv.ParseBool(value)
			if err != nil {
				invalidParameters = append(invalidParameters, key)
				return
			}
			*p = &boolVal
		}
	}
	parseParameter("accessGroupName", &req.AccessGroup)
	parseParameter("accesspointEnableRam", &req.EnabledRam)
	parseParameter("accesspointOwnerGroupId", &req.OwnerGroupId)
	parseParameter("accesspointOwnerUserId", &req.OwnerUserId)
	parseParameter("accesspointPermission", &req.Permission)
	parseParameter("accesspointPosixUserId", &req.PosixUserId)
	parseParameter("accesspointPosixGroupId", &req.PosixGroupId)
	parseParameter("accesspointPosixSecondaryGroupIds", &req.PosixSecondaryGroupIds)

	if len(invalidParameters) > 0 {
		key := invalidParameters[0]
		return nil, status.Errorf(codes.InvalidArgument, "storageclass parameters.%s is invalid: %q", key, parameters[key])
	}

	resp, err := c.nasClient.CreateAccesspoint(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "nas:CreateAccesspoint: %v", err)
	}
	return resp, nil
}

func (c *accesspointController) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	cnfsName := attributes["containerNetworkFileSystem"]
	if cnfsName == "" {
		return nil, status.Error(codes.InvalidArgument, "missing containerNetworkFileSystem in volume attributes")
	}
	cnfs, err := c.config.CNFSGetter.GetCNFS(ctx, cnfsName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
	}
	filesystemId := cnfs.Status.FsAttributes.FilesystemID
	if filesystemId == "" {
		return nil, status.Error(codes.InvalidArgument, "empty filesystemId")
	}
	accesspointId := attributes["accesspointId"]
	if accesspointId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing accesspointId in VolumeAttributes")
	}

	// cancel dir quota
	if attributes["volumeCapacity"] == "true" {
		apInfo, err := c.nasClient.DescribeAccesspoint(filesystemId, accesspointId)
		if err != nil {
			if cloud.IsAccessPointNotFoundError(err) {
				logrus.Infof("accesspoint %s already deleted", accesspointId)
				return &csi.DeleteVolumeResponse{}, nil
			}
			return nil, status.Errorf(codes.Internal, "nas:DescribeAccesspoint failed: %v", err)
		}
		if err := c.nasClient.CancelDirQuota(&sdk.CancelDirQuotaRequest{
			FileSystemId: &filesystemId,
			Path:         apInfo.Body.AccessPoint.RootPath,
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:CancelDirQuota failed: %v", err)
		}
	}

	// delete accesspoint
	err = c.nasClient.DeleteAccesspoint(filesystemId, accesspointId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "nas:DeleteAccesspoint failed: %v", err)
	}
	return &csi.DeleteVolumeResponse{}, nil
}

func (c *accesspointController) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	cnfsName := attributes["containerNetworkFileSystem"]
	if cnfsName == "" {
		return nil, status.Error(codes.InvalidArgument, "missing containerNetworkFileSystem in volume attributes")
	}
	cnfs, err := c.config.CNFSGetter.GetCNFS(ctx, cnfsName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
	}
	filesystemId := cnfs.Status.FsAttributes.FilesystemID
	if filesystemId == "" {
		return nil, status.Error(codes.InvalidArgument, "empty filesystemId")
	}
	accesspointId := attributes["accesspointId"]
	if accesspointId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing accesspointId in VolumeAttributes")
	}

	capacity := req.GetCapacityRange().GetRequiredBytes()
	if attributes["volumeCapacity"] == "true" {
		quota := (capacity + GiB - 1) >> 30
		apInfo, err := c.nasClient.DescribeAccesspoint(filesystemId, accesspointId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "nas:DescribeAccesspoint failed: %v", err)
		}
		if err := c.nasClient.SetDirQuota(&sdk.SetDirQuotaRequest{
			FileSystemId: &filesystemId,
			Path:         apInfo.Body.AccessPoint.RootPath,
			SizeLimit:    &quota,
			QuotaType:    tea.String("Enforcement"),
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:SetDirQuota failed: %v", err)
		}
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: quota << 30}, nil
	}
	logrus.Warn("volume capacity not enabled when provision, skip quota expandsion")
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: capacity}, nil
}
