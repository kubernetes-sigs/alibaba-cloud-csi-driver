// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhysicalConnectionId(v string) *CreatePhysicalConnectionResponseBody
	GetPhysicalConnectionId() *string
	SetRequestId(v string) *CreatePhysicalConnectionResponseBody
	GetRequestId() *string
}

type CreatePhysicalConnectionResponseBody struct {
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionResponseBody) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreatePhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhysicalConnectionResponseBody) SetPhysicalConnectionId(v string) *CreatePhysicalConnectionResponseBody {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionResponseBody) SetRequestId(v string) *CreatePhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
