// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateAccessGroupResponseBody
	GetAccessGroupName() *string
	SetRequestId(v string) *CreateAccessGroupResponseBody
	GetRequestId() *string
}

type CreateAccessGroupResponseBody struct {
	// The name of the permission group.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189F4F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponseBody) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateAccessGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessGroupResponseBody) SetAccessGroupName(v string) *CreateAccessGroupResponseBody {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupResponseBody) SetRequestId(v string) *CreateAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
