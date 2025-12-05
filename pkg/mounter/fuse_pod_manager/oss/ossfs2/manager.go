package ossfs2

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

func init() {
	ossfpm.RegisterFuseMounter(ossfpm.OssFs2Type, NewFuseOssfs)
	ossfpm.RegisterFuseMounterPath(ossfpm.OssFs2Type, "/usr/local/bin/ossfs2")
	ossfpm.RegisterFuseInterceptors(ossfpm.OssFs2Type, []mounter.MountInterceptor{interceptors.Ossfs2SecretInterceptor})
}

var defaultOssfs2Dbglevel = fpm.DebugLevelInfo

type fuseOssfs struct {
	config fpm.FuseContainerConfig
}

var ossfs2Dbglevels = map[string]string{
	fpm.DebugLevelDebug: "debug",
	fpm.DebugLevelInfo:  "info",
}

func NewFuseOssfs(configmap *corev1.ConfigMap, m metadata.MetadataProvider) ossfpm.OSSFuseMounterType {
	config := fpm.ExtractFuseContainerConfig(configmap, ossfpm.OssFs2Type)
	// set default image
	ossfpm.SetDefaultImage(ossfpm.OssFs2Type, m, &config)
	// set default memory request
	if _, ok := config.Resources.Requests[corev1.ResourceMemory]; !ok {
		config.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("50Mi")
	}

	return &fuseOssfs{config: config}
}

func (f *fuseOssfs) Name() string {
	return ossfpm.OssFs2Type
}
func (f *fuseOssfs) PrecheckAuthConfig(o *ossfpm.Options, onNode bool) error {

	if o.AuthType != ossfpm.AuthTypeRRSA && o.AssumeRoleArn != "" {
		return fmt.Errorf("only support access OSS through STS AssumeRole when authType is RRSA")
	}

	switch o.AuthType {
	case ossfpm.AuthTypeRRSA:
		if err := ossfpm.CheckRRSAParams(o); err != nil {
			return err
		}
	case ossfpm.AuthTypeSTS:
		// rolename may retrieve from metadata service
		if onNode && o.RoleName == "" {
			return fmt.Errorf("missing roleName or ramRole in volume attributes")
		}
	case "":
		if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
			return nil
		}
		if o.SecretRef != "" {
			if o.AkID != "" || o.AkSecret != "" {
				return fmt.Errorf("AK and secretRef cannot be set at the same time")
			}
			if o.SecretRef == mounterutils.GetCredientialsSecretName(ossfpm.OssFsType) {
				return fmt.Errorf("invalid SecretRef name")
			}
			return nil
		}
		// aksk may retrieve from ENV
		if onNode && (o.AkID == "" || o.AkSecret == "") {
			return fmt.Errorf("missing access key in node publish secret")
		}
	default:
		return fmt.Errorf("%s do not support authType: %s", f.Name(), o.AuthType)
	}
	return nil
}

func (f *fuseOssfs) MakeAuthConfig(o *ossfpm.Options, m metadata.MetadataProvider) (authCfg *fpm.AuthConfig, err error) {
	authCfg = &fpm.AuthConfig{AuthType: o.AuthType}
	switch o.AuthType {
	case ossfpm.AuthTypeRRSA:
		rrsaCfg, err := ossfpm.GetRRSAConfig(o, m)
		if err != nil {
			return nil, fmt.Errorf("Get RoleArn and OidcProviderArn for RRSA error: %v", err)
		}
		authCfg.RrsaConfig = rrsaCfg
	case ossfpm.AuthTypeSTS:
		authCfg.RoleName = o.RoleName
	case "":
		if o.SecretRef != "" {
			authCfg.SecretRef = o.SecretRef
			return
		}
		authCfg.Secrets = map[string]string{
			mounterutils.GetPasswdFileName(f.Name()): fmt.Sprintf("--oss_access_key_id=%s\n--oss_access_key_secret=%s", o.AkID, o.AkSecret),
		}
	default:
		return nil, fmt.Errorf("%s do not support authType: %s", f.Name(), o.AuthType)
	}
	return authCfg, nil
}

