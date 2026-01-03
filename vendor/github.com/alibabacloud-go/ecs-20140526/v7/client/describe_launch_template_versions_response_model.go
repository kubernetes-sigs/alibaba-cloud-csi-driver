// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplateVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLaunchTemplateVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLaunchTemplateVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLaunchTemplateVersionsResponseBody) *DescribeLaunchTemplateVersionsResponse
	GetBody() *DescribeLaunchTemplateVersionsResponseBody
}

type DescribeLaunchTemplateVersionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLaunchTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLaunchTemplateVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLaunchTemplateVersionsResponse) GetBody() *DescribeLaunchTemplateVersionsResponseBody {
	return s.Body
}

func (s *DescribeLaunchTemplateVersionsResponse) SetHeaders(v map[string]*string) *DescribeLaunchTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponse) SetStatusCode(v int32) *DescribeLaunchTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponse) SetBody(v *DescribeLaunchTemplateVersionsResponseBody) *DescribeLaunchTemplateVersionsResponse {
	s.Body = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
