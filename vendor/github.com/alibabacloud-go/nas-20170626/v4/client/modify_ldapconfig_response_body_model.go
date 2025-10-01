// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLDAPConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLDAPConfigResponseBody
	GetRequestId() *string
}

type ModifyLDAPConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529DA963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLDAPConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLDAPConfigResponseBody) SetRequestId(v string) *ModifyLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLDAPConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
