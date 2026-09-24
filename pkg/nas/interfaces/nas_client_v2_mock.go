package interfaces

import (
	"context"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

type MockNasClientV2Interface struct {
	client *cloud.MockNasInterface
}

// Compile-time proof that the hand-written delegating mock stays in sync with the
// interface it stands in for: if NasClientV2Interface gains a method this mock does
// not implement, the package fails to build here instead of only at the call site.
var _ NasClientV2Interface = (*MockNasClientV2Interface)(nil)

func NewMockNasClientV2Interface(ctrl *gomock.Controller) *MockNasClientV2Interface {
	return &MockNasClientV2Interface{client: cloud.NewMockNasInterface(ctrl)}
}

func (n *MockNasClientV2Interface) CreateDir(ctx context.Context, req *sdk.CreateDirRequest) error {
	_, err := n.client.CreateDir(req)
	return err
}

func (n *MockNasClientV2Interface) SetDirQuota(ctx context.Context, req *sdk.SetDirQuotaRequest) error {
	_, err := n.client.SetDirQuota(req)
	return err
}

func (n *MockNasClientV2Interface) CancelDirQuota(ctx context.Context, req *sdk.CancelDirQuotaRequest) error {
	_, err := n.client.CancelDirQuota(req)
	return err
}

func (n *MockNasClientV2Interface) GetRecycleBinAttribute(ctx context.Context, filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	return n.client.GetRecycleBinAttribute(&sdk.GetRecycleBinAttributeRequest{
		FileSystemId: &filesystemId,
	})
}

func (n *MockNasClientV2Interface) CreateAccesspoint(ctx context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	return n.client.CreateAccessPoint(req)
}

func (n *MockNasClientV2Interface) DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error {
	_, err := n.client.DeleteAccessPointWithContext(ctx, &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}, &dara.RuntimeOptions{})
	return err
}

func (n *MockNasClientV2Interface) DescribeAccesspoint(ctx context.Context, filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	return n.client.DescribeAccessPoint(&sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
}

func (n *MockNasClientV2Interface) ListAccesspoints(ctx context.Context, req *sdk.ListAccessPointsRequest) (*sdk.ListAccessPointsResponse, error) {
	return n.client.ListAccessPoints(req)
}

func (n *MockNasClientV2Interface) DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error) {
	return n.client.DescribeFileSystems(&sdk.DescribeFileSystemsRequest{
		FileSystemId: &filesystemID,
	})
}

func (n *MockNasClientV2Interface) CreateAgenticSpace(ctx context.Context, req *sdk.CreateAgenticSpaceRequest) (*sdk.CreateAgenticSpaceResponse, error) {
	return n.client.CreateAgenticSpace(req)
}

func (n *MockNasClientV2Interface) GetAgenticSpace(ctx context.Context, req *sdk.GetAgenticSpaceRequest) (*sdk.GetAgenticSpaceResponse, error) {
	return n.client.GetAgenticSpace(req)
}

func (n *MockNasClientV2Interface) DescribeAgenticSpaces(ctx context.Context, req *sdk.DescribeAgenticSpacesRequest) (*sdk.DescribeAgenticSpacesResponse, error) {
	return n.client.DescribeAgenticSpacesWithContext(ctx, req, &dara.RuntimeOptions{})
}

func (n *MockNasClientV2Interface) DeleteAgenticSpace(ctx context.Context, req *sdk.DeleteAgenticSpaceRequest) (*sdk.DeleteAgenticSpaceResponse, error) {
	return n.client.DeleteAgenticSpace(req)
}

func (n *MockNasClientV2Interface) SetAgenticSpaceQuota(ctx context.Context, req *sdk.SetAgenticSpaceQuotaRequest) (*sdk.SetAgenticSpaceQuotaResponse, error) {
	return n.client.SetAgenticSpaceQuota(req)
}
