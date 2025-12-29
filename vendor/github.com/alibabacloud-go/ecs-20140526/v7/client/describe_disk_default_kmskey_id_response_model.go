// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskDefaultKMSKeyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskDefaultKMSKeyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskDefaultKMSKeyIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskDefaultKMSKeyIdResponseBody) *DescribeDiskDefaultKMSKeyIdResponse
	GetBody() *DescribeDiskDefaultKMSKeyIdResponseBody
}

type DescribeDiskDefaultKMSKeyIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskDefaultKMSKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskDefaultKMSKeyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskDefaultKMSKeyIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) GetBody() *DescribeDiskDefaultKMSKeyIdResponseBody {
	return s.Body
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) SetHeaders(v map[string]*string) *DescribeDiskDefaultKMSKeyIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) SetStatusCode(v int32) *DescribeDiskDefaultKMSKeyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) SetBody(v *DescribeDiskDefaultKMSKeyIdResponseBody) *DescribeDiskDefaultKMSKeyIdResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
