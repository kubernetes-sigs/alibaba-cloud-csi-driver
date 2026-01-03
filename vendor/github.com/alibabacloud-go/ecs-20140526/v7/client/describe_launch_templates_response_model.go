// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLaunchTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLaunchTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLaunchTemplatesResponseBody) *DescribeLaunchTemplatesResponse
	GetBody() *DescribeLaunchTemplatesResponseBody
}

type DescribeLaunchTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLaunchTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLaunchTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLaunchTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLaunchTemplatesResponse) GetBody() *DescribeLaunchTemplatesResponseBody {
	return s.Body
}

func (s *DescribeLaunchTemplatesResponse) SetHeaders(v map[string]*string) *DescribeLaunchTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLaunchTemplatesResponse) SetStatusCode(v int32) *DescribeLaunchTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLaunchTemplatesResponse) SetBody(v *DescribeLaunchTemplatesResponseBody) *DescribeLaunchTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeLaunchTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
