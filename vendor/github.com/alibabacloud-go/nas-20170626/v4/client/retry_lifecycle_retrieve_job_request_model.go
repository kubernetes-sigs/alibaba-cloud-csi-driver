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
	// The data retrieval task ID.
	//
	// **Scenarios**
	//
	// Call this operation to retry a data retrieval task that has entered the `failed` state. Common causes for a task to enter the `failed` state include:
	//
	// - A backend error occurred during data retrieval from the InfrequentAccess or Archive storage tier.
	//
	// - The data retrieval request timed out.
	//
	// - A temporary storage tier failure or network exception occurred.
	//
	// **Before you begin**
	//
	// Before calling this operation, call [ListLifecycleRetrieveJobs](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-listlifecycleretrievejobs) to query the task list, confirm that the target task is in the `failed` state, and obtain the JobId of the task you want to retry.
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
