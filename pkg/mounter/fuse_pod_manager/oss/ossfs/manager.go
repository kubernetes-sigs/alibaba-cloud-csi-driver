package ossfs

import (
	"fmt"
	"maps"
	"os"
	"path"
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
	ossfpm.RegisterFuseMounter(mounterutils.OssFsType, NewFuseOssfs)
	ossfpm.RegisterFuseMounterPath(mounterutils.OssFsType, "/usr/local/bin/ossfs")
	ossfpm.RegisterFuseInterceptors(mounterutils.OssFsType, []mounter.MountInterceptor{interceptors.OssfsSecretInterceptor})
}

var defaultOssfsDbglevel = fpm.DebugLevelWarn

const (
	hostPrefix                = "/host"
	OssfsDefMimeTypesFilePath = "/etc/mime.types"
	// The ossfs image includes a default MIME configuration located at /csi/mime.types
	OssfsCsiMimeTypesFilePath = "/csi/mime.types"

	CsiSecretStoreDriver   = "secrets-store.csi.k8s.io"
	SecretProviderClassKey = "secretProviderClass"
)

type fuseOssfs struct {
	config fpm.FuseContainerConfig
}

var ossfsDbglevels = map[string]string{
	fpm.DebugLevelDebug: "debug",
	fpm.DebugLevelInfo:  "info",
	fpm.DebugLevelWarn:  "warn",
	fpm.DebugLevelError: "err",
	fpm.DebugLevelFatal: "crit",
}

func NewFuseOssfs(configmap *corev1.ConfigMap, m metadata.MetadataProvider) ossfpm.OSSFuseMounterType {
	config := fpm.ExtractFuseContainerConfig(configmap, mounterutils.OssFsType)

	// set default image
	ossfpm.SetDefaultImage(mounterutils.OssFsType, m, &config)
	// set default memory request
	if _, ok := config.Resources.Requests[corev1.ResourceMemory]; !ok {
		config.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("50Mi")
	}

	return &fuseOssfs{config: config}
}

func (f *fuseOssfs) Name() string {
	return mounterutils.OssFsType
}

func (f *fuseOssfs) PrecheckAuthConfig(o *ossfpm.Options, onNode bool) error {

	if o.AuthType != ossfpm.AuthTypeRRSA && o.AssumeRoleArn != "" {
		return fmt.Errorf("only support access OSS through STS AssumeRole when authType is RRSA")
	}

	switch o.AuthType {
	case ossfpm.AuthTypePublic:
	case ossfpm.AuthTypeSTS:
		// rolename may retrieve from metadata service
		if onNode && o.RoleName == "" {
			return fmt.Errorf("missing roleName or ramRole in volume attributes")
		}
	case ossfpm.AuthTypeRRSA:
		if err := ossfpm.CheckRRSAParams(o); err != nil {
			return err
		}
	case ossfpm.AuthTypeCSS:
		if o.SecretProviderClass == "" {
			return fmt.Errorf("use CsiSecretStore but secretProviderClass is empty")
		}
	default:
		if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
			return nil
		}
		// conflict by secret-volume and republish-token-rotate
		sv := o.SecretRef != ""
		rtr := o.SecurityToken != ""
		if sv && rtr {
			return fmt.Errorf("token and secretRef cannot be set at the same time")
		}
		if sv || rtr {
			if o.AkID != "" || o.AkSecret != "" {
				return fmt.Errorf("AK and secretRef cannot be set at the same time")
			}
			if o.SecretRef == mounterutils.GetCredientialsSecretName(mounterutils.OssFsType) {
				return fmt.Errorf("invalid SecretRef name")
			}
			return nil
		}
		// aksk may retrieve from ENV
		if onNode && (o.AkID == "" || o.AkSecret == "") {
			return fmt.Errorf("missing access key in node publish secret")
		}
	}
	return nil
}

