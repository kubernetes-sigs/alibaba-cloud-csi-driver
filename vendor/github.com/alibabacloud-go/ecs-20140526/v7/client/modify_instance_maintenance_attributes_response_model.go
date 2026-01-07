// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintenanceAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMaintenanceAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMaintenanceAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMaintenanceAttributesResponseBody) *ModifyInstanceMaintenanceAttributesResponse
	GetBody() *ModifyInstanceMaintenanceAttributesResponseBody
}

type ModifyInstanceMaintenanceAttributesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMaintenanceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMaintenanceAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintenanceAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintenanceAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMaintenanceAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMaintenanceAttributesResponse) GetBody() *ModifyInstanceMaintenanceAttributesResponseBody {
	return s.Body
}

func (s *ModifyInstanceMaintenanceAttributesResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintenanceAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesResponse) SetStatusCode(v int32) *ModifyInstanceMaintenanceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesResponse) SetBody(v *ModifyInstanceMaintenanceAttributesResponseBody) *ModifyInstanceMaintenanceAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
