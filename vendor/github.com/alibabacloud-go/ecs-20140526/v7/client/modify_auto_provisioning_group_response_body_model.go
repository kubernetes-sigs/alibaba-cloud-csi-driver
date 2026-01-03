// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoProvisioningGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoProvisioningGroupResponseBody
	GetRequestId() *string
}

type ModifyAutoProvisioningGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B48A12CD-1295-4A38-A8F0-0E92C937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoProvisioningGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoProvisioningGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoProvisioningGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoProvisioningGroupResponseBody) SetRequestId(v string) *ModifyAutoProvisioningGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
