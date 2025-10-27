// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToFilesystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachVscToFilesystemsResponseBody
	GetRequestId() *string
}

type AttachVscToFilesystemsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVscToFilesystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachVscToFilesystemsResponseBody) SetRequestId(v string) *AttachVscToFilesystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachVscToFilesystemsResponseBody) Validate() error {
	return dara.Validate(s)
}
