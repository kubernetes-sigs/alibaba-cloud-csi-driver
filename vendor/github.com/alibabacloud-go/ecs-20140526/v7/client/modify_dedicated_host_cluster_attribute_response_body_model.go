// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDedicatedHostClusterAttributeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedHostClusterAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D02FF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedHostClusterAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedHostClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
