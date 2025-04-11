package oss

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"k8s.io/component-base/featuregate"
)

func TestSetDefaultImage(t *testing.T) {
	tests := []struct {
		name          string
		fuseType      string
		config        *mounterutils.FuseContainerConfig
		region        string
		registryUrl   string
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
			name:          "RegistryAndRegionSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			registryUrl:   "custom-registry",
			expectedImage: fmt.Sprintf("custom-registry/acs/csi-ossfs:%s", defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegionSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			region:        "cn-hangzhou",
			expectedImage: fmt.Sprintf("registry-cn-hangzhou-vpc.ack.aliyuncs.com/acs/csi-ossfs:%s", defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegionNotSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			expectedImage: fmt.Sprintf("%s/acs/csi-ossfs:%s", defaultRegistry, defaultOssfsUpdatedImageTag),
		},
		{
			name:          "UnknownFuseType",
			fuseType:      "UnknownType",
			config:        &mounterutils.FuseContainerConfig{},
			expectedImage: fmt.Sprintf("%s/acs/csi-UnknownType:latest", defaultRegistry),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("REGION_ID", test.region)
			t.Setenv("DEFAULT_REGISTRY", test.registryUrl)
			fakeMate := metadata.NewMetadata()
			setDefaultImage(test.fuseType, fakeMate, test.config)
			assert.Equal(t, test.expectedImage, test.config.Image)
		})
	}
}

func Test_checkRRSAParams(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		wantErr bool
	}{
		{
			"rolename",
			&Options{RoleName: "test-role-name"},
			false,
		},
		{
			"arns",
			&Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
		{
			"arn",
			&Options{RoleArn: "test-role-arn"},
			true,
		},
		{
			"arn-and-rolename",
			&Options{RoleName: "test-role-name", OidcProviderArn: "test-oidc-provider-arn"},
			true,
		},
		{
			"arns-and-rolename",
			&Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkRRSAParams(tt.opts)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_getRRSAConfig(t *testing.T) {
	m := metadata.NewMetadata()
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "112233445566")
	t.Setenv("CLUSTER_ID", "c12345678")
	tests := []struct {
		name    string
		opt     Options
		wantCfg *utils.RrsaConfig
	}{
		{
			"rolename",
			Options{RoleName: "test-role-name"},
			&utils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: utils.FuseServiceAccountName},
		},
		{
			"specified-arns",
			Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&utils.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: utils.FuseServiceAccountName},
		},
		{
			"arns-first",
			Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&utils.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: utils.FuseServiceAccountName},
		},
		{
			"serviceaccount",
			Options{RoleName: "test-role-name", ServiceAccountName: "test-service-account"},
			&utils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: "test-service-account"},
		},
		{
			"assumeRoleArn",
			Options{RoleName: "test-role-name", AssumeRoleArn: "acs:ram::112233445566:role/assume-role-name"},
			&utils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: utils.FuseServiceAccountName, AssumeRoleArn: "acs:ram::112233445566:role/assume-role-name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := getRRSAConfig(&tt.opt, m)
			assert.Nil(t, err)
			if !reflect.DeepEqual(cfg, tt.wantCfg) {
				t.Errorf("getRRSAConfig() = %v, want %v", cfg, tt.wantCfg)
			}
		})
	}
}

func Test_getSTSEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		region      string
		envEndpoint string
		expected    string
	}{
		{
			name:        "region and env are both empty",
			region:      "",
			envEndpoint: "",
			expected:    "https://sts.aliyuncs.com",
		},
		{
			name:        "env is empty",
			region:      "cn-beijing",
			envEndpoint: "",
			expected:    "https://sts-vpc.cn-beijing.aliyuncs.com",
		},
		{
			name:        "With STS_ENDPOINT environment variable",
			region:      "cn-hangzhou",
			envEndpoint: "sts-vpc.cn-beijing.aliyuncs.com",
			expected:    "sts-vpc.cn-beijing.aliyuncs.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envEndpoint != "" {
				t.Setenv("STS_ENDPOINT", tt.envEndpoint)
			}

			endpoint := getSTSEndpoint(tt.region)
			if endpoint != tt.expected {
				t.Errorf("Expected endpoint to be %s, got %s", tt.expected, endpoint)
			}
		})
	}
}
