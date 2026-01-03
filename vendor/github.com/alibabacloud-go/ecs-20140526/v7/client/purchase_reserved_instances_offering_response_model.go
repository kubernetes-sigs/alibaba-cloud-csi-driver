// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseReservedInstancesOfferingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseReservedInstancesOfferingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseReservedInstancesOfferingResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseReservedInstancesOfferingResponseBody) *PurchaseReservedInstancesOfferingResponse
	GetBody() *PurchaseReservedInstancesOfferingResponseBody
}

type PurchaseReservedInstancesOfferingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseReservedInstancesOfferingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseReservedInstancesOfferingResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseReservedInstancesOfferingResponse) GoString() string {
	return s.String()
}

func (s *PurchaseReservedInstancesOfferingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseReservedInstancesOfferingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseReservedInstancesOfferingResponse) GetBody() *PurchaseReservedInstancesOfferingResponseBody {
	return s.Body
}

func (s *PurchaseReservedInstancesOfferingResponse) SetHeaders(v map[string]*string) *PurchaseReservedInstancesOfferingResponse {
	s.Headers = v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponse) SetStatusCode(v int32) *PurchaseReservedInstancesOfferingResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponse) SetBody(v *PurchaseReservedInstancesOfferingResponseBody) *PurchaseReservedInstancesOfferingResponse {
	s.Body = v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
