// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVSwitchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVSwitchesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse
	GetBody() *DescribeVSwitchesResponseBody
}

type DescribeVSwitchesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVSwitchesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVSwitchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVSwitchesResponse) GetBody() *DescribeVSwitchesResponseBody {
	return s.Body
}

func (s *DescribeVSwitchesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchesResponse) SetStatusCode(v int32) *DescribeVSwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

func (s *DescribeVSwitchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
