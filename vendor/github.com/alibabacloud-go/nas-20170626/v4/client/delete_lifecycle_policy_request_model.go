// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *DeleteLifecyclePolicyRequest
	GetLifecyclePolicyId() *string
	SetLifecyclePolicyName(v string) *DeleteLifecyclePolicyRequest
	GetLifecyclePolicyName() *string
}

type DeleteLifecyclePolicyRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the lifecycle policy. Required for the CPFS for Lingjun file systems.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// The name of the lifecycle policy. The name must be 3 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// >  Optional for General-purpose NAS file systems. If this parameter is provided, it takes precedence over LifecyclePolicyId. If left empty, LifecyclePolicyId is used.
	//
	// example:
	//
	// lifecyclepolicy1
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
}

func (s DeleteLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteLifecyclePolicyRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *DeleteLifecyclePolicyRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DeleteLifecyclePolicyRequest) SetFileSystemId(v string) *DeleteLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLifecyclePolicyRequest) SetLifecyclePolicyId(v string) *DeleteLifecyclePolicyRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *DeleteLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *DeleteLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DeleteLifecyclePolicyRequest) Validate() error {
	return dara.Validate(s)
}
