package pov

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	DocumentURL = "http://100.100.100.200/latest/dynamic/instance-identity/document"

	defaultRetryCount = 5
)

type MetadataService interface {
	GetDoc() (*InstanceDocument, error)
}

// MetadataService ...
type metadataService struct {
	retryCount int
}

func NewMetadataService() MetadataService {
	return metadataService{
		retryCount: defaultRetryCount,
	}
}

type InstanceDocument struct {
	RegionID   string `json:"region-id"`
	InstanceID string `json:"instance-id"`
	ZoneID     string `json:"zone-id"`
}

func (ms metadataService) getDoc() (*InstanceDocument, error) {
	resp, err := http.Get(DocumentURL)
	if err != nil {
		return &InstanceDocument{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &InstanceDocument{}, err
	}

	result := &InstanceDocument{}
	if err = json.Unmarshal(body, result); err != nil {
		return &InstanceDocument{}, err
	}
	return result, nil
}

func (ms metadataService) GetDoc() (*InstanceDocument, error) {
	var err error
	var doc *InstanceDocument
	for i := 0; i < ms.retryCount; i++ {
		doc, err = ms.getDoc()
		if err != nil {
			continue
		}
		return doc, nil
	}
	return doc, err
}

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
	CreateVolume(ctx context.Context, volumeName string, diskOptions *PovOptions) (fsID, requestID string, err error)
	DeleteVolume(ctx context.Context, volumeName string) (reuqestID string, err error)
	CreateVolumeMountPoint(ctx context.Context, filesystemID string) (mpID string, err error)
	AttachVscMountPoint(ctx context.Context, mpID, fsID, instanceID string) (requestID string, err error)
	DescribeVscMountPoints(ctx context.Context, fsID, mpID string) (dvmpr *dfs.DescribeVscMountPointsResponse, err error)
	DetachVscMountPoint(ctx context.Context, mpID, filesystemID, instanceID string) (requestID string, err error)
}

type cloud struct {
	dbfc *dfs.Client
}

func newCloud() (Cloud, error) {
	// Init ECS Client
	ac := utils.GetAccessControl()
	dfsClient, err := dfs.NewClientWithOptions(GlobalConfigVar.regionID, ac.Config, ac.Credential)
	if err != nil {
		return nil, err
	}
	return &cloud{dbfc: dfsClient}, nil
}

func (c *cloud) CreateVolume(ctx context.Context, volumeName string, diskOptions *PovOptions) (fsID, requestID string, err error) {
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
		return "", "", err
	}
	return resp.FileSystemId, resp.RequestId, nil
}

func (c *cloud) DeleteVolume(ctx context.Context, filesystemID string) (reqeustID string, err error) {

	cdfsr := dfs.CreateDeleteFileSystemRequest()
	cdfsr.FileSystemId = filesystemID
	cdfsr.InputRegionId = GlobalConfigVar.regionID

	resp, err := c.dbfc.DeleteFileSystem(cdfsr)
	if err != nil {
		return "", err
	}
	return resp.RequestId, nil
}

func (c *cloud) CreateVolumeMountPoint(ctx context.Context, filesystemID string) (mpID string, err error) {
	cmp := dfs.CreateCreateVscMountPointRequest()
	cmp.FileSystemId = filesystemID
	cmp.InputRegionId = GlobalConfigVar.regionID
	resp, err := c.dbfc.CreateVscMountPoint(cmp)
	if err != nil {
		return "", err
	}
	return resp.MountPointId, nil
}

func (c *cloud) AttachVscMountPoint(ctx context.Context, mpID, fsID, instanceID string) (requestID string, err error) {
	cavmpr := dfs.CreateAttachVscMountPointRequest()
	jStr, err := json.Marshal([]string{instanceID})
	if err != nil {
		return "", err
	}
	cavmpr.InstanceIds = string(jStr)
	cavmpr.MountPointId = mpID
	cavmpr.FileSystemId = fsID
	cavmpr.InputRegionId = GlobalConfigVar.regionID
	resp, err := c.dbfc.AttachVscMountPoint(cavmpr)
	if err != nil {
		return "", err
	}

	return resp.RequestId, nil
}

func (c *cloud) DetachVscMountPoint(ctx context.Context, mpID, filesystemID, instanceID string) (requestID string, err error) {

	cdvmpr := dfs.CreateDetachVscMountPointRequest()
	cdvmpr.InputRegionId = GlobalConfigVar.regionID
	cdvmpr.MountPointId = mpID
	cdvmpr.FileSystemId = filesystemID
	jStr, err := json.Marshal([]string{instanceID})
	if err != nil {
		return "", err
	}
	cdvmpr.InstanceIds = string(jStr)

	resp, err := c.dbfc.DetachVscMountPoint(cdvmpr)
	if err != nil {
		return "", err
	}
	return resp.RequestId, nil
}

func (c *cloud) DescribeVscMountPoints(ctx context.Context, fsID, mpID string) (*dfs.DescribeVscMountPointsResponse, error) {

	dvmp := dfs.CreateDescribeVscMountPointsRequest()
	dvmp.InputRegionId = GlobalConfigVar.regionID
	dvmp.FileSystemId = fsID
	if mpID != "" {
		dvmp.MountPointId = mpID
	}

	resp, err := c.dbfc.DescribeVscMountPoints(dvmp)
	if err != nil {
		return dfs.CreateDescribeVscMountPointsResponse(), err
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
