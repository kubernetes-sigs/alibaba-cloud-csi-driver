// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLifecyclePolicyExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *StopLifecyclePolicyExecutionRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *StopLifecyclePolicyExecutionRequest
	GetLifecyclePolicyId() *string
}

type StopLifecyclePolicyExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
}

func (s StopLifecyclePolicyExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLifecyclePolicyExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopLifecyclePolicyExecutionRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StopLifecyclePolicyExecutionRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *StopLifecyclePolicyExecutionRequest) SetFileSystemId(v string) *StopLifecyclePolicyExecutionRequest {
	s.FileSystemId = &v
	return s
}

func (s *StopLifecyclePolicyExecutionRequest) SetLifecyclePolicyId(v string) *StopLifecyclePolicyExecutionRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *StopLifecyclePolicyExecutionRequest) Validate() error {
	return dara.Validate(s)
}
