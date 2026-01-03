// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskDefaultKMSKeyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDiskDefaultKMSKeyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDiskDefaultKMSKeyIdResponse
	GetStatusCode() *int32
	SetBody(v *ResetDiskDefaultKMSKeyIdResponseBody) *ResetDiskDefaultKMSKeyIdResponse
	GetBody() *ResetDiskDefaultKMSKeyIdResponseBody
}

type ResetDiskDefaultKMSKeyIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDiskDefaultKMSKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDiskDefaultKMSKeyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskDefaultKMSKeyIdResponse) GoString() string {
	return s.String()
}

func (s *ResetDiskDefaultKMSKeyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDiskDefaultKMSKeyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDiskDefaultKMSKeyIdResponse) GetBody() *ResetDiskDefaultKMSKeyIdResponseBody {
	return s.Body
}

func (s *ResetDiskDefaultKMSKeyIdResponse) SetHeaders(v map[string]*string) *ResetDiskDefaultKMSKeyIdResponse {
	s.Headers = v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdResponse) SetStatusCode(v int32) *ResetDiskDefaultKMSKeyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdResponse) SetBody(v *ResetDiskDefaultKMSKeyIdResponseBody) *ResetDiskDefaultKMSKeyIdResponse {
	s.Body = v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
