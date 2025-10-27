// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinRestoreJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateRecycleBinRestoreJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateRecycleBinRestoreJobResponseBody
	GetRequestId() *string
}

type CreateRecycleBinRestoreJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// rb-10****491ff-r-162****165400
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinRestoreJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateRecycleBinRestoreJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetJobId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetRequestId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponseBody) Validate() error {
	return dara.Validate(s)
}
