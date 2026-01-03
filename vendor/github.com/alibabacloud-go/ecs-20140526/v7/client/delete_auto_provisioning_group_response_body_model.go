// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoProvisioningGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoProvisioningGroupResponseBody
	GetRequestId() *string
}

type DeleteAutoProvisioningGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B48A12CD-1295-4A38-A8F0-0E92C937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoProvisioningGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoProvisioningGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoProvisioningGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoProvisioningGroupResponseBody) SetRequestId(v string) *DeleteAutoProvisioningGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoProvisioningGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
