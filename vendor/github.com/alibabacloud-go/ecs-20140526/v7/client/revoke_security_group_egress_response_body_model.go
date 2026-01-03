// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupEgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeSecurityGroupEgressResponseBody
	GetRequestId() *string
}

type RevokeSecurityGroupEgressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSecurityGroupEgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeSecurityGroupEgressResponseBody) SetRequestId(v string) *RevokeSecurityGroupEgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeSecurityGroupEgressResponseBody) Validate() error {
	return dara.Validate(s)
}
