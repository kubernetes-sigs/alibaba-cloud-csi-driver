package testing

import (
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
)

type FakeClient struct {
	Disks    sync.Map
	Requests [][]string
}

func (c *FakeClient) Describe(ids []string) (desc.Response[ecs.Disk], error) {
	c.Requests = append(c.Requests, ids)
	resp := desc.Response[ecs.Disk]{}
	for _, id := range ids {
		if disk, ok := c.Disks.Load(id); ok {
			resp.Resources = append(resp.Resources, *disk.(*ecs.Disk))
		}
	}
	return resp, nil
}

func (c *FakeClient) GetID(resource *ecs.Disk) string {
	return resource.DiskId
}

func (c *FakeClient) Type() string {
	return "fakeDisk"
}

func (c *FakeClient) BatchSize() int {
	return 100
}
