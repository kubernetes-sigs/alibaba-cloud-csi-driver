// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmbAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySmbAclResponseBody
	GetRequestId() *string
}

type ModifySmbAclResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySmbAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmbAclResponseBody) SetRequestId(v string) *ModifySmbAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmbAclResponseBody) Validate() error {
	return dara.Validate(s)
}
