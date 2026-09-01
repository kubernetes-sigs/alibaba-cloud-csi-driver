// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataInsightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DisableDataInsightRequest
	GetFileSystemId() *string
}

type DisableDataInsightRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64y*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableDataInsightRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDataInsightRequest) GoString() string {
	return s.String()
}

func (s *DisableDataInsightRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DisableDataInsightRequest) SetFileSystemId(v string) *DisableDataInsightRequest {
	s.FileSystemId = &v
	return s
}

func (s *DisableDataInsightRequest) Validate() error {
	return dara.Validate(s)
}
