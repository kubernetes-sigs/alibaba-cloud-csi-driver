// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDedicatedHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseDedicatedHostResponseBody
	GetRequestId() *string
}

type ReleaseDedicatedHostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B15AC8-E6F6-49A4-8985-8C07104B9199
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseDedicatedHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseDedicatedHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseDedicatedHostResponseBody) SetRequestId(v string) *ReleaseDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseDedicatedHostResponseBody) Validate() error {
	return dara.Validate(s)
}
