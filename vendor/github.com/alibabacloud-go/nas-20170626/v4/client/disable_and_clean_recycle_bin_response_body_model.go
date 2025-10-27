// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAndCleanRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAndCleanRecycleBinResponseBody
	GetRequestId() *string
}

type DisableAndCleanRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9E15E394-38A6-457A-A62A-D9797C9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAndCleanRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAndCleanRecycleBinResponseBody) SetRequestId(v string) *DisableAndCleanRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAndCleanRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}
