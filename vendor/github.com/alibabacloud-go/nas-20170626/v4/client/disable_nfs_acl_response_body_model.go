// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNfsAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableNfsAclResponseBody
	GetRequestId() *string
}

type DisableNfsAclResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 24487C24-AE54-57EC-B4E4-4EDEEEB83B01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableNfsAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableNfsAclResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNfsAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableNfsAclResponseBody) SetRequestId(v string) *DisableNfsAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableNfsAclResponseBody) Validate() error {
	return dara.Validate(s)
}
