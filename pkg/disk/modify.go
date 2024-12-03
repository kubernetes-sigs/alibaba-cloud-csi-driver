package disk

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/go-logr/logr"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/waitstatus"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

type ModifyParameters struct {
	Category         Category
	PerformanceLevel PerformanceLevel
	ProvisionedIops  *int
	BurstingEnabled  *bool
	Tags             []ecs.TagResourcesTag
	RemoveTags       []string
}

type ModifyServer struct {
	ecsClient      cloud.ECSInterface
	taskWaiter     waitstatus.StatusWaiter[ecs.Task]
	runningTaskIDs sync.Map
}

type errTaskNotFinished struct {
	TaskID string
}

func (e errTaskNotFinished) Error() string {
	return fmt.Sprintf("task %s is not finished yet", e.TaskID)
}

func (errTaskNotFinished) Is(target error) bool {
	return target == context.DeadlineExceeded
}

func (m *ModifyServer) waitForTask(ctx context.Context, logger logr.Logger, taskID string) (*ecs.Task, error) {
	task, err := m.taskWaiter.WaitFor(ctx, taskID, desc.TaskSattled)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return task, errTaskNotFinished{taskID}
		}
		return task, fmt.Errorf("error while waiting for task %s to finish: %w", taskID, err)
	}
	logger.V(2).Info("task finished", "task", task.TaskId, "status", task.TaskStatus, "finishedTime", task.FinishedTime)
	return task, nil
}

func (m *ModifyServer) retrieveTask(logger logr.Logger, diskID string) (*ecs.Task, error) {
	req := ecs.CreateDescribeTasksRequest()
	req.PageSize = requests.NewInteger(1)
	req.ResourceIds = ptr.To([]string{diskID})
	req.TaskAction = "ModifyDiskSpec"

	resp, err := wrap.V1(logger, m.ecsClient.DescribeTasks)(req)
	if err != nil {
		return nil, err
	}
	var task *ecs.Task
	if len(resp.TaskSet.Task) > 0 {
		task = &resp.TaskSet.Task[0]
	}
	return task, nil
}

func (m *ModifyServer) modifyDiskSpec(ctx context.Context, logger logr.Logger, diskID, category, performanceLevel string, ProvisionedIops *int) error {
	req := ecs.CreateModifyDiskSpecRequest()
	req.DiskId = diskID
	req.DiskCategory = category
	req.PerformanceLevel = performanceLevel
	if ProvisionedIops != nil {
		req.ProvisionedIops = requests.NewInteger(*ProvisionedIops)
	}

	resp, err := wrap.V1(logger, m.ecsClient.ModifyDiskSpec)(req)
	if err != nil {
		if errors.Is(err, wrap.ErrorCode("NoChangeInDiskCategoryAndPerformanceLevel")) {
			logger.V(2).Info("disk spec not changed")
			return nil
		}
		return fmt.Errorf("error while modifying disk %s: %w", diskID, err)
	}

	logger.V(2).Info("modifying disk spec", "task", resp.TaskId)
	m.runningTaskIDs.Store(diskID, resp.TaskId)
	task, err := m.waitForTask(ctx, logger, resp.TaskId)
	if err != nil {
		return err
	}
	m.runningTaskIDs.Delete(diskID)
	if task.TaskStatus != "Finished" {
		return fmt.Errorf("unexpected task status %s", task.TaskStatus)
	}
	return nil
}

func (m *ModifyServer) modifyDiskAttribute(ctx context.Context, logger logr.Logger, diskID string, burstingEnabled bool) error {
	var err error
	for range 3 {
		req := ecs.CreateModifyDiskAttributeRequest()
		req.DiskId = diskID
		req.BurstingEnabled = requests.NewBoolean(burstingEnabled)

		_, err = wrap.V1(logger, m.ecsClient.ModifyDiskAttribute)(req)
		if err == nil {
			logger.V(2).Info("modified disk bursting enabled", "enabled", burstingEnabled)
			return nil
		}
		if errors.Is(err, wrap.ErrorCode("BurstingEnabledForModifyingDiskUnsupported")) {
			// TODO: ECS will optimize this, and we may remove this retry in the future
			logger.V(2).Info("disk still modifying, retrying")
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(2 * time.Second):
				continue
			}
		} else {
			return fmt.Errorf("error while modifying disk %s bursting enabled: %w", diskID, err)
		}
	}
	return err
}

func (m *ModifyServer) Modify(ctx context.Context, diskID string, params ModifyParameters) error {
	logger := klog.FromContext(ctx)

waitPrevious:
	taskObj, ok := m.runningTaskIDs.Load(diskID)
	if ok {
		taskID := taskObj.(string)
		logger.V(2).Info("wait for running task", "task", taskID)
		_, err := m.waitForTask(ctx, logger, taskID)
		if err != nil {
			return err
		}
		m.runningTaskIDs.Delete(diskID)
	}

	if len(params.Category) > 0 || len(params.PerformanceLevel) > 0 || params.ProvisionedIops != nil {
		err := m.modifyDiskSpec(ctx, logger, diskID, string(params.Category), string(params.PerformanceLevel), params.ProvisionedIops)
		if err != nil {
			if errors.Is(err, wrap.ErrorCode(IncorrectDiskStatus)) {
				// Already modifying? This can happen if CSI was restarted and runningTaskIDs lost
				logger.V(2).Info("maybe disk is being modified, recovering previous task")
				task, errTask := m.retrieveTask(logger, diskID)
				if errTask != nil {
					return fmt.Errorf("IncorrectDiskStatus, then retrieve task failed: %v", errTask)
				}
				if task == nil || task.TaskStatus != "Processing" {
					return fmt.Errorf("IncorrectDiskStatus, and no processing task found")
				}
				m.runningTaskIDs.Store(diskID, task.TaskId)
				goto waitPrevious
			}
			return err
		}
	}

	// Should goes after ModifyDiskSpec, because we may need to modify category to cloud_auto first
	if params.BurstingEnabled != nil {
		err := m.modifyDiskAttribute(ctx, logger, diskID, *params.BurstingEnabled)
		if err != nil {
			return err
		}
	}

	// Remove tags before adding new ones, to avoid maximum tag count limit.
	if len(params.RemoveTags) > 0 {
		req := ecs.CreateUntagResourcesRequest()
		req.ResourceType = "disk"
		req.ResourceId = ptr.To([]string{diskID})
		req.TagKey = &params.RemoveTags

		_, err := wrap.V1(logger, m.ecsClient.UntagResources)(req)
		if err != nil {
			return fmt.Errorf("error while untagging disk %s: %w", diskID, err)
		}
		logger.V(2).Info("untagged disk")
	}

	if len(params.Tags) > 0 {
		req := ecs.CreateTagResourcesRequest()
		req.ResourceType = "disk"
		req.ResourceId = ptr.To([]string{diskID})
		req.Tag = &params.Tags

		_, err := wrap.V1(logger, m.ecsClient.TagResources)(req)
		if err != nil {
			return fmt.Errorf("error while tagging disk %s: %w", diskID, err)
		}
		logger.V(2).Info("tagged disk")
	}
	return nil
}
