// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientFromBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveClientFromBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveClientFromBlackListResponse
	GetStatusCode() *int32
	SetBody(v *RemoveClientFromBlackListResponseBody) *RemoveClientFromBlackListResponse
	GetBody() *RemoveClientFromBlackListResponseBody
}

type RemoveClientFromBlackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClientFromBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClientFromBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientFromBlackListResponse) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveClientFromBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveClientFromBlackListResponse) GetBody() *RemoveClientFromBlackListResponseBody {
	return s.Body
}

func (s *RemoveClientFromBlackListResponse) SetHeaders(v map[string]*string) *RemoveClientFromBlackListResponse {
	s.Headers = v
	return s
}

func (s *RemoveClientFromBlackListResponse) SetStatusCode(v int32) *RemoveClientFromBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClientFromBlackListResponse) SetBody(v *RemoveClientFromBlackListResponseBody) *RemoveClientFromBlackListResponse {
	s.Body = v
	return s
}

func (s *RemoveClientFromBlackListResponse) Validate() error {
	return dara.Validate(s)
}
