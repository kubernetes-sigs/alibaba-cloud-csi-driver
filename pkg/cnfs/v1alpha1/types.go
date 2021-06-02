package v1alpha1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//GVR is cnfs version
var GVR = schema.GroupVersionResource{
	Group:    "storage.alibabacloud.com",
	Version:  "v1alpha1",
	Resource: "containernetworkfilesystems",
}

//ContainerNetworkFileSystem define cnfs crd
type ContainerNetworkFileSystem struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerNetworkFileSystemSpec   `json:"spec,omitempty"`
	Status ContainerNetworkFileSystemStatus `json:"status,omitempty"`
}

//ContainerNetworkFileSystemStatus define cnfs status field
type ContainerNetworkFileSystemStatus struct {
	Status       string                                `json:"status,omitempty"`
	FsAttributes FsAttributes                          `json:"fsAttributes,omitempty"`
	Conditions   []ContainerNetworkFileSystemCondition `json:"conditions,omitempty"`
}

//ContainerNetworkFileSystemCondition define cnfs condition field
type ContainerNetworkFileSystemCondition struct {
	LastProbeTime string `json:"lastProbeTime,omitempty"`
	Status        string `json:"status,omitempty"`
	Reason        string `json:"reason,omitempty"`
}

//ContainerNetworkFileSystemSpec define cnfs spec field
type ContainerNetworkFileSystemSpec struct {
	StorageType   string     `json:"type,omitempty"`
	ReclaimPolicy string     `json:"reclaimPolicy,omitempty"`
	Description   string     `json:"description,omitempty"`
	Parameters    Parameters `json:"parameters,omitempty"`
}

//FsAttributes define cnfs status FsAttributes field
type FsAttributes struct {
	RegionID        string    `json:"regionId,omitempty"`
	ZoneID          string    `json:"zoneId,omitempty"`
	StorageType     string    `json:"storageType,omitempty"`
	ProtocolType    string    `json:"protocolType,omitempty"`
	EncryptType     string    `json:"encryptType,omitempty"`
	AccessGroupName string    `json:"accessGroupName,omitempty"`
	VpcID           string    `json:"vpcId,omitempty"`
	VSwitchID       string    `json:"vSwitchId,omitempty"`
	FilesystemID    string    `json:"filesystemId,omitempty"`
	FilesystemType  string    `json:"filesystemType,omitempty"`
	Server          string    `json:"server,omitempty"`
	BucketName      string    `json:"bucketName,omitempty"`
	EndPoint        *EndPoint `json:"endPoint,omitempty"`
	ACLType         string    `json:"aclType,omitempty"`
	EnableTrashCan  string    `json:"enableTrashCan,omitempty"`
}

//EndPoint define cnfs endpoint field when cnfs type is oss
type EndPoint struct {
	Internal string `json:"internal,omitempty"`
	Extranet string `json:"extranet,omitempty"`
}

//Parameters define cnfs parameters field
type Parameters struct {
	Secret         *Secret `json:"secret,omitempty"`
	StorageType    string  `json:"storageType,omitempty"`
	ProtocolType   string  `json:"protocolType,omitempty"`
	EncryptType    string  `json:"encryptType,omitempty"`
	VSwitchID      string  `json:"vSwitchId,omitempty"`
	Server         string  `json:"server,omitempty"`
	FileSystemType string  `json:"filesystemType,omitempty"`
	BucketName     string  `json:"bucketName,omitempty"`
	ACLType        string  `json:"aclType,omitempty"`
	EnableTrashCan string  `json:"enableTrashCan,omitempty"`
}

//Secret define secret field
type Secret struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,,omitempty"`
}

//ContainerNetworkFileSystemList define cnfs list
type ContainerNetworkFileSystemList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []ContainerNetworkFileSystem `json:"items"`
}
