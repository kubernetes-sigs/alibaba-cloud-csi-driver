// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithSAMLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeRoleWithSAMLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeRoleWithSAMLResponse
	GetStatusCode() *int32
	SetBody(v *AssumeRoleWithSAMLResponseBody) *AssumeRoleWithSAMLResponse
	GetBody() *AssumeRoleWithSAMLResponseBody
}

type AssumeRoleWithSAMLResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleWithSAMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleWithSAMLResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeRoleWithSAMLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeRoleWithSAMLResponse) GetBody() *AssumeRoleWithSAMLResponseBody {
	return s.Body
}

func (s *AssumeRoleWithSAMLResponse) SetHeaders(v map[string]*string) *AssumeRoleWithSAMLResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetStatusCode(v int32) *AssumeRoleWithSAMLResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetBody(v *AssumeRoleWithSAMLResponseBody) *AssumeRoleWithSAMLResponse {
	s.Body = v
	return s
}

func (s *AssumeRoleWithSAMLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
