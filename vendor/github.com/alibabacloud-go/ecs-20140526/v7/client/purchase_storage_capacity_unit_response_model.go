// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseStorageCapacityUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseStorageCapacityUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseStorageCapacityUnitResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseStorageCapacityUnitResponseBody) *PurchaseStorageCapacityUnitResponse
	GetBody() *PurchaseStorageCapacityUnitResponseBody
}

type PurchaseStorageCapacityUnitResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseStorageCapacityUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseStorageCapacityUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseStorageCapacityUnitResponse) GoString() string {
	return s.String()
}

func (s *PurchaseStorageCapacityUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseStorageCapacityUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseStorageCapacityUnitResponse) GetBody() *PurchaseStorageCapacityUnitResponseBody {
	return s.Body
}

func (s *PurchaseStorageCapacityUnitResponse) SetHeaders(v map[string]*string) *PurchaseStorageCapacityUnitResponse {
	s.Headers = v
	return s
}

func (s *PurchaseStorageCapacityUnitResponse) SetStatusCode(v int32) *PurchaseStorageCapacityUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseStorageCapacityUnitResponse) SetBody(v *PurchaseStorageCapacityUnitResponseBody) *PurchaseStorageCapacityUnitResponse {
	s.Body = v
	return s
}

func (s *PurchaseStorageCapacityUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
