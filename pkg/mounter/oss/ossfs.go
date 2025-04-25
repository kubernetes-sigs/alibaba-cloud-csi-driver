package oss

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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	csiutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

var defaultOssfsImageTag = "v1.88.4-80d165c-aliyun"
var defaultOssfsUpdatedImageTag = "v1.91.5.ack.1-ed398f6-aliyun"
var defaultOssfsDbglevel = utils.DebugLevelError

const (
	hostPrefix                = "/host"
	OssfsDefMimeTypesFilePath = "/etc/mime.types"
	// The ossfs image includes a default MIME configuration located at /csi/mime.types
	OssfsCsiMimeTypesFilePath = "/csi/mime.types"

	defaultRegistry = "registry-cn-hangzhou.ack.aliyuncs.com"

	CsiSecretStoreDriver   = "secrets-store.csi.k8s.io"
	SecretProviderClassKey = "secretProviderClass"
)

type fuseOssfs struct {
	config utils.FuseContainerConfig
}

var ossfsDbglevels = map[string]string{
	utils.DebugLevelDebug: "debug",
	utils.DebugLevelInfo:  "info",
	utils.DebugLevelWarn:  "warn",
	utils.DebugLevelError: "err",
	utils.DebugLevelFatal: "crit",
}

func NewFuseOssfs(configmap *corev1.ConfigMap, m metadata.MetadataProvider) OSSFuseMounterType {
	config := utils.ExtractFuseContainerConfig(configmap, OssFsType)

	// set default image
	setDefaultImage(OssFsType, m, &config)
	// set default memory request
	if _, ok := config.Resources.Requests[corev1.ResourceMemory]; !ok {
		config.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("50Mi")
	}

	return &fuseOssfs{config: config}
}

func (f *fuseOssfs) Name() string {
	return OssFsType
}

func (f *fuseOssfs) PrecheckAuthConfig(o *Options, onNode bool) error {

	if o.AuthType != AuthTypeRRSA && o.AssumeRoleArn != "" {
		return fmt.Errorf("only support access OSS through STS AssumeRole when authType is RRSA")
	}

	switch o.AuthType {
	case AuthTypePublic:
	case AuthTypeSTS:
		// rolename may retrieve from metadata service
		if onNode && o.RoleName == "" {
			return fmt.Errorf("missing roleName or ramRole in volume attributes")
		}
	case AuthTypeRRSA:
		if err := checkRRSAParams(o); err != nil {
			return err
		}
	case AuthTypeCSS:
		if o.SecretProviderClass == "" {
			return fmt.Errorf("use CsiSecretStore but secretProviderClass is empty")
		}
	default:
		if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
			return nil
		}
		if o.SecretRef != "" {
			if o.AkID != "" || o.AkSecret != "" {
				return fmt.Errorf("AK and secretRef cannot be set at the same time")
			}
			if o.SecretRef == utils.GetCredientialsSecretName(OssFsType) {
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

func (f *fuseOssfs) MakeAuthConfig(o *Options, m metadata.MetadataProvider) (*utils.AuthConfig, error) {
	authCfg := &utils.AuthConfig{AuthType: o.AuthType}
	switch o.AuthType {
	case AuthTypePublic:
	case AuthTypeRRSA:
		rrsaCfg, err := getRRSAConfig(o, m)
		if err != nil {
			return nil, fmt.Errorf("Get RoleArn and OidcProviderArn for RRSA error: %v", err)
		}
		authCfg.RrsaConfig = rrsaCfg
	case AuthTypeCSS:
		authCfg.SecretProviderClassName = o.SecretProviderClass
	case AuthTypeSTS:
		authCfg.RoleName = o.RoleName
	default:
		if o.SecretRef != "" {
			authCfg.SecretRef = o.SecretRef
		} else {
			authCfg.Secrets = map[string]string{
				utils.GetPasswdFileName(f.Name()): fmt.Sprintf("%s:%s:%s", o.Bucket, o.AkID, o.AkSecret),
			}
		}
	}
	return authCfg, nil
}

func (f *fuseOssfs) PodTemplateSpec(c *utils.FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
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

func (f *fuseOssfs) buildPodSpec(c *utils.FusePodContext, target string) (spec corev1.PodSpec, _ error) {
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
	metricsDirVolume := corev1.Volume{
		Name: "metrics-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/run/ossfs",
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
	socketPath := utils.GetMountProxySocketPath(c.VolumeId)
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

	buildOssfsAuthSpec(c, target, &spec, &container)

	container.Args = []string{"--socket=" + socketPath, "-v=4"}

	spec.Containers = []corev1.Container{container}
	spec.NodeName = c.NodeName
	spec.HostNetwork = true
	spec.PriorityClassName = "system-node-critical"
	spec.Tolerations = []corev1.Toleration{{Operator: corev1.TolerationOpExists}}
	return
}

func (f *fuseOssfs) MakeMountOptions(o *Options, m metadata.MetadataProvider) (mountOptions []string, err error) {

	region, _ := m.Get(metadata.RegionID)

	mountOptions = append(mountOptions, fmt.Sprintf("url=%s", o.URL))
	if o.ReadOnly {
		mountOptions = append(mountOptions, "ro")
	}

	switch o.Encrypted {
	case EncryptedTypeAes256:
		mountOptions = append(mountOptions, "use_sse")
	case EncryptedTypeKms:
		if o.KmsKeyId == "" {
			mountOptions = append(mountOptions, "use_sse=kmsid")
		} else {
			mountOptions = append(mountOptions, fmt.Sprintf("use_sse=kmsid:%s", o.KmsKeyId))
		}
	}

	// set use_metrics to enabled monitoring by default
	if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
		mountOptions = append(mountOptions, "use_metrics")
	}
	if o.MetricsTop != "" {
		mountOptions = append(mountOptions, fmt.Sprintf("metrics_top=%s", o.MetricsTop))
	}

	switch o.SigVersion {
	case SigV1:
		mountOptions = append(mountOptions, "sigv1")
	case SigV4:
		if o.Region == "" {
			return nil, fmt.Errorf("SigV4 is not supported without region")
		}
		mountOptions = append(mountOptions, "sigv4")
		mountOptions = append(mountOptions, fmt.Sprintf("region=%s", o.Region))
	}

	authOptions := o.getAuthOptions(region)
	mountOptions = append(mountOptions, authOptions...)

	return mountOptions, nil

}

func (o *Options) getAuthOptions(region string) (mountOptions []string) {
	switch o.AuthType {
	case AuthTypePublic:
		mountOptions = append(mountOptions, "public_bucket=1")
	case AuthTypeRRSA:
		mountOptions = append(mountOptions, fmt.Sprintf("rrsa_endpoint=%s", getSTSEndpoint(region)))
		if o.AssumeRoleArn != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("assume_role_arn=%s", o.AssumeRoleArn))
			if o.ExternalId != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("assume_role_external_id=%s", o.ExternalId))
			}
		}
	case AuthTypeCSS:
		mountOptions = append(mountOptions, "secret_store_dir=/etc/ossfs/secrets-store")
	case AuthTypeSTS:
		if o.RoleName != "" {
			mountOptions = append(mountOptions, "ram_role="+o.RoleName)
		}
	default:
		if o.SecretRef != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("passwd_file=%s", filepath.Join(utils.GetConfigDir(o.FuseType), utils.GetPasswdFileName(o.FuseType))))
			mountOptions = append(mountOptions, "use_session_token")
		}
		// publishSecretRef will make option in mount-proxy server
	}
	return
}

