// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlanMaintenanceWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanWindowId(v string) *CreatePlanMaintenanceWindowResponseBody
	GetPlanWindowId() *string
	SetRequestId(v string) *CreatePlanMaintenanceWindowResponseBody
	GetRequestId() *string
}

type CreatePlanMaintenanceWindowResponseBody struct {
	// example:
	//
	// pw-bp1a9yavgq3dgttvowun
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// example:
	//
	// 7D5B1188-3F08-56D1-A6B2-91B267452633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePlanMaintenanceWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowResponseBody) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *CreatePlanMaintenanceWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePlanMaintenanceWindowResponseBody) SetPlanWindowId(v string) *CreatePlanMaintenanceWindowResponseBody {
	s.PlanWindowId = &v
	return s
}

func (s *CreatePlanMaintenanceWindowResponseBody) SetRequestId(v string) *CreatePlanMaintenanceWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePlanMaintenanceWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
