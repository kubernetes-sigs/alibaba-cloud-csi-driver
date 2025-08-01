package oss

import (
	"context"
	"fmt"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func Test_buildAuthSpec_ossfs(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &mounterutils.AuthConfig{}
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
	rrsaCfg := mounterutils.RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	authCfg.AuthType = AuthTypeRRSA
	fakeOssfs := &fuseOssfs{}
	fakeOssfs.buildAuthSpec(&mounterutils.FusePodContext{
		Context:    context.Background(),
		Namespace:  mounterutils.LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
		FuseType:   OssFsType,
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
		opts    *Options
		wantErr bool
	}{
		{
			name: "empty aksk",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with accessKey",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "akid",
					AkSecret: "aksecret",
				},
				FuseType: OssFsType,
			},
			wantErr: false,
		},
		{
			name: "success with secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				FuseType:  OssFsType,
			},
			wantErr: false,
		},
		{
			name: "conflict between accessKey and secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				AccessKey: AccessKey{
					AkID: "akid",
				},
				FuseType: OssFsType,
			},
			wantErr: true,
		},
		{
			name: "invalid secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: utils.GetCredientialsSecretName(OssFsType),
				FuseType:  OssFsType,
			},
			wantErr: true,
		},
		{
			name: "use assumeRole with non-RRSA authType",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				AccessKey: AccessKey{
					AkID: "akid",
				},
				AssumeRoleArn: "test-assume-role-arn",
				FuseType:      OssFsType,
			},
			wantErr: true,
		},
		{
			name: "empty roleName, ARNs",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: AuthTypeRRSA,
				FuseType: OssFsType,
			},
			wantErr: true,
		},
		{
			name: "empty roleName, ARN",
			opts: &Options{
				URL:             "1.1.1.1",
				Bucket:          "aliyun",
				Path:            "/path",
				AuthType:        AuthTypeRRSA,
				OidcProviderArn: "test-oidc-provider-arn",
				FuseType:        OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with csi-secret-store",
			opts: &Options{
				URL:                 "1.1.1.1",
				Bucket:              "aliyun",
				Path:                "/path",
				AuthType:            AuthTypeCSS,
				SecretProviderClass: "test-secret-provider-class",
				FuseType:            OssFsType,
			},
			wantErr: false,
		},
		{
			name: "empty secretProviderClass",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: AuthTypeCSS,
				FuseType: OssFsType,
			},
			wantErr: true,
		},
		{
			name: "success with public authType",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: AuthTypePublic,
				FuseType: OssFsType,
			},
			wantErr: false,
		},
		{
			name: "token republish",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				TokenSecret: TokenSecret{
					AccessKeyId:     "akid",
					AccessKeySecret: "aksecret",
					Expiration:      "expiration",
					SecurityToken:   "securitytoken",
				},
				FuseType: OssFsType,
			},
			wantErr: false,
		},
		{
			name: "conflicts token",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				TokenSecret: TokenSecret{
					AccessKeyId:     "akid",
					AccessKeySecret: "aksecret",
					Expiration:      "expiration",
					SecurityToken:   "securitytoken",
				},
				SecretRef: "non-empty",
				FuseType:  OssFsType,
			},
			wantErr: true,
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
		options        *Options
		expectedConfig *utils.AuthConfig
		expectedError  error
	}{
		{
			name: "public",
			options: &Options{
				AuthType: AuthTypePublic,
			},
			expectedConfig: &utils.AuthConfig{
				AuthType: AuthTypePublic,
			},
			expectedError: nil,
		},
		{
			name: "AuthTypeRRSA_Success",
			options: &Options{
				AuthType:           AuthTypeRRSA,
				ServiceAccountName: "test-sa",
				RoleName:           "test-role",
				AssumeRoleArn:      "test-assume-role-arn",
			},
			expectedConfig: &utils.AuthConfig{
				AuthType: AuthTypeRRSA,
				RrsaConfig: &utils.RrsaConfig{
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
			options: &Options{
				AuthType:            AuthTypeCSS,
				SecretProviderClass: "secret-provider-class",
			},
			expectedConfig: &utils.AuthConfig{
				AuthType:                AuthTypeCSS,
				SecretProviderClassName: "secret-provider-class",
			},
			expectedError: nil,
		},
		{
			name: "AuthTypeSTS",
			options: &Options{
				AuthType: AuthTypeSTS,
				RoleName: "role-name",
			},
			expectedConfig: &utils.AuthConfig{
				AuthType: AuthTypeSTS,
				RoleName: "role-name",
			},
			expectedError: nil,
		},
		{
			name: "OtherAuthType_SecretRef",
			options: &Options{
				AuthType:  "",
				SecretRef: "secret-ref",
			},
			expectedConfig: &utils.AuthConfig{
				AuthType:  "",
				SecretRef: "secret-ref",
			},
			expectedError: nil,
		},
		{
			name: "OtherAuthType_Secrets",
			options: &Options{
				AuthType: "",
				Bucket:   "bucket",
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
				FuseType: OssFsType,
			},
			expectedConfig: &utils.AuthConfig{
				AuthType: "",
				Secrets: map[string]string{
					utils.GetPasswdFileName(OssFsType): "bucket:ak-id:ak-secret",
				},
			},
			expectedError: nil,
		},
		{
			name: "OtherAuthType_TokenSecrets",
			options: &Options{
				AuthType: "",
				Bucket:   "bucket",
				TokenSecret: TokenSecret{
					AccessKeyId:     "ak-id",
					AccessKeySecret: "ak-secret",
					Expiration:      "expiration",
					SecurityToken:   "security-token",
				},
				FuseType: OssFsType,
			},
			expectedConfig: &utils.AuthConfig{
				AuthType: "",
				Secrets: map[string]string{
					utils.GetPasswdFileName(OssFsType) + "/" + KeyAccessKeyId:     "ak-id",
					utils.GetPasswdFileName(OssFsType) + "/" + KeyAccessKeySecret: "ak-secret",
					utils.GetPasswdFileName(OssFsType) + "/" + KeySecurityToken:   "security-token",
					utils.GetPasswdFileName(OssFsType) + "/" + KeyExpiration:      "expiration",
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
		opts          *Options
		region        string
		expected      []string
		expectedError bool
	}{
		{
			name: "Basic Options",
			opts: &Options{
				URL: "oss://bucket",
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"use_metrics",
			},
		},
		{
			name: "ReadOnly Option",
			opts: &Options{
				URL:      "oss://bucket",
				ReadOnly: true,
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"ro",
				"use_metrics",
			},
		},
		{
			name: "EncryptedTypeAes256",
			opts: &Options{
				URL:       "oss://bucket",
				Encrypted: EncryptedTypeAes256,
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"use_sse",
				"use_metrics",
			},
		},
		{
			name: "EncryptedTypeKms",
			opts: &Options{
				URL:       "oss://bucket",
				Encrypted: EncryptedTypeKms,
				KmsKeyId:  "1234",
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"use_sse=kmsid:1234",
				"use_metrics",
			},
		},
		{
			name: "Metrics Enabled, use rrsa",
			opts: &Options{
				URL:           "oss://bucket",
				MetricsTop:    "5",
				AuthType:      AuthTypeRRSA,
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
			opts: &Options{
				URL:        "oss://bucket",
				SigVersion: SigV4,
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			region: "us-east-1",
			expected: []string{
				"url=oss://bucket",
				"use_metrics",
				"sigv4",
				"region=us-east-1",
			},
			expectedError: false,
		},
		{
			name: "SigV4 Without Region",
			opts: &Options{
				URL:        "oss://bucket",
				SigVersion: SigV4,
			},
			expectedError: true,
		},
		{
			name: "AuthTypeCSS",
			opts: &Options{
				URL:      "oss://bucket",
				AuthType: AuthTypeCSS,
			},
			expected: []string{
				"url=oss://bucket",
				"secret_store_dir=/etc/ossfs/secrets-store",
				"use_metrics",
			},
		},
		{
			name: "AuthTypeSTS",
			opts: &Options{
				URL:      "oss://bucket",
				AuthType: AuthTypeSTS,
				RoleName: "role-name",
			},
			expected: []string{
				"url=oss://bucket",
				"ram_role=role-name",
				"use_metrics",
			},
		},
		{
			name: "DefaultAuthType",
			opts: &Options{
				URL:      "oss://bucket",
				FuseType: "ossfs",
				AccessKey: AccessKey{
					AkID:     "ak-id",
					AkSecret: "ak-secret",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"use_metrics",
			},
		},
		{
			name: "SecretRef-AuthType",
			opts: &Options{
				URL:       "oss://bucket",
				SecretRef: "secret",
				FuseType:  "ossfs",
			},
			expected: []string{
				"url=oss://bucket",
				"passwd_file=/etc/ossfs/passwd-ossfs",
				"use_session_token",
				"use_metrics",
			},
		},
		{
			name: "TokenSecret-AuthType",
			opts: &Options{
				URL:      "oss://bucket",
				FuseType: "ossfs",
				TokenSecret: TokenSecret{
					AccessKeyId: "akid",
				},
			},
			expected: []string{
				"url=oss://bucket",
				"use_session_token",
				"use_metrics",
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
		opts        *Options
		wantOptions []string
	}{
		{
			name: "public",
			opts: &Options{
				AuthType: AuthTypePublic,
			},
			wantOptions: []string{"public_bucket=1"},
		},
		{
			name:   "rrsa",
			region: "us-east-1",
			opts: &Options{
				AuthType:      AuthTypeRRSA,
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
			opts: &Options{
				AuthType:   AuthTypeRRSA,
				ExternalId: "test-external-id",
			},
			wantOptions: []string{
				"rrsa_endpoint=https://sts.aliyuncs.com",
			},
		},
		{
			name: "css",
			opts: &Options{
				AuthType: AuthTypeCSS,
			},
			wantOptions: []string{
				"secret_store_dir=/etc/ossfs/secrets-store",
			},
		},
		{
			name: "sts",
			opts: &Options{
				AuthType: AuthTypeSTS,
				RoleName: "test-role",
			},
			wantOptions: []string{
				"ram_role=test-role",
			},
		},
		{
			name: "secretref",
			opts: &Options{
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
			opts: &Options{
				FuseType: "ossfs",
				AccessKey: AccessKey{
					AkID:     "test-akid",
					AkSecret: "test-aksecret",
				},
			},
		},
		{
			name: "token",
			opts: &Options{
				FuseType: "ossfs",
				TokenSecret: TokenSecret{
					AccessKeyId:     "test-akid",
					AccessKeySecret: "test-aksecret",
					Expiration:      "2023-01-01T00:00:00Z",
					SecurityToken:   "test-token",
				},
			},
			wantOptions: []string{
				"use_session_token",
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
			want:    []string{"others", "dbglevel=warn", "allow_other", "listobjectsv2"},
		},
		{
			name:    "set option",
			options: []string{"others", "dbglevel=debug", "others"},
			want:    []string{"others", "dbglevel=debug", "others", "allow_other", "listobjectsv2"},
		},
		{
			name:     "set option, set config",
			cfglevel: "info",
			options:  []string{"others", "dbglevel=debug", "others"},
			want:     []string{"others", "dbglevel=debug", "others", "allow_other", "listobjectsv2"},
		},
		{
			name:     "empty option, set config",
			cfglevel: "debug",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=debug", "allow_other", "listobjectsv2"},
		},
		{
			name:     "empty option, invalid config",
			cfglevel: "invalid",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=warn", "allow_other", "listobjectsv2"},
		},
		{
			name:        "mime-support=true",
			enabledMime: true,
			options:     []string{"others"},
			want:        []string{"others", "dbglevel=warn", "allow_other", "mime=" + OssfsCsiMimeTypesFilePath, "listobjectsv2"},
		},
		{
			name:    "listobjectsv2 has set",
			options: []string{"others", "listobjectsv2"},
			want:    []string{"others", "listobjectsv2", "dbglevel=warn", "allow_other"},
		},
		{
			name:    "allow other",
			options: []string{"others", "allow_other"},
			want:    []string{"others", "allow_other", "dbglevel=warn", "listobjectsv2"},
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
