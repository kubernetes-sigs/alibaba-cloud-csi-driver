// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryLifecycleRetrieveJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *RetryLifecycleRetrieveJobRequest
	GetJobId() *string
}

type RetryLifecycleRetrieveJobRequest struct {
	// The ID of the data retrieval task.
	//
	// This parameter is required.
	//
	// example:
	//
	// lrj-nfstest-ia-160****853-hshvw
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RetryLifecycleRetrieveJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *RetryLifecycleRetrieveJobRequest) SetJobId(v string) *RetryLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

func (s *RetryLifecycleRetrieveJobRequest) Validate() error {
	return dara.Validate(s)
}
