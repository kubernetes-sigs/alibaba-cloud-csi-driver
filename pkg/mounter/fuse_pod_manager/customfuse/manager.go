package customfuse

import (
	"context"
	"fmt"
	"maps"
	"path/filepath"
	"strings"
	"time"

	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

const (
	configMapName      = "csi-plugin"
	configMapNamespace = "kube-system"
)

type CustomFuse struct {
	defaultConfig fpm.FuseContainerConfig
	client        kubernetes.Interface
}

func NewCustomFuse(csiCfg utils.Config, client kubernetes.Interface) *CustomFuse {
	defaultConfig := fpm.ExtractFuseContainerConfig(csiCfg, mounterutils.CustomFuseType)
	defaultConfig.Image = extractImage(csiCfg.ConfigMap, mounterutils.CustomFuseType)

	return &CustomFuse{
		defaultConfig: defaultConfig,
		client:        client,
	}
}

// resolveConfig reads the csi-plugin ConfigMap and extracts the FuseContainerConfig
// for the given fuseType. The ConfigMap key is "fuse-{fuseType}".
// This is called per-mount to pick up image changes without restarting.
// Falls back to defaultConfig if the ConfigMap read fails or the fuseType key is missing.
//
// Unlike ossfs, customfuse requires a full image path (e.g. "registry.cn-hangzhou.cr.aliyuncs.com/acs/csi-fuse-juicefs:v1.0")
// because different fuseTypes may use completely different registries and image names.
// The "image" key is parsed here directly; ExtractFuseContainerConfig ignores it for ossfs backward compat.
func (f *CustomFuse) resolveConfig(fuseType string) fpm.FuseContainerConfig {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cm, err := f.client.CoreV1().ConfigMaps(configMapNamespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil {
		klog.Warningf("Failed to read configmap %s/%s: %v, using default config", configMapNamespace, configMapName, err)
		return f.defaultConfig
	}

	cfg := utils.Config{ConfigMap: cm.Data}
	config := fpm.ExtractFuseContainerConfig(cfg, fuseType)
	config.Image = extractImage(cm.Data, fuseType)

	if config.Image == "" {
		klog.V(2).Infof("No image configured for fuseType %q in configmap, using default config", fuseType)
		return f.defaultConfig
	}

	return config
}

// extractImage parses the "image" key from the "fuse-{fuseType}" entry in the ConfigMap data.
func extractImage(data map[string]string, fuseType string) string {
	content, ok := data["fuse-"+fuseType]
	if !ok {
		return ""
	}
	for line := range strings.SplitSeq(content, "\n") {
		key, value, _ := strings.Cut(strings.TrimSpace(line), "=")
		if key == "image" {
			return value
		}
	}
	return ""
}

func (f *CustomFuse) Name() string {
	return mounterutils.CustomFuseType
}

func (f *CustomFuse) PodTemplateSpec(c *fpm.FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	fuseType := c.FuseType
	if fuseType == "" {
		fuseType = mounterutils.CustomFuseType
	}

	config := f.resolveConfig(fuseType)

	spec, err := f.buildPodSpec(config, c, target)
	if err != nil {
		return nil, err
	}

	pod := new(corev1.PodTemplateSpec)
	pod.Spec = spec
	pod.Annotations = maps.Clone(config.Annotations)
	pod.Labels = maps.Clone(config.Labels)
	return pod, nil
}

func (f *CustomFuse) AddDefaultMountOptions(options []string) []string {
	return options
}

func (f *CustomFuse) buildPodSpec(config fpm.FuseContainerConfig, c *fpm.FusePodContext, target string) (spec corev1.PodSpec, _ error) {
	if config.Image == "" {
		return spec, fmt.Errorf("no image configured for fuseType %q; set image in configmap %s/%s under key fuse-%s",
			c.FuseType, configMapNamespace, configMapName, c.FuseType)
	}

	targetDir := filepath.Dir(target)
	targetDirVolume := corev1.Volume{
		Name: "target-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: targetDir,
				Type: ptr.To(corev1.HostPathDirectoryOrCreate),
			},
		},
	}

	metricsDir := utils.GetFuseMetricsMountDir("/var/run/customfuse", c.VolumeId)
	metricsDirVolume := corev1.Volume{
		Name: "metrics-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: metricsDir,
				Type: ptr.To(corev1.HostPathDirectoryOrCreate),
			},
		},
	}

	spec.Volumes = []corev1.Volume{targetDirVolume, metricsDirVolume}

	bidirectional := corev1.MountPropagationBidirectional
	socketPath := mounterutils.GetMountProxySocketPath(c.VolumeId, true)

	volumeMounts := []corev1.VolumeMount{
		{
			Name:             targetDirVolume.Name,
			MountPath:        targetDir,
			MountPropagation: &bidirectional,
		},
		{
			Name:      metricsDirVolume.Name,
			MountPath: metricsDirVolume.HostPath.Path,
		},
	}

	if c.EntrypointConfig != "" {
		execMode := int32(0755)
		entrypointKey := c.EntrypointKey
		if entrypointKey == "" {
			entrypointKey = "entrypoint.sh"
		}
		configVolume := corev1.Volume{
			Name: "entrypoint-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: c.EntrypointConfig,
					},
					Items: []corev1.KeyToPath{{
						Key:  entrypointKey,
						Path: "entrypoint.sh",
					}},
					DefaultMode: &execMode,
				},
			},
		}
		spec.Volumes = append(spec.Volumes, configVolume)
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      configVolume.Name,
			MountPath: "/etc/fuse-config",
			ReadOnly:  true,
		})
	}

	container := corev1.Container{
		Name:         f.Name(),
		Image:        config.Image,
		Resources:    config.Resources,
		VolumeMounts: volumeMounts,
		SecurityContext: &corev1.SecurityContext{
			Privileged: ptr.To(true),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				Exec: &corev1.ExecAction{
					Command: []string{"test", "-S", socketPath},
				},
			},
			PeriodSeconds:    2,
			FailureThreshold: 5,
		},
	}

	container.Args = []string{"--socket=" + socketPath, "-v=4"}

	spec.Containers = []corev1.Container{container}
	spec.NodeName = c.NodeName
	spec.HostNetwork = true
	spec.DNSPolicy = c.PodTemplateConfig.DnsPolicy
	spec.PriorityClassName = "system-node-critical"
	spec.Tolerations = []corev1.Toleration{{Operator: corev1.TolerationOpExists}}
	return
}
