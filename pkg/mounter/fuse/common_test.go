package fuse

import (
	"fmt"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/component-base/featuregate"
)

// MockMetadataProvider 模拟 metadata.MetadataProvider 接口
type MockMetadataProvider struct {
	mock.Mock
}

func (m *MockMetadataProvider) Get(key metadata.MetadataKey) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func TestSetDefaultImage(t *testing.T) {
	tests := []struct {
		name          string
		fuseType      string
		config        *mounterutils.FuseContainerConfig
		metadata      metadata.MetadataProvider
		featureGate   *featuregate.FeatureGate
		expectedImage string
	}{
		{
			name:     "ImageAlreadySet",
			fuseType: OssFsType,
			config: &mounterutils.FuseContainerConfig{
				Image: "custom-image",
			},
			expectedImage: "custom-image",
		},
		{
			name:     "RegistryAndRegionSet",
			fuseType: OssFsType,
			config:   &mounterutils.FuseContainerConfig{},
			metadata: &MockMetadataProvider{
				mock.Mock{
					ExpectedCalls: []*mock.Call{
						{Method: "Get", Arguments: mock.Arguments{metadata.RegistryURL}, ReturnArguments: mock.Arguments{"custom-registry", nil}},
						{Method: "Get", Arguments: mock.Arguments{metadata.RegionID}, ReturnArguments: mock.Arguments{"", fmt.Errorf("region not found")}},
					},
				},
			},
			expectedImage: fmt.Sprintf("custom-registry/acs/csi-ossfs:%s", defaultOssfsUpdatedImageTag),
		},
		{
			name:     "RegionSet",
			fuseType: OssFsType,
			config:   &mounterutils.FuseContainerConfig{},
			metadata: &MockMetadataProvider{
				mock.Mock{
					ExpectedCalls: []*mock.Call{
						{Method: "Get", Arguments: mock.Arguments{metadata.RegistryURL}, ReturnArguments: mock.Arguments{"", nil}},
						{Method: "Get", Arguments: mock.Arguments{metadata.RegionID}, ReturnArguments: mock.Arguments{"cn-hangzhou", nil}},
					},
				},
			},
			expectedImage: fmt.Sprintf("registry-cn-hangzhou-vpc.ack.aliyuncs.com/acs/csi-ossfs:%s", defaultOssfsUpdatedImageTag),
		},
		{
			name:     "RegionNotSet",
			fuseType: OssFsType,
			config:   &mounterutils.FuseContainerConfig{},
			metadata: &MockMetadataProvider{
				mock.Mock{
					ExpectedCalls: []*mock.Call{
						{Method: "Get", Arguments: mock.Arguments{metadata.RegistryURL}, ReturnArguments: mock.Arguments{"", nil}},
						{Method: "Get", Arguments: mock.Arguments{metadata.RegionID}, ReturnArguments: mock.Arguments{"", fmt.Errorf("region not found")}},
					},
				},
			},
			expectedImage: fmt.Sprintf("%s/acs/csi-ossfs:%s", defaultRegistry, defaultOssfsUpdatedImageTag),
		},
		{
			name:     "UnknownFuseType",
			fuseType: "UnknownType",
			config:   &mounterutils.FuseContainerConfig{},
			metadata: &MockMetadataProvider{
				mock.Mock{
					ExpectedCalls: []*mock.Call{
						{Method: "Get", Arguments: mock.Arguments{metadata.RegistryURL}, ReturnArguments: mock.Arguments{"", nil}},
						{Method: "Get", Arguments: mock.Arguments{metadata.RegionID}, ReturnArguments: mock.Arguments{"", fmt.Errorf("region not found")}},
					},
				},
			},
			expectedImage: fmt.Sprintf("%s/acs/csi-UnknownType:latest", defaultRegistry),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var meta metadata.MetadataProvider
			if test.metadata != nil {
				meta = test.metadata
			}
			setDefaultImage(test.fuseType, meta, test.config)
			assert.Equal(t, test.expectedImage, test.config.Image)
		})
	}
}

func Test_checkRRSAParams(t *testing.T) {
	tests := []struct {
		name     string
		roleName string
		rc       *mounterutils.RrsaConfig
		wantErr  bool
	}{
		{
			"rolename",
			"test-role-name",
			nil,
			false,
		},
		{
			"arns",
			"",
			&mounterutils.RrsaConfig{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
		{
			"arn",
			"",
			&mounterutils.RrsaConfig{RoleArn: "test-role-arn"},
			true,
		},
		{
			"arn-and-rolename",
			"test-role-name",
			&mounterutils.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn"},
			true,
		},
		{
			"arns-and-rolename",
			"test-role-name",
			&mounterutils.RrsaConfig{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &mounterutils.AuthConfig{
				RoleName:   tt.roleName,
				RrsaConfig: tt.rc,
			}
			err := checkRRSAParams(ac)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
