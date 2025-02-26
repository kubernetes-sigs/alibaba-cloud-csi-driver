package desc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"k8s.io/apimachinery/pkg/util/sets"
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

type Disk struct {
	*ecs.Client
}

func (c Disk) Describe(ids []string) (Response[ecs.Disk], error) {
	req := ecs.CreateDescribeDisksRequest()
	req.DiskIds = encodeIDs(ids)
	req.PageSize = requests.NewInteger(batchSize)

	ret := Response[ecs.Disk]{}
	resp, err := c.Client.DescribeDisks(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.Disks.Disk

	if len(ret.Resources) < len(ids) {
		// Maybe we have some shared disks?
		gotIDs := make(sets.Set[string], len(ret.Resources))
		for _, disk := range ret.Resources {
			gotIDs.Insert(disk.DiskId)
		}
		var sharedIDs []string
		for _, id := range ids {
			if !gotIDs.Has(id) {
				sharedIDs = append(sharedIDs, id)
			}
		}
		if len(sharedIDs) > 0 {
			req.DiskIds = encodeIDs(sharedIDs)
			req.EnableShared = requests.NewBoolean(true)
			resp, err := c.Client.DescribeDisks(req)
			if err != nil {
				return ret, err
			}
			ret.Resources = append(ret.Resources, resp.Disks.Disk...)
		}
	}
	return ret, nil
}

func (c Disk) GetID(resource *ecs.Disk) string {
	return resource.DiskId
}

func (c Disk) Type() string {
	return "disk"
}

func (c Disk) BatchSize() int {
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
