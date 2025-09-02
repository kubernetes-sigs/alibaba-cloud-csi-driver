// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReimageNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReimageNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReimageNodesResponse
	GetStatusCode() *int32
	SetBody(v *ReimageNodesResponseBody) *ReimageNodesResponse
	GetBody() *ReimageNodesResponseBody
}

type ReimageNodesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReimageNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReimageNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesResponse) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReimageNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReimageNodesResponse) GetBody() *ReimageNodesResponseBody {
	return s.Body
}

func (s *ReimageNodesResponse) SetHeaders(v map[string]*string) *ReimageNodesResponse {
	s.Headers = v
	return s
}

func (s *ReimageNodesResponse) SetStatusCode(v int32) *ReimageNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReimageNodesResponse) SetBody(v *ReimageNodesResponseBody) *ReimageNodesResponse {
	s.Body = v
	return s
}

func (s *ReimageNodesResponse) Validate() error {
	return dara.Validate(s)
}
