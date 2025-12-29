// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintenanceAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMaintenanceAttributesResponseBody
	GetRequestId() *string
}

type ModifyInstanceMaintenanceAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMaintenanceAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintenanceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintenanceAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMaintenanceAttributesResponseBody) SetRequestId(v string) *ModifyInstanceMaintenanceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
