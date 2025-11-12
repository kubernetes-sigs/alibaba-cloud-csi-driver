// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVscResponseBody
	GetRequestId() *string
	SetVscId(v string) *CreateVscResponseBody
	GetVscId() *string
}

type CreateVscResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s CreateVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVscResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVscResponseBody) GetVscId() *string {
	return s.VscId
}

func (s *CreateVscResponseBody) SetRequestId(v string) *CreateVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVscResponseBody) SetVscId(v string) *CreateVscResponseBody {
	s.VscId = &v
	return s
}

func (s *CreateVscResponseBody) Validate() error {
	return dara.Validate(s)
}
