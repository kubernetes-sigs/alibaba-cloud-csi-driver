package oss

import (
	"context"
	"fmt"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func TestPrecheckAuthConfig_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
	tests := []struct {
		name    string
		opts    *Options
		wantErr bool
	}{
		{
			"invalid authtype",
			&Options{
				AuthType: AuthTypeCSS,
			},
			true,
		},
		{
			"valid sts",
			&Options{
				AuthType: AuthTypeSTS,
				RoleName: "test",
			},
			false,
		},
		{
			"invalid sts",
			&Options{
				AuthType: AuthTypeSTS,
			},
			true,
		},
		{
			"invalid rrsa",
			&Options{
				AuthType: AuthTypeRRSA,
			},
			true,
		},
		{
			"valid rrsa",
			&Options{
				AuthType: AuthTypeRRSA,
				RoleName: "test",
			},
			false,
		},
		{
			"use assumeRole with non-RRSA authType",
			&Options{
				AssumeRoleArn: "test-assume-role-arn",
			},
			true,
		},
		{
			"empty aksecret",
			&Options{
				AkID:     "test-ak",
				AkSecret: "",
			},
			true,
		},
		{
			"success - aksk",
			&Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			false,
		},
		{
			"success - secretref",
			&Options{
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
	fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
	tests := []struct {
		name    string
		options *Options
		wantCfg *utils.AuthConfig
		wantErr bool
	}{
		{
			"sts",
			&Options{
				AuthType: AuthTypeSTS,
				RoleName: "test-role-name",
			},
			&utils.AuthConfig{
				AuthType: AuthTypeSTS,
				RoleName: "test-role-name",
			},
			false,
		},
		{
			"aksk",
			&Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			&utils.AuthConfig{
				Secrets: map[string]string{
					utils.GetPasswdFileName(fakeOssfs.Name()): fmt.Sprintf("--oss_access_key_id=%s\n--oss_access_key_secret=%s", "test-ak", "test-ak-secret"),
				},
			},
			false,
		},
		{
			"secretref",
			&Options{
				SecretRef: "test-secretref",
			},
			&utils.AuthConfig{
				SecretRef: "test-secretref",
			},
			false,
		},
		{
			"rrsa with rolename",
			&Options{
				AuthType: AuthTypeRRSA,
				RoleName: "test-role",
			},
			&utils.AuthConfig{
				AuthType: AuthTypeRRSA,
				RrsaConfig: &utils.RrsaConfig{
					ServiceAccountName: mounterutils.FuseServiceAccountName,
					RoleArn:            "acs:ram::account-id:role/test-role",
					OidcProviderArn:    "acs:ram::account-id:oidc-provider/ack-rrsa-cluster-id",
				},
			},
			false,
		},
		{
			"rrsa with arns",
			&Options{
				AuthType:           AuthTypeRRSA,
				ServiceAccountName: "test-sa",
				RoleArn:            "test-role-arn",
				OidcProviderArn:    "test-oidc-provider-arn",
				AssumeRoleArn:      "test-assume-role-arn",
			},
			&utils.AuthConfig{
				AuthType: AuthTypeRRSA,
				RrsaConfig: &utils.RrsaConfig{
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
		opts          *Options
		region        string
		expected      []string
		expectedError bool
	}{
		{
			name: "ro",
			opts: &Options{
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
			opts: &Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: SigV4,
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
			opts: &Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: SigV4,
			},
			expectedError: true,
		},
		{
			name: "secretref",
			opts: &Options{
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
			opts: &Options{
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
				"metrics_top=5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("REGION_ID", tt.region)
			fakeMeta := metadata.NewMetadata()
			fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
			opts, err := fakeOssfs.MakeMountOptions(tt.opts, fakeMeta)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.expected, opts)
		})
	}
}

func TestAddDefaultMountOptions_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeInter := NewFuseOssfs2(nil, fakeMeta)
	fakeOssfs, ok := fakeInter.(*fuseOssfs2)
	if !ok {
		t.Fatalf("failed to cast to fuseOssfs2")
	}
	tests := []struct {
		name        string
		options     []string
		cfglevel    string
		defaultOpts string
		want        []string
	}{
		{
			name:    "empty option, empty config",
			options: []string{"others"},
			want:    []string{"others", "log_level=info", "log_dir=/dev/stdout", "use_metrics=basic"},
		},
		{
			name:    "set option",
			options: []string{"others", "log_level=debug", "log_dir=/tmp/ossfs2", "others", "use_metrics=advanced"},
			want:    []string{"others", "log_level=debug", "log_dir=/tmp/ossfs2", "others", "use_metrics=advanced"},
		},
		{
			name:     "set option, set config",
			cfglevel: "info",
			options:  []string{"others", "log_level=debug", "others"},
			want:     []string{"others", "log_level=debug", "others", "log_dir=/dev/stdout", "use_metrics=basic"},
		},
		{
			name:     "empty option, set config",
			cfglevel: "debug",
			options:  []string{"others"},
			want:     []string{"others", "log_level=debug", "log_dir=/dev/stdout", "use_metrics=basic"},
		},
		{
			name:     "empty option, invalid config",
			cfglevel: "invalid",
			options:  []string{"others"},
			want:     []string{"others", "log_level=info", "log_dir=/dev/stdout", "use_metrics=basic"},
		},
		{
			name:        "default options",
			cfglevel:    "",
			options:     nil,
			defaultOpts: "others,log_dir=/tmp/ossfs2",
			want:        []string{"others", "log_dir=/tmp/ossfs2", "log_level=info", "use_metrics=basic"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.defaultOpts != "" {
				t.Setenv("DEFAULT_OSSFS2_OPTIONS", tt.defaultOpts)
			}
			fakeOssfs.config.Dbglevel = tt.cfglevel
			got := fakeOssfs.AddDefaultMountOptions(tt.options)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetAuthOpttions_ossfs2(t *testing.T) {
	tests := []struct {
		name        string
		opts        *Options
		wantOptions []string
	}{
		{
			name: "secretref",
			opts: &Options{
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
			opts: &Options{
				FuseType: "ossfs2",
			},
		},
		{
			name: "rrsa",
			opts: &Options{
				FuseType:   "ossfs2",
				AuthType:   AuthTypeRRSA,
				ExternalId: "test-id",
			},
			wantOptions: []string{
				"rrsa_endpoint=https://sts-vpc.cn-hangzhou.aliyuncs.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOssfs := &fuseOssfs2{}
			opts := fakeOssfs.getAuthOptions(tt.opts, "cn-hangzhou")
			assert.Equal(t, tt.wantOptions, opts)
		})
	}
}

func TestBuildAuthSpec_ossfs2(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &utils.AuthConfig{
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
	fakeOssfs := &fuseOssfs2{}
	fakeOssfs.buildAuthSpec(&mounterutils.FusePodContext{
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
	rrsaCfg := mounterutils.RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	authCfg.AuthType = AuthTypeRRSA
	fakeOssfs.buildAuthSpec(&mounterutils.FusePodContext{
		Context:    context.Background(),
		Namespace:  mounterutils.LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
		FuseType:   OssFsType,
	}, "target", &spec, &container)

	assert.Equal(t, "rrsa-oidc-token", spec.Volumes[len(spec.Volumes)-1].Name)
	volumeMount = container.VolumeMounts[len(container.VolumeMounts)-1]
	assert.Contains(t, "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens", volumeMount.MountPath)
	assert.Contains(t, "rrsa-oidc-token", volumeMount.Name)
}

func TestTranslateMetricsModeToOption_ossfs2(t *testing.T) {
	fakeOssfs := &fuseOssfs2{}
	tests := []struct {
		name string
		mode string
		want string
	}{
		{"default", "", "use_metrics=basic"},
		{"advanced", mounterutils.MetricsModeAdvanced, "use_metrics=advanced"},
		{"disabled", mounterutils.MetricsModeDisabled, ""},
		{"enabled", mounterutils.MetricsModeEnabled, "use_metrics=basic"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := fakeOssfs.translateMetricsModeToOption(tt.mode)
			assert.Equal(t, tt.want, actual)
		})
	}
}
