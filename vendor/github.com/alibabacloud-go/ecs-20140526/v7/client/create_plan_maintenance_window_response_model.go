// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlanMaintenanceWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePlanMaintenanceWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePlanMaintenanceWindowResponse
	GetStatusCode() *int32
	SetBody(v *CreatePlanMaintenanceWindowResponseBody) *CreatePlanMaintenanceWindowResponse
	GetBody() *CreatePlanMaintenanceWindowResponseBody
}

type CreatePlanMaintenanceWindowResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlanMaintenanceWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePlanMaintenanceWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowResponse) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePlanMaintenanceWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePlanMaintenanceWindowResponse) GetBody() *CreatePlanMaintenanceWindowResponseBody {
	return s.Body
}

func (s *CreatePlanMaintenanceWindowResponse) SetHeaders(v map[string]*string) *CreatePlanMaintenanceWindowResponse {
	s.Headers = v
	return s
}

func (s *CreatePlanMaintenanceWindowResponse) SetStatusCode(v int32) *CreatePlanMaintenanceWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlanMaintenanceWindowResponse) SetBody(v *CreatePlanMaintenanceWindowResponseBody) *CreatePlanMaintenanceWindowResponse {
	s.Body = v
	return s
}

func (s *CreatePlanMaintenanceWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