func (f *fuseOssfs) MakeMountOptions(o *ossfpm.Options, m metadata.MetadataProvider) (mountOptions []string, err error) {
	mountOptions = append(mountOptions,
		"oss_endpoint="+o.URL,
		"oss_bucket="+o.Bucket,
		"oss_bucket_prefix="+o.Path,
	)

	if o.ReadOnly {
		mountOptions = append(mountOptions, "ro=true")
	}

	if o.MetricsTop != "" {
		mountOptions = append(mountOptions, "use_metrics=true")
		mountOptions = append(mountOptions, fmt.Sprintf("metrics_top=%s", o.MetricsTop))
	}

	if o.SigVersion == ossfpm.SigV4 {
		region, _ := m.Get(metadata.RegionID)
		if region == "" {
			return nil, fmt.Errorf("SigV4 is not supported without region")
		}
		mountOptions = append(mountOptions, fmt.Sprintf("oss_region=%s", region))
	}
	region, _ := m.Get(metadata.RegionID)
	authOptions := f.getAuthOptions(o, region)
	mountOptions = append(mountOptions, authOptions...)

	return
}

func (f *fuseOssfs) PodTemplateSpec(c *fpm.FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	spec, err := f.buildPodSpec(c, target)
	if err != nil {
		return nil, err
	}

	pod := new(corev1.PodTemplateSpec)
	pod.Spec = spec

	pod.Annotations = maps.Clone(f.config.Annotations)
	pod.Labels = maps.Clone(f.config.Labels)
	return pod, nil
}

func (f *fuseOssfs) getAuthOptions(o *ossfpm.Options, region string) (mountOptions []string) {
	switch o.AuthType {
	case ossfpm.AuthTypeRRSA:
		mountOptions = append(mountOptions, fmt.Sprintf("rrsa_endpoint=%s", ossfpm.GetSTSEndpoint(region)))
		if o.AssumeRoleArn != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("assume_role_arn=%s", o.AssumeRoleArn))
			if o.ExternalId != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("assume_role_external_id=%s", o.ExternalId))
			}
		}
	case ossfpm.AuthTypeSTS:
		if o.RoleName != "" {
			mountOptions = append(mountOptions, "ram_role="+o.RoleName)
		}
	case "":
		if o.SecretRef != "" {
			mountOptions = append(mountOptions,
				fmt.Sprintf("oss_sts_multi_conf_ak_file=%s", filepath.Join(mounterutils.GetConfigDir(o.FuseType), mounterutils.GetPasswdFileName(o.FuseType), ossfpm.KeyAccessKeyId)),
				fmt.Sprintf("oss_sts_multi_conf_sk_file=%s", filepath.Join(mounterutils.GetConfigDir(o.FuseType), mounterutils.GetPasswdFileName(o.FuseType), ossfpm.KeyAccessKeySecret)),
				fmt.Sprintf("oss_sts_multi_conf_token_file=%s", filepath.Join(mounterutils.GetConfigDir(o.FuseType), mounterutils.GetPasswdFileName(o.FuseType), ossfpm.KeySecurityToken)),
			)
		}
		// publishSecretRef will make option in mount-proxy server
	default:
		return nil
	}
	return
}

