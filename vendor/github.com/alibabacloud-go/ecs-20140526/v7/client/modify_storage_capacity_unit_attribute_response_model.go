// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageCapacityUnitAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStorageCapacityUnitAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStorageCapacityUnitAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStorageCapacityUnitAttributeResponseBody) *ModifyStorageCapacityUnitAttributeResponse
	GetBody() *ModifyStorageCapacityUnitAttributeResponseBody
}

type ModifyStorageCapacityUnitAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStorageCapacityUnitAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStorageCapacityUnitAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageCapacityUnitAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyStorageCapacityUnitAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStorageCapacityUnitAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStorageCapacityUnitAttributeResponse) GetBody() *ModifyStorageCapacityUnitAttributeResponseBody {
	return s.Body
}

func (s *ModifyStorageCapacityUnitAttributeResponse) SetHeaders(v map[string]*string) *ModifyStorageCapacityUnitAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeResponse) SetStatusCode(v int32) *ModifyStorageCapacityUnitAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeResponse) SetBody(v *ModifyStorageCapacityUnitAttributeResponseBody) *ModifyStorageCapacityUnitAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
