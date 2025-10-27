// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleRetrieveJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateLifecycleRetrieveJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateLifecycleRetrieveJobResponseBody
	GetRequestId() *string
}

type CreateLifecycleRetrieveJobResponseBody struct {
	// The ID of the data retrieval task.
	//
	// example:
	//
	// lrj-nfstest-ia-160****853-hshvw
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLifecycleRetrieveJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateLifecycleRetrieveJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetJobId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponseBody) Validate() error {
	return dara.Validate(s)
}
