package oss

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type OSSFuseMounterType interface {
	fpm.FuseMounterType
	PrecheckAuthConfig(o *Options, onNode bool) error
	MakeAuthConfig(o *Options, m metadata.MetadataProvider) (*fpm.AuthConfig, error)
	MakeMountOptions(o *Options, m metadata.MetadataProvider) ([]string, error)
}

type OSSFusePodManager struct {
	fpm.FusePodManager
	OSSFuseMounterType
}

func NewOSSFusePodManager(fuseType OSSFuseMounterType, client kubernetes.Interface) *OSSFusePodManager {
	manager := fpm.NewFusePodManager(fuseType, client)
	return &OSSFusePodManager{
		*manager,
		fuseType,
	}
}

// SigVersion is the version of signature for ossfs
type SigVersion string

const (
	// default version for ossfs
	SigV1 SigVersion = "v1"
	// need set region
	SigV4 SigVersion = "v4"
)

const (
	EncryptedTypeKms    = "kms"
	EncryptedTypeAes256 = "aes256"
)

const (
	AuthTypeSTS    = "sts"
	AuthTypeRRSA   = "rrsa"
	AuthTypeCSS    = "csi-secret-store"
	AuthTypePublic = "public"
)

type AccessKey struct {
	AkID     string `json:"akId"`
	AkSecret string `json:"akSecret"`
}
type TokenSecret struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Expiration      string `json:"expiration"`
	SecurityToken   string `json:"securityToken"`
}

// Options contains options for target oss
type Options struct {
	// DirectAssigned indicates the volume should be directly assigned to the pod.
	// When DirectAssigned is True, it means the runtime is either rund or coco:
	// - Controller server should skip the controller publish phase
	// - Node server should check the skipAttach field to determine if it's rund (skipAttach=true) or coco (skipAttach=false)
	// Otherwise, the runtime is runc
	DirectAssigned bool
	CNFSName       string

	// oss options
	Bucket string `json:"bucket"`
	URL    string `json:"url"`
	Path   string `json:"path"`

	// authorization options
	// accesskey
	AccessKey   `json:",inline"`
	TokenSecret `json:",inline"`
	SecretRef   string `json:"secretRef"`
	// RRSA
	RoleName           string `json:"roleName"` // also for STS
	RoleArn            string `json:"roleArn"`
	OidcProviderArn    string `json:"oidcProviderArn"`
	ServiceAccountName string `json:"serviceAccountName"`
	// assume role
	AssumeRoleArn string `json:"assumeRoleArn"`
	ExternalId    string `json:"externalId"`
	// csi secret store
	SecretProviderClass string `json:"secretProviderClass"`

	// ossfs options
	OtherOpts  string     `json:"otherOpts"`
	MetricsTop string     `json:"metricsTop"`
	Encrypted  string     `json:"encrypted"`
	KmsKeyId   string     `json:"kmsKeyId"`
	SigVersion SigVersion `json:"sigVersion"`

	// mount options
	UseSharedPath bool   `json:"useSharedPath"`
	AuthType      string `json:"authType"`
	FuseType      string `json:"fuseType"`
	ReadOnly      bool   `json:"readOnly"`

	// pod template
	DnsPolicy corev1.DNSPolicy `json:"dnsPolicy"`
}
