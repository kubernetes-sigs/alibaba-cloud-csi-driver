// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeFileSystemResponseBody
	GetRequestId() *string
}

type UpgradeFileSystemResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeFileSystemResponseBody) SetRequestId(v string) *UpgradeFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