func (f *fuseOssfs) MakeAuthConfig(o *ossfpm.Options, m metadata.MetadataProvider) (*fpm.AuthConfig, error) {
	authCfg := &fpm.AuthConfig{AuthType: o.AuthType}
	switch o.AuthType {
	case ossfpm.AuthTypePublic:
	case ossfpm.AuthTypeRRSA:
		rrsaCfg, err := ossfpm.GetRRSAConfig(o, m)
		if err != nil {
			return nil, fmt.Errorf("Get RoleArn and OidcProviderArn for RRSA error: %v", err)
		}
		authCfg.RrsaConfig = rrsaCfg
	case ossfpm.AuthTypeCSS:
		authCfg.SecretProviderClassName = o.SecretProviderClass
	case ossfpm.AuthTypeSTS:
		authCfg.RoleName = o.RoleName
	default:
		// fixed credentials
		if o.AkID != "" && o.AkSecret != "" {
			authCfg.Secrets = map[string]string{
				mounterutils.GetPasswdFileName(f.Name()): fmt.Sprintf("%s:%s:%s", o.Bucket, o.AkID, o.AkSecret),
			}
			return authCfg, nil
		}

		// secret volume for STS.Token
		if o.SecretRef != "" {
			authCfg.SecretRef = o.SecretRef
			return authCfg, nil
		}

		// republish token retoate for STS.Token
		authCfg.Secrets = map[string]string{
			mounterutils.KeyAccessKeyId:     o.AccessKeyId,
			mounterutils.KeyAccessKeySecret: o.AccessKeySecret,
			mounterutils.KeySecurityToken:   o.SecurityToken,
			mounterutils.KeyExpiration:      o.Expiration,
		}
	}
	return authCfg, nil
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
	// mime.types/csi-mime.types in /etc is used
	etcDirVolume := corev1.Volume{
		Name: "host-etc",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc",
				Type: ptr.To(corev1.HostPathDirectory),
			},
		},
	}

	spec.Volumes = []corev1.Volume{targetDirVolume, metricsDirVolume, etcDirVolume}

	bidirectional := corev1.MountPropagationBidirectional
	socketPath := mounterutils.GetMountProxySocketPath(c.VolumeId)
	container := corev1.Container{
		Name:      f.Name(),
		Image:     f.config.Image,
		Resources: f.config.Resources,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:             targetDirVolume.Name,
				MountPath:        targetDir,
				MountPropagation: &bidirectional,
			}, {
				Name:      metricsDirVolume.Name,
				MountPath: metricsDirVolume.HostPath.Path,
			}, {
				Name:      etcDirVolume.Name,
				MountPath: path.Join(hostPrefix, etcDirVolume.HostPath.Path),
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

func (f *fuseOssfs) MakeMountOptions(o *ossfpm.Options, m metadata.MetadataProvider) (mountOptions []string, err error) {

	region, _ := m.Get(metadata.RegionID)

	mountOptions = append(mountOptions, fmt.Sprintf("url=%s", o.URL))
	if o.ReadOnly {
		mountOptions = append(mountOptions, "ro")
	}

	switch o.Encrypted {
	case ossfpm.EncryptedTypeAes256:
		mountOptions = append(mountOptions, "use_sse")
	case ossfpm.EncryptedTypeKms:
		if o.KmsKeyId == "" {
			mountOptions = append(mountOptions, "use_sse=kmsid")
		} else {
			mountOptions = append(mountOptions, fmt.Sprintf("use_sse=kmsid:%s", o.KmsKeyId))
		}
	}

	if o.MetricsTop != "" {
		mountOptions = append(mountOptions, "use_metrics")
		mountOptions = append(mountOptions, fmt.Sprintf("metrics_top=%s", o.MetricsTop))
	}

	switch o.SigVersion {
	case ossfpm.SigV1:
		mountOptions = append(mountOptions, "sigv1")
	case ossfpm.SigV4:
		if region == "" {
			return nil, fmt.Errorf("SigV4 is not supported without region")
		}
		mountOptions = append(mountOptions, "sigv4")
		mountOptions = append(mountOptions, fmt.Sprintf("region=%s", region))
	}

	authOptions := f.getAuthOptions(o, region)
	mountOptions = append(mountOptions, authOptions...)

	return mountOptions, nil

}

func (f *fuseOssfs) getAuthOptions(o *ossfpm.Options, region string) (mountOptions []string) {
	switch o.AuthType {
	case ossfpm.AuthTypePublic:
		mountOptions = append(mountOptions, "public_bucket=1")
	case ossfpm.AuthTypeRRSA:
		mountOptions = append(mountOptions, fmt.Sprintf("rrsa_endpoint=%s", ossfpm.GetSTSEndpoint(region)))
		if o.AssumeRoleArn != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("assume_role_arn=%s", o.AssumeRoleArn))
			if o.ExternalId != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("assume_role_external_id=%s", o.ExternalId))
			}
		}
	case ossfpm.AuthTypeCSS:
		mountOptions = append(mountOptions, "secret_store_dir=/etc/ossfs/secrets-store")
	case ossfpm.AuthTypeSTS:
		if o.RoleName != "" {
			mountOptions = append(mountOptions, "ram_role="+o.RoleName)
		}
	default:
		// fixed credentials
		if o.AkID != "" && o.AkSecret != "" {
			// it will make passwd_file option in mount-proxy server as it's under a tempdir
			return
		}

		// secret volume for STS.Token
		if o.SecretRef != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("passwd_file=%s", filepath.Join(mounterutils.GetConfigDir(o.FuseType), mounterutils.GetPasswdFileName(o.FuseType))))
		}

		// republish token retoate for STS.Token
		// it will make passwd_file option in mount-proxy server as it's under a tempdir

		// for both STS.Token
		mountOptions = append(mountOptions, "use_session_token")
	}
	return
}

