// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetFileSystemResponseBody
	GetRequestId() *string
}

type ResetFileSystemResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetFileSystemResponseBody) SetRequestId(v string) *ResetFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
