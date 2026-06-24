package labeler

import (
	"context"
	"errors"
	"testing"

	ecsv2 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	clienttesting "k8s.io/client-go/testing"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/ptr"
)

func describeInstanceTypesV2Resp(instanceType string, diskQty int32) *ecsv2.DescribeInstanceTypesResponse {
	return &ecsv2.DescribeInstanceTypesResponse{
		Body: &ecsv2.DescribeInstanceTypesResponseBody{
			InstanceTypes: &ecsv2.DescribeInstanceTypesResponseBodyInstanceTypes{
				InstanceType: []*ecsv2.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType{
					{InstanceTypeId: &instanceType, DiskQuantity: &diskQty},
				},
			},
		},
	}
}

func makeAvailResp(zoneID, diskType string) *ecsv2.DescribeAvailableResourceResponse {
	return &ecsv2.DescribeAvailableResourceResponse{
		Body: &ecsv2.DescribeAvailableResourceResponseBody{
			AvailableZones: &ecsv2.DescribeAvailableResourceResponseBodyAvailableZones{
				AvailableZone: []*ecsv2.DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone{
					{
						ZoneId: &zoneID,
						AvailableResources: &ecsv2.DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources{
							AvailableResource: []*ecsv2.DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource{
								{
									Type: new("DataDisk"),
									SupportedResources: &ecsv2.DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources{
										SupportedResource: []*ecsv2.DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource{
											{Value: &diskType},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func TestCollectManagedVolumes(t *testing.T) {
	node := &v1.Node{
		Status: v1.NodeStatus{
			VolumesInUse: []v1.UniqueVolumeName{
				"kubernetes.io/csi/diskplugin.csi.alibabacloud.com^d-bp1in1",
				"kubernetes.io/csi/other.csi.k8s.io^d-bp1in2",
				"kubernetes.io/other/foo",
			},
			VolumesAttached: []v1.AttachedVolume{
				{Name: "kubernetes.io/csi/diskplugin.csi.alibabacloud.com^d-bp1at1"},
				{Name: "kubernetes.io/csi/other.csi.k8s.io^d-bp1at2"},
			},
		},
	}
	managed := disk.CollectManagedVolumes(node)
	assert.True(t, managed.Has("d-bp1in1"))
	assert.True(t, managed.Has("d-bp1at1"))
	assert.False(t, managed.Has("d-bp1in2"))
	assert.False(t, managed.Has("d-bp1at2"))
	assert.Equal(t, 2, managed.Len())
}

func TestBuildPatch(t *testing.T) {
	t.Run("no-op when up to date", func(t *testing.T) {
		node := &v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{disk.MaxDiskAnnotation: "8"},
				Labels: map[string]string{
					disk.NodeDiskTypeLabelPrefix + "cloud_essd": "available",
				},
			},
		}
		got := disk.BuildNodePatch(node, 8, []string{"cloud_essd"}, disk.MaxDiskAnnotation)
		assert.Nil(t, got)
	})

	t.Run("patch when missing", func(t *testing.T) {
		node := &v1.Node{}
		got := disk.BuildNodePatch(node, 5, []string{"cloud_essd", "cloud_ssd"}, disk.MaxDiskAnnotation)
		require.NotNil(t, got)
		require.JSONEq(t, `{
			"metadata": {
				"annotations": {"node.csi.alibabacloud.com/max-disk": "5"},
				"labels": {"node.csi.alibabacloud.com/disktype.cloud_essd": "available", "node.csi.alibabacloud.com/disktype.cloud_ssd": "available"}
			}
		}`, string(got))
	})

	t.Run("patch when annotation mismatch", func(t *testing.T) {
		node := &v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{disk.MaxDiskAnnotation: "3"},
				Labels: map[string]string{
					disk.NodeDiskTypeLabelPrefix + "cloud_essd": "available",
				},
			},
		}
		got := disk.BuildNodePatch(node, 8, []string{"cloud_essd"}, disk.MaxDiskAnnotation)
		require.NotNil(t, got)
	})
}

type testFixture struct {
	r      *Reconciler
	client *fake.Clientset
	ecs    *cloud.MockECSv2Interface
	eflo   *cloud.MockEFLOInterface
	ctrl   *gomock.Controller
}

func newTestFixture(t *testing.T, nodes ...*v1.Node) *testFixture {
	t.Helper()

	objs := make([]runtime.Object, 0, len(nodes))
	for _, n := range nodes {
		objs = append(objs, n)
	}
	client := fake.NewSimpleClientset(objs...)

	factory := informers.NewSharedInformerFactory(client, 0)
	nodeInformer := factory.Core().V1().Nodes()
	// Populate indexer synchronously so the lister returns nodes immediately
	// without needing to start the informer goroutine.
	for _, n := range nodes {
		require.NoError(t, nodeInformer.Informer().GetIndexer().Add(n))
	}

	ctrl := gomock.NewController(t)
	ecsMock := cloud.NewMockECSv2Interface(ctrl)
	efloMock := cloud.NewMockEFLOInterface(ctrl)

	q := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[string]())
	t.Cleanup(func() { q.ShutDown() })

	return &testFixture{
		r: &Reconciler{
			Client:            client,
			ECS:               ecsMock,
			EFLO:              efloMock,
			RegionID:          "cn-test",
			NodeLister:        nodeInformer.Lister(),
			Queue:             q,
			instanceTypeCache: NewTTLCache[string, int32](defaultInstanceTypeTTL),
			nodeTypeCache:     NewTTLCache[string, int32](defaultInstanceTypeTTL),
			diskTypesCache:    NewTTLCache[diskTypeCacheKey, []string](defaultInstanceTypeTTL),
		},
		client: client,
		ecs:    ecsMock,
		eflo:   efloMock,
		ctrl:   ctrl,
	}
}

func assertNoPatch(t *testing.T, f *testFixture) {
	t.Helper()
	for _, a := range f.client.Actions() {
		assert.NotEqual(t, "patch", a.GetVerb(), "unexpected patch: %+v", a)
	}
}

func TestReconcile_SkipAlreadyAnnotated(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "node-a",
			Annotations: map[string]string{disk.MaxDiskAnnotation: "8"},
		},
		Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1abc"},
	}
	f := newTestFixture(t, node)

	logger, ctx := ktesting.NewTestContext(t)
	require.NoError(t, f.r.reconcile(ctx, logger, "node-a"))
	assertNoPatch(t, f)
}

func TestReconcile_Lingjun(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "node-lj",
			Labels: map[string]string{"alibabacloud.com/lingjun-worker": "true"},
		},
		Spec: v1.NodeSpec{ProviderID: "cn-test.lingjun-01"},
	}
	f := newTestFixture(t, node)

	f.eflo.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{
		Body: &eflo_controller20221215.DescribeNodeResponseBody{
			NodeType: new("ebs-enhanced"),
		},
	}, nil)
	f.eflo.EXPECT().DescribeNodeType(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeTypeResponse{
		Body: &eflo_controller20221215.DescribeNodeTypeResponseBody{
			DiskQuantity: new(int32(10)),
		},
	}, nil)

	f.ecs.EXPECT().DescribeDisksWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&ecsv2.DescribeDisksResponse{
			Body: &ecsv2.DescribeDisksResponseBody{
				Disks: &ecsv2.DescribeDisksResponseBodyDisks{
					Disk: []*ecsv2.DescribeDisksResponseBodyDisksDisk{
						{DiskId: new("d-system"), Category: new("cloud_essd"), Status: new("In_use")},
					},
				},
			},
		}, nil)

	logger, ctx := ktesting.NewTestContext(t)
	require.NoError(t, f.r.reconcile(ctx, logger, "node-lj"))

	var patchBody []byte
	for _, a := range f.client.Actions() {
		if pa, ok := a.(clienttesting.PatchAction); ok {
			patchBody = pa.GetPatch()
		}
	}
	require.NotNil(t, patchBody, "expected a patch action for Lingjun node")
	require.JSONEq(t, `{
		"metadata": {
			"annotations": {"node.csi.alibabacloud.com/max-disk": "9"},
			"labels": {"node.csi.alibabacloud.com/disktype.cloud_auto": "available", "node.csi.alibabacloud.com/disktype.cloud_essd": "available"}
		}
	}`, string(patchBody))
}

