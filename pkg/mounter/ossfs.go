package mounter

import (
	"errors"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	mountutils "k8s.io/mount-utils"
	"k8s.io/utils/pointer"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

var FuseOssfs FuseMounterType = &fuseOssfs{}

type fuseOssfs struct{}

func (f *fuseOssfs) name() string {
	return "ossfs"
}

func (f *fuseOssfs) buildPodSpec(
	source, target, fstype string, options, mountFlags []string, nodeName string, config FuseContainerConfig,
) (spec corev1.PodSpec, _ error) {
	if config.Image == "" {
		return spec, errors.New("missing image configuration")
	}

	hostPathDirectoryType := corev1.HostPathDirectory
	hostPathFileOrCreate := corev1.HostPathFileOrCreate
	kubeletDirVolume := corev1.Volume{
		Name: "kubelet-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: utils.KubeletRootDir,
				Type: &hostPathDirectoryType,
			},
		},
	}
	metricsDirVolume := corev1.Volume{
		Name: "metrics-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/run/ossfs",
				Type: &hostPathDirectoryType,
			},
		},
	}
	passwdFileVolume := corev1.Volume{
		Name: "passwd-ossfs",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc/passwd-ossfs",
				Type: &hostPathFileOrCreate,
			},
		},
	}
	spec.Volumes = []corev1.Volume{kubeletDirVolume, metricsDirVolume, passwdFileVolume}

	switch dbglevel := config.Extra["dbglevel"]; dbglevel {
	case "info", "warn", "err", "crit":
		alreadySet := false
		for _, option := range options {
			if strings.Contains(option, "dbglevel") {
				alreadySet = true
				break
			}
		}
		if !alreadySet {
			options = append(options, "dbglevel="+dbglevel)
		}
	default:
		return spec, fmt.Errorf("unknown ossfs dbglevel: %q", dbglevel)
	}
	args := mountutils.MakeMountArgs(source, target, "", options)
	args = append(args, mountFlags...)
	propagationContainerToHost := corev1.MountPropagationHostToContainer
	container := corev1.Container{
		Name:  "fuse-mounter",
		Image: config.Image,
		Args:  args,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:             kubeletDirVolume.Name,
				MountPath:        utils.KubeletRootDir,
				MountPropagation: &propagationContainerToHost,
			}, {
				Name:      metricsDirVolume.Name,
				MountPath: metricsDirVolume.HostPath.Path,
			}, {
				Name:      passwdFileVolume.Name,
				MountPath: passwdFileVolume.HostPath.Path,
			},
		},
		Resources: config.Resources,
		SecurityContext: &corev1.SecurityContext{
			Privileged: pointer.BoolPtr(true),
		},
	}
	spec.Containers = []corev1.Container{container}
	spec.RestartPolicy = corev1.RestartPolicyOnFailure
	spec.NodeName = nodeName
	spec.HostNetwork = true
	spec.PriorityClassName = "system-node-critical"
	return
}
