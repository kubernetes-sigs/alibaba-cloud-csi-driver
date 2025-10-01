// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLDAPConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLDAPConfigResponseBody
	GetRequestId() *string
}

type DeleteLDAPConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529DA963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLDAPConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLDAPConfigResponseBody) SetRequestId(v string) *DeleteLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLDAPConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
