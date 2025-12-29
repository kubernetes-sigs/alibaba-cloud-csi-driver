// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageSetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStorageSetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStorageSetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStorageSetAttributeResponseBody) *ModifyStorageSetAttributeResponse
	GetBody() *ModifyStorageSetAttributeResponseBody
}

type ModifyStorageSetAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStorageSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStorageSetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyStorageSetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStorageSetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStorageSetAttributeResponse) GetBody() *ModifyStorageSetAttributeResponseBody {
	return s.Body
}

func (s *ModifyStorageSetAttributeResponse) SetHeaders(v map[string]*string) *ModifyStorageSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyStorageSetAttributeResponse) SetStatusCode(v int32) *ModifyStorageSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStorageSetAttributeResponse) SetBody(v *ModifyStorageSetAttributeResponseBody) *ModifyStorageSetAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyStorageSetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
