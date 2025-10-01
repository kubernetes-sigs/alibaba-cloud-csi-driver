// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAndCleanRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAndCleanRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAndCleanRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *DisableAndCleanRecycleBinResponseBody) *DisableAndCleanRecycleBinResponse
	GetBody() *DisableAndCleanRecycleBinResponseBody
}

type DisableAndCleanRecycleBinResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAndCleanRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAndCleanRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAndCleanRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAndCleanRecycleBinResponse) GetBody() *DisableAndCleanRecycleBinResponseBody {
	return s.Body
}

func (s *DisableAndCleanRecycleBinResponse) SetHeaders(v map[string]*string) *DisableAndCleanRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetStatusCode(v int32) *DisableAndCleanRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetBody(v *DisableAndCleanRecycleBinResponseBody) *DisableAndCleanRecycleBinResponse {
	s.Body = v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) Validate() error {
	return dara.Validate(s)
}
