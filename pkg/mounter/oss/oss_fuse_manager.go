package oss

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/client-go/kubernetes"
)

type OSSFuseMounterType interface {
	utils.FuseMounterType
	PrecheckAuthConfig(o *Options, onNode bool) error
	MakeAuthConfig(o *Options, m metadata.MetadataProvider) (*utils.AuthConfig, error)
	MakeMountOptions(o *Options, m metadata.MetadataProvider) ([]string, error)
}

type OSSFusePodManager struct {
	utils.FusePodManager
	OSSFuseMounterType
}

func NewOSSFusePodManager(fuseType OSSFuseMounterType, client kubernetes.Interface) *OSSFusePodManager {
	manager := utils.NewFusePodManager(fuseType, client)
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

// Options contains options for target oss
type Options struct {
	DirectAssigned bool
	CNFSName       string

	// oss options
	Bucket string `json:"bucket"`
	Region string `json:"region"`
	URL    string `json:"url"`
	Path   string `json:"path"`

	// authorization options
	// accesskey
	AkID      string `json:"akId"`
	AkSecret  string `json:"akSecret"`
	SecretRef string `json:"secretRef"`
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
}
