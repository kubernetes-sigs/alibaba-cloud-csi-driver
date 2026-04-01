// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLifecyclePolicyRequest
	GetDescription() *string
	SetFileSystemId(v string) *UpdateLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *UpdateLifecyclePolicyRequest
	GetLifecyclePolicyId() *string
	SetPaths(v []*string) *UpdateLifecyclePolicyRequest
	GetPaths() []*string
	SetRetrieveRules(v []*UpdateLifecyclePolicyRequestRetrieveRules) *UpdateLifecyclePolicyRequest
	GetRetrieveRules() []*UpdateLifecyclePolicyRequestRetrieveRules
	SetStorageType(v string) *UpdateLifecyclePolicyRequest
	GetStorageType() *string
	SetTransitRules(v []*UpdateLifecyclePolicyRequestTransitRules) *UpdateLifecyclePolicyRequest
	GetTransitRules() []*UpdateLifecyclePolicyRequestTransitRules
}

type UpdateLifecyclePolicyRequest struct {
	// example:
	//
	// Lifecycle policy description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64y*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lsp-bp1234567890ab****
	LifecyclePolicyId *string                                      `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	Paths             []*string                                    `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	RetrieveRules     []*UpdateLifecyclePolicyRequestRetrieveRules `json:"RetrieveRules,omitempty" xml:"RetrieveRules,omitempty" type:"Repeated"`
	// example:
	//
	// InfrequentAccess
	StorageType  *string                                     `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	TransitRules []*UpdateLifecyclePolicyRequestTransitRules `json:"TransitRules,omitempty" xml:"TransitRules,omitempty" type:"Repeated"`
}

func (s UpdateLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *UpdateLifecyclePolicyRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *UpdateLifecyclePolicyRequest) GetPaths() []*string {
	return s.Paths
}

func (s *UpdateLifecyclePolicyRequest) GetRetrieveRules() []*UpdateLifecyclePolicyRequestRetrieveRules {
	return s.RetrieveRules
}

func (s *UpdateLifecyclePolicyRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *UpdateLifecyclePolicyRequest) GetTransitRules() []*UpdateLifecyclePolicyRequestTransitRules {
	return s.TransitRules
}

func (s *UpdateLifecyclePolicyRequest) SetDescription(v string) *UpdateLifecyclePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetFileSystemId(v string) *UpdateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetLifecyclePolicyId(v string) *UpdateLifecyclePolicyRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetPaths(v []*string) *UpdateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetRetrieveRules(v []*UpdateLifecyclePolicyRequestRetrieveRules) *UpdateLifecyclePolicyRequest {
	s.RetrieveRules = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetStorageType(v string) *UpdateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetTransitRules(v []*UpdateLifecyclePolicyRequestTransitRules) *UpdateLifecyclePolicyRequest {
	s.TransitRules = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) Validate() error {
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

type UpdateLifecyclePolicyRequestRetrieveRules struct {
	// example:
	//
	// RetrieveType
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// All
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLifecyclePolicyRequestRetrieveRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequestRetrieveRules) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) GetThreshold() *string {
	return s.Threshold
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) SetAttribute(v string) *UpdateLifecyclePolicyRequestRetrieveRules {
	s.Attribute = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) SetThreshold(v string) *UpdateLifecyclePolicyRequestRetrieveRules {
	s.Threshold = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) Validate() error {
	return dara.Validate(s)
}

type UpdateLifecyclePolicyRequestTransitRules struct {
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// 3
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLifecyclePolicyRequestTransitRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequestTransitRules) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequestTransitRules) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateLifecyclePolicyRequestTransitRules) GetThreshold() *string {
	return s.Threshold
}

func (s *UpdateLifecyclePolicyRequestTransitRules) SetAttribute(v string) *UpdateLifecyclePolicyRequestTransitRules {
	s.Attribute = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestTransitRules) SetThreshold(v string) *UpdateLifecyclePolicyRequestTransitRules {
	s.Threshold = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestTransitRules) Validate() error {
	return dara.Validate(s)
}
