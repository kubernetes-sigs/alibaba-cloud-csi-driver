// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLifecycleRetrieveJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelLifecycleRetrieveJobResponseBody
	GetRequestId() *string
}

type CancelLifecycleRetrieveJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelLifecycleRetrieveJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CancelLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelLifecycleRetrieveJobResponseBody) Validate() error {
	return dara.Validate(s)
}
