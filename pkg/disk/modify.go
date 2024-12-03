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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

const (
	NoChangeInDiskCategoryAndPerformanceLevel wrap.ErrorCode = "NoChangeInDiskCategoryAndPerformanceLevel"
)

type ModifyParameters struct {
	Category         Category
	PerformanceLevel PerformanceLevel
	ProvisionedIops  *int
	BurstingEnabled  *bool
	Tags             []ecs.TagResourcesTag
	RemoveTags       []string
}

func (p *ModifyParameters) buildModifySpecRequest(diskID string) *ecs.ModifyDiskSpecRequest {
	if p.Category == "" && p.PerformanceLevel == "" && p.ProvisionedIops == nil {
		return nil
	}
	req := ecs.CreateModifyDiskSpecRequest()
	req.DiskId = diskID
	req.DiskCategory = string(p.Category)
	req.PerformanceLevel = string(p.PerformanceLevel)
	if p.ProvisionedIops != nil {
		req.ProvisionedIops = requests.NewInteger(*p.ProvisionedIops)
	}
	return req
}

type ModifyServer struct {
	ecsClient      cloud.ECSInterface
	taskWaiter     waitstatus.StatusWaiter[ecs.Task]
	runningTaskIDs sync.Map
}

type errTaskProcessing struct {
	TaskID string
}

func (e errTaskProcessing) Error() string {
	return fmt.Sprintf("task %s is processing", e.TaskID)
}

func (errTaskProcessing) Is(target error) bool {
	return target == context.DeadlineExceeded
}

func (m *ModifyServer) waitForTask(ctx context.Context, logger logr.Logger, taskID string) (*ecs.Task, error) {
	task, err := m.taskWaiter.WaitFor(ctx, taskID, desc.TaskSattled)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return task, errTaskProcessing{taskID}
		}
		return task, fmt.Errorf("error while waiting for task %s to finish: %w", taskID, err)
	}
	if task != nil {
		logger.V(2).Info("task finished", "task", task.TaskId, "status", task.TaskStatus, "finishedTime", task.FinishedTime)
	} else {
		logger.V(2).Info("task disappeared", "task", taskID)
	}
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

func mapModifySpecCode(code wrap.ErrorCode) codes.Code {
	switch code {
	case "InvalidDiskId.NotFound":
		return codes.NotFound
	case "InvalidParameter",
		"InvalidDiskCategory.NotSupported",
		"InvalidDiskCategory.ValueNotSupported",
		"InvalidProvisionedIops.LimitExceed":
		return codes.InvalidArgument
	case "OperationDenied.PerformanceLevelNotMatch",
		"DiskInCoolingPeriod",
		"ModifyingDiskCategoryLimitExceed":
		return codes.FailedPrecondition
	case "InvalidStatus.DiskNotReady", IncorrectDiskStatus:
		return codes.Aborted
	default:
		return codes.Internal
	}
}

