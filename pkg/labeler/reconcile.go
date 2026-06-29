/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

package labeler

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

// Reconciler holds the labeler's shared state for processing nodes.
type Reconciler struct {
	Client     kubernetes.Interface
	ECS        cloud.ECSv2Interface
	EFLO       cloud.EFLOInterface
	RegionID   string
	NodeLister corev1listers.NodeLister
	Queue      workqueue.TypedRateLimitingInterface[string]

	instanceTypeCache *TTLCache[string, int32]
	nodeTypeCache     *TTLCache[string, int32]
	diskTypesCache    *TTLCache[diskTypeCacheKey, []string]
}

type diskTypeCacheKey struct {
	instanceType string
	zoneID       string
}

func (r *Reconciler) enqueue(obj any) {
	node, ok := obj.(*v1.Node)
	if !ok {
		return
	}
	if node.Name == "" {
		return
	}
	r.Queue.Add(node.Name)
}

func (r *Reconciler) runWorker(ctx context.Context) {
	for r.processNextItem(ctx) {
	}
}

func (r *Reconciler) processNextItem(ctx context.Context) bool {
	key, shutdown := r.Queue.Get()
	if shutdown {
		return false
	}
	defer r.Queue.Done(key)

	logger := klog.FromContext(ctx)
	err := r.reconcile(ctx, logger, key)
	if err == nil {
		r.Queue.Forget(key)
		return true
	}
	logger.Error(err, "reconcile failed", "node", key)
	r.Queue.AddRateLimited(key)
	return true
}

// reconcile handles a single Node.
// Skip conditions (in order):
//  1. Node already has disk.MaxDiskAnnotation. This avoids a thundering herd of
//     OpenAPI calls when labeler restarts / leader changes. Admins who need
//     recomputation can delete the annotation.
//  2. ProviderID / labels do not yield an instance id.
func (r *Reconciler) reconcile(ctx context.Context, logger klog.Logger, nodeName string) error {
	node, err := r.NodeLister.Get(nodeName)
	if apierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}

	logger = logger.WithValues("node", nodeName)
	ctx = klog.NewContext(ctx, logger)

	if _, ok := node.Annotations[disk.MaxDiskAnnotation]; ok {
		logger.V(5).Info("skip node with existing annotation")
		return nil
	}

	m := metadata.NewMetadata()
	m.AddKubernetesNode(node)
	m.EnableEFLO(r.EFLO)
	m.EnableOpenAPI(r.ECS)
	mp := m.WithSession(ctx)

	instanceID, err := mp.Get(metadata.InstanceID)
	if err != nil {
		if errors.Is(err, metadata.ErrUnknownMetadataKey) { // Edge node?
			logger.V(4).Info("skip node without instance id", "providerID", node.Spec.ProviderID, "err", err)
			return nil
		}
		return err
	}

	machineKind, mkErr := m.MachineKind()
	if mkErr != nil && !errors.Is(mkErr, metadata.ErrUnknownMetadataKey) {
		return fmt.Errorf("determine machine kind for %s: %w", nodeName, mkErr)
	}

	var diskQuantity int32
	var diskTypes []string
	switch machineKind {
	case metadata.MachineKindLingjun:
		diskQuantity, diskTypes, err = r.resolveLingjun(ctx, mp)
	case metadata.MachineKindECS:
		diskQuantity, diskTypes, err = r.resolveECS(ctx, mp)
	default:
		logger.V(3).Info("skip unknown machine kind", "kind", machineKind)
		return nil
	}
	if err != nil {
		return err
	}

	attachedDisks, err := disk.GetAttachedCloudDisks(ctx, r.ECS, instanceID, r.RegionID)
	if err != nil {
		return fmt.Errorf("describe disks for %s: %w", instanceID, err)
	}

	managed := disk.CollectManagedVolumes(node)
	var nonManaged int32 = 0
	for _, d := range attachedDisks {
		if !managed.Has(d) {
			nonManaged++
		}
	}
	maxVolumes := diskQuantity - nonManaged

	patch := disk.BuildNodePatch(node, int(maxVolumes), diskTypes, disk.MaxDiskAnnotation)
	if patch == nil {
		logger.V(5).Info("node already up to date")
		return nil
	}
	_, err = r.Client.CoreV1().Nodes().Patch(ctx, nodeName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
	if err != nil {
		return fmt.Errorf("patch node: %w", err)
	}
	logger.V(2).Info("patched node",
		"maxVolumes", maxVolumes,
		"diskQuantity", diskQuantity,
		"diskTypes", diskTypes,
	)
	return nil
}

func (r *Reconciler) resolveECS(ctx context.Context, mp metadata.MetadataProvider) (int32, []string, error) {
	instanceType, err := mp.Get(metadata.InstanceType)
	if err != nil {
		return 0, nil, err
	}

	diskQuantity, err := r.resolveDiskQuantity(ctx, mp, instanceType)
	if err != nil {
		return 0, nil, fmt.Errorf("get disk quantity for %s: %w", instanceType, err)
	}

	zoneID, err := mp.Get(metadata.ZoneID)
	if err != nil {
		return 0, nil, err
	}
	diskTypes, err := r.getDiskTypes(ctx, instanceType, zoneID)
	if err != nil {
		return 0, nil, fmt.Errorf("get disk types for %s/%s: %w", instanceType, zoneID, err)
	}

	return diskQuantity, diskTypes, nil
}

func (r *Reconciler) resolveLingjun(ctx context.Context, mp metadata.MetadataProvider) (int32, []string, error) {
	nodeType, err := mp.Get(metadata.LingjunNodeType)
	if err != nil {
		return 0, nil, fmt.Errorf("get nodeType for Lingjun node: %w", err)
	}

	diskQuantity, err := r.resolveLingjunDiskQuantity(ctx, mp, nodeType)
	if err != nil {
		return 0, nil, fmt.Errorf("get disk quantity for Lingjun node: %w", err)
	}

	diskTypes := disk.ParseLingjunNodeDiskTypes(os.Getenv("EBS_CATEGORY_LINGJUN_SUPPORTED"))
	return diskQuantity, diskTypes, nil
}

// resolveDiskQuantity returns the cached disk quantity for instanceType,
// computing it via nm.DiskQuantity() if not cached.
func (r *Reconciler) resolveDiskQuantity(ctx context.Context, nm metadata.MetadataProvider, instanceType string) (int32, error) {
	return r.instanceTypeCache.Get(ctx, instanceType, func() (int32, error) {
		return nm.DiskQuantity()
	})
}

func (r *Reconciler) getDiskTypes(ctx context.Context, instanceType, zoneID string) ([]string, error) {
	return r.diskTypesCache.Get(ctx, diskTypeCacheKey{instanceType, zoneID}, func() ([]string, error) {
		return disk.GetAvailableDiskTypes(ctx, r.ECS, instanceType, zoneID, r.RegionID)
	})
}

// resolveLingjunDiskQuantity returns the cached disk quantity for a Lingjun nodeType,
// computing it via mp.DiskQuantity() if not cached.
func (r *Reconciler) resolveLingjunDiskQuantity(ctx context.Context, mp metadata.MetadataProvider, nodeType string) (int32, error) {
	return r.nodeTypeCache.Get(ctx, nodeType, func() (int32, error) {
		return mp.DiskQuantity()
	})
}
