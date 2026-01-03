// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreatePhysicalConnectionResponseBody) *CreatePhysicalConnectionResponse
	GetBody() *CreatePhysicalConnectionResponseBody
}

type CreatePhysicalConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePhysicalConnectionResponse) GetBody() *CreatePhysicalConnectionResponseBody {
	return s.Body
}

func (s *CreatePhysicalConnectionResponse) SetHeaders(v map[string]*string) *CreatePhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalConnectionResponse) SetStatusCode(v int32) *CreatePhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhysicalConnectionResponse) SetBody(v *CreatePhysicalConnectionResponseBody) *CreatePhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *CreatePhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
