// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LeaveSecurityGroupResponseBody
	GetRequestId() *string
}

type LeaveSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LeaveSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LeaveSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LeaveSecurityGroupResponseBody) SetRequestId(v string) *LeaveSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *LeaveSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
