// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientToBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddClientToBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddClientToBlackListResponse
	GetStatusCode() *int32
	SetBody(v *AddClientToBlackListResponseBody) *AddClientToBlackListResponse
	GetBody() *AddClientToBlackListResponseBody
}

type AddClientToBlackListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddClientToBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddClientToBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddClientToBlackListResponse) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddClientToBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddClientToBlackListResponse) GetBody() *AddClientToBlackListResponseBody {
	return s.Body
}

func (s *AddClientToBlackListResponse) SetHeaders(v map[string]*string) *AddClientToBlackListResponse {
	s.Headers = v
	return s
}

func (s *AddClientToBlackListResponse) SetStatusCode(v int32) *AddClientToBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClientToBlackListResponse) SetBody(v *AddClientToBlackListResponseBody) *AddClientToBlackListResponse {
	s.Body = v
	return s
}

func (s *AddClientToBlackListResponse) Validate() error {
	return dara.Validate(s)
}
