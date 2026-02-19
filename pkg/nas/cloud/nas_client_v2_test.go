package cloud

import (
	"testing"

	nas "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

const nasV2Region = "cn-hangzhou"

func TestNewNasClientV2(t *testing.T) {
	actual, err := NewNasClientV2(nasV2Region)
	assert.NotNil(t, actual)
	assert.NoError(t, err)
}

func TestCreateDirSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CreateDir(gomock.Any()).Return(
			&nas.CreateDirResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body:       &nas.CreateDirResponseBody{RequestId: new("")},
			}, nil)
	})
	err := client.CreateDir(&nas.CreateDirRequest{})
	assert.NoError(t, err)
}

func newNasClientV2ForTest(t *testing.T, mockExpects func(*interfaces.MockNasV2Interface)) *NasClientV2 {
	ctrl := gomock.NewController(t)
	mockNas := interfaces.NewMockNasV2Interface(ctrl)
	mockExpects(mockNas)
	return &NasClientV2{
		region:  nasV2Region,
		limiter: rate.NewLimiter(2, 10),
		client:  mockNas,
	}
}

func TestCreateDirError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CreateDir(gomock.Any()).Return(
			&nas.CreateDirResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(403),
				Body:       &nas.CreateDirResponseBody{RequestId: new("")},
			}, &tea.SDKError{
				Code:       new("InvalidProtocolType.NotSupported"),
				StatusCode: new(403),
				Message:    new("The specified protocol type does not supported."),
			})
	})
	err := client.CreateDir(&nas.CreateDirRequest{})
	assert.Error(t, err)
}

func TestSetDirQuotaSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().SetDirQuota(gomock.Any()).Return(
			&nas.SetDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body: &nas.SetDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(true),
				}}, nil)
	})
	err := client.SetDirQuota(&nas.SetDirQuotaRequest{})
	assert.NoError(t, err)
}

func TestSetDirQuotaFailure(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().SetDirQuota(gomock.Any()).Return(
			&nas.SetDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body: &nas.SetDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(false),
				}}, nil)
	})
	err := client.SetDirQuota(&nas.SetDirQuotaRequest{})
	assert.Error(t, err)
}

func TestSetDirQuotaError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().SetDirQuota(gomock.Any()).Return(
			&nas.SetDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body: &nas.SetDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(false),
				},
			}, &tea.SDKError{
				Code:       new("InvalidParameter.DirPathNotExist"),
				StatusCode: new(404),
				Message:    new("The Dir Path does not exist"),
			})
	})
	err := client.SetDirQuota(&nas.SetDirQuotaRequest{})
	assert.Error(t, err)
}

func TestCancelDirQuotaSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CancelDirQuota(gomock.Any()).Return(
			&nas.CancelDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body: &nas.CancelDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(true),
				},
			}, nil)
	})
	err := client.CancelDirQuota(&nas.CancelDirQuotaRequest{})
	assert.NoError(t, err)
}

func TestCancelDirQuotaFailure(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CancelDirQuota(gomock.Any()).Return(
			&nas.CancelDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body: &nas.CancelDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(false),
				},
			}, nil)
	})
	err := client.CancelDirQuota(&nas.CancelDirQuotaRequest{})
	assert.Error(t, err)
}

func TestCancelDirQuotaError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CancelDirQuota(gomock.Any()).Return(
			&nas.CancelDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body: &nas.CancelDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(false),
				},
			}, &tea.SDKError{
				Code:       new("InvalidParameter.DirPathNotExist"),
				StatusCode: new(404),
				Message:    new("The Dir Path does not exist"),
			})
	})
	err := client.CancelDirQuota(&nas.CancelDirQuotaRequest{})
	assert.Error(t, err)
}

func TestCancelDirQuotaIgnoreQuotaNotExistsError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CancelDirQuota(gomock.Any()).Return(
			&nas.CancelDirQuotaResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body: &nas.CancelDirQuotaResponseBody{
					RequestId: new(""),
					Success:   new(false),
				},
			}, &tea.SDKError{
				Code:       new("InvalidParameter.QuotaNotExistOnPath"),
				StatusCode: new(404),
				Message:    new("The specified path does not have quota."),
			})
	})
	err := client.CancelDirQuota(&nas.CancelDirQuotaRequest{})
	assert.NoError(t, err)
}

func TestGetRecycleBinAttributeSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().GetRecycleBinAttribute(gomock.Any()).Return(
			&nas.GetRecycleBinAttributeResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body:       &nas.GetRecycleBinAttributeResponseBody{RequestId: new("")},
			}, nil)
	})
	_, err := client.GetRecycleBinAttribute("")
	assert.NoError(t, err)
}

func TestGetRecycleBinAttributeError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().GetRecycleBinAttribute(gomock.Any()).Return(
			&nas.GetRecycleBinAttributeResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(400),
				Body:       &nas.GetRecycleBinAttributeResponseBody{RequestId: new("")},
			}, &tea.SDKError{
				Code:       new("MissingFileSystemId"),
				StatusCode: new(400),
				Message:    new("FileSystemId is mandatory for this action.\t"),
			})
	})
	_, err := client.GetRecycleBinAttribute("")
	assert.Error(t, err)
}

func TestCreateAccessPointSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CreateAccessPoint(gomock.Any()).Return(
			&nas.CreateAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body:       &nas.CreateAccessPointResponseBody{RequestId: new("")},
			}, nil)
	})
	_, err := client.CreateAccesspoint(&nas.CreateAccessPointRequest{})
	assert.NoError(t, err)
}

func TestCreateAccessPointError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().CreateAccessPoint(gomock.Any()).Return(
			&nas.CreateAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(403),
				Body:       &nas.CreateAccessPointResponseBody{RequestId: new("")},
			}, &tea.SDKError{
				Code:       new("OperationDenied.AccessPointCountsExceeded\n"),
				StatusCode: new(403),
				Message:    new("The maximum number of access point has reached its limits.\t"),
			})
	})
	_, err := client.CreateAccesspoint(&nas.CreateAccessPointRequest{})
	assert.Error(t, err)
}

func TestDeleteAccessPointSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().DeleteAccessPoint(gomock.Any()).Return(
			&nas.DeleteAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body:       &nas.DeleteAccessPointResponseBody{RequestId: new("")},
			}, nil)
	})
	err := client.DeleteAccesspoint("", "")
	assert.NoError(t, err)
}

func TestDeleteAccessPointError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().DeleteAccessPoint(gomock.Any()).Return(
			&nas.DeleteAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body:       &nas.DeleteAccessPointResponseBody{RequestId: new("")},
			}, &tea.SDKError{
				Code:       new("InvalidFileSystem.NotFound"),
				StatusCode: new(404),
				Message:    new("The specified file system does not exist."),
			})
	})
	err := client.DeleteAccesspoint("", "")
	assert.Error(t, err)
}

func TestDescribeAccessPointSuccess(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().DescribeAccessPoint(gomock.Any()).Return(
			&nas.DescribeAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(201),
				Body:       &nas.DescribeAccessPointResponseBody{RequestId: new("")},
			}, nil)
	})
	_, err := client.DescribeAccesspoint("", "")
	assert.NoError(t, err)
}

func TestDescribeAccessPointError(t *testing.T) {
	t.Parallel()
	client := newNasClientV2ForTest(t, func(mockNas *interfaces.MockNasV2Interface) {
		mockNas.EXPECT().DescribeAccessPoint(gomock.Any()).Return(
			&nas.DescribeAccessPointResponse{
				Headers:    make(map[string]*string),
				StatusCode: tea.Int32(404),
				Body:       &nas.DescribeAccessPointResponseBody{RequestId: new("")},
			}, &tea.SDKError{
				Code:       new("InvalidAccessPointId.NotFound"),
				StatusCode: new(404),
				Message:    new("The access point id does not exist."),
			})
	})
	_, err := client.DescribeAccesspoint("", "")
	assert.Error(t, err)
}

func TestIsAccessPointNotFoundErrorNil(t *testing.T) {
	actual := IsAccessPointNotFoundError(nil)
	assert.False(t, actual)
}

func TestIsAccessPointNotFoundErrorNotFound(t *testing.T) {
	t.Parallel()
	actual := IsAccessPointNotFoundError(&tea.SDKError{Code: new("NotFound")})
	assert.True(t, actual)
}
