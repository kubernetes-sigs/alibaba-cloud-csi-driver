package disk

import (
	"context"
	"errors"
	"testing"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/waitstatus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2/ktesting"
)

type fakeTaskWaiter struct {
	fn func(ctx context.Context, id string, pred waitstatus.StatusPredicate[**ecs20140526.DescribeTasksResponseBodyTaskSetTask]) (**ecs20140526.DescribeTasksResponseBodyTaskSetTask, error)
}

func (f *fakeTaskWaiter) WaitFor(ctx context.Context, id string, pred waitstatus.StatusPredicate[**ecs20140526.DescribeTasksResponseBodyTaskSetTask]) (**ecs20140526.DescribeTasksResponseBodyTaskSetTask, error) {
	return f.fn(ctx, id, pred)
}

func finishedTaskWaiter() *fakeTaskWaiter {
	return &fakeTaskWaiter{
		fn: func(ctx context.Context, id string, pred waitstatus.StatusPredicate[**ecs20140526.DescribeTasksResponseBodyTaskSetTask]) (**ecs20140526.DescribeTasksResponseBodyTaskSetTask, error) {
			task := &ecs20140526.DescribeTasksResponseBodyTaskSetTask{
				TaskId:     new(id),
				TaskStatus: new(desc.TaskFinished),
			}
			return &task, nil
		},
	}
}

func newTestModifyServer(t *testing.T) (*cloud.MockECSv2Interface, *ModifyServer) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSv2Interface(ctrl)
	return c, &ModifyServer{
		ecsClient:  c,
		taskWaiter: finishedTaskWaiter(),
	}
}

func sdkError(code string) error {
	return &tea.SDKError{
		Code:    new(code),
		Message: new(code),
	}
}

func TestModify_HappyPath(t *testing.T) {
	c, m := newTestModifyServer(t)

	// 1. verifyModifyDiskSpec: dry-run returns DryRunOperation → dirty=true.
	// Assert the exact request: the empty PerformanceLevel must be omitted,
	// otherwise ECS rejects it with InvalidDiskCategory.NotSupported.
	c.EXPECT().ModifyDiskSpec(&ecs20140526.ModifyDiskSpecRequest{
		DiskId:          new("d-test"),
		DiskCategory:    new("cloud_auto"),
		ProvisionedIops: new(int64(1000)),
		DryRun:          new(true),
	}).Return(nil, sdkError("DryRunOperation"))

	// 2. UntagResources
	c.EXPECT().UntagResources(gomock.Any()).Return(&ecs20140526.UntagResourcesResponse{}, nil)

	// 3. TagResources
	c.EXPECT().TagResources(gomock.Any()).Return(&ecs20140526.TagResourcesResponse{}, nil)

	// 4. modifyDiskSpec: actual call returns a task ID
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(&ecs20140526.ModifyDiskSpecResponse{
		Body: &ecs20140526.ModifyDiskSpecResponseBody{
			TaskId: new("t-123"),
		},
	}, nil)

	// 5. modifyDiskAttribute
	c.EXPECT().ModifyDiskAttribute(gomock.Any()).Return(&ecs20140526.ModifyDiskAttributeResponse{}, nil)

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Category:        DiskESSDAuto,
		ProvisionedIops: new(1000),
		BurstingEnabled: new(true),
		Tags: []*ecs20140526.TagResourcesRequestTag{
			{Key: new("tag2"), Value: new("value2")},
		},
		RemoveTags: []*string{new("tag1")},
	})
	require.NoError(t, err)
	var remaining []any
	m.runningTaskIDs.Range(func(key, value any) bool {
		remaining = append(remaining, key)
		return true
	})
	assert.Empty(t, remaining, "runningTaskIDs should be empty after successful modify")
}

func TestModify_NoChange(t *testing.T) {
	c, m := newTestModifyServer(t)

	// verifyModifyDiskSpec: dry-run returns NoChange → dirty=false, skip modifySpec
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError(string(NoChangeInDiskCategoryAndPerformanceLevel)))

	// UntagResources
	c.EXPECT().UntagResources(gomock.Any()).Return(&ecs20140526.UntagResourcesResponse{}, nil)

	// TagResources
	c.EXPECT().TagResources(gomock.Any()).Return(&ecs20140526.TagResourcesResponse{}, nil)

	// modifyDiskAttribute
	c.EXPECT().ModifyDiskAttribute(gomock.Any()).Return(&ecs20140526.ModifyDiskAttributeResponse{}, nil)

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Category:        DiskESSDAuto,
		ProvisionedIops: new(1000),
		BurstingEnabled: new(true),
		Tags: []*ecs20140526.TagResourcesRequestTag{
			{Key: new("tag2"), Value: new("value2")},
		},
		RemoveTags: []*string{new("tag1")},
	})
	require.NoError(t, err)
}

