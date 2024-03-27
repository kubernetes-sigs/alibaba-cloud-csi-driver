package waitstatus

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func encodeIDs(ids []string) string {
	res, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	return string(res)
}

type DescribeDisks struct {
	*ecs.Client
}

func (c DescribeDisks) Describe(ids []string) (DescribeResourceResponse[ecs.Disk], error) {
	req := ecs.CreateDescribeDisksRequest()
	req.DiskIds = encodeIDs(ids)
	req.PageSize = requests.NewInteger(batchSize)

	ret := DescribeResourceResponse[ecs.Disk]{}
	resp, err := c.Client.DescribeDisks(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.Disks.Disk
	return ret, nil
}

func (c DescribeDisks) GetID(resource *ecs.Disk) string {
	return resource.DiskId
}

func (c DescribeDisks) Type() string {
	return "disk"
}

type DescribeSnapshots struct {
	*ecs.Client
}

func (c DescribeSnapshots) Describe(ids []string) (DescribeResourceResponse[ecs.Snapshot], error) {
	req := ecs.CreateDescribeSnapshotsRequest()
	req.SnapshotIds = encodeIDs(ids)
	req.PageSize = requests.NewInteger(batchSize)

	ret := DescribeResourceResponse[ecs.Snapshot]{}
	resp, err := c.Client.DescribeSnapshots(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.Snapshots.Snapshot
	return ret, nil
}

func (c DescribeSnapshots) GetID(resource *ecs.Snapshot) string {
	return resource.SnapshotId
}

func (c DescribeSnapshots) Type() string {
	return "snapshot"
}
