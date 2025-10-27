// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRecycleBinJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelRecycleBinJobRequest
	GetJobId() *string
}

type CancelRecycleBinJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rb-15****ed-r-1625****2441
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelRecycleBinJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRecycleBinJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelRecycleBinJobRequest) SetJobId(v string) *CancelRecycleBinJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelRecycleBinJobRequest) Validate() error {
	return dara.Validate(s)
}
