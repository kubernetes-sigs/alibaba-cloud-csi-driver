// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *CreateDiskResponseBody
	GetDiskId() *string
	SetOrderId(v string) *CreateDiskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDiskResponseBody
	GetRequestId() *string
}

type CreateDiskResponseBody struct {
	// The disk ID.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The order ID.
	//
	// > The order ID is returned only when you create a subscription disk.
	//
	// example:
	//
	// 20413515388****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskResponseBody) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateDiskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiskResponseBody) SetDiskId(v string) *CreateDiskResponseBody {
	s.DiskId = &v
	return s
}

func (s *CreateDiskResponseBody) SetOrderId(v string) *CreateDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskResponseBody) SetRequestId(v string) *CreateDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
