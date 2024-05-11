package pov

import (
	"context"
	"encoding/json"
	"os"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

type PovOptions struct {
	zoneID                       string
	dataRedundancyType           string
	protocolType                 string
	storageType                  string
	capacity                     int64
	fsName                       string
	throughputMode               string
	provisionedThroughputInmibps string
}

type Cloud interface {
	CreateVolume(ctx context.Context, volumeName string, diskOptions *PovOptions) (fsId, requestID string, err error)
	DeleteVolume(ctx context.Context, volumeName string) (reuqestID string, err error)
	CreateVolumeMountPoint(ctx context.Context, filesystemID string) (mpId string, err error)
	AttachVscMountPoint(ctx context.Context, mpId, fsId, instanceID string) (requestID string, err error)
	DescribeVscMountPoints(ctx context.Context, fsId, mpId string) (dvmpr *dfs.DescribeVscMountPointsResponse, err error)
	DetachVscMountPoint(ctx context.Context, mpId, filesystemID, instanceID string) (requestID string, err error)
}

type cloud struct {
	dbfc *dfs.Client
}

func newCloud() (Cloud, error) {
	// Init ECS Client
	ac := utils.GetAccessControl()
	klog.Infof("newCloud: ac: %+v", ac)
	dfsClient, err := dfs.NewClientWithOptions(GlobalConfigVar.regionID, ac.Config, ac.Credential)
	if err != nil {
		return nil, err
	}
	setPovEndPoint(GlobalConfigVar.regionID)
	return &cloud{dbfc: dfsClient}, nil
}

func (c *cloud) updateToken() error {
	ac := utils.GetAccessControl()
	dfsClient, err := dfs.NewClientWithOptions(GlobalConfigVar.regionID, ac.Config, ac.Credential)
	if err != nil {
		return err
	}
	c.dbfc = dfsClient
	return nil
}

// setPovEndPoint Set Endpoint for pov
func setPovEndPoint(regionID string) {
	aliyunep.AddEndpointMapping(regionID, "Dfs", "dfs-vpc."+regionID+".aliyuncs.com")

	// use environment endpoint setting first;
	if ep := os.Getenv("POV_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Dfs", ep)
	}
}

func (c *cloud) CreateVolume(ctx context.Context, volumeName string, diskOptions *PovOptions) (fsId, requestID string, err error) {
	cfsr := dfs.CreateCreateFileSystemRequest()
	cfsr.FileSystemName = volumeName
	cfsr.InputRegionId = GlobalConfigVar.regionID
	cfsr.ZoneId = diskOptions.zoneID
	cfsr.DataRedundancyType = diskOptions.dataRedundancyType
	cfsr.ProtocolType = diskOptions.protocolType
	cfsr.StorageType = diskOptions.storageType
	cfsr.SpaceCapacity = requests.NewInteger64(diskOptions.capacity)

	resp, err := c.dbfc.CreateFileSystem(cfsr)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.CreateFileSystem(cfsr)
		if err != nil {
			return "", "", err
		}
	}
	return resp.FileSystemId, resp.RequestId, nil
}

func (c *cloud) DeleteVolume(ctx context.Context, filesystemID string) (reqeustID string, err error) {

	cdfsr := dfs.CreateDeleteFileSystemRequest()
	cdfsr.FileSystemId = filesystemID
	cdfsr.InputRegionId = GlobalConfigVar.regionID

	resp, err := c.dbfc.DeleteFileSystem(cdfsr)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.DeleteFileSystem(cdfsr)
		if err != nil {
			return "", err
		}
	}
	return resp.RequestId, nil
}

func (c *cloud) CreateVolumeMountPoint(ctx context.Context, filesystemID string) (mpId string, err error) {
	cmp := dfs.CreateCreateVscMountPointRequest()
	cmp.FileSystemId = filesystemID
	cmp.InputRegionId = GlobalConfigVar.regionID
	resp, err := c.dbfc.CreateVscMountPoint(cmp)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.CreateVscMountPoint(cmp)
		if err != nil {
			return "", err
		}
	}
	return resp.MountPointId, nil
}

func (c *cloud) AttachVscMountPoint(ctx context.Context, mpId, fsId, instanceID string) (requestID string, err error) {
	cavmpr := dfs.CreateAttachVscMountPointRequest()
	jStr, err := json.Marshal([]string{instanceID})
	if err != nil {
		return "", err
	}
	cavmpr.InstanceIds = string(jStr)
	cavmpr.MountPointId = mpId
	cavmpr.FileSystemId = fsId
	cavmpr.InputRegionId = GlobalConfigVar.regionID
	resp, err := c.dbfc.AttachVscMountPoint(cavmpr)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.AttachVscMountPoint(cavmpr)
		if err != nil {
			return "", err
		}
	}

	return resp.RequestId, nil
}

func (c *cloud) DetachVscMountPoint(ctx context.Context, mpId, filesystemID, instanceID string) (requestID string, err error) {

	cdvmpr := dfs.CreateDetachVscMountPointRequest()
	cdvmpr.InputRegionId = GlobalConfigVar.regionID
	cdvmpr.MountPointId = mpId
	cdvmpr.FileSystemId = filesystemID
	jStr, err := json.Marshal([]string{instanceID})
	if err != nil {
		return "", err
	}
	cdvmpr.InstanceIds = string(jStr)

	resp, err := c.dbfc.DetachVscMountPoint(cdvmpr)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.DetachVscMountPoint(cdvmpr)
		if err != nil {
			return "", err
		}
	}
	return resp.RequestId, nil
}

func (c *cloud) DescribeVscMountPoints(ctx context.Context, fsId, mpId string) (*dfs.DescribeVscMountPointsResponse, error) {

	dvmp := dfs.CreateDescribeVscMountPointsRequest()
	dvmp.InputRegionId = GlobalConfigVar.regionID
	dvmp.FileSystemId = fsId
	if mpId != "" {
		dvmp.MountPointId = mpId
	}

	resp, err := c.dbfc.DescribeVscMountPoints(dvmp)
	if err != nil {
		c.updateToken()
		resp, err = c.dbfc.DescribeVscMountPoints(dvmp)
		if err != nil {
			return dfs.CreateDescribeVscMountPointsResponse(), err
		}
	}
	return resp, nil
}

type VscMountPointResp struct {
	RequestID   string
	TotalCount  string
	MountPoints []*VscMountPoint
}

type VscMountPoint struct {
	MountPointId       string
	InstanceTotalCount int
	MountPointInstance []*MPInstance
}

type MPInstance struct {
	InstanceId string
	Status     PovStatus
	Vscs       []*Vsc
}

type Vsc struct {
	Id     string
	Status PovStatus
	Type   string
}

type PovStatus int

const (
	NORMAL   PovStatus = iota // NORMAL = 0
	CREATING                  // CREATING = 1
	INVALID                   // INVALID = 2
)

func (vs PovStatus) String() string {
	return []string{"NORMAL", "CREATING", "INVALID"}[vs]
}
