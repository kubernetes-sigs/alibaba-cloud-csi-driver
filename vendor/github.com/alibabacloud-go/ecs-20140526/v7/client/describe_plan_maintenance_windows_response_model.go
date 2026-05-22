// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlanMaintenanceWindowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlanMaintenanceWindowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlanMaintenanceWindowsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlanMaintenanceWindowsResponseBody) *DescribePlanMaintenanceWindowsResponse
	GetBody() *DescribePlanMaintenanceWindowsResponseBody
}

type DescribePlanMaintenanceWindowsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlanMaintenanceWindowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlanMaintenanceWindowsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlanMaintenanceWindowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlanMaintenanceWindowsResponse) GetBody() *DescribePlanMaintenanceWindowsResponseBody {
	return s.Body
}

func (s *DescribePlanMaintenanceWindowsResponse) SetHeaders(v map[string]*string) *DescribePlanMaintenanceWindowsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponse) SetStatusCode(v int32) *DescribePlanMaintenanceWindowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponse) SetBody(v *DescribePlanMaintenanceWindowsResponseBody) *DescribePlanMaintenanceWindowsResponse {
	s.Body = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
