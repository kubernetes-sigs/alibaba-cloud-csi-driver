// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryOrFilePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectoryOrFilePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectoryOrFilePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectoryOrFilePropertiesResponseBody) *GetDirectoryOrFilePropertiesResponse
	GetBody() *GetDirectoryOrFilePropertiesResponseBody
}

type GetDirectoryOrFilePropertiesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryOrFilePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectoryOrFilePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectoryOrFilePropertiesResponse) GetBody() *GetDirectoryOrFilePropertiesResponseBody {
	return s.Body
}

func (s *GetDirectoryOrFilePropertiesResponse) SetHeaders(v map[string]*string) *GetDirectoryOrFilePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) SetStatusCode(v int32) *GetDirectoryOrFilePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) SetBody(v *GetDirectoryOrFilePropertiesResponseBody) *GetDirectoryOrFilePropertiesResponse {
	s.Body = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) Validate() error {
	return dara.Validate(s)
}
