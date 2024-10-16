package waitstatus

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

const (
	pollInterval = 2 * time.Second
	batchSize    = 100 // DescribeDisks supports up to 100 disk IDs at once
)

type DiskStatusWaiter interface {
	WaitForDisk(ctx context.Context, diskID string, pred DiskStatusPredicate) (*ecs.Disk, error)
}

// a function that asserts the disk should quit waiting
type DiskStatusPredicate func(disk *ecs.Disk) bool

type ECSDescribeDisks interface {
	DescribeDisks(request *ecs.DescribeDisksRequest) (response *ecs.DescribeDisksResponse, err error)
}

// Utils for implementors of DiskStatusPredicate

func IsInstanceAttached(disk *ecs.Disk, instanceID string) bool {
	for _, a := range disk.Attachments.Attachment {
		if a.InstanceId == instanceID {
			return true
		}
	}
	return false
}
