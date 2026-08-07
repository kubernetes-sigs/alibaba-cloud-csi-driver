// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLifecyclePolicyRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest
	GetLifecyclePolicyName() *string
	SetLifecyclePolicyType(v string) *CreateLifecyclePolicyRequest
	GetLifecyclePolicyType() *string
	SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest
	GetLifecycleRuleName() *string
	SetPath(v string) *CreateLifecyclePolicyRequest
	GetPath() *string
	SetPaths(v []*string) *CreateLifecyclePolicyRequest
	GetPaths() []*string
	SetRetrieveRules(v []*CreateLifecyclePolicyRequestRetrieveRules) *CreateLifecyclePolicyRequest
	GetRetrieveRules() []*CreateLifecyclePolicyRequestRetrieveRules
	SetStorageType(v string) *CreateLifecyclePolicyRequest
	GetStorageType() *string
	SetTransitRules(v []*CreateLifecyclePolicyRequestTransitRules) *CreateLifecyclePolicyRequest
	GetTransitRules() []*CreateLifecyclePolicyRequestTransitRules
}

type CreateLifecyclePolicyRequest struct {
	// The description of the lifecycle policy.
	//
	// Format:
	//
	// The description must be 3 to 64 characters in length, start with a letter, and can contain letters, digits, underscores (_), or hyphens (-).
	//
	// >Only CPFS for Lingjun is supported.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The Policy Name of the lifecycle management policy. The name must be 3 to 64 characters in length, start with an uppercase letter or lowercase letter, and can contain letters, digits, underscores (_), or hyphens (-).
	//
	// >This parameter is required for General-purpose NAS but not required for CPFS for Lingjun.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The policy type.
	//
	// - Auto (default): automatic execution.
	//
	// - OnDemand: on-demand execution.
	//
	// example:
	//
	// Auto
	LifecyclePolicyType *string `json:"LifecyclePolicyType,omitempty" xml:"LifecyclePolicyType,omitempty"`
	// The management rule associated with the lifecycle management policy. Only General-purpose NAS is supported.
	//
	// Valid values:
	//
	// - DEFAULT_ATIME_14: files that have not been accessed for 14 days.
	//
	// - DEFAULT_ATIME_30: files that have not been accessed for 30 days.
	//
	// - DEFAULT_ATIME_60: files that have not been accessed for 60 days.
	//
	// - DEFAULT_ATIME_90: files that have not been accessed for 90 days.
	//
	// - DEFAULT_ATIME_180: files that have not been accessed for 180 days. DEFAULT_ATIME_180 is supported only when StorageType is set to Archive.
	//
	// >- If an IA storage class policy has already been configured for the directory, the archive policy duration must be longer than the IA storage class policy duration.
	//
	// > - Only General-purpose NAS supports this parameter.
	//
	// example:
	//
	// DEFAULT_ATIME_14
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of the directory associated with the lifecycle management policy. Only General-purpose NAS is supported.
	//
	// - General-purpose NAS supports associating only a single directory. The path must start with a forward slash (/) and must be an existing path in the mount target.
	//
	// > Only General-purpose NAS is supported. For General-purpose NAS, use Paths.N to associate multiple directories at the same time.
	//
	// >- Only one of Path and Paths can be specified.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The absolute paths of directories associated with the lifecycle management policy.
	//
	// example:
	//
	// "/path1", "/path2"
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The file data retrieval rules. A maximum of one rule can be configured.
	//
	// >Only CPFS for Lingjun file systems are supported.
	//
	// >When LifecyclePolicyType is set to OnDemand, at least one of TransitRules or RetrieveRules must be specified.
	RetrieveRules []*CreateLifecyclePolicyRequestRetrieveRules `json:"RetrieveRules,omitempty" xml:"RetrieveRules,omitempty" type:"Repeated"`
	// The storage type.
	//
	// - InfrequentAccess: IA storage class.
	//
	// - Archive: Archive storage.
	//
	// >General-purpose NAS supports InfrequentAccess and Archive. CPFS for Lingjun supports only InfrequentAccess.
	//
	// This parameter is required.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The file data transit rules. A maximum of one rule can be configured.
	//
	// >Supported only when LifecyclePolicyType is set to Auto for CPFS for Lingjun file systems.
	TransitRules []*CreateLifecyclePolicyRequestTransitRules `json:"TransitRules,omitempty" xml:"TransitRules,omitempty" type:"Repeated"`
}

