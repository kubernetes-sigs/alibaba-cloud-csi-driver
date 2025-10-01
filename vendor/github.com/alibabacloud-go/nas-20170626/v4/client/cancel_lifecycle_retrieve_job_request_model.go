// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLifecycleRetrieveJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelLifecycleRetrieveJobRequest
	GetJobId() *string
}

type CancelLifecycleRetrieveJobRequest struct {
	// The ID of the data retrieval task.
	//
	// This parameter is required.
	//
	// example:
	//
	// lrj-nfstest-ia-160****853-hshvw
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelLifecycleRetrieveJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelLifecycleRetrieveJobRequest) SetJobId(v string) *CancelLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelLifecycleRetrieveJobRequest) Validate() error {
	return dara.Validate(s)
}
