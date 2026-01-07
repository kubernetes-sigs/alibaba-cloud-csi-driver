// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedeployInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedeployInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RedeployInstanceResponseBody) *RedeployInstanceResponse
	GetBody() *RedeployInstanceResponseBody
}

type RedeployInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedeployInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedeployInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RedeployInstanceResponse) GoString() string {
	return s.String()
}

func (s *RedeployInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedeployInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedeployInstanceResponse) GetBody() *RedeployInstanceResponseBody {
	return s.Body
}

func (s *RedeployInstanceResponse) SetHeaders(v map[string]*string) *RedeployInstanceResponse {
	s.Headers = v
	return s
}

func (s *RedeployInstanceResponse) SetStatusCode(v int32) *RedeployInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RedeployInstanceResponse) SetBody(v *RedeployInstanceResponseBody) *RedeployInstanceResponse {
	s.Body = v
	return s
}

func (s *RedeployInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