func TestReconcile_SkipBadProviderID(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{Name: "node-x"},
		Spec:       v1.NodeSpec{ProviderID: ""},
	}
	f := newTestFixture(t, node)

	logger, ctx := ktesting.NewTestContext(t)
	require.NoError(t, f.r.reconcile(ctx, logger, "node-x"))
	assertNoPatch(t, f)
}

func TestReconcile_FullPatch(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-ok",
			Labels: map[string]string{
				"node.kubernetes.io/instance-type": "ecs.g7.large",
				"topology.kubernetes.io/zone":      "cn-test-a",
			},
		},
		Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1ok"},
	}
	f := newTestFixture(t, node)

	f.ecs.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(describeInstanceTypesV2Resp("ecs.g7.large", 8), nil)

	// zonal call: expect ZoneId == "cn-test-a" (from node label, not providerID).
	f.ecs.EXPECT().DescribeAvailableResourceWithContext(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, req *ecsv2.DescribeAvailableResourceRequest, _ *dara.RuntimeOptions) (*ecsv2.DescribeAvailableResourceResponse, error) {
			assert.Equal(t, "ecs.g7.large", ptr.Deref(req.InstanceType, ""))
			assert.Equal(t, "DataDisk", ptr.Deref(req.DestinationResource, ""))
			assert.Equal(t, "disk", ptr.Deref(req.ResourceType, ""))
			assert.Equal(t, "cn-test-a", ptr.Deref(req.ZoneId, ""), "zonal ZoneId must come from node label, not providerID region")
			assert.NotEqual(t, "region", ptr.Deref(req.Scope, ""))
			return makeAvailResp("cn-test-a", "cloud_essd"), nil
		})
	f.ecs.EXPECT().DescribeAvailableResourceWithContext(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, req *ecsv2.DescribeAvailableResourceRequest, _ *dara.RuntimeOptions) (*ecsv2.DescribeAvailableResourceResponse, error) {
			assert.Equal(t, "region", ptr.Deref(req.Scope, ""))
			assert.Equal(t, "", ptr.Deref(req.ZoneId, ""), "regional call must not carry ZoneId")
			return makeAvailResp("", "cloud_regional_disk_auto"), nil
		})

	f.ecs.EXPECT().DescribeDisksWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&ecsv2.DescribeDisksResponse{
			Body: &ecsv2.DescribeDisksResponseBody{
				Disks: &ecsv2.DescribeDisksResponseBodyDisks{
					Disk: []*ecsv2.DescribeDisksResponseBodyDisksDisk{
						{DiskId: new("d-system"), Category: new("cloud_essd"), Status: new("In_use")},
					},
				},
			},
		}, nil)

	logger, ctx := ktesting.NewTestContext(t)
	require.NoError(t, f.r.reconcile(ctx, logger, "node-ok"))

	var patchBody []byte
	for _, a := range f.client.Actions() {
		if pa, ok := a.(clienttesting.PatchAction); ok {
			patchBody = pa.GetPatch()
		}
	}
	require.NotNil(t, patchBody, "expected a patch action")
	require.JSONEq(t, `{
		"metadata": {
			"annotations": {"node.csi.alibabacloud.com/max-disk": "7"},
			"labels": {"node.csi.alibabacloud.com/disktype.cloud_essd": "available", "node.csi.alibabacloud.com/disktype.cloud_regional_disk_auto": "available"}
		}
	}`, string(patchBody))
}

