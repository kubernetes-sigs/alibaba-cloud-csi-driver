// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDefaultKMSKeyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskDefaultKMSKeyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskDefaultKMSKeyIdResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskDefaultKMSKeyIdResponseBody) *ModifyDiskDefaultKMSKeyIdResponse
	GetBody() *ModifyDiskDefaultKMSKeyIdResponseBody
}

type ModifyDiskDefaultKMSKeyIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskDefaultKMSKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskDefaultKMSKeyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDefaultKMSKeyIdResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) GetBody() *ModifyDiskDefaultKMSKeyIdResponseBody {
	return s.Body
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) SetHeaders(v map[string]*string) *ModifyDiskDefaultKMSKeyIdResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) SetStatusCode(v int32) *ModifyDiskDefaultKMSKeyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) SetBody(v *ModifyDiskDefaultKMSKeyIdResponseBody) *ModifyDiskDefaultKMSKeyIdResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
