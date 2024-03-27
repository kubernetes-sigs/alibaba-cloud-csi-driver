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

type StatusWaiter[T any] interface {
	WaitFor(ctx context.Context, id string, pred StatusPredicate[*T]) (*T, error)
}

// a function that asserts the disk/snapshot should quit waiting
type StatusPredicate[T any] func(resource T) bool

type DescribeResourceResponse[T any] struct {
	RequestID string
	Resources []T
}

type ECSDescribeResources[T any] interface {
	Describe(ids []string) (resources DescribeResourceResponse[T], err error)
	GetID(resource *T) string
	Type() string
}

// Utils for implementors of StatusPredicate

func IsInstanceAttached(disk *ecs.Disk, instanceID string) bool {
	for _, a := range disk.Attachments.Attachment {
		if a.InstanceId == instanceID {
			return true
		}
	}
	return false
}

func SnapshotCut(snapshot *ecs.Snapshot) bool {
	return snapshot.CreationTime != ""
}
