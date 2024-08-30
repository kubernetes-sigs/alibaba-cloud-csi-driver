package mounter

import (
	"context"
	"reflect"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func Test_buildAuthSpec(t *testing.T) {
	nodeName := "test-node-name"
	volumeId := "test-pv-name"
	authCfg := &AuthConfig{}
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
				Type: new(corev1.HostPathType),
			},
		},
	}
	spec := corev1.PodSpec{}
	spec.Volumes = []corev1.Volume{targetVolume}
	rrsaCfg := RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	authCfg.AuthType = AuthTypeRRSA
	buildAuthSpec(&FusePodContext{
		Context:    context.Background(),
		Namespace:  LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
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
				config: FuseContainerConfig{
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