func TestModify_RecoverTask(t *testing.T) {
	c, m := newTestModifyServer(t)

	// Round 1: dry-run returns IncorrectDiskStatus → try to recover task
	dryRun := c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError(IncorrectDiskStatus))

	// retrieveTask finds a processing task
	c.EXPECT().DescribeTasks(gomock.Any()).Return(&ecs20140526.DescribeTasksResponse{
		Body: &ecs20140526.DescribeTasksResponseBody{
			TaskSet: &ecs20140526.DescribeTasksResponseBodyTaskSet{
				Task: []*ecs20140526.DescribeTasksResponseBodyTaskSetTask{
					{TaskId: new("t-recover"), TaskStatus: new(desc.TaskProcessing)},
				},
			},
		},
	}, nil).After(dryRun)

	// Round 2: waitForTask on the recovered task — use a waiter that returns Finished
	// (the default finishedTaskWaiter handles this)

	// Round 2: dry-run now passes
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError("DryRunOperation"))

	// Actual ModifyDiskSpec
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(&ecs20140526.ModifyDiskSpecResponse{
		Body: &ecs20140526.ModifyDiskSpecResponseBody{
			TaskId: new("t-new"),
		},
	}, nil)

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Category: DiskESSDAuto,
	})
	require.NoError(t, err)
}

func TestModify_DiskInCoolingPeriod(t *testing.T) {
	c, m := newTestModifyServer(t)

	// verifyModifyDiskSpec: dry-run returns DiskInCoolingPeriod → FailedPrecondition
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError("DiskInCoolingPeriod"))

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Category: DiskESSDAuto,
	})
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.FailedPrecondition, st.Code())
}

func TestModify_ModifyDiskSpecErrors(t *testing.T) {
	tests := []struct {
		name         string
		errorCode    string
		expectedCode codes.Code
	}{
		{"NoChange", string(NoChangeInDiskCategoryAndPerformanceLevel), codes.OK},
		{"NotFound", "InvalidDiskId.NotFound", codes.NotFound},
		{"InvalidParameter", "InvalidParameter", codes.InvalidArgument},
		{"InvalidDiskCategory", "InvalidDiskCategory.NotSupported", codes.InvalidArgument},
		{"IopsLimitExceed", "InvalidProvisionedIops.LimitExceed", codes.InvalidArgument},
		{"PerformanceLevelNotMatch", "OperationDenied.PerformanceLevelNotMatch", codes.FailedPrecondition},
		{"CoolingPeriod", "DiskInCoolingPeriod", codes.FailedPrecondition},
		{"LimitExceed", "ModifyingDiskCategoryLimitExceed", codes.FailedPrecondition},
		{"DiskNotReady", "InvalidStatus.DiskNotReady", codes.Aborted},
		{"IncorrectDiskStatus", IncorrectDiskStatus, codes.Aborted},
		{"Unknown", "SomeUnexpectedError", codes.Internal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, m := newTestModifyServer(t)

			// dry-run passes
			c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError("DryRunOperation"))
			// actual call fails
			c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError(tt.errorCode))

			_, ctx := ktesting.NewTestContext(t)
			err := m.Modify(ctx, "d-test", ModifyParameters{
				Category: DiskESSDAuto,
			})
			st, _ := status.FromError(err)
			assert.Equal(t, tt.expectedCode, st.Code())
		})
	}
}

