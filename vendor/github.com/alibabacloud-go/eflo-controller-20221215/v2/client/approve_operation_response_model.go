// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveOperationResponse
	GetStatusCode() *int32
	SetBody(v *ApproveOperationResponseBody) *ApproveOperationResponse
	GetBody() *ApproveOperationResponseBody
}

type ApproveOperationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveOperationResponse) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveOperationResponse) GetBody() *ApproveOperationResponseBody {
	return s.Body
}

func (s *ApproveOperationResponse) SetHeaders(v map[string]*string) *ApproveOperationResponse {
	s.Headers = v
	return s
}

func (s *ApproveOperationResponse) SetStatusCode(v int32) *ApproveOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOperationResponse) SetBody(v *ApproveOperationResponseBody) *ApproveOperationResponse {
	s.Body = v
	return s
}

func (s *ApproveOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
