package oss

import (
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	k8sver "k8s.io/apimachinery/pkg/util/version"
)

func TestGetFuseMounterPath(t *testing.T) {
	// Test registering a path
	testType := "test-fuse-type"
	testPath := "/usr/local/bin/test-fuse"

	// Register a test path
	RegisterFuseMounterPath(testType, testPath)

	// Test getting the path
	path, ok := GetFuseMounterPath(testType)
	assert.True(t, ok, "should find registered path")
	assert.Equal(t, testPath, path)

	// Test unknown type
	_, ok = GetFuseMounterPath("unknown-type")
	assert.False(t, ok, "should not find unregistered path")
}

func TestGetAllFuseMounterPaths(t *testing.T) {
	// Register a test path
	testType := "test-fuse-type-2"
	testPath := "/usr/local/bin/test-fuse-2"
	RegisterFuseMounterPath(testType, testPath)

	paths := GetAllFuseMounterPaths()
	assert.GreaterOrEqual(t, len(paths), 1, "Should have at least 1 registered path")
	assert.Equal(t, testPath, paths[testType], "Should contain registered path")
}

func TestGetFuseMounterInterceptors(t *testing.T) {
	testType := "test-fuse-type-1"
	testInterceptors := []mounter.MountInterceptor{interceptors.OssfsSecretInterceptor}

	RegisterFuseInterceptors(testType, testInterceptors)

	interceptors, ok := GetFuseMountInterceptors(testType)
	assert.True(t, ok, "should find registered interceptors")
	assert.Equal(t, interceptors, testInterceptors)
}

func TestGetAllRegisteredFuseTypes(t *testing.T) {
	// Register a test factory
	testType := "test-fuse-type-3"
	testFactory := func(utils.Config, metadata.MetadataProvider) OSSFuseMounterType {
		return nil
	}
	RegisterFuseMounter(testType, testFactory)

	types := GetAllRegisteredFuseTypes()
	assert.GreaterOrEqual(t, len(types), 1, "Should have at least 1 registered type")

	// Check that test type is in the list
	typeMap := make(map[string]bool)
	for _, typ := range types {
		typeMap[typ] = true
	}
	assert.True(t, typeMap[testType], "test type should be registered")
}

func TestShouldConstrainResourceVersion(t *testing.T) {
	tests := []struct {
		name           string
		k8sVersion     *k8sver.Version
		envValue       string
		expectedResult bool
	}{
		{
			name:           "K8s 1.31 should not constrain RV",
			k8sVersion:     k8sver.MajorMinor(1, 31),
			expectedResult: false,
		},
		{
			name:           "K8s 1.32 should not constrain RV",
			k8sVersion:     k8sver.MajorMinor(1, 32),
			expectedResult: false,
		},
		{
			name:           "K8s 1.30 should constrain RV",
			k8sVersion:     k8sver.MajorMinor(1, 30),
			expectedResult: true,
		},
		{
			name:           "K8s 1.28 should constrain RV",
			k8sVersion:     k8sver.MajorMinor(1, 28),
			expectedResult: true,
		},
		{
			name:           "nil version should constrain RV",
			k8sVersion:     nil,
			expectedResult: true,
		},
		{
			name:           "ENV=true should override K8s 1.31",
			k8sVersion:     k8sver.MajorMinor(1, 31),
			envValue:       "true",
			expectedResult: true,
		},
		{
			name:           "ENV=false should override K8s 1.30",
			k8sVersion:     k8sver.MajorMinor(1, 30),
			envValue:       "false",
			expectedResult: false,
		},
		{
			name:           "ENV invalid value should be ignored",
			k8sVersion:     k8sver.MajorMinor(1, 30),
			envValue:       "invalid",
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set ENV if specified
			if tt.envValue != "" {
				t.Setenv("OSS_FUSE_CONSTRAIN_RV", tt.envValue)
			}

			result := ShouldConstrainResourceVersion(tt.k8sVersion)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestGetAllOSSFusePodManagers(t *testing.T) {
	fakeMeta := metadata.NewMetadata()

	// Register a test factory that returns a valid mounter
	testType := "test-fuse-type-4"
	testFactory := func(utils.Config, metadata.MetadataProvider) OSSFuseMounterType {
		return &testFuseMounter{name: testType}
	}
	RegisterFuseMounter(testType, testFactory)

	// Test with nil configmap and nil client (CSI agent mode)
	managers := GetAllOSSFusePodManagers(utils.Config{}, fakeMeta, nil, nil)

	// Should have at least the test manager
	assert.GreaterOrEqual(t, len(managers), 1, "Should have at least 1 fuse pod manager")

	// Check that manager is created for registered type
	assert.NotNil(t, managers[testType], "test manager should be created")

	// Verify manager type
	assert.IsType(t, &OSSFusePodManager{}, managers[testType])
}

// testFuseMounter is a minimal implementation for testing
type testFuseMounter struct {
	name string
}

func (t *testFuseMounter) Name() string {
	return t.name
}

func (t *testFuseMounter) PodTemplateSpec(c *fpm.FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	return nil, nil
}

func (t *testFuseMounter) AddDefaultMountOptions(options []string) []string {
	return options
}

func (t *testFuseMounter) PrecheckAuthConfig(o *Options, onNode bool) error {
	return nil
}

func (t *testFuseMounter) MakeAuthConfig(o *Options, m metadata.MetadataProvider) (*fpm.AuthConfig, error) {
	return nil, nil
}

func (t *testFuseMounter) MakeMountOptions(o *Options, m metadata.MetadataProvider) ([]string, error) {
	return nil, nil
}
