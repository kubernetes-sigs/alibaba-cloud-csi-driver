// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlanMaintenanceWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePlanMaintenanceWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePlanMaintenanceWindowResponse
	GetStatusCode() *int32
	SetBody(v *DeletePlanMaintenanceWindowResponseBody) *DeletePlanMaintenanceWindowResponse
	GetBody() *DeletePlanMaintenanceWindowResponseBody
}

type DeletePlanMaintenanceWindowResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlanMaintenanceWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlanMaintenanceWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePlanMaintenanceWindowResponse) GoString() string {
	return s.String()
}

func (s *DeletePlanMaintenanceWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePlanMaintenanceWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePlanMaintenanceWindowResponse) GetBody() *DeletePlanMaintenanceWindowResponseBody {
	return s.Body
}

func (s *DeletePlanMaintenanceWindowResponse) SetHeaders(v map[string]*string) *DeletePlanMaintenanceWindowResponse {
	s.Headers = v
	return s
}

func (s *DeletePlanMaintenanceWindowResponse) SetStatusCode(v int32) *DeletePlanMaintenanceWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlanMaintenanceWindowResponse) SetBody(v *DeletePlanMaintenanceWindowResponseBody) *DeletePlanMaintenanceWindowResponse {
	s.Body = v
	return s
}

func (s *DeletePlanMaintenanceWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
