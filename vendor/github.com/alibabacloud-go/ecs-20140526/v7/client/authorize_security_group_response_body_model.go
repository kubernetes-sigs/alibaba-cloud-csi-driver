// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeSecurityGroupResponseBody
	GetRequestId() *string
}

type AuthorizeSecurityGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeSecurityGroupResponseBody) SetRequestId(v string) *AuthorizeSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
