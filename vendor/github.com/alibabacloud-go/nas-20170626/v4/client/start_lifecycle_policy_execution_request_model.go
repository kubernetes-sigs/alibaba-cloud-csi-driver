// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLifecyclePolicyExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *StartLifecyclePolicyExecutionRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *StartLifecyclePolicyExecutionRequest
	GetLifecyclePolicyId() *string
}

type StartLifecyclePolicyExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-0015****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
}

func (s StartLifecyclePolicyExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLifecyclePolicyExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartLifecyclePolicyExecutionRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StartLifecyclePolicyExecutionRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *StartLifecyclePolicyExecutionRequest) SetFileSystemId(v string) *StartLifecyclePolicyExecutionRequest {
	s.FileSystemId = &v
	return s
}

func (s *StartLifecyclePolicyExecutionRequest) SetLifecyclePolicyId(v string) *StartLifecyclePolicyExecutionRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *StartLifecyclePolicyExecutionRequest) Validate() error {
	return dara.Validate(s)
}
