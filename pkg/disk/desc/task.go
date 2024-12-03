package desc

import (
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type Task struct {
	*ecs.Client
}

func (c Task) Describe(ids []string) (Response[ecs.Task], error) {
	req := ecs.CreateDescribeTasksRequest()
	req.TaskIds = strings.Join(ids, ",")
	req.PageSize = requests.NewInteger(batchSize)

	ret := Response[ecs.Task]{}
	resp, err := c.Client.DescribeTasks(req)
	if err != nil {
		return ret, err
	}
	ret.RequestID = resp.RequestId
	ret.Resources = resp.TaskSet.Task
	return ret, nil
}

func (c Task) GetID(resource *ecs.Task) string {
	return resource.TaskId
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

func TaskSattled(t *ecs.Task) bool {
	return t.TaskStatus == TaskFinished || t.TaskStatus == TaskFailed
}
