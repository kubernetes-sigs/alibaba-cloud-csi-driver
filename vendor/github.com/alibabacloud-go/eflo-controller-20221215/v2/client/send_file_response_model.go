// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendFileResponse
	GetStatusCode() *int32
	SetBody(v *SendFileResponseBody) *SendFileResponse
	GetBody() *SendFileResponseBody
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendFileResponse) GetBody() *SendFileResponseBody {
	return s.Body
}

func (s *SendFileResponse) SetHeaders(v map[string]*string) *SendFileResponse {
	s.Headers = v
	return s
}

func (s *SendFileResponse) SetStatusCode(v int32) *SendFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFileResponse) SetBody(v *SendFileResponseBody) *SendFileResponse {
	s.Body = v
	return s
}

func (s *SendFileResponse) Validate() error {
	return dara.Validate(s)
}
