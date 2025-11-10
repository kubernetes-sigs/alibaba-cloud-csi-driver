package desc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

const batchSize = 100

type Response[T any] struct {
	RequestID string
	Resources []T
}

type Client[T any] interface {
	Describe(ids []string) (resources Response[T], err error)
	GetID(resource *T) string
	Type() string
	BatchSize() int
}

func encodeIDs(ids []string) string {
	res, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	return string(res)
}

type disk struct {
	cloud.ECSInterface
}

func Disk(client cloud.ECSInterface) Client[ecs.Disk] {
	return disk{client}
}

func (c disk) Describe(ids []string) (Response[ecs.Disk], error) {
	req := ecs.CreateDescribeDisksRequest()
	req.DiskIds = encodeIDs(ids)
	req.PageSize = requests.NewInteger(batchSize)

	ret := Response[ecs.Disk]{}
	resp, err := c.DescribeDisks(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.Disks.Disk
	return ret, nil
}

func (c disk) GetID(resource *ecs.Disk) string {
	return resource.DiskId
}

func (c disk) Type() string {
	return "disk"
}

func (c disk) BatchSize() int {
	return batchSize
}

type Snapshots struct {
	*ecs.Client
}

func (c Snapshots) Describe(ids []string) (Response[ecs.Snapshot], error) {
	req := ecs.CreateDescribeSnapshotsRequest()
	req.SnapshotIds = encodeIDs(ids)
	req.PageSize = requests.NewInteger(batchSize)

	ret := Response[ecs.Snapshot]{}
	resp, err := c.Client.DescribeSnapshots(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.Snapshots.Snapshot
	return ret, nil
}

func (c Snapshots) GetID(resource *ecs.Snapshot) string {
	return resource.SnapshotId
}

func (c Snapshots) Type() string {
	return "snapshot"
}

func (c Snapshots) BatchSize() int {
	return batchSize
}
