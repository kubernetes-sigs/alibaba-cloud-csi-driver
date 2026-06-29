package interfaces

import (
	"context"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

type MockNasClientV2Interface struct {
	client *cloud.MockNasInterface
}

func NewMockNasClientV2Interface(ctrl *gomock.Controller) *MockNasClientV2Interface {
	return &MockNasClientV2Interface{client: cloud.NewMockNasInterface(ctrl)}
}

func (n *MockNasClientV2Interface) CreateDir(ctx context.Context, req *sdk.CreateDirRequest) error {
	_, err := n.client.CreateDirWithContext(ctx, req, nil)
	return err
}

func (n *MockNasClientV2Interface) SetDirQuota(ctx context.Context, req *sdk.SetDirQuotaRequest) error {
	_, err := n.client.SetDirQuotaWithContext(ctx, req, nil)
	return err
}

func (n *MockNasClientV2Interface) CancelDirQuota(ctx context.Context, req *sdk.CancelDirQuotaRequest) error {
	_, err := n.client.CancelDirQuotaWithContext(ctx, req, nil)
	return err
}

func (n *MockNasClientV2Interface) GetRecycleBinAttribute(ctx context.Context, filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	return n.client.GetRecycleBinAttributeWithContext(ctx, &sdk.GetRecycleBinAttributeRequest{
		FileSystemId: &filesystemId,
	}, nil)
}

func (n *MockNasClientV2Interface) CreateAccesspoint(ctx context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	return n.client.CreateAccessPointWithContext(ctx, req, nil)
}

func (n *MockNasClientV2Interface) DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error {
	_, err := n.client.DeleteAccessPointWithContext(ctx, &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}, nil)
	return err
}

func (n *MockNasClientV2Interface) DescribeAccesspoint(ctx context.Context, filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	return n.client.DescribeAccessPointWithContext(ctx, &sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}, nil)
}

func (n *MockNasClientV2Interface) DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error) {
	return n.client.DescribeFileSystemsWithContext(ctx, &sdk.DescribeFileSystemsRequest{
		FileSystemId: &filesystemID,
	}, nil)
}
