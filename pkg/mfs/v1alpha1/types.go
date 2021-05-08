package v1alpha1

import (
	"github.com/looplab/fsm"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sync"
)

var GVR = schema.GroupVersionResource{
	Group:    "storage.alibabacloud.com",
	Version:  "v1alpha1",
	Resource: "managedfilesystems",
}

type ManagedFileSystem struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedFileSystemSpec   `json:"spec,omitempty"`
	Status ManagedFileSystemStatus `json:"status,omitempty"`
	mux    sync.Mutex
	FSM    *fsm.FSM
}

type ManagedFileSystemStatus struct {
	Status     string                       `json:"status,omitempty"`
	Conditions []ManagedFileSystemCondition `json:"conditions,omitempty"`
}

type ManagedFileSystemCondition struct {
	LastProbeTime      string `json:"lastProbeTime,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	Status             string `json:"status,omitempty"`
	Reason             string `json:"reason,omitempty"`
}

type ManagedFileSystemSpec struct {
	StorageType   string       `json:"type,omitempty"`
	ReclaimPolicy string       `json:"reclaimPolicy,omitempty"`
	Description   string       `json:"description,omitempty"`
	Parameters    Parameters   `json:"parameters,omitempty"`
	FsAttributes  FsAttributes `json:"fsAttributes,omitempty"`
}

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
	Server          string    `json:"server,,omitempty"`
	BucketName      string    `json:"bucketName,omitempty"`
	EndPoint        *EndPoint `json:"endPoint,omitempty"`
}

type EndPoint struct {
	Internal string `json:"internal,omitempty"`
	Extranet string `json:"extranet,omitempty"`
}

//
type Parameters struct {
	Secret         *Secret `json:"secret,omitempty"`
	StorageType    string  `json:"storageType,omitempty"`
	ProtocolType   string  `json:"protocolType,omitempty"`
	EncryptType    string  `json:"encryptType,omitempty"`
	VSwitchID      string  `json:"vSwitchId,omitempty"`
	FilesystemID   string  `json:"filesystemId,omitempty"`
	Server         string  `json:"server,omitempty"`
	FileSystemType string  `json:"filesystemType,omitempty"`
}

type Secret struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,,omitempty"`
}

type ManagedFileSystemList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedFileSystem `json:"items"`
}


