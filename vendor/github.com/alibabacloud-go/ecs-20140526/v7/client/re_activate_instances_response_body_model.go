// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReActivateInstancesResponseBody
	GetRequestId() *string
}

type ReActivateInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51AB7717-6E1A-4D1D-A44D-54CB123ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReActivateInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReActivateInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ReActivateInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReActivateInstancesResponseBody) SetRequestId(v string) *ReActivateInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReActivateInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
