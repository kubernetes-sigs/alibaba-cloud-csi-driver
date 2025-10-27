// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateRecycleBinDeleteJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateRecycleBinDeleteJobResponseBody
	GetRequestId() *string
}

type CreateRecycleBinDeleteJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 8C****C54
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinDeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateRecycleBinDeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetJobId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetRequestId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
