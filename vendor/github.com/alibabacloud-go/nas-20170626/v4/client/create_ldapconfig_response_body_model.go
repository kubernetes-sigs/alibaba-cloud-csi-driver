// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLDAPConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLDAPConfigResponseBody
	GetRequestId() *string
}

type CreateLDAPConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLDAPConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLDAPConfigResponseBody) SetRequestId(v string) *CreateLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLDAPConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
