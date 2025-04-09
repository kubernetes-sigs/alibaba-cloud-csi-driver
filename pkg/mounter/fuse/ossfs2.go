package fuse

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

var defaultOssfs2ImageTag = "v0.0.0-052ed2a-aliyun"

type fuseOssfs2 struct {
	config utils.FuseContainerConfig
}

var ossfs2Dbglevels = map[string]string{
	utils.DebugLevelDebug: "debug",
	utils.DebugLevelInfo:  "info",
}

const defaultDbglevel = utils.DebugLevelInfo

func NewFuseOssfs2(configmap *corev1.ConfigMap, m metadata.MetadataProvider) utils.FuseMounterType {
	config := utils.ExtractFuseContainerConfig(configmap, OssFs2Type)
	// set default image
	setDefaultImage(OssFs2Type, m, &config)
	// set default memory request
	if _, ok := config.Resources.Requests[corev1.ResourceMemory]; !ok {
		config.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("50Mi")
	}

	return &fuseOssfs2{config: config}
}

func (f *fuseOssfs2) Name() string {
	return OssFs2Type
}

func (f *fuseOssfs2) CheckAuthConfig(c *utils.AuthConfig) error {
	if c.AuthType != "" {
		return fmt.Errorf("%s do not support authType: %s", OssFs2Type, c.AuthType)
	}
	return nil
}

func (f *fuseOssfs2) PodTemplateSpec(c *utils.FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	spec, err := f.buildPodSpec(c, target)
	if err != nil {
		return nil, err
	}

	pod := new(corev1.PodTemplateSpec)
	pod.Spec = spec

	pod.Annotations = make(map[string]string)
	for k, v := range f.config.Annotations {
		pod.Annotations[k] = v
	}
	pod.Labels = make(map[string]string)
	for k, v := range f.config.Labels {
		pod.Labels[k] = v
	}
	return pod, nil
}

func (f *fuseOssfs2) buildPodSpec(c *utils.FusePodContext, target string) (spec corev1.PodSpec, _ error) {
	targetDir := filepath.Dir(target)
	targetDirVolume := corev1.Volume{
		Name: "target-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: targetDir,
				Type: new(corev1.HostPathType),
			},
		},
	}
	*targetDirVolume.HostPath.Type = corev1.HostPathDirectoryOrCreate
	spec.Volumes = []corev1.Volume{targetDirVolume}

	bidirectional := corev1.MountPropagationBidirectional
	socketPath := utils.GetMountProxySocketPath(c.VolumeId)
	container := corev1.Container{
		Name:  f.Name(),
		Image: f.config.Image,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:             targetDirVolume.Name,
				MountPath:        targetDir,
				MountPropagation: &bidirectional,
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Privileged: tea.Bool(true),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				Exec: &corev1.ExecAction{
					Command: []string{
						"test", "-S", socketPath,
					},
				},
			},
			PeriodSeconds:    2,
			FailureThreshold: 5,
		},
	}

	container.Args = []string{"-socket=" + socketPath, "-v=4"}

	spec.Containers = []corev1.Container{container}
	spec.NodeName = c.NodeName
	spec.HostNetwork = true
	spec.PriorityClassName = "system-node-critical"
	spec.Tolerations = []corev1.Toleration{{Operator: corev1.TolerationOpExists}}
	return
}

func (f *fuseOssfs2) AddDefaultMountOptions(options []string) []string {
	alreadySet := false
	for _, option := range options {
		if strings.Contains(option, "log_level") {
			alreadySet = true
			break
		}
	}
	if !alreadySet {
		level, ok := ossfs2Dbglevels[f.config.Dbglevel]
		if ok {
			options = append(options, fmt.Sprintf("log_level=%s", level))
		} else {
			if f.config.Dbglevel != "" {
				klog.Warningf("invalid dbglevel for ossfs: %q, use default dbglevel %s", f.config.Dbglevel, defaultDbglevel)
			}
			options = append(options, fmt.Sprintf("log_level=%s", ossfs2Dbglevels[defaultDbglevel]))
		}
	}

	return options
}
