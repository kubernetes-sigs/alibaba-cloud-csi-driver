// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityGroupResponseBody
	GetRequestId() *string
}

type DeleteSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityGroupResponseBody) SetRequestId(v string) *DeleteSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
