package ossfs

import (
	"context"
	"fmt"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func Test_buildPodSpec_dnsPolicy(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &fpm.AuthConfig{}
	ptCfg := &fpm.PodTemplateConfig{}
	authCfg.AuthType = ossfpm.AuthTypeSTS
	fakeOssfs := &fuseOssfs{}
	spec, err := fakeOssfs.buildPodSpec(&fpm.FusePodContext{
		Context:           context.Background(),
		Namespace:         mounterutils.LegacyFusePodNamespace,
		NodeName:          nodeName,
		VolumeId:          volumeId,
		AuthConfig:        authCfg,
		PodTemplateConfig: ptCfg,
		FuseType:          ossfpm.OssFsType,
	}, "target")

	require.NoError(t, err)
	assert.Equal(t, corev1.DNSPolicy(""), spec.DNSPolicy)

	ptCfg = &fpm.PodTemplateConfig{
		DnsPolicy: corev1.DNSPolicy("ClusterFirstWithHostNet"),
	}
	spec, err = fakeOssfs.buildPodSpec(&fpm.FusePodContext{
		Context:           context.Background(),
		Namespace:         mounterutils.LegacyFusePodNamespace,
		NodeName:          nodeName,
		VolumeId:          volumeId,
		AuthConfig:        authCfg,
		PodTemplateConfig: ptCfg,
		FuseType:          ossfpm.OssFsType,
	}, "target")

	require.NoError(t, err)
	assert.Equal(t, corev1.DNSClusterFirstWithHostNet, spec.DNSPolicy)
}

func Test_buildAuthSpec_ossfs(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &fpm.AuthConfig{}
	container := corev1.Container{
		Name:  "fuse-mounter",
		Image: "test-image",
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "test-mounts",
				MountPath: "target",
			},
		},
	}
	targetVolume := corev1.Volume{
		Name: "test-mounts",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "target",
				Type: ptr.To(corev1.HostPathDirectoryOrCreate),
			},
		},
	}
	spec := corev1.PodSpec{}
	spec.Volumes = []corev1.Volume{targetVolume}
	rrsaCfg := fpm.RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	authCfg.AuthType = ossfpm.AuthTypeRRSA
	fakeOssfs := &fuseOssfs{}
	fakeOssfs.buildAuthSpec(&fpm.FusePodContext{
		Context:    context.Background(),
		Namespace:  mounterutils.LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
		FuseType:   ossfpm.OssFsType,
	}, "target", &spec, &container)

	assert.Equal(t, "rrsa-oidc-token", spec.Volumes[len(spec.Volumes)-1].Name)
	volumeMount := container.VolumeMounts[len(container.VolumeMounts)-1]
	assert.Contains(t, "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens", volumeMount.MountPath)
	assert.Contains(t, "rrsa-oidc-token", volumeMount.Name)
}

