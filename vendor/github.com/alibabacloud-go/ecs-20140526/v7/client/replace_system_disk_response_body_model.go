// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSystemDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ReplaceSystemDiskResponseBody
	GetDiskId() *string
	SetRequestId(v string) *ReplaceSystemDiskResponseBody
	GetRequestId() *string
}

type ReplaceSystemDiskResponseBody struct {
	// The ID of the new system disk.
	//
	// example:
	//
	// d-bp67acfmxazb4ph****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceSystemDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSystemDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceSystemDiskResponseBody) GetDiskId() *string {
	return s.DiskId
}

func (s *ReplaceSystemDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceSystemDiskResponseBody) SetDiskId(v string) *ReplaceSystemDiskResponseBody {
	s.DiskId = &v
	return s
}

func (s *ReplaceSystemDiskResponseBody) SetRequestId(v string) *ReplaceSystemDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceSystemDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