func TestReconcile_APIErrors(t *testing.T) {
	errAPI := errors.New("api error")

	cases := []struct {
		name           string
		node           *v1.Node
		setupMocks     func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface)
		wantErr        bool
		wantErrContain string // substring to match in error message
	}{
		{
			name: "ecs_describe_instance_types_error",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node-ecs",
					Labels: map[string]string{
						"node.kubernetes.io/instance-type": "ecs.g7.large",
						"topology.kubernetes.io/zone":      "cn-test-a",
					},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1ecs"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				ecs.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errAPI)
			},
			wantErr:        true,
			wantErrContain: "get disk quantity for ecs.g7.large",
		},
		{
			name: "ecs_describe_available_resource_error",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node-ecs2",
					Labels: map[string]string{
						"node.kubernetes.io/instance-type": "ecs.g7.large",
						"topology.kubernetes.io/zone":      "cn-test-a",
					},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1ecs2"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				ecs.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(describeInstanceTypesV2Resp("ecs.g7.large", 8), nil)
				ecs.EXPECT().DescribeAvailableResourceWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errAPI).AnyTimes()
			},
			wantErr:        true,
			wantErrContain: "get disk types for ecs.g7.large/cn-test-a",
		},
		{
			name: "ecs_describe_disks_error",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node-ecs3",
					Labels: map[string]string{
						"node.kubernetes.io/instance-type": "ecs.g7.large",
						"topology.kubernetes.io/zone":      "cn-test-a",
					},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.i-bp1ecs3"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				ecs.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(describeInstanceTypesV2Resp("ecs.g7.large", 8), nil)
				ecs.EXPECT().DescribeAvailableResourceWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(makeAvailResp("cn-test-a", "cloud_essd"), nil)
				ecs.EXPECT().DescribeAvailableResourceWithContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(makeAvailResp("", "cloud_essd"), nil)
				ecs.EXPECT().DescribeDisksWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errAPI)
			},
			wantErr:        true,
			wantErrContain: "describe disks for",
		},
		{
			name: "eflo_describe_node_error",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "node-lj1",
					Labels: map[string]string{"alibabacloud.com/lingjun-worker": "true"},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.lingjun-01"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				eflo.EXPECT().DescribeNode(gomock.Any()).Return(nil, errAPI)
			},
			wantErr:        true,
			wantErrContain: "get nodeType for Lingjun node",
		},
		{
			name: "eflo_describe_node_type_error",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "node-lj2",
					Labels: map[string]string{"alibabacloud.com/lingjun-worker": "true"},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.lingjun-02"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				eflo.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{
					Body: &eflo_controller20221215.DescribeNodeResponseBody{
						NodeType: new("ebs-enhanced"),
					},
				}, nil)
				eflo.EXPECT().DescribeNodeType(gomock.Any()).Return(nil, errAPI)
			},
			wantErr:        true,
			wantErrContain: "get disk quantity for Lingjun node",
		},
		{
			name: "eflo_empty_node_type",
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "node-lj3",
					Labels: map[string]string{"alibabacloud.com/lingjun-worker": "true"},
				},
				Spec: v1.NodeSpec{ProviderID: "cn-test.lingjun-03"},
			},
			setupMocks: func(ecs *cloud.MockECSv2Interface, eflo *cloud.MockEFLOInterface) {
				eflo.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{
					Body: &eflo_controller20221215.DescribeNodeResponseBody{},
				}, nil)
			},
			wantErr:        true,
			wantErrContain: "get disk quantity for Lingjun node",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newTestFixture(t, c.node)
			c.setupMocks(f.ecs, f.eflo)

			logger, ctx := ktesting.NewTestContext(t)
			err := f.r.reconcile(ctx, logger, c.node.Name)

			if c.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), c.wantErrContain)
				assertNoPatch(t, f)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