// verifyModifyDiskSpec waits for the current modify task to complete, and dry runs the modify to verify the request is valid.
//
// req is modified in-place, do not reuse it.
// returns whether the volume is already at the requested status.
func (m *ModifyServer) verifyModifyDiskSpec(ctx context.Context, logger logr.Logger, req *ecs.ModifyDiskSpecRequest) (dirty bool, err error) {
	diskID := req.DiskId
	req.DryRun = requests.NewBoolean(true)

	var taskID string
	taskObj, ok := m.runningTaskIDs.Load(diskID)
	if ok {
		taskID = taskObj.(string)
	}

	for range 2 {
		if taskID != "" {
			logger.V(2).Info("wait for previous running task", "task", taskID)
			_, err := m.waitForTask(ctx, logger, taskID)
			if err != nil {
				return true, err
			}
			m.runningTaskIDs.Delete(diskID)
		}

		_, err = wrap.V1(logger, m.ecsClient.ModifyDiskSpec)(req)
		if err == nil { // should not happen for dry run
			return true, status.Errorf(codes.Internal, "dry run ModifyDiskSpec succeeds unexpectedly")
		}
		var code wrap.ErrorCode
		if errors.As(err, &code) {
			switch code {
			case NoChangeInDiskCategoryAndPerformanceLevel:
				logger.V(2).Info("disk spec clean")
				return false, nil
			case "DryRunOperation": // server-side verification passed
				return true, nil
			case "InvalidStatus.DiskNotReady", IncorrectDiskStatus:
				// Already modifying? This can happen if CSI was restarted and runningTaskIDs lost
				logger.V(2).Info("maybe disk is being modified, recovering previous task")
				task, taskErr := m.retrieveTask(logger, diskID)
				if taskErr != nil {
					return true, fmt.Errorf("retrieve task failed: %w", taskErr)
				}
				if task == nil || task.TaskStatus != desc.TaskProcessing {
					logger.V(2).Info("no processing task found")
					return true, status.Errorf(codes.FailedPrecondition, "disk %s has incorrect status: %v", req.DiskId, err)
				}
				taskID = task.TaskId
				m.runningTaskIDs.Store(diskID, taskID)
				// task is still running, check again
				continue
			default:
				return true, status.Error(mapModifySpecCode(code), err.Error())
			}
		}
		return true, status.Errorf(codes.Internal, "modify disk %s failed: %v", req.DiskId, err)
	}
	return true, status.Errorf(codes.Aborted, "disk %s is being modified with task %s", diskID, taskID)
}

func (m *ModifyServer) modifyDiskSpec(ctx context.Context, logger logr.Logger, req *ecs.ModifyDiskSpecRequest) error {
	resp, err := wrap.V1(logger, m.ecsClient.ModifyDiskSpec)(req)
	if err != nil {
		var code wrap.ErrorCode
		var grpcCode = codes.Internal
		if errors.As(err, &code) {
			switch code {
			case NoChangeInDiskCategoryAndPerformanceLevel:
				logger.V(2).Info("disk spec not changed")
				return nil
			default:
				grpcCode = mapModifySpecCode(code)
			}
		}
		return status.Errorf(grpcCode, "failed to modify disk %s: %v", req.DiskId, err)
	}

	logger.V(2).Info("modifying disk spec", "task", resp.TaskId)
	m.runningTaskIDs.Store(req.DiskId, resp.TaskId)
	task, err := m.waitForTask(ctx, logger, resp.TaskId)
	if err != nil {
		return err
	}
	m.runningTaskIDs.Delete(req.DiskId)
	if task == nil {
		return status.Errorf(codes.Internal, "task %s not found", resp.TaskId)
	}
	if task.TaskStatus != desc.TaskFinished {
		return status.Errorf(codes.Internal, "unexpected task status %s", task.TaskStatus)
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

	if params.Category != "" {
		category, ok := AllCategories[params.Category]
		if !ok {
			logger.Info("unknown disk category", "category", params.Category)
		} else {
			// Only do cross-request validate here. Most of the validations is done through dry-run below.
			if params.BurstingEnabled != nil && !category.Bursting {
				return status.Errorf(codes.InvalidArgument, "disk category %s does not support bursting enabled", params.Category)
			}
		}
	}

	specReq := params.buildModifySpecRequest(diskID)
	if specReq != nil {
		dirty, err := m.verifyModifyDiskSpec(ctx, logger, specReq)
		if err != nil {
			return err
		}
		if !dirty {
			specReq = nil
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

	if specReq != nil {
		specReq = params.buildModifySpecRequest(diskID) // not dry-run
		if err := m.modifyDiskSpec(ctx, logger, specReq); err != nil {
			return fmt.Errorf("error while modifying disk spec: %w", err)
		}
	}

	// Should goes after ModifyDiskSpec, because we may need to modify category to cloud_auto first
	if params.BurstingEnabled != nil {
		err := m.modifyDiskAttribute(ctx, logger, diskID, *params.BurstingEnabled)
		if err != nil {
			return err
		}
	}
	return nil
}