const (
	KeyDbgLevel      = "dbglevel"
	KeyAllowOther    = "allow_other"
	KeyUseMetrics    = "use_metrics"
	KeyMime          = "mime"
	KeyListObjectsV2 = "listobjectsv2"
)

func (f *fuseOssfs) AddDefaultMountOptions(options []string) []string {
	defaultOSSFSOptions := os.Getenv("DEFAULT_OSSFS_OPTIONS")
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

	// set default dbg level
	if _, ok := tm[KeyDbgLevel]; !ok {
		level, ok := ossfsDbglevels[f.config.Dbglevel]
		if ok {
			options = append(options, fmt.Sprintf("dbglevel=%s", level))
		} else {
			if f.config.Dbglevel != "" {
				klog.Warningf("invalid dbglevel for ossfs: %q, use default dbglevel %s", f.config.Dbglevel, defaultOssfsDbglevel)
			}
			options = append(options, fmt.Sprintf("dbglevel=%s", ossfsDbglevels[defaultOssfsDbglevel]))
		}
	}

	// set default allow_other
	if _, ok := tm[KeyAllowOther]; !ok {
		options = append(options, "allow_other")
	}

	// set use_metrics to enabled monitoring by default
	if _, ok := tm[KeyUseMetrics]; !ok {
		if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
			if f.config.MetricsMode != fpm.MetricsModeDisabled {
				options = append(options, "use_metrics")
			}
		}
	}

	// set mime
	if _, ok := tm[KeyMime]; !ok {
		if !utils.IsFileExisting(filepath.Join(hostPrefix, OssfsDefMimeTypesFilePath)) && strings.ToLower(f.config.Extra["mime-support"]) == "true" {
			// mime.types not exists, use csi-mime.types
			options = append(options, fmt.Sprintf("mime=%s", OssfsCsiMimeTypesFilePath))
		}
	}

	// set listobjectsv2
	// Note: OSS officially recommends using v2 API as the preferred version,
	// but ossfs hasn't enabled listobjectv2 by default yet.
	// This is a temporary workaround added by CSI driver.
	// TODO: Remove this logic when ossfs enables it by default.
	if _, ok := tm[KeyListObjectsV2]; !ok {
		options = append(options, "listobjectsv2")
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
	case ossfpm.AuthTypeSTS, ossfpm.AuthTypePublic:
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
	case ossfpm.AuthTypeCSS:
		secretStoreMountDir := "/etc/ossfs/secrets-store"
		secretStoreVolume := corev1.Volume{
			Name: "secrets-store",
			VolumeSource: corev1.VolumeSource{
				CSI: &corev1.CSIVolumeSource{
					Driver:           CsiSecretStoreDriver,
					ReadOnly:         tea.Bool(true),
					VolumeAttributes: map[string]string{SecretProviderClassKey: authCfg.SecretProviderClassName},
				},
			},
		}
		spec.Volumes = append(spec.Volumes, secretStoreVolume)
		secretStoreVolumeMount := corev1.VolumeMount{
			Name:      secretStoreVolume.Name,
			MountPath: secretStoreMountDir,
			ReadOnly:  true,
		}
		container.VolumeMounts = append(container.VolumeMounts, secretStoreVolumeMount)
	default:
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
