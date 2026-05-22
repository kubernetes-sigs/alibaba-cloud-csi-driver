// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlanMaintenanceWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPlanMaintenanceWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPlanMaintenanceWindowResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPlanMaintenanceWindowResponseBody) *ModifyPlanMaintenanceWindowResponse
	GetBody() *ModifyPlanMaintenanceWindowResponseBody
}

type ModifyPlanMaintenanceWindowResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlanMaintenanceWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlanMaintenanceWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlanMaintenanceWindowResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlanMaintenanceWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPlanMaintenanceWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPlanMaintenanceWindowResponse) GetBody() *ModifyPlanMaintenanceWindowResponseBody {
	return s.Body
}

func (s *ModifyPlanMaintenanceWindowResponse) SetHeaders(v map[string]*string) *ModifyPlanMaintenanceWindowResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlanMaintenanceWindowResponse) SetStatusCode(v int32) *ModifyPlanMaintenanceWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlanMaintenanceWindowResponse) SetBody(v *ModifyPlanMaintenanceWindowResponseBody) *ModifyPlanMaintenanceWindowResponse {
	s.Body = v
	return s
}

func (s *ModifyPlanMaintenanceWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
