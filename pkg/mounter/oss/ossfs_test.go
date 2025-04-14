package oss

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func Test_buildOssfsAuthSpec(t *testing.T) {
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
	buildOssfsAuthSpec(&mounterutils.FusePodContext{
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
			secret := getPasswdSecretVolume(tt.secretRef)

			assert.Equal(t, tt.expectedEmpty, secret == nil)
			if secret != nil {
				assert.Equal(t, tt.expectedName, secret.SecretName)
				assert.Equal(t, tt.expectedItems, secret.Items)
			}
		})
	}
}

func Test_AddDefaultMountOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []string
		dbglevel string
		mime     string
		want     []string
	}{
		{
			name:     "Debug level not set",
			options:  []string{},
			dbglevel: "",
			mime:     "false",
			want:     []string{"dbglevel=err"},
		},
		{
			name:     "Debug level set by config",
			options:  []string{},
			dbglevel: "warn",
			mime:     "false",
			want:     []string{"dbglevel=warn"},
		},
		{
			name:     "Debug level set by mount options",
			options:  []string{"dbglevel=info"},
			dbglevel: "warn",
			mime:     "false",
			want:     []string{"dbglevel=info"},
		},
		{
			name:     "Invalid debug level",
			options:  []string{},
			dbglevel: "unknown",
			mime:     "false",
			want:     []string{"dbglevel=err"},
		},
		{
			name:     "Mime support enabled without existing mime.types",
			options:  []string{},
			dbglevel: "",
			mime:     "true",
			want:     []string{"dbglevel=err", "mime=" + OssfsCsiMimeTypesFilePath},
		},
		{
			name:     "Mime support disabled with existing mime.types",
			options:  []string{},
			dbglevel: "",
			mime:     "false",
			want:     []string{"dbglevel=err"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fuseOssfs{
				config: mounterutils.FuseContainerConfig{
					Dbglevel: tt.dbglevel,
					Extra: map[string]string{
						"mime-support": tt.mime,
					},
				},
			}
			got := f.AddDefaultMountOptions(tt.options)
			if len(got) != len(tt.want) {
				t.Errorf("AddDefaultMountOptions() got = %v, want %v", got, tt.want)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDefaultMountOptions() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
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
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
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
				AkID:      "11111",
				FuseType:  OssFsType,
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
				URL:           "1.1.1.1",
				Bucket:        "aliyun",
				Path:          "/path",
				SecretRef:     "secret",
				AkID:          "11111",
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
				AkID:     "ak-id",
				AkSecret: "ak-secret",
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

func TestGetAuthOpttions(t *testing.T) {
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := tt.opts.getAuthOptions(tt.region)
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
			want:    []string{"others", "dbglevel=err"},
		},
		{
			name:    "set option",
			options: []string{"others", "dbglevel=debug", "others"},
			want:    []string{"others", "dbglevel=debug", "others"},
		},
		{
			name:     "set option, set config",
			cfglevel: "info",
			options:  []string{"others", "dbglevel=debug", "others"},
			want:     []string{"others", "dbglevel=debug", "others"},
		},
		{
			name:     "empty option, set config",
			cfglevel: "debug",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=debug"},
		},
		{
			name:     "empty option, invalid config",
			cfglevel: "invalid",
			options:  []string{"others"},
			want:     []string{"others", "dbglevel=err"},
		},
		{
			name:        "mime-support=true",
			enabledMime: true,
			options:     []string{"others"},
			want:        []string{"others", "dbglevel=err", "mime=" + OssfsCsiMimeTypesFilePath},
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
