// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlanMaintenanceWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPlanMaintenanceWindowResponseBody
	GetRequestId() *string
}

type ModifyPlanMaintenanceWindowResponseBody struct {
	// example:
	//
	// F3CD6886-D8D0-4FEE-B93E-1B732396****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlanMaintenanceWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPlanMaintenanceWindowResponseBody) SetRequestId(v string) *ModifyPlanMaintenanceWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
