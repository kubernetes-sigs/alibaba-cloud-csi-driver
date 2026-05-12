package disk

import (
	"context"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type diskRepo struct{}

func (ad *diskRepo) WaitRootBlock(ctx context.Context, serial string) (string, error) {
	panic("not implemented")
}

func (ad *diskRepo) GetAttached(logger klog.Logger, diskID string) (string, error) {
	panic("not implemented")
}

func (ad *diskRepo) DeleteAttached(diskID string) {
	panic("not implemented")
}

func (ad *diskRepo) ListBlocks() (sets.Set[string], error) {
	panic("not implemented")
}

func (ad *diskRepo) findDevice(ctx context.Context, diskID, serial string, before sets.Set[string]) (string, error) {
	panic("not implemented")
}