func TestPrecheckAuthConfig_ossfs(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs(nil, fakeMeta)
	tests := []struct {
		name    string
		opts    *ossfpm.Options
		wantErr bool
	}{
		{
			name: "empty aksk",
			opts: &ossfpm.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with accessKey",
			opts: &ossfpm.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
				FuseType: ossfpm.OssFsType,
			},
			wantErr: false,
		},
		{
			name: "success with secretRef",
			opts: &ossfpm.Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				FuseType:  ossfpm.OssFsType,
			},
			wantErr: false,
		},
		{
			name: "conflict between accessKey and secretRef",
			opts: &ossfpm.Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				AkID:      "11111",
				FuseType:  ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "invalid secretRef",
			opts: &ossfpm.Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: mounterutils.GetCredientialsSecretName(ossfpm.OssFsType),
				FuseType:  ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "use assumeRole with non-RRSA authType",
			opts: &ossfpm.Options{
				URL:           "1.1.1.1",
				Bucket:        "aliyun",
				Path:          "/path",
				SecretRef:     "secret",
				AkID:          "11111",
				AssumeRoleArn: "test-assume-role-arn",
				FuseType:      ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "empty roleName, ARNs",
			opts: &ossfpm.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: ossfpm.AuthTypeRRSA,
				FuseType: ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "empty roleName, ARN",
			opts: &ossfpm.Options{
				URL:             "1.1.1.1",
				Bucket:          "aliyun",
				Path:            "/path",
				AuthType:        ossfpm.AuthTypeRRSA,
				OidcProviderArn: "test-oidc-provider-arn",
				FuseType:        ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with csi-secret-store",
			opts: &ossfpm.Options{
				URL:                 "1.1.1.1",
				Bucket:              "aliyun",
				Path:                "/path",
				AuthType:            ossfpm.AuthTypeCSS,
				SecretProviderClass: "test-secret-provider-class",
				FuseType:            ossfpm.OssFsType,
			},
			wantErr: false,
		},
		{
			name: "empty secretProviderClass",
			opts: &ossfpm.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: ossfpm.AuthTypeCSS,
				FuseType: ossfpm.OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with public authType",
			opts: &ossfpm.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: ossfpm.AuthTypePublic,
				FuseType: ossfpm.OssFsType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fakeOssfs.PrecheckAuthConfig(tt.opts, true)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestMakeAuthConfig_ossfs(t *testing.T) {
	t.Setenv("CLUSTER_ID", "cluster-id")
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "account-id")
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs(nil, fakeMeta)
	tests := []struct {
		name           string
		options        *ossfpm.Options
		expectedConfig *fpm.AuthConfig
		expectedError  error
	}{
		{
			name: "public",
			options: &ossfpm.Options{
				AuthType: ossfpm.AuthTypePublic,
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType: ossfpm.AuthTypePublic,
			},
			expectedError: nil,
		},
		{
			name: "AuthTypeRRSA_Success",
			options: &ossfpm.Options{
				AuthType:           ossfpm.AuthTypeRRSA,
				ServiceAccountName: "test-sa",
				RoleName:           "test-role",
				AssumeRoleArn:      "test-assume-role-arn",
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType: ossfpm.AuthTypeRRSA,
				RrsaConfig: &fpm.RrsaConfig{
					ServiceAccountName: "test-sa",
					RoleArn:            "acs:ram::account-id:role/test-role",
					OidcProviderArn:    "acs:ram::account-id:oidc-provider/ack-rrsa-cluster-id",
					AssumeRoleArn:      "test-assume-role-arn",
				},
			},
			expectedError: nil,
		},
		{
			name: "AuthTypeCSS",
			options: &ossfpm.Options{
				AuthType:            ossfpm.AuthTypeCSS,
				SecretProviderClass: "secret-provider-class",
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType:                ossfpm.AuthTypeCSS,
				SecretProviderClassName: "secret-provider-class",
			},
			expectedError: nil,
		},
		{
			name: "AuthTypeSTS",
			options: &ossfpm.Options{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "role-name",
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "role-name",
			},
			expectedError: nil,
		},
		{
			name: "OtherAuthType_SecretRef",
			options: &ossfpm.Options{
				AuthType:  "",
				SecretRef: "secret-ref",
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType:  "",
				SecretRef: "secret-ref",
			},
			expectedError: nil,
		},
		{
			name: "OtherAuthType_Secrets",
			options: &ossfpm.Options{
				AuthType: "",
				Bucket:   "bucket",
				AkID:     "ak-id",
				AkSecret: "ak-secret",
				FuseType: ossfpm.OssFsType,
			},
			expectedConfig: &fpm.AuthConfig{
				AuthType: "",
				Secrets: map[string]string{
					mounterutils.GetPasswdFileName(ossfpm.OssFsType): "bucket:ak-id:ak-secret",
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authConfig, err := fakeOssfs.MakeAuthConfig(tt.options, fakeMeta)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedConfig, authConfig)
			}
		})
	}
}

func TestMakeMountOptions_ossfs(t *testing.T) {
	tests := []struct {
		name          string
		opts          *ossfpm.Options
		region        string
		expected      []string
		expectedError bool
	}{
		{
			name: "Basic Options",
			opts: &ossfpm.Options{
				URL: "oss://bucket",
			},
			expected: []string{
				"url=oss://bucket",
			},
		},
		{
			name: "ReadOnly Option",
			opts: &ossfpm.Options{
				URL:      "oss://bucket",
				ReadOnly: true,
			},
			expected: []string{
				"url=oss://bucket",
				"ro",
			},
		},
		{
			name: "EncryptedTypeAes256",
			opts: &ossfpm.Options{
				URL:       "oss://bucket",
				Encrypted: ossfpm.EncryptedTypeAes256,
			},
			expected: []string{
				"url=oss://bucket",
				"use_sse",
			},
		},
		{
			name: "EncryptedTypeKms",
			opts: &ossfpm.Options{
				URL:       "oss://bucket",
				Encrypted: ossfpm.EncryptedTypeKms,
				KmsKeyId:  "1234",
			},
			expected: []string{
				"url=oss://bucket",
				"use_sse=kmsid:1234",
			},
		},
		{
			name: "Metrics Enabled, use rrsa",
			opts: &ossfpm.Options{
				URL:           "oss://bucket",
				MetricsTop:    "5",
				AuthType:      ossfpm.AuthTypeRRSA,
				AssumeRoleArn: "arn:acs:ram::123456789012:role/role-name",
			},
			region: "us-east-1",
			expected: []string{
				"url=oss://bucket",
				"use_metrics",
				"metrics_top=5",
				"rrsa_endpoint=https://sts-vpc.us-east-1.aliyuncs.com",
				"assume_role_arn=arn:acs:ram::123456789012:role/role-name",
			},
		},
		{
			name: "SigV4",
			opts: &ossfpm.Options{
				URL:        "oss://bucket",
				SigVersion: ossfpm.SigV4,
			},
			region: "us-east-1",
			expected: []string{
				"url=oss://bucket",
				"sigv4",
				"region=us-east-1",
			},
			expectedError: false,
		},
		{
			name: "SigV4 Without Region",
			opts: &ossfpm.Options{
				URL:        "oss://bucket",
				SigVersion: ossfpm.SigV4,
			},
			expectedError: true,
		},
		{
			name: "AuthTypeCSS",
			opts: &ossfpm.Options{
				URL:      "oss://bucket",
				AuthType: ossfpm.AuthTypeCSS,
			},
			expected: []string{
				"url=oss://bucket",
				"secret_store_dir=/etc/ossfs/secrets-store",
			},
		},
		{
			name: "AuthTypeSTS",
			opts: &ossfpm.Options{
				URL:      "oss://bucket",
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "role-name",
			},
			expected: []string{
				"url=oss://bucket",
				"ram_role=role-name",
			},
		},
		{
			name: "DefaultAuthType",
			opts: &ossfpm.Options{
				URL:       "oss://bucket",
				SecretRef: "secret",
				FuseType:  "ossfs",
			},
			expected: []string{
				"url=oss://bucket",
				"passwd_file=/etc/ossfs/passwd-ossfs",
				"use_session_token",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("REGION_ID", tt.region)
			fakeMeta := metadata.NewMetadata()
			fakeOssfs := NewFuseOssfs(nil, fakeMeta)
			mountOptions, err := fakeOssfs.MakeMountOptions(tt.opts, fakeMeta)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.ElementsMatch(t, tt.expected, mountOptions)
		})
	}
}

func TestGetAuthOpttions_ossfs(t *testing.T) {
	tests := []struct {
		name        string
		region      string
		opts        *ossfpm.Options
		wantOptions []string
	}{
		{
			name: "public",
			opts: &ossfpm.Options{
				AuthType: ossfpm.AuthTypePublic,
			},
			wantOptions: []string{"public_bucket=1"},
		},
		{
			name:   "rrsa",
			region: "us-east-1",
			opts: &ossfpm.Options{
				AuthType:      ossfpm.AuthTypeRRSA,
				AssumeRoleArn: "acs:ram::account-id:role/test-role",
				ExternalId:    "test-external-id",
			},
			wantOptions: []string{
				"rrsa_endpoint=https://sts-vpc.us-east-1.aliyuncs.com",
				"assume_role_arn=acs:ram::account-id:role/test-role",
				"assume_role_external_id=test-external-id",
			},
		},
		{
			name: "rrsa - empty region",
			opts: &ossfpm.Options{
				AuthType:   ossfpm.AuthTypeRRSA,
				ExternalId: "test-external-id",
			},
			wantOptions: []string{
				"rrsa_endpoint=https://sts.aliyuncs.com",
			},
		},
		{
			name: "css",
			opts: &ossfpm.Options{
				AuthType: ossfpm.AuthTypeCSS,
			},
			wantOptions: []string{
				"secret_store_dir=/etc/ossfs/secrets-store",
			},
		},
		{
			name: "sts",
			opts: &ossfpm.Options{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "test-role",
			},
			wantOptions: []string{
				"ram_role=test-role",
			},
		},
		{
			name: "secretref",
			opts: &ossfpm.Options{
				SecretRef: "test-secret",
				FuseType:  "ossfs",
			},
			wantOptions: []string{
				"passwd_file=/etc/ossfs/passwd-ossfs",
				"use_session_token",
			},
		},
		{
			name: "aksk",
			opts: &ossfpm.Options{
				FuseType: "ossfs",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOssfs := &fuseOssfs{}
			opts := fakeOssfs.getAuthOptions(tt.opts, tt.region)
			assert.Equal(t, tt.wantOptions, opts)
		})
	}
}

func TestAddDefaultMountOptions_ossfs(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeInter := NewFuseOssfs(nil, fakeMeta)
	fakeOssfs, ok := fakeInter.(*fuseOssfs)
	if !ok {
		t.Fatalf("failed to cast to fuseOssfs")
	}
	tests := []struct {
		name        string
		options     []string
		cfglevel    string
		enabledMime bool
		want        []string
	}{
		{
			name:    "empty option, empty config",
			options: []string{"others"},
			want:    []string{"others", "dbglevel=warn", "allow_other", "use_metrics", "listobjectsv2"},
		},
		{
			name:    "set option",
			options: []string{"others", "dbglevel=debug", "others"},
			want:    []string{"others", "dbglevel=debug", "others", "allow_other", "use_metrics", "listobjectsv2"},
		},
		{
			name:     "set option, set config",
			cfglevel: "info",
			options:  []string{"others", "dbglevel=debug", "others"},
			want:     []string{"others", "dbglevel=debug", "others", "allow_other", "use_metrics", "listobjectsv2"},
		},
		{
			name:     "empty option, set config",
			cfglevel: "debug",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=debug", "allow_other", "use_metrics", "listobjectsv2"},
		},
		{
			name:     "empty option, invalid config",
			cfglevel: "invalid",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=warn", "allow_other", "use_metrics", "listobjectsv2"},
		},
		{
			name:        "mime-support=true",
			enabledMime: true,
			options:     []string{"others"},
			want:        []string{"others", "dbglevel=warn", "allow_other", "use_metrics", "mime=" + OssfsCsiMimeTypesFilePath, "listobjectsv2"},
		},
		{
			name:    "listobjectsv2 has set",
			options: []string{"others", "listobjectsv2"},
			want:    []string{"others", "listobjectsv2", "dbglevel=warn", "allow_other", "use_metrics"},
		},
		{
			name:    "allow other",
			options: []string{"others", "allow_other"},
			want:    []string{"others", "allow_other", "dbglevel=warn", "use_metrics", "listobjectsv2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOssfs.config.Dbglevel = tt.cfglevel
			fakeOssfs.config.Extra = map[string]string{
				"mime-support": fmt.Sprintf("%t", tt.enabledMime),
			}
			got := fakeOssfs.AddDefaultMountOptions(tt.options)
			assert.Equal(t, tt.want, got)
		})
	}
}
