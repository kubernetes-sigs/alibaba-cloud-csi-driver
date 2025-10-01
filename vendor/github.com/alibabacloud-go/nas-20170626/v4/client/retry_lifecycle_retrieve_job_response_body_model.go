// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryLifecycleRetrieveJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryLifecycleRetrieveJobResponseBody
	GetRequestId() *string
}

type RetryLifecycleRetrieveJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryLifecycleRetrieveJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryLifecycleRetrieveJobResponseBody) SetRequestId(v string) *RetryLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryLifecycleRetrieveJobResponseBody) Validate() error {
	return dara.Validate(s)
}
