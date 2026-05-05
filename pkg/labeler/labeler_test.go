/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

package labeler

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	ecsv2 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/klog/v2/ktesting"
)

// TestRunOnce_FullPipeline exercises the full labeler pipeline from runOnce
// (informer + cache sync + workers + reconcile) through to the node patch,
// using synctest for deterministic time control.
func TestRunOnce_FullPipeline(t *testing.T) { synctest.Test(t, testRunOnce_FullPipelineSync) }
func testRunOnce_FullPipelineSync(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsMock := cloud.NewMockECSv2Interface(ctrl)

	ecsMock.EXPECT().DescribeInstanceTypes(gomock.Any()).
		Return(describeInstanceTypesV2Resp("ecs.g7.large", 8), nil).AnyTimes()
	ecsMock.EXPECT().DescribeAvailableResource(gomock.Any()).DoAndReturn(
		func(req *ecsv2.DescribeAvailableResourceRequest) (*ecsv2.DescribeAvailableResourceResponse, error) {
			scope := ""
			if req.Scope != nil {
				scope = *req.Scope
			}
			if scope == "region" {
				return makeAvailResp("", "cloud_regional_disk_auto"), nil
			}
			zoneID := ""
			if req.ZoneId != nil {
				zoneID = *req.ZoneId
			}
			return makeAvailResp(zoneID, "cloud_essd"), nil
		}).AnyTimes()
	ecsMock.EXPECT().DescribeDisks(gomock.Any()).Return(&ecsv2.DescribeDisksResponse{
		Body: &ecsv2.DescribeDisksResponseBody{
			Disks: &ecsv2.DescribeDisksResponseBodyDisks{
				Disk: []*ecsv2.DescribeDisksResponseBodyDisksDisk{
					{DiskId: new("d-sys"), Category: new("cloud_essd"), Status: new("In_use")},
				},
			},
		},
	}, nil).AnyTimes()

	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-1",
			Labels: map[string]string{
				"node.kubernetes.io/instance-type": "ecs.g7.large",
				"topology.kubernetes.io/zone":      "cn-test-a",
			},
		},
		Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1node1"},
	}

	client := fake.NewSimpleClientset(node)

	_, ctx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		assert.ErrorIs(t, Run(ctx, client, ecsMock, "cn-test", Options{
			NumWorkers: 1,
			Identity:   "test-labeler",
		}), ctx.Err())
	}()

	time.Sleep(100 * time.Millisecond) // wait for WaitForCacheSync
	synctest.Wait()
	patchedNode, err := client.CoreV1().Nodes().Get(ctx, "node-1", metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, "7", patchedNode.Annotations[disk.MaxDiskAnnotation], "8 - 1 non-managed = 7")
	assert.Equal(t, "available", patchedNode.Labels[disk.NodeDiskTypeLabelPrefix+"cloud_essd"])
	assert.Equal(t, "available", patchedNode.Labels[disk.NodeDiskTypeLabelPrefix+"cloud_regional_disk_auto"])

	cancel()
	synctest.Wait() // for error assertion in goroutine
}
