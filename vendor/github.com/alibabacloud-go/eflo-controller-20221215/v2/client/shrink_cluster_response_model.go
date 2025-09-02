// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShrinkClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShrinkClusterResponse
	GetStatusCode() *int32
	SetBody(v *ShrinkClusterResponseBody) *ShrinkClusterResponse
	GetBody() *ShrinkClusterResponseBody
}

type ShrinkClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShrinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShrinkClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterResponse) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShrinkClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShrinkClusterResponse) GetBody() *ShrinkClusterResponseBody {
	return s.Body
}

func (s *ShrinkClusterResponse) SetHeaders(v map[string]*string) *ShrinkClusterResponse {
	s.Headers = v
	return s
}

func (s *ShrinkClusterResponse) SetStatusCode(v int32) *ShrinkClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ShrinkClusterResponse) SetBody(v *ShrinkClusterResponseBody) *ShrinkClusterResponse {
	s.Body = v
	return s
}

func (s *ShrinkClusterResponse) Validate() error {
	return dara.Validate(s)
}
