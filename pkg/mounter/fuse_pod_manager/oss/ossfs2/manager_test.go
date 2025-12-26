package ossfs2

import (
	"context"
	"fmt"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func TestPrecheckAuthConfig_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs(nil, fakeMeta)
	tests := []struct {
		name    string
		opts    *ossfpm.Options
		wantErr bool
	}{
		{
			"invalid authtype",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeCSS,
			},
			true,
		},
		{
			"valid sts",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "test",
			},
			false,
		},
		{
			"invalid sts",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeSTS,
			},
			true,
		},
		{
			"invalid rrsa",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeRRSA,
			},
			true,
		},
		{
			"valid rrsa",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeRRSA,
				RoleName: "test",
			},
			false,
		},
		{
			"use assumeRole with non-RRSA authType",
			&ossfpm.Options{
				AssumeRoleArn: "test-assume-role-arn",
			},
			true,
		},
		{
			"empty aksecret",
			&ossfpm.Options{
				AkID:     "test-ak",
				AkSecret: "",
			},
			true,
		},
		{
			"success - aksk",
			&ossfpm.Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			false,
		},
		{
			"success - secretref",
			&ossfpm.Options{
				SecretRef: "test-secret-ref",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fakeOssfs.PrecheckAuthConfig(tt.opts, true)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestMakeAuthConfig_ossfs2(t *testing.T) {
	t.Setenv("CLUSTER_ID", "cluster-id")
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "account-id")
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs(nil, fakeMeta)
	tests := []struct {
		name    string
		options *ossfpm.Options
		wantCfg *fpm.AuthConfig
		wantErr bool
	}{
		{
			"sts",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "test-role-name",
			},
			&fpm.AuthConfig{
				AuthType: ossfpm.AuthTypeSTS,
				RoleName: "test-role-name",
			},
			false,
		},
		{
			"aksk",
			&ossfpm.Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			&fpm.AuthConfig{
				Secrets: map[string]string{
					mounterutils.GetPasswdFileName(fakeOssfs.Name()): fmt.Sprintf("--oss_access_key_id=%s\n--oss_access_key_secret=%s", "test-ak", "test-ak-secret"),
				},
			},
			false,
		},
		{
			"secretref",
			&ossfpm.Options{
				SecretRef: "test-secretref",
			},
			&fpm.AuthConfig{
				SecretRef: "test-secretref",
			},
			false,
		},
		{
			"rrsa with rolename",
			&ossfpm.Options{
				AuthType: ossfpm.AuthTypeRRSA,
				RoleName: "test-role",
			},
			&fpm.AuthConfig{
				AuthType: ossfpm.AuthTypeRRSA,
				RrsaConfig: &fpm.RrsaConfig{
					ServiceAccountName: fpm.FuseServiceAccountName,
					RoleArn:            "acs:ram::account-id:role/test-role",
					OidcProviderArn:    "acs:ram::account-id:oidc-provider/ack-rrsa-cluster-id",
				},
			},
			false,
		},
		{
			"rrsa with arns",
			&ossfpm.Options{
				AuthType:           ossfpm.AuthTypeRRSA,
				ServiceAccountName: "test-sa",
				RoleArn:            "test-role-arn",
				OidcProviderArn:    "test-oidc-provider-arn",
				AssumeRoleArn:      "test-assume-role-arn",
			},
			&fpm.AuthConfig{
				AuthType: ossfpm.AuthTypeRRSA,
				RrsaConfig: &fpm.RrsaConfig{
					ServiceAccountName: "test-sa",
					RoleArn:            "test-role-arn",
					OidcProviderArn:    "test-oidc-provider-arn",
					AssumeRoleArn:      "test-assume-role-arn",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := fakeOssfs.MakeAuthConfig(tt.options, fakeMeta)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantCfg, gotCfg)
		})
	}
}

