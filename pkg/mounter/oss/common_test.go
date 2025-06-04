package oss

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/component-base/featuregate"
)

func TestSetDefaultImage(t *testing.T) {
	tests := []struct {
		name             string
		fuseType         string
		config           *mounterutils.FuseContainerConfig
		region           string
		registryUrl      string
		repositoryPrefix string
		featureGate      *featuregate.FeatureGate
		expectedImage    string
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
			name:             "PrefixSet",
			fuseType:         OssFsType,
			config:           &mounterutils.FuseContainerConfig{},
			repositoryPrefix: "custom-registry/ns",
			expectedImage:    fmt.Sprintf("custom-registry/ns/csi-ossfs:%s", defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegistryAndRegionSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			registryUrl:   "custom-registry",
			expectedImage: fmt.Sprintf("custom-registry/%s/csi-ossfs:%s", utils.DefImageNamespace, defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegionSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			region:        "cn-hangzhou",
			expectedImage: fmt.Sprintf("registry-cn-hangzhou-vpc.ack.aliyuncs.com/%s/csi-ossfs:%s", utils.DefImageNamespace, defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegionNotSet",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			expectedImage: fmt.Sprintf("%s/%s/csi-ossfs:%s", utils.DefImageRegistry, utils.DefImageNamespace, defaultOssfsUpdatedImageTag),
		},
		{
			name:             "PrefixAndUrlConflict",
			fuseType:         OssFsType,
			config:           &mounterutils.FuseContainerConfig{},
			registryUrl:      "custom-registry",
			repositoryPrefix: "anotehr-registry/acs/",
			expectedImage:    fmt.Sprintf("anotehr-registry/%s/csi-ossfs:%s", utils.DefImageNamespace, defaultOssfsUpdatedImageTag),
		},
		{
			name:          "RegionAndUrlConflict",
			fuseType:      OssFsType,
			config:        &mounterutils.FuseContainerConfig{},
			region:        "cn-hangzhou",
			registryUrl:   "registry-cn-beijing-vpc.ack.aliyuncs.com/",
			expectedImage: fmt.Sprintf("registry-cn-beijing-vpc.ack.aliyuncs.com/%s/csi-ossfs:%s", utils.DefImageNamespace, defaultOssfsUpdatedImageTag),
		},
		{
			name:          "UnknownFuseType",
			fuseType:      "UnknownType",
			config:        &mounterutils.FuseContainerConfig{},
			expectedImage: fmt.Sprintf("%s/%s/csi-UnknownType:latest", utils.DefImageRegistry, utils.DefImageNamespace),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("REGION_ID", test.region)
			t.Setenv("DEFAULT_REGISTRY", test.registryUrl)
			t.Setenv("IMAGE_REPOSITORY_PREFIX", test.repositoryPrefix)
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
		wantCfg *mounterutils.RrsaConfig
	}{
		{
			"rolename",
			Options{RoleName: "test-role-name"},
			&mounterutils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: mounterutils.FuseServiceAccountName},
		},
		{
			"specified-arns",
			Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounterutils.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: mounterutils.FuseServiceAccountName},
		},
		{
			"arns-first",
			Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounterutils.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: mounterutils.FuseServiceAccountName},
		},
		{
			"serviceaccount",
			Options{RoleName: "test-role-name", ServiceAccountName: "test-service-account"},
			&mounterutils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: "test-service-account"},
		},
		{
			"assumeRoleArn",
			Options{RoleName: "test-role-name", AssumeRoleArn: "acs:ram::112233445566:role/assume-role-name"},
			&mounterutils.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: mounterutils.FuseServiceAccountName, AssumeRoleArn: "acs:ram::112233445566:role/assume-role-name"},
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

func Test_getPasswdSecretVolume(t *testing.T) {
	tests := []struct {
		name          string
		secretRef     string
		expectedEmpty bool
		expectedName  string
		expectedItems []corev1.KeyToPath
	}{
		{
			name:          "TestEmptySecretRef",
			secretRef:     "",
			expectedEmpty: true,
		},
		{
			name:          "TestNonEmptySecretRef",
			secretRef:     "my-secret",
			expectedEmpty: false,
			expectedName:  "my-secret",
			expectedItems: []corev1.KeyToPath{
				{
					Key:  "AccessKeyId",
					Path: "passwd-ossfs/AccessKeyId",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "AccessKeySecret",
					Path: "passwd-ossfs/AccessKeySecret",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "Expiration",
					Path: "passwd-ossfs/Expiration",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "SecurityToken",
					Path: "passwd-ossfs/SecurityToken",
					Mode: tea.Int32(0600),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := getPasswdSecretVolume(tt.secretRef, "ossfs")

			assert.Equal(t, tt.expectedEmpty, secret == nil)
			if secret != nil {
				assert.Equal(t, tt.expectedName, secret.SecretName)
				assert.Equal(t, tt.expectedItems, secret.Items)
			}
		})
	}
}
