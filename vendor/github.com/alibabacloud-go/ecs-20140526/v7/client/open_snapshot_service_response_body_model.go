// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSnapshotServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenSnapshotServiceResponseBody
	GetRequestId() *string
}

type OpenSnapshotServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C36234E8-4C67-5F6C-8C07-F51B2EE2C560
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenSnapshotServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenSnapshotServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSnapshotServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenSnapshotServiceResponseBody) SetRequestId(v string) *OpenSnapshotServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenSnapshotServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