func (f *fuseOssfs) buildPodSpec(c *fpm.FusePodContext, target string) (spec corev1.PodSpec, _ error) {
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
	// In runc scenarios, CSI currently only supports shared mounting mode.
	// Therefore, narrowing down the mount path to the sha256(volumeId) directory
	// does not affect.
	// Note: Reuse the same metrics directory path as ossfs like targetDir.
	// Since ossfs and ossfs2 always run in separate pods (one per volume),
	// they won't interfere with each other even when using the same path pattern.
	metricsDir := utils.GetFuseMetricsMountDir("/var/run/ossfs", c.VolumeId)
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
	socketPath := mounterutils.GetMountProxySocketPath(c.VolumeId)
	container := corev1.Container{
		Name:  f.Name(),
		Image: f.config.Image,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:             targetDirVolume.Name,
				MountPath:        targetDir,
				MountPropagation: &bidirectional,
			}, {
				Name:      metricsDirVolume.Name,
				MountPath: metricsDirVolume.HostPath.Path,
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

	f.buildAuthSpec(c, target, &spec, &container)

	container.Args = []string{"--socket=" + socketPath, "-v=4"}

	if strings.ToLower(f.config.Extra["set-dumpable"]) == "true" {
		container.Env = append(container.Env, corev1.EnvVar{Name: "SET_DUMPABLE", Value: "true"})
	}

	spec.Containers = []corev1.Container{container}
	spec.NodeName = c.NodeName
	spec.HostNetwork = true
	spec.DNSPolicy = c.PodTemplateConfig.DnsPolicy
	spec.PriorityClassName = "system-node-critical"
	spec.Tolerations = []corev1.Toleration{{Operator: corev1.TolerationOpExists}}
	return
}

const (
	KeyLogLevel = "log_level"
	KeyLogDir   = "log_dir"
)

func (f *fuseOssfs) AddDefaultMountOptions(options []string) []string {
	defaultOSSFSOptions := os.Getenv("DEFAULT_OSSFS2_OPTIONS")
	if defaultOSSFSOptions != "" {
		options = append(options, strings.Split(defaultOSSFSOptions, ",")...)
	}

	tm := map[string]string{}
	for _, option := range options {
		if option == "" {
			continue
		}
		k, v, _ := strings.Cut(option, "=")
		tm[k] = v
	}

	// set default log level
	if _, ok := tm[KeyLogLevel]; !ok {
		level, ok := ossfs2Dbglevels[f.config.Dbglevel]
		if ok {
			options = append(options, fmt.Sprintf("log_level=%s", level))
		} else {
			if f.config.Dbglevel != "" {
				klog.Warningf("invalid dbglevel for ossfs: %q, use default dbglevel %s", f.config.Dbglevel, defaultOssfs2Dbglevel)
			}
			options = append(options, fmt.Sprintf("log_level=%s", ossfs2Dbglevels[defaultOssfs2Dbglevel]))
		}
	}

	// set default log dir
	if _, ok := tm[KeyLogDir]; !ok {
		options = append(options, "log_dir=/dev/stdout")
	}

	// set use_metrics to enabled monitoring by default
	if _, ok := tm["use_metrics"]; !ok {
		// TODO: After stabilization, change the default value to true
		// Currently, metrics collection is controlled by the MetricsMode configuration
		// Once the feature is stable, we should enable metrics by default
		if f.config.MetricsMode == fpm.MetricsModeEnabled {
			options = append(options, "use_metrics=true")
		}
	}

	return options
}

func (f *fuseOssfs) buildAuthSpec(c *fpm.FusePodContext, target string, spec *corev1.PodSpec, container *corev1.Container) {
	if spec == nil || container == nil {
		return
	}
	authCfg := c.AuthConfig
	if authCfg == nil {
		return
	}

	switch authCfg.AuthType {
	case ossfpm.AuthTypeSTS:
	case ossfpm.AuthTypeRRSA:
		if authCfg.RrsaConfig == nil {
			return
		}
		spec.ServiceAccountName = authCfg.RrsaConfig.ServiceAccountName
		rrsaMountDir := "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens"
		rrsaVolume := corev1.Volume{
			Name: "rrsa-oidc-token",
			VolumeSource: corev1.VolumeSource{
				Projected: &corev1.ProjectedVolumeSource{
					DefaultMode: tea.Int32(0600),
					Sources: []corev1.VolumeProjection{
						{
							ServiceAccountToken: &corev1.ServiceAccountTokenProjection{
								Audience:          "sts.aliyuncs.com",
								ExpirationSeconds: tea.Int64(3600),
								Path:              "token",
							},
						},
					},
				},
			},
		}
		spec.Volumes = append(spec.Volumes, rrsaVolume)
		rrsaVolumeMount := corev1.VolumeMount{
			Name:      rrsaVolume.Name,
			MountPath: rrsaMountDir,
		}
		container.VolumeMounts = append(container.VolumeMounts, rrsaVolumeMount)
		envs := []corev1.EnvVar{
			{
				Name:  "ALIBABA_CLOUD_ROLE_ARN",
				Value: authCfg.RrsaConfig.RoleArn,
			},
			{
				Name:  "ALIBABA_CLOUD_OIDC_PROVIDER_ARN",
				Value: authCfg.RrsaConfig.OidcProviderArn,
			},
			{
				Name:  "ALIBABA_CLOUD_OIDC_TOKEN_FILE",
				Value: rrsaMountDir + "/token",
			},
			{
				Name:  "ROLE_SESSION_NAME",
				Value: mounterutils.GetRoleSessionName(c.VolumeId, target, c.FuseType),
			},
		}
		container.Env = append(container.Env, envs...)
	case "":
		secretVolumeSource := ossfpm.GetPasswdSecretVolume(authCfg.SecretRef, c.FuseType)
		if secretVolumeSource != nil {
			passwdSecretVolume := corev1.Volume{
				Name: mounterutils.GetPasswdFileName(c.FuseType),
				VolumeSource: corev1.VolumeSource{
					Secret: secretVolumeSource,
				},
			}
			spec.Volumes = append(spec.Volumes, passwdSecretVolume)
			passwdVolumeMont := corev1.VolumeMount{
				Name:      passwdSecretVolume.Name,
				MountPath: mounterutils.GetConfigDir(c.FuseType),
			}
			container.VolumeMounts = append(container.VolumeMounts, passwdVolumeMont)
		}
	}

}
