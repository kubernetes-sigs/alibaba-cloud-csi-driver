package interfaces

import (
	sdk "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/golang/mock/gomock"
)

type MockNasClientV2Interface struct {
	client *MockNasV2Interface
}

func NewMockNasClientV2Interface(ctrl *gomock.Controller) *MockNasClientV2Interface {
	return &MockNasClientV2Interface{client: NewMockNasV2Interface(ctrl)}
}

func (n *MockNasClientV2Interface) CreateDir(req *sdk.CreateDirRequest) error {
	_, err := n.client.CreateDir(req)
	return err
}

func (n *MockNasClientV2Interface) SetDirQuota(req *sdk.SetDirQuotaRequest) error {
	_, err := n.client.SetDirQuota(req)
	return err
}

func (n *MockNasClientV2Interface) CancelDirQuota(req *sdk.CancelDirQuotaRequest) error {
	_, err := n.client.CancelDirQuota(req)
	return err
}

func (n *MockNasClientV2Interface) GetRecycleBinAttribute(filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	return n.client.GetRecycleBinAttribute(&sdk.GetRecycleBinAttributeRequest{
		FileSystemId: &filesystemId,
	})
}

func (n *MockNasClientV2Interface) CreateAccesspoint(req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	return n.client.CreateAccessPoint(req)
}

func (n *MockNasClientV2Interface) DeleteAccesspoint(filesystemId, accessPointId string) error {
	_, err := n.client.DeleteAccessPoint(&sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
	return err
}

func (n *MockNasClientV2Interface) DescribeAccesspoint(filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	return n.client.DescribeAccessPoint(&sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
}