func (f *fuseOssfs) AddDefaultMountOptions(options []string) []string {
	alreadySet := false
	for _, option := range options {
		if strings.Contains(option, "dbglevel") {
			alreadySet = true
			break
		}
	}
	if !alreadySet {
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

	if !csiutils.IsFileExisting(filepath.Join(hostPrefix, OssfsDefMimeTypesFilePath)) && strings.ToLower(f.config.Extra["mime-support"]) == "true" {
		// mime.types not exists, use csi-mime.types
		options = append(options, fmt.Sprintf("mime=%s", OssfsCsiMimeTypesFilePath))
	}

	defaultOSSFSOptions := os.Getenv("DEFAULT_OSSFS_OPTIONS")
	if defaultOSSFSOptions != "" {
		options = append(options, strings.Split(defaultOSSFSOptions, ",")...)
	}
	return options
}

func buildOssfsAuthSpec(c *utils.FusePodContext, target string, spec *corev1.PodSpec, container *corev1.Container) {
	if spec == nil || container == nil {
		return
	}
	authCfg := c.AuthConfig
	if authCfg == nil {
		return
	}

	switch authCfg.AuthType {
	case AuthTypeSTS, AuthTypePublic:
	case AuthTypeRRSA:
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
				Value: utils.GetRoleSessionName(c.VolumeId, target, c.FuseType),
			},
		}
		container.Env = append(container.Env, envs...)
	case AuthTypeCSS:
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
		secretVolumeSource := getPasswdSecretVolume(authCfg.SecretRef)
		if secretVolumeSource != nil {
			passwdSecretVolume := corev1.Volume{
				Name: "passwd-ossfs",
				VolumeSource: corev1.VolumeSource{
					Secret: secretVolumeSource,
				},
			}
			spec.Volumes = append(spec.Volumes, passwdSecretVolume)
			passwdVolumeMont := corev1.VolumeMount{
				Name:      passwdSecretVolume.Name,
				MountPath: utils.GetConfigDir(c.FuseType),
			}
			container.VolumeMounts = append(container.VolumeMounts, passwdVolumeMont)
		}
	}
}

// keep consistent with RAM response
var secretRefKeysToParse []string = []string{
	"AccessKeyId",
	"AccessKeySecret",
	"Expiration",
	"SecurityToken",
}

func getPasswdSecretVolume(secretRef string) (secret *corev1.SecretVolumeSource) {
	passwdFilename := "passwd-ossfs"
	if secretRef == "" {
		return nil
	}
	items := []corev1.KeyToPath{}
	for _, key := range secretRefKeysToParse {
		item := corev1.KeyToPath{
			Key:  key,
			Path: fmt.Sprintf("%s/%s", passwdFilename, key),
			Mode: tea.Int32(0600),
		}
		items = append(items, item)
	}
	secret = &corev1.SecretVolumeSource{
		SecretName: secretRef,
		Items:      items,
	}
	return
}
