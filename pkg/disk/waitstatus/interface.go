// waitstatus package provides a simple way to wait for a resource to reach a desired state.
//
// It differs from batcher in that this allows specific optimizations like adaptive poll intervals.
package waitstatus

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

const (
	pollInterval = 2 * time.Second
)

type StatusWaiter[T any] interface {
	// WaitFor waits until the pred returns true for the specified id.
	WaitFor(ctx context.Context, id string, pred StatusPredicate[*T]) (*T, error)
}

// a function that asserts the disk/snapshot should quit waiting
type StatusPredicate[T any] func(resource T) bool

// Utils for implementors of StatusPredicate

func IsInstanceAttached(disk *ecs.Disk, instanceID string) bool {
	for _, a := range disk.Attachments.Attachment {
		if a.InstanceId == instanceID {
			return true
		}
	}
	return false
}

func SnapshotAvailable(snapshot *ecs.Snapshot) bool {
	return snapshot.Available
}

func SnapshotCut(snapshot *ecs.Snapshot) bool {
	return snapshot.CreationTime != ""
}
