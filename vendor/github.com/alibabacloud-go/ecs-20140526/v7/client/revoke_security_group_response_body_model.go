// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeSecurityGroupResponseBody
	GetRequestId() *string
}

type RevokeSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeSecurityGroupResponseBody) SetRequestId(v string) *RevokeSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
