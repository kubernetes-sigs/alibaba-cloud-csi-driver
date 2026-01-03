// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEncryptionByDefaultStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskEncryptionByDefaultStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskEncryptionByDefaultStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskEncryptionByDefaultStatusResponseBody) *DescribeDiskEncryptionByDefaultStatusResponse
	GetBody() *DescribeDiskEncryptionByDefaultStatusResponseBody
}

type DescribeDiskEncryptionByDefaultStatusResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskEncryptionByDefaultStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskEncryptionByDefaultStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEncryptionByDefaultStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) GetBody() *DescribeDiskEncryptionByDefaultStatusResponseBody {
	return s.Body
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) SetHeaders(v map[string]*string) *DescribeDiskEncryptionByDefaultStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) SetStatusCode(v int32) *DescribeDiskEncryptionByDefaultStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) SetBody(v *DescribeDiskEncryptionByDefaultStatusResponseBody) *DescribeDiskEncryptionByDefaultStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
