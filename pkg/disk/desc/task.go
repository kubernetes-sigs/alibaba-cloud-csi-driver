package desc

import (
	"context"
	"strings"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/utils/ptr"
)

type Task struct {
	Client cloud.ECSv2Interface
}

func (c Task) Describe(ids []string) (Response[*ecs20140526.DescribeTasksResponseBodyTaskSetTask], error) {
	req := &ecs20140526.DescribeTasksRequest{
		TaskIds:  new(strings.Join(ids, ",")),
		PageSize: new(int32(batchSize)),
	}

	ret := Response[*ecs20140526.DescribeTasksResponseBodyTaskSetTask]{}
	// Task polling is shared across multiple waiters via the batched waiter, so it
	// is not tied to any single request's context.
	resp, err := c.Client.DescribeTasksWithContext(context.Background(), req, &dara.RuntimeOptions{})
	if err != nil {
		return ret, err
	}
	if resp.Body == nil {
		return ret, nil
	}
	ret.RequestID = ptr.Deref(resp.Body.RequestId, "")
	if resp.Body.TaskSet != nil {
		ret.Resources = resp.Body.TaskSet.Task
	}
	return ret, nil
}

func (c Task) GetID(resource **ecs20140526.DescribeTasksResponseBodyTaskSetTask) string {
	r := *resource
	if r == nil {
		return ""
	}
	return ptr.Deref(r.TaskId, "")
}

func (c Task) Type() string {
	return "task"
}

func (c Task) BatchSize() int {
	return batchSize
}

const (
	TaskFinished   = "Finished"
	TaskFailed     = "Failed"
	TaskProcessing = "Processing"
)

func TaskSattled(pt **ecs20140526.DescribeTasksResponseBodyTaskSetTask) bool {
	t := *pt
	if t == nil {
		return false
	}
	status := ptr.Deref(t.TaskStatus, "")
	return status == TaskFinished || status == TaskFailed
}
