// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmbAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSmbAclResponseBody
	GetRequestId() *string
}

type DisableSmbAclResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSmbAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSmbAclResponseBody) SetRequestId(v string) *DisableSmbAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSmbAclResponseBody) Validate() error {
	return dara.Validate(s)
}
