/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

// Package labeler implements a centralized metadata labeler that writes
// per-node max-disk annotation and disk type labels by calling ECS OpenAPI
// from a single controller pod. This removes the need for each node to
// hold OpenAPI credentials.
package labeler

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const (
	defaultLockName        = "metadata-labeler-alibabacloud-csi-driver"
	defaultNumWorkers      = 10
	defaultLeaseDuration   = 30 * time.Second
	defaultRenewDeadline   = 20 * time.Second
	defaultRetryPeriod     = 5 * time.Second
	defaultInstanceTypeTTL = 24 * time.Hour
)

// nodeLabelKeys is the set of label keys that must be retained in the informer
// cache for metadata extraction and disk-type patching.
var nodeLabelKeys sets.Set[string]

// nodeAnnotationKeys is the set of annotation keys that must be retained in the
// informer cache for reconcile.
var nodeAnnotationKeys sets.Set[string]

func init() {
	nodeLabelKeys = metadata.NodeMetadataLabelKeys()
	nodeAnnotationKeys = metadata.NodeMetadataAnnotationKeys().
		Insert(disk.MaxDiskAnnotation)
}

// Options configures the labeler runtime.
type Options struct {
	// Namespace holds the leader election lease.
	Namespace string
	// LockName overrides the leader election lock name.
	LockName string
	// NumWorkers is the number of reconcile workers.
	NumWorkers int
	// Identity is the unique identity used by leader election.
	// Defaults to POD_NAME or hostname if empty.
	Identity string
}

// Run starts the labeler. It blocks until ctx is done.
// It performs leader election, then runs a node informer + reconcile workqueue.
func Run(ctx context.Context, client kubernetes.Interface, ecsClient cloud.ECSv2Interface, efloClient cloud.EFLOInterface, regionID string, opts Options) error {
	if opts.Namespace == "" {
		opts.Namespace = os.Getenv("POD_NAMESPACE")
	}
	if opts.Namespace == "" {
		opts.Namespace = "kube-system"
	}
	if opts.LockName == "" {
		opts.LockName = defaultLockName
	}
	if opts.NumWorkers <= 0 {
		opts.NumWorkers = defaultNumWorkers
	}
	if opts.Identity == "" {
		if name := os.Getenv("POD_NAME"); name != "" {
			opts.Identity = name
		} else {
			hostname, err := os.Hostname()
			if err != nil {
				return fmt.Errorf("cannot determine identity: %w", err)
			}
			opts.Identity = hostname
		}
	}

	lock, err := resourcelock.New(
		resourcelock.LeasesResourceLock,
		opts.Namespace,
		opts.LockName,
		client.CoreV1(),
		client.CoordinationV1(),
		resourcelock.ResourceLockConfig{Identity: opts.Identity},
	)
	if err != nil {
		return fmt.Errorf("create resource lock: %w", err)
	}

	logger := klog.FromContext(ctx)
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	le, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: defaultLeaseDuration,
		RenewDeadline: defaultRenewDeadline,
		RetryPeriod:   defaultRetryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(leaderCtx context.Context) {
				logger.Info("became leader", "identity", opts.Identity)
				if err := runOnce(leaderCtx, client, ecsClient, efloClient, regionID, opts); err != nil {
					logger.Error(err, "exited with error")
					cancel() // release lease so another instance can take over
				}
			},
			OnStoppedLeading: func() {
				logger.Info("lost leadership", "identity", opts.Identity)
				// abort outer ctx so main can exit and get restarted by kubelet
				cancel()
			},
			OnNewLeader: func(identity string) {
				if identity != opts.Identity {
					logger.Info("another instance is leader", "leader", identity)
				}
			},
		},
		Name:            opts.LockName,
		ReleaseOnCancel: true,
	})
	if err != nil {
		return fmt.Errorf("create leader elector: %w", err)
	}

	le.Run(runCtx)
	return runCtx.Err()
}

// runOnce runs the labeler loop assuming leadership is held.
func runOnce(ctx context.Context, client kubernetes.Interface, ecsClient cloud.ECSv2Interface, efloClient cloud.EFLOInterface, regionID string, opts Options) error {
	factory := informers.NewSharedInformerFactoryWithOptions(
		client, 0,
		informers.WithTransform(nodeTransform),
	)
	nodeInformer := factory.Core().V1().Nodes()

	queue := workqueue.NewTypedRateLimitingQueueWithConfig(
		workqueue.DefaultTypedControllerRateLimiter[string](),
		workqueue.TypedRateLimitingQueueConfig[string]{Name: "labeler"},
	)
	defer queue.ShutDown()

	r := &Reconciler{
		Client:            client,
		ECS:               ecsClient,
		EFLO:              efloClient,
		RegionID:          regionID,
		NodeLister:        nodeInformer.Lister(),
		Queue:             queue,
		instanceTypeCache: NewTTLCache[string, int32](defaultInstanceTypeTTL),
		nodeTypeCache:     NewTTLCache[string, int32](defaultInstanceTypeTTL),
		diskTypesCache:    NewTTLCache[diskTypeCacheKey, []string](defaultInstanceTypeTTL),
	}

	_, err := nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj any) { r.enqueue(obj) },
		UpdateFunc: func(_, newObj any) { r.enqueue(newObj) },
	})
	if err != nil {
		return fmt.Errorf("add event handler: %w", err)
	}

	factory.Start(ctx.Done())
	logger := klog.FromContext(ctx)
	if !cache.WaitForNamedCacheSync("labeler", ctx.Done(), nodeInformer.Informer().HasSynced) {
		return fmt.Errorf("informer cache did not sync")
	}
	logger.Info("informer cache synced, starting workers", "workers", opts.NumWorkers)

	for range opts.NumWorkers {
		go r.runWorker(ctx)
	}

	<-ctx.Done()
	logger.Info("shutting down")
	return nil
}

// nodeTransform strips unused Node fields to reduce informer memory footprint
// on clusters with many nodes. Only the fields relevant to max-disk/disk-type
// computation are retained.
func nodeTransform(obj any) (any, error) {
	node, ok := obj.(*v1.Node)
	if !ok {
		return obj, nil
	}

	// Filter labels: only keep those used by metadata extraction or disk-type patching.
	filteredLabels := filterMap(node.Labels, func(k string) bool {
		return nodeLabelKeys.Has(k) || strings.HasPrefix(k, disk.NodeDiskTypeLabelPrefix)
	})

	// Filter annotations: only keep those checked or written during reconcile.
	filteredAnnotations := filterMap(node.Annotations, func(k string) bool {
		return nodeAnnotationKeys.Has(k)
	})

	trimmed := &v1.Node{
		TypeMeta: node.TypeMeta,
		ObjectMeta: metav1.ObjectMeta{
			Name:            node.Name,
			ResourceVersion: node.ResourceVersion,
			UID:             node.UID,
			Labels:          filteredLabels,
			Annotations:     filteredAnnotations,
		},
		Spec: v1.NodeSpec{
			ProviderID: node.Spec.ProviderID,
		},
		Status: v1.NodeStatus{
			VolumesInUse:    node.Status.VolumesInUse,
			VolumesAttached: node.Status.VolumesAttached,
		},
	}
	return trimmed, nil
}

// filterMap returns a new map containing only entries whose keys satisfy pred.
func filterMap(m map[string]string, pred func(string) bool) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		if pred(k) {
			result[k] = v
		}
	}
	return result
}
