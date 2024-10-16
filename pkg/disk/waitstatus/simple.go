package waitstatus

import (
	"context"
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"k8s.io/utils/clock"
)

type Simple struct {
	client ECSDescribeDisks
	clk    clock.WithTicker

	ClientFactory func(context.Context) (ECSDescribeDisks, error)
}

func NewSimple(client ECSDescribeDisks, clk clock.WithTicker) *Simple {
	return &Simple{
		client: client,
		clk:    clk,
	}
}

func (w *Simple) WaitForDisk(ctx context.Context, diskID string, pred DiskStatusPredicate) (*ecs.Disk, error) {
	client := w.client
	if w.ClientFactory != nil {
		var err error
		client, err = w.ClientFactory(ctx)
		if err != nil {
			return nil, err
		}
	}
	ticker := w.clk.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		req := ecs.CreateDescribeDisksRequest()
		req.DiskIds = "[\"" + diskID + "\"]"
		req.PageSize = requests.NewInteger(batchSize)
		resp, err := client.DescribeDisks(req)
		if err != nil {
			return nil, err
		}
		if len(resp.Disks.Disk) == 0 {
			return nil, nil
		}
		if len(resp.Disks.Disk) > 1 {
			return nil, errors.New("too many disks")
		}
		disk := &resp.Disks.Disk[0]
		if pred(disk) {
			return disk, nil
		}
		select {
		case <-ticker.C():
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
