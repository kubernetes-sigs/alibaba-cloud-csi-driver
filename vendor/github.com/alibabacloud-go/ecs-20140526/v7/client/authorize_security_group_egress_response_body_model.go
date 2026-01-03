// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupEgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeSecurityGroupEgressResponseBody
	GetRequestId() *string
}

type AuthorizeSecurityGroupEgressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeSecurityGroupEgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeSecurityGroupEgressResponseBody) SetRequestId(v string) *AuthorizeSecurityGroupEgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressResponseBody) Validate() error {
	return dara.Validate(s)
}
