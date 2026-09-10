// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCpfsAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreateCpfsAccessPointResponseBody
	GetAccessPointId() *string
	SetRequestId(v string) *CreateCpfsAccessPointResponseBody
	GetRequestId() *string
}

type CreateCpfsAccessPointResponseBody struct {
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCpfsAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCpfsAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCpfsAccessPointResponseBody) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreateCpfsAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCpfsAccessPointResponseBody) SetAccessPointId(v string) *CreateCpfsAccessPointResponseBody {
	s.AccessPointId = &v
	return s
}

func (s *CreateCpfsAccessPointResponseBody) SetRequestId(v string) *CreateCpfsAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCpfsAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