func TestMakeMountOptions_ossfs2(t *testing.T) {
	tests := []struct {
		name          string
		opts          *ossfpm.Options
		region        string
		expected      []string
		expectedError bool
	}{
		{
			name: "ro",
			opts: &ossfpm.Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
				Bucket:   "test-bucket",
				Path:     "/",
				URL:      "oss://test-bucket/",
				ReadOnly: true,
			},
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"ro=true",
			},
		},
		{
			name: "sigv4",
			opts: &ossfpm.Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: ossfpm.SigV4,
			},
			region: "test-region",
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"oss_region=test-region",
			},
		},
		{
			name: "sigv4 with empty region",
			opts: &ossfpm.Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: ossfpm.SigV4,
			},
			expectedError: true,
		},
		{
			name: "secretref",
			opts: &ossfpm.Options{
				SecretRef: "test-secretref",
				FuseType:  "ossfs2",
				Bucket:    "test-bucket",
				Path:      "/",
				URL:       "oss://test-bucket/",
			},
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"oss_sts_multi_conf_ak_file=/etc/ossfs2/passwd-ossfs2/AccessKeyId",
				"oss_sts_multi_conf_sk_file=/etc/ossfs2/passwd-ossfs2/AccessKeySecret",
				"oss_sts_multi_conf_token_file=/etc/ossfs2/passwd-ossfs2/SecurityToken",
			},
			expectedError: false,
		},
		{
			name: "metrics top",
			opts: &ossfpm.Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				MetricsTop: "5",
			},
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"use_metrics=true",
				"metrics_top=5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("REGION_ID", tt.region)
			fakeMeta := metadata.NewMetadata()
			fakeOssfs := NewFuseOssfs(nil, fakeMeta)
			opts, err := fakeOssfs.MakeMountOptions(tt.opts, fakeMeta)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.expected, opts)
		})
	}
}

