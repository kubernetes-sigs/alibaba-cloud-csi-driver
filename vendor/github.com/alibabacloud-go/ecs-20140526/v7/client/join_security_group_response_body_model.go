// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinSecurityGroupResponseBody
	GetRequestId() *string
}

type JoinSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinSecurityGroupResponseBody) SetRequestId(v string) *JoinSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
