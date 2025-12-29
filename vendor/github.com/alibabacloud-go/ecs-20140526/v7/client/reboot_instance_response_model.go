// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse
	GetBody() *RebootInstanceResponseBody
}

type RebootInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootInstanceResponse) GetBody() *RebootInstanceResponseBody {
	return s.Body
}

func (s *RebootInstanceResponse) SetHeaders(v map[string]*string) *RebootInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootInstanceResponse) SetStatusCode(v int32) *RebootInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstanceResponse) SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse {
	s.Body = v
	return s
}

func (s *RebootInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
