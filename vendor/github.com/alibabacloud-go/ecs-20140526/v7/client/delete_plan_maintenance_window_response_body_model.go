// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlanMaintenanceWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePlanMaintenanceWindowResponseBody
	GetRequestId() *string
}

type DeletePlanMaintenanceWindowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePlanMaintenanceWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePlanMaintenanceWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlanMaintenanceWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePlanMaintenanceWindowResponseBody) SetRequestId(v string) *DeletePlanMaintenanceWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePlanMaintenanceWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
