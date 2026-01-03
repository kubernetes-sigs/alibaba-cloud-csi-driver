// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CancelPhysicalConnectionResponseBody) *CancelPhysicalConnectionResponse
	GetBody() *CancelPhysicalConnectionResponseBody
}

type CancelPhysicalConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *CancelPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPhysicalConnectionResponse) GetBody() *CancelPhysicalConnectionResponseBody {
	return s.Body
}

func (s *CancelPhysicalConnectionResponse) SetHeaders(v map[string]*string) *CancelPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *CancelPhysicalConnectionResponse) SetStatusCode(v int32) *CancelPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPhysicalConnectionResponse) SetBody(v *CancelPhysicalConnectionResponseBody) *CancelPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *CancelPhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
