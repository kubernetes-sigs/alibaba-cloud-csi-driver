package mounter

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
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
	// ak
	passwdMountDir := "/etc/ossfs"
	passwdFilename := "passwd-ossfs"
	passwdSecretVolume := corev1.Volume{
		Name: "passwd-ossfs",
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: OssfsCredentialSecretName,
				Items: []corev1.KeyToPath{
					{
						Key:  fmt.Sprintf("%s.%s", nodeName, volumeId),
						Path: passwdFilename,
						Mode: pointer.Int32Ptr(0600),
					},
				},
				Optional: pointer.BoolPtr(true),
			},
		},
	}
	passwdVolumeMont := corev1.VolumeMount{
		Name:      passwdSecretVolume.Name,
		MountPath: passwdMountDir,
	}
	rrsaCfg := RrsaConfig{
		OidcProviderArn: "test-oidc-provider-arn",
		RoleArn:         "test-role-arn",
	}
	authCfg.RrsaConfig = &rrsaCfg
	buildAuthSpec(&FusePodContext{
		Context:    context.Background(),
		Namespace:  LegacyFusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   volumeId,
		AuthConfig: authCfg,
	}, "target", &spec, &container)
	assert.Contains(t, spec.Volumes, passwdSecretVolume)
	assert.Contains(t, container.VolumeMounts, passwdVolumeMont)
}