func (s CreateLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateLifecyclePolicyRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *CreateLifecyclePolicyRequest) GetLifecyclePolicyType() *string {
	return s.LifecyclePolicyType
}

func (s *CreateLifecyclePolicyRequest) GetLifecycleRuleName() *string {
	return s.LifecycleRuleName
}

func (s *CreateLifecyclePolicyRequest) GetPath() *string {
	return s.Path
}

func (s *CreateLifecyclePolicyRequest) GetPaths() []*string {
	return s.Paths
}

func (s *CreateLifecyclePolicyRequest) GetRetrieveRules() []*CreateLifecyclePolicyRequestRetrieveRules {
	return s.RetrieveRules
}

func (s *CreateLifecyclePolicyRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateLifecyclePolicyRequest) GetTransitRules() []*CreateLifecyclePolicyRequestTransitRules {
	return s.TransitRules
}

func (s *CreateLifecyclePolicyRequest) SetDescription(v string) *CreateLifecyclePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetFileSystemId(v string) *CreateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecyclePolicyType(v string) *CreateLifecyclePolicyRequest {
	s.LifecyclePolicyType = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPath(v string) *CreateLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPaths(v []*string) *CreateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetRetrieveRules(v []*CreateLifecyclePolicyRequestRetrieveRules) *CreateLifecyclePolicyRequest {
	s.RetrieveRules = v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetStorageType(v string) *CreateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetTransitRules(v []*CreateLifecyclePolicyRequestTransitRules) *CreateLifecyclePolicyRequest {
	s.TransitRules = v
	return s
}

func (s *CreateLifecyclePolicyRequest) Validate() error {
	if s.RetrieveRules != nil {
		for _, item := range s.RetrieveRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransitRules != nil {
		for _, item := range s.TransitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLifecyclePolicyRequestRetrieveRules struct {
	// The rule attribute. Valid values:
	//
	// - RetrieveType: the retrieval method.
	//
	// example:
	//
	// RetrieveType
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The rule threshold. Valid values:
	//
	// - RetrieveType
	//
	//     - AfterVisit: supported when LifecyclePolicyType is set to Auto. Indicates best-effort recall on visit.
	//
	//     - All: supported when LifecyclePolicyType is set to OnDemand. Indicates retrieving all data.
	//
	// example:
	//
	// All
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateLifecyclePolicyRequestRetrieveRules) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyRequestRetrieveRules) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequestRetrieveRules) GetAttribute() *string {
	return s.Attribute
}

func (s *CreateLifecyclePolicyRequestRetrieveRules) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateLifecyclePolicyRequestRetrieveRules) SetAttribute(v string) *CreateLifecyclePolicyRequestRetrieveRules {
	s.Attribute = &v
	return s
}

func (s *CreateLifecyclePolicyRequestRetrieveRules) SetThreshold(v string) *CreateLifecyclePolicyRequestRetrieveRules {
	s.Threshold = &v
	return s
}

func (s *CreateLifecyclePolicyRequestRetrieveRules) Validate() error {
	return dara.Validate(s)
}

type CreateLifecyclePolicyRequestTransitRules struct {
	// The rule attribute.
	//
	// Valid values:
	//
	// - Atime: the access time of the file.
	//
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The rule threshold.
	//
	// Valid values:
	//
	// - When Attribute is set to Atime, the value indicates the number of days since the file was last accessed. Valid values: 0 to 365.
	//
	// example:
	//
	// 3
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateLifecyclePolicyRequestTransitRules) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyRequestTransitRules) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequestTransitRules) GetAttribute() *string {
	return s.Attribute
}

func (s *CreateLifecyclePolicyRequestTransitRules) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateLifecyclePolicyRequestTransitRules) SetAttribute(v string) *CreateLifecyclePolicyRequestTransitRules {
	s.Attribute = &v
	return s
}

func (s *CreateLifecyclePolicyRequestTransitRules) SetThreshold(v string) *CreateLifecyclePolicyRequestTransitRules {
	s.Threshold = &v
	return s
}

func (s *CreateLifecyclePolicyRequestTransitRules) Validate() error {
	return dara.Validate(s)
}