func TestAddDefaultMountOptions_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeInter := NewFuseOssfs(nil, fakeMeta)
	fakeOssfs, ok := fakeInter.(*fuseOssfs)
	if !ok {
		t.Fatalf("failed to cast to fuseOssfs")
	}
	tests := []struct {
		name        string
		options     []string
		config      *fpm.FuseContainerConfig
		defaultOpts string
		want        []string
	}{
		{
			name:    "empty option, empty config",
			options: []string{"others"},
			want:    []string{"others", "log_level=info", "log_dir=/dev/stdout", "use_metrics=true"},
		},
		{
			name:    "set option",
			options: []string{"others", "log_level=debug", "log_dir=/tmp/ossfs2", "others", "use_metrics=false"},
			want:    []string{"others", "log_level=debug", "log_dir=/tmp/ossfs2", "others", "use_metrics=false"},
		},
		{
			name: "set option, set config",
			config: &fpm.FuseContainerConfig{
				Dbglevel: fpm.DebugLevelInfo,
			},
			options: []string{"others", "log_level=debug", "others"},
			want:    []string{"others", "log_level=debug", "others", "log_dir=/dev/stdout", "use_metrics=true"},
		},
		{
			name: "empty option, set config",
			config: &fpm.FuseContainerConfig{
				Dbglevel: fpm.DebugLevelDebug,
			},
			options: []string{"others"},
			want:    []string{"others", "log_level=debug", "log_dir=/dev/stdout", "use_metrics=true"},
		},
		{
			name: "empty option, invalid config",
			config: &fpm.FuseContainerConfig{
				Dbglevel: "invalid",
			},
			options: []string{"others"},
			want:    []string{"others", "log_level=info", "log_dir=/dev/stdout", "use_metrics=true"},
		},
		{
			name:        "default options",
			options:     nil,
			defaultOpts: "others,log_dir=/tmp/ossfs2",
			want:        []string{"others", "log_dir=/tmp/ossfs2", "log_level=info", "use_metrics=true"},
		},
		{
			name: "set metrics mode to disabled",
			config: &fpm.FuseContainerConfig{
				MetricsMode: fpm.MetricsModeDisabled,
			},
			want: []string{"log_level=info", "log_dir=/dev/stdout"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.defaultOpts != "" {
				t.Setenv("DEFAULT_OSSFS2_OPTIONS", tt.defaultOpts)
			}
			if tt.config != nil {
				fakeOssfs.config = *tt.config
			}
			got := fakeOssfs.AddDefaultMountOptions(tt.options)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetAuthOpttions_ossfs2(t *testing.T) {
	tests := []struct {
		name        string
		opts        *ossfpm.Options
		wantOptions []string
	}{
		{
			name: "secretref",
			opts: &ossfpm.Options{
				SecretRef: "test-secret",
				FuseType:  "ossfs2",
			},
			wantOptions: []string{
				"oss_sts_multi_conf_ak_file=/etc/ossfs2/passwd-ossfs2/AccessKeyId",
				"oss_sts_multi_conf_sk_file=/etc/ossfs2/passwd-ossfs2/AccessKeySecret",
				"oss_sts_multi_conf_token_file=/etc/ossfs2/passwd-ossfs2/SecurityToken",
			},
		},
		{
			name: "aksk",
			opts: &ossfpm.Options{
				FuseType: "ossfs2",
			},
		},
		{
			name: "rrsa",
			opts: &ossfpm.Options{
				FuseType:   "ossfs2",
				AuthType:   ossfpm.AuthTypeRRSA,
				ExternalId: "test-id",
			},
			wantOptions: []string{
				"rrsa_endpoint=https://sts-vpc.cn-hangzhou.aliyuncs.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOssfs := &fuseOssfs{}
			opts := fakeOssfs.getAuthOptions(tt.opts, "cn-hangzhou")
			assert.Equal(t, tt.wantOptions, opts)
		})
	}
}

func TestBuildAuthSpec_ossfs2(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &fpm.AuthConfig{
		SecretRef: "test-secret-ref",
	}
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
	fakeOssfs := &fuseOssfs{}
	fakeOssfs.buildAuthSpec(&fpm.FusePodContext{
		Context:    context.Background(),
		Namespace:  mounterutils.LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
		FuseType:   "ossfs2",
	}, "target", &spec, &container)

	volumeMount := container.VolumeMounts[len(container.VolumeMounts)-1]
	assert.Equal(t, "passwd-ossfs2", volumeMount.Name)

	volume := spec.Volumes[len(spec.Volumes)-1]
	assert.Equal(t, "passwd-ossfs2", volume.Name)
	expectedItems := []corev1.KeyToPath{
		{
			Key:  "AccessKeyId",
			Path: "passwd-ossfs2/AccessKeyId",
			Mode: tea.Int32(0600),
		},
		{
			Key:  "AccessKeySecret",
			Path: "passwd-ossfs2/AccessKeySecret",
			Mode: tea.Int32(0600),
		},
		{
			Key:  "Expiration",
			Path: "passwd-ossfs2/Expiration",
			Mode: tea.Int32(0600),
		},
		{
			Key:  "SecurityToken",
			Path: "passwd-ossfs2/SecurityToken",
			Mode: tea.Int32(0600),
		},
	}
	assert.Equal(t, expectedItems, volume.Secret.Items)
	assert.Equal(t, "test-secret-ref", volume.Secret.SecretName)

	spec = corev1.PodSpec{}
	spec.Volumes = []corev1.Volume{targetVolume}
	rrsaCfg := fpm.RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	authCfg.AuthType = ossfpm.AuthTypeRRSA
	fakeOssfs.buildAuthSpec(&fpm.FusePodContext{
		Context:    context.Background(),
		Namespace:  mounterutils.LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
		FuseType:   ossfpm.OssFsType,
	}, "target", &spec, &container)

	assert.Equal(t, "rrsa-oidc-token", spec.Volumes[len(spec.Volumes)-1].Name)
	volumeMount = container.VolumeMounts[len(container.VolumeMounts)-1]
	assert.Contains(t, "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens", volumeMount.MountPath)
	assert.Contains(t, "rrsa-oidc-token", volumeMount.Name)
}