func TestModify_TaskTimeoutThenRetry(t *testing.T) {
	c, m := newTestModifyServer(t)

	// --- First call: task times out ---
	// dry-run passes
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError("DryRunOperation"))
	// actual call returns task
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(&ecs20140526.ModifyDiskSpecResponse{
		Body: &ecs20140526.ModifyDiskSpecResponseBody{
			TaskId: new("t-slow"),
		},
	}, nil)

	// waitForTask: deadline exceeded, returns nil double-pointer (same as real Batched.WaitFor on ctx.Done())
	taskStillProcessing := &fakeTaskWaiter{
		fn: func(ctx context.Context, id string, pred waitstatus.StatusPredicate[**ecs20140526.DescribeTasksResponseBodyTaskSetTask]) (**ecs20140526.DescribeTasksResponseBodyTaskSetTask, error) {
			return nil, context.DeadlineExceeded
		},
	}
	m.taskWaiter = taskStillProcessing

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Category: DiskESSDAuto,
	})
	// errTaskProcessing Is DeadlineExceeded, so CSI framework will retry
	require.Error(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded))

	// --- Second call (CSI retry): recovers the running task ---
	// runningTaskIDs should have "t-slow" stored
	val, ok := m.runningTaskIDs.Load("d-test")
	require.True(t, ok)
	assert.Equal(t, "t-slow", val)

	// switch waiter: task now finishes
	m.taskWaiter = finishedTaskWaiter()

	// verifyModifyDiskSpec: waitForTask on "t-slow" succeeds (finished), then dry-run passes
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(nil, sdkError("DryRunOperation"))
	// actual call returns a new task
	c.EXPECT().ModifyDiskSpec(gomock.Any()).Return(&ecs20140526.ModifyDiskSpecResponse{
		Body: &ecs20140526.ModifyDiskSpecResponseBody{
			TaskId: new("t-new"),
		},
	}, nil)

	err = m.Modify(ctx, "d-test", ModifyParameters{
		Category: DiskESSDAuto,
	})
	require.NoError(t, err)

	// runningTaskIDs should be cleaned up
	var remaining []any
	m.runningTaskIDs.Range(func(key, value any) bool {
		remaining = append(remaining, key)
		return true
	})
	assert.Empty(t, remaining)
}

func TestModify_TagsOnly(t *testing.T) {
	c, m := newTestModifyServer(t)

	// Only UntagResources + TagResources, no ModifyDiskSpec or ModifyDiskAttribute
	c.EXPECT().UntagResources(gomock.Any()).Return(&ecs20140526.UntagResourcesResponse{}, nil)
	c.EXPECT().TagResources(gomock.Any()).Return(&ecs20140526.TagResourcesResponse{}, nil)

	_, ctx := ktesting.NewTestContext(t)
	err := m.Modify(ctx, "d-test", ModifyParameters{
		Tags: []*ecs20140526.TagResourcesRequestTag{
			{Key: new("tag2"), Value: new("value2")},
		},
		RemoveTags: []*string{new("tag1")},
	})
	require.NoError(t, err)
}

func TestBuildModifySpecRequest(t *testing.T) {
	// Empty Category/PerformanceLevel must NOT be sent: the v2 SDK serializes
	// pointers to empty strings, and passing e.g. an empty PerformanceLevel
	// alongside DiskCategory=cloud_auto makes ECS reject the request with
	// InvalidDiskCategory.NotSupported.
	cases := []struct {
		name     string
		params   ModifyParameters
		wantNil  bool
		category *string
		level    *string
		iops     *int64
	}{
		{
			name:    "all empty",
			params:  ModifyParameters{},
			wantNil: true,
		},
		{
			name:   "iops only",
			params: ModifyParameters{ProvisionedIops: new(10)},
			iops:   new(int64(10)),
		},
		{
			name:     "category only",
			params:   ModifyParameters{Category: DiskESSDAuto},
			category: new("cloud_auto"),
		},
		{
			name:   "performance level only",
			params: ModifyParameters{PerformanceLevel: PERFORMANCE_LEVEL1},
			level:  new("PL1"),
		},
		{
			name:     "category and level",
			params:   ModifyParameters{Category: DiskESSD, PerformanceLevel: PERFORMANCE_LEVEL2},
			category: new("cloud_essd"),
			level:    new("PL2"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.params.buildModifySpecRequest("d-test")
			if c.wantNil {
				assert.Nil(t, req)
				return
			}
			require.NotNil(t, req)
			assert.Equal(t, new("d-test"), req.DiskId)
			assert.Equal(t, c.category, req.DiskCategory)
			assert.Equal(t, c.level, req.PerformanceLevel)
			assert.Equal(t, c.iops, req.ProvisionedIops)
		})
	}
}
